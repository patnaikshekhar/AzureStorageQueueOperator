package main

import (
	"flag"
	"log"
	"time"

	"github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/clientset/versioned"
	informers "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/informers/externalversions"
	controller "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/controller"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Location of the kubeconfig file")
	flag.Parse()

	var (
		config *rest.Config
		err    error
	)

	if kubeconfig == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		log.Fatal(err)
	}

	client := versioned.NewForConfigOrDie(config)

	sharedinformers := informers.NewSharedInformerFactory(client, 10*time.Minute)
	controller := controller.NewAzureQueueController(client, sharedinformers.Azure().V1alpha1().AzureQueues())

	sharedinformers.Start(nil)
	log.Println(controller)
}
