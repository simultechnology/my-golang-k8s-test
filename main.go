package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	var namespace *string
	kubeconfig = flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	namespace = flag.String("namespace", "", "(optional) specify namespace to get pods in particular namespace")
	flag.Parse()

	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//for _, ns := range namespaces.Items {
	//	fmt.Printf("namespace: %s\n", ns.Name)
	//}

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	// pods, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	pods, err := clientset.CoreV1().Pods(*namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// list all pods
	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %s, Pod Name: %s\n", pod.Namespace, pod.Name)
	}
}
