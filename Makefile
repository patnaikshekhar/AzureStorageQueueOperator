run: build
	./azure-storage-queue-operator --kubeconfig=$$HOME/.kube/config

build:
	dep ensure
	go fmt
	go build -o azure-storage-queue-operator

generate:
	$$GOPATH/src/k8s.io/code-generator/generate-groups.sh all \
    	"github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client" \
    	"github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/apis" \
    	azurequeue:v1alpha1