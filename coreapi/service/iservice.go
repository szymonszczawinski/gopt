package service

const NEW_FUNCTION string = "New"

type ComponentType string

const (
	ComponentTypeMessenger          = "Messenger"
	ComponentTypeTypeStorageService = "StorageService"
	ComponentTypeHttpServerService  = "HttpServerService"
	ComponentTypeHttpClientService  = "HttpClientService"
	ComponentTypeBunDatabase        = "BunDatabase"
	ComponentTypeSqlDatabase        = "SqlDatabase"
	ComponentTypeProjectRepository  = "ProjectRepository"
	ComponentTypeProjectService     = "ProjectService"
	ComponentTypeAuthRepository     = "AuthRepository"
	ComponentTypeAuthService        = "AuthService"
)

type IServiceManager interface {
	MustGetComponent(componentType ComponentType) IComponent
	StartComponent(componentType ComponentType, component IComponent) error
	RegisterComponent(componentType ComponentType, component IComponent)
}

type IComponent interface {
	StartComponent()
}
