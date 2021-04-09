package resources

type MockResources struct{}

func NewMockResources() ResourceManager {
	log.Info("Connecting to mock resources")
	return &MockResources{}
}

func (mp *MockResources) Init() error {
	return nil
}

func (mp *MockResources) GetSecret(namespace string, name string, key string) (string, error) {
	return "", nil
}
