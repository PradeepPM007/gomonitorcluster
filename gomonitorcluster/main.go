package main

import (
	"flag"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "", "My kubeconfig path")
	mynamespace := flag.String("mynamespace", "", "My Namespace")
	mypod := flag.String("mypod", "", "My Pod")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}
	ClientSet := kubernetes.NewForConfigOrDie(config)

	if ClientSet == nil {
		log.Println("ClientSet is empty")
		return
	}

	ListPodsOnFilters(ClientSet, *mynamespace, *mypod)

}
