generate:
	$$GOPATH/src/k8s.io/code-generator/generate-groups.sh all \
    	"github.com/patnaikshekhar/azure_queue_operator/pkg/client" \
    	"github.com/patnaikshekhar/azure_queue_operator/pkg/apis" \
    	azurequeue:v1alpha1