package main

import (
	"flag"
	"log"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig string
	var namespace string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to Kubernetes config file. Defaults to in-cluster config.")
	flag.StringVar(&namespace, "namespace", "", "Namespace where to search for ")
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	config := connect(kubeconfig)

	k8sSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
	}
	factory := informers.NewSharedInformerFactory(k8sSet, 0)
	watchSvc(factory, namespace)
	time.Sleep(50 * time.Second)

	// certSet, err := clientset.NewForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }

	// informer := v1.NewDeploymentInformer(k8sSet, namespace, 60*time.Second, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	// informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc:    func(obj interface{}) { print(obj, "add") },
	// 	UpdateFunc: func(obj interface{}, update interface{}) { print(obj, "update") },
	// 	DeleteFunc: func(obj interface{}) { print(obj, "del") },
	// })
	// go EndpointServer()
	// informer.Run(wait.NeverStop)

}

func connect(kubeconfig string) *rest.Config {
	var config *rest.Config
	var err error
	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Panic(err)
		}
	}

	return config
}
