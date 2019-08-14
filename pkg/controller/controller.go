package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	v1alpha1 "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/apis/azurequeue/v1alpha1"
	azurev1alpha1 "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/clientset/versioned"
	informers "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/informers/externalversions/azurequeue/v1alpha1"
	listers "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/listers/azurequeue/v1alpha1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type AzureStorageQueueController struct {
	azqueueLister   listers.AzureQueueLister
	azqueueInformer informers.AzureQueueInformer
	queue           workqueue.RateLimitingInterface
}

func NewAzureQueueController(client *azurev1alpha1.Clientset, queueInformer informers.AzureQueueInformer) AzureStorageQueueController {

	c := AzureStorageQueueController{
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "azqueuesync"),
	}

	queueInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Printf("AZ Queue added %v", obj)
			c.enqueue(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			log.Print("AZ Queue Updated")
			c.enqueue(newObj)
		},
		DeleteFunc: func(obj interface{}) {
			log.Print("AZ Queue deleted")
			c.enqueue(obj)
		},
	})

	return c
}

func (c *AzureStorageQueueController) enqueue(obj interface{}) {
	var key string
	var err error

	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(fmt.Errorf("Enqueue failed with: %v", err))
	}

	c.queue.Add(key)
}

func (c *AzureStorageQueueController) Run(stop <-chan struct{}) {
	var wg sync.WaitGroup

	defer func() {
		c.queue.ShutDown()
		wg.Wait()
	}()

	go func() {
		wait.Until(c.runWorker, time.Second, stop)
		wg.Done()
	}()

	log.Println("Waiting for stop signal")
	<-stop
	log.Println("Stop signal")
}

func (c *AzureStorageQueueController) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *AzureStorageQueueController) processNextWorkItem() bool {
	key, quit := c.queue.Get()

	if quit {
		return false
	}

	defer c.queue.Done(key)

	err := c.doSync()

	if err == nil {
		c.queue.Forget(key)
		return true
	}

	runtime.HandleError(fmt.Errorf("doSync failed with: %v", err))

	c.queue.AddRateLimited(key)

	return true
}

func (c *AzureStorageQueueController) doSync(key string) error {

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("doSync failed with: %v", err))
		return err
	}

	azqueue, err := c.azqueueLister.AzureQueues(namespace).Get(name)
	if err != nil {
		runtime.HandleError(fmt.Errorf("doSync failed with: %v", err))
		return err
	}

	err = c.createQueue(azqueue)
	return err
}

func (c *AzureStorageQueueController) createQueue(obj *v1alpha1.AzureQueue) error {

	connectionString := c.getSecret(
		obj.Spec.ConnectionStringSecretRef.Name, obj.Spec.ConnectionStringSecretRef.Key)

	accountName, accountKey, err := c.parseConnectionString(connectionString)

	u, err := url.Parse(fmt.Sprintf("https://%s.queue.core.windows.net/queue5", accountName))
	if err != nil {
		return err
	}

	credential, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return err
	}

	pipeline := azqueue.NewPipeline(credential, azqueue.PipelineOptions{})
	queueURL := azqueue.NewQueueURL(*u, pipeline)

	ctx := context.Background()
	_, err = queueURL.Create(ctx, azqueue.Metadata{})

	return err
}

func (c *AzureStorageQueueController) parseConnectionString(connectionString string) (string, string, error) {
	parts := strings.Split(connectionString, ";")

	var endpointProtocol, name, key, endpointSuffix string
	for _, v := range parts {
		if strings.HasPrefix(v, "DefaultEndpointsProtocol") {
			protocolParts := strings.SplitN(v, "=", 2)
			if len(protocolParts) == 2 {
				endpointProtocol = protocolParts[1]
			}
		} else if strings.HasPrefix(v, "AccountName") {
			accountParts := strings.SplitN(v, "=", 2)
			if len(accountParts) == 2 {
				name = accountParts[1]
			}
		} else if strings.HasPrefix(v, "AccountKey") {
			keyParts := strings.SplitN(v, "=", 2)
			if len(keyParts) == 2 {
				key = keyParts[1]
			}
		} else if strings.HasPrefix(v, "EndpointSuffix") {
			suffixParts := strings.SplitN(v, "=", 2)
			if len(suffixParts) == 2 {
				endpointSuffix = suffixParts[1]
			}
		}
	}
	if name == "" || key == "" || endpointProtocol == "" || endpointSuffix == "" {
		return "", "", errors.New("Can't parse storage connection string")
	}

	return name, key, nil
}
