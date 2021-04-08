package k8s

import (
	"k8s.io/client-go/kubernetes"
)

type MockKubernetes struct {
	client     *kubernetes.Clientset
}

func NewMockKubernetes() KubernetesClient {
	log.Info("Connecting to mock kubernetes")
	return &MockKubernetes{}
}

func (mp *MockKubernetes) Init() error {
	return nil
}

func (mp *MockKubernetes) GetSecret(name string, key string) (string, error) {
	return "", nil
}
