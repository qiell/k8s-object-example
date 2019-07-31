package main

import (
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	// config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientSet.AppsV1().Deployments("weave").Get("weave-scope-app", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	ds, err := clientSet.AppsV1().DaemonSets("weave").Get("weave-scope-agent", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Deployment name ", deployment.GetName(), "\nDS Name ", ds.GetName())
	GetObject(deployment)
	GetObject(ds)
}

func GetObject(object metav1.Object) {
	fmt.Println(object.GetName())
}
