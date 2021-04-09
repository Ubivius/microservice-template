package resources

type ResourcesManager interface {
	GetSecret(namespace string, name string, key string) (string, error)
	Init() error
}
