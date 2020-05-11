package main

import (
	"flag"
	"log"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	clientset "github.com/jetstack/cert-manager/pkg/client/clientset/versioned"
)

func main() {
	var kubeconfig string
	var namespace string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to Kubernetes config file. Defaults to in-cluster config.")
	flag.StringVar(&namespace, "namespace", "", "Namespace where to search for. Default is all namespaces.")
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	config := connect(kubeconfig)

	k8sSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
	}
	factory := informers.NewSharedInformerFactory(k8sSet, 0)
	// watchSvc(factory, namespace)

	certSet, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	certInterface := certSet.CertmanagerV1alpha3()
	if err != nil {
		log.Println(err)
	}

	go EndpointServer()

	informer := factory.Core().V1().Services().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	defer runtime.HandleCrash()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) { listService(informer.GetIndexer().List(), certInterface) },
		UpdateFunc: func(obj interface{}, update interface{}) { listService(informer.GetIndexer().List(), certInterface) },
		// DeleteFunc: func(obj interface{}) { certList = Convert(informer.GetIndexer().List()) },
	})
	informer.Run(stopper)
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
