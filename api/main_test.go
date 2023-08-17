package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	k8stesting "k8s.io/client-go/testing"
	"net/http"
	"testing"
)

func TestGetPodsInfo(t *testing.T) {
	// モックPodのリストを作成
	mockPods := &corev1.PodList{
		Items: []corev1.Pod{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-1",
					Namespace: "default",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-2",
					Namespace: "default",
				},
			},
		},
	}

	// fake clientsetを作成
	clientset := fake.NewSimpleClientset()

	// Podのリストを取得するリクエストに対してモックPodのリストを返すように設定
	clientset.PrependReactor("list", "pods", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, mockPods, nil
	})

	// getPodsInfo関数をテスト
	podNames, err := getPodsInfoWithClientset(clientset.CoreV1(), "")
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	expectedPodNames := []string{"pod-1", "pod-2"}
	for i, name := range podNames {
		if name != expectedPodNames[i] {
			t.Errorf("Expected pod name %s, but got %s", expectedPodNames[i], name)
		}
	}
}

// clientsetを引数として受け取る関数
func getPodsInfoWithClientset(clientset v1.CoreV1Interface, namespace string) ([]string, error) {
	// get pods in the specified namespace or all the namespaces by omitting namespace
	pods, err := clientset.Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podNames []string
	// list all pods
	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames, nil
}

func TestHandleGetPodsInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-pods-info", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)

	// TODO: リクエストクエリの設定、レスポンスの取得、レスポンスの確認を行います...
}
