package main

import (
	"context"
	"fmt"

	"github.com/rodaine/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPodsOnFilters(clientset *kubernetes.Clientset, namespace string, podname string) {

	listops := metav1.ListOptions{}
	if len(podname) > 0 {
		listops.FieldSelector = fmt.Sprintf("metadata.name=%s", podname)
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), listops)
	if err != nil {
		panic(err)
	}
	tbl := table.New("NameSpace", "PodName", "Count")
	tbl.AddRow("------", "------", "------")
	for _, item := range pods.Items {
		var count int32
		for _, status := range item.Status.ContainerStatuses {
			count += status.RestartCount
		}
		tbl.AddRow(item.GetNamespace(), item.GetName(), count)
	}
	tbl.Print()
}
