package main

import (
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
)

// main.goの内容から抽出した関数
func getPodsInfo(clientset kubernetes.Interface, namespace *string) ([]string, error) {

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	// pods, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	pods, err := clientset.CoreV1().Pods(*namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	var podNames []string
	// list all pods
	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %s, Pod Name: %s\n", pod.Namespace, pod.Name)
		podNames = append(podNames, fmt.Sprintf("%s/%s", pod.Namespace, pod.Name))
	}
	return podNames, nil // podNamesは取得したPodの名前のリスト
}

func handleGetPodsInfo(w http.ResponseWriter, r *http.Request) {
	//kubeconfig := r.URL.Query().Get("kubeconfig")
	ns := r.URL.Query().Get("namespace")

	// デフォルトのkubeconfigを使用してconfigを取得
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// var clientset kubernetes.Interface
	// clientset = ...
	clientset, err := kubernetes.NewForConfig(config)

	pods, err := getPodsInfo(clientset, &ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pods)
}

func main() {
	http.HandleFunc("/get-pods-info", handleGetPodsInfo)
	http.ListenAndServe(":8080", nil)
}
