package k8s

import (
	"os"

	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/rest"
)

type Kubernetes struct {
	client     *kubernetes.Clientset
}

func NewKubernetes() KubernetesClient {
	mp := &Kubernetes{}
	err := mp.Init()
	// If Init fails, kill the program
	if err != nil {
		log.Error(err, "Kubernetes setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *Kubernetes) Init() error {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	mp.client = client  
	return nil
}

func (mp *Kubernetes) GetSecret(name string, key string) (string,error) {
	res, err := mp.client.Core().Secrets("default").Get(name)
	if err!= nil{
		return "", err
	}
	return res.Data[key], nil
}
