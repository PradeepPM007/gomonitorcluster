package main

import (
	"flag"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "", "My kubeconfig path")
	mynamespace := flag.String("mynamespace", "", "My Namespace Name")
	mypod := flag.String("mypod", "", "My Pod Name")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("creating ClientSet")
	ClientSet := kubernetes.NewForConfigOrDie(config)

	if ClientSet == nil {
		panic("ClientSet is empty")
	}
	ListPodsOnFilters(ClientSet, *mynamespace, *mypod)
}
