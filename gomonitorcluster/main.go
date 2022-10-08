package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	homepath, _ := os.UserHomeDir()
	kubeconfig := flag.String("kubeconfig", filepath.Join(homepath, ".kube", "config"), "My kubeconfig path")
	mynamespace := flag.String("mynamespace", "", "My Namespace Name")
	mypod := flag.String("mypod", "", "My Pod Name")
	flag.Parse()

	log.Println("---------------> config path:", *kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("---------------> creating ClientSet")
	ClientSet := kubernetes.NewForConfigOrDie(config)

	if ClientSet == nil {
		panic("ClientSet is empty")
	}
	ListPodsOnFilters(ClientSet, *mynamespace, *mypod)
}
