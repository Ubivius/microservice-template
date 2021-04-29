package resources

type ResourceManager interface {
	GetSecret(namespace string, name string, key string) (string, error)
	Init() error
}
