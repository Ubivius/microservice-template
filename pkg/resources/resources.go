package resources

import (
	"context"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Resources struct {
	k8sClient *kubernetes.Clientset
}

func NewResources() ResourceManager {
	mp := &Resources{}
	err := mp.Init()
	// If Init fails, kill the program
	if err != nil {
		log.Error(err, "Kubernetes setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *Resources) Init() error {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	mp.k8sClient = clientset
	return nil
}

func (mp *Resources) GetSecret(namespace string, name string, key string) (string,error) {
	res, err := mp.k8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err!= nil {
		return "", err
	}
	return string(res.Data[key]), nil
}
