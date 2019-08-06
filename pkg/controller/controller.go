package controller

import (
	azurev1alpha1 "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/clientset/versioned"
	informers "github.com/patnaikshekhar/AzureStorageQueueOperator/pkg/client/informers/externalversions/azurequeue/v1alpha1"
)

type AzureStorageQueueController struct {

}

func NewAzureQueueController(client *azurev1alpha1.Clientset, queueInformer informers.AzureQueueInformer) AzureStorageQueueController {
	return AzureStorageQueueController{}
}

