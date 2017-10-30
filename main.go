package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/appscode/go/log"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"github.com/tamalsaha/go-oneliners"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	masterURL := ""
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}

	kc := kubernetes.NewForConfigOrDie(config)
	nodes, err := kc.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	if nodes != nil {
		for _, node := range nodes.Items {
			oneliners.FILE(node.Name)
		}
	}

	crdClient := crd_cs.NewForConfigOrDie(config)
	crds, err := crdClient.ApiextensionsV1beta1().CustomResourceDefinitions().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, crd := range crds.Items {
		oneliners.FILE(crd.Name)
	}
}
