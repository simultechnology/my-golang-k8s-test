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

// getPodsInfo retrieves the names of all pods in the specified namespace.
// If no namespace is specified (empty string), it retrieves pods from all namespaces.
// It uses the provided Kubernetes clientset to make API calls.
// Returns a list of pod names in the format "namespace/pod-name".
func getPodsInfo(clientset kubernetes.Interface, namespace *string) ([]string, error) {
	pods, err := clientset.CoreV1().Pods(*namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	var podNames []string
	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %s, Pod Name: %s\n", pod.Namespace, pod.Name)
		podNames = append(podNames, fmt.Sprintf("%s/%s", pod.Namespace, pod.Name))
	}
	return podNames, nil
}

// handleGetPodsInfo is an HTTP handler that responds with a list of pod names.
// The namespace can be specified as a query parameter. If not specified, it fetches pods from all namespaces.
func handleGetPodsInfo(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")

	// Create a Kubernetes clientset using the default kubeconfig.
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pods, err := getPodsInfo(clientset, &ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pods)
}

// main initializes the HTTP server and sets up routes.
func main() {
	http.HandleFunc("/get-pods-info", handleGetPodsInfo)
	http.ListenAndServe(":8080", nil)
}