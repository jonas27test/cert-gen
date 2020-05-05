package main

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func watchSvc(factory informers.SharedInformerFactory, namespace string) {
	informer := factory.Core().V1().Services().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	defer runtime.HandleCrash()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) { informer.GetIndexer().List() },
		UpdateFunc: func(obj interface{}, update interface{}) { certList = Convert(informer.GetIndexer().List()) },
		DeleteFunc: func(obj interface{}) { certList = Convert(informer.GetIndexer().List()) },
	})
	go informer.Run(stopper)
}
