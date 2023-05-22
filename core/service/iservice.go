package service

const NEW_FUNCTION string = "New"

type ServiceType string
type IServiceManager interface {
	GetService(serviceType ServiceType) (any, error)
	StartService(serviceType ServiceType, service IService) error
	RegisterService(serviceType ServiceType, service IService)
}

type IService interface {
	StartService()
}
