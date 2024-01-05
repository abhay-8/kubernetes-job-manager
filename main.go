package main

import (
	"flag"
	"job-manager/jobs"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientset = &kubernetes.Clientset{}

func init() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", " ", "path to kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}

	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	jobs.CreateJob()
}
