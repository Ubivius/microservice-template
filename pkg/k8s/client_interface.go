package k8s

type KubernetesClient interface {
	GetSecret(name string, key string) (string, error)
	Init() error
}
