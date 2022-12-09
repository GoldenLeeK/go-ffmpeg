package driver

type Driverable interface {
	GetName() string
	Create(configtion map[string]string) Driverable
}
