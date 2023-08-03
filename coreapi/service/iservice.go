package service

const NEW_FUNCTION string = "New"

type ComponentType string

const (
	ComponentTypeMessenger          = "Messenger"
	ComponentTypeTypeStorageService = "StorageService"
	ComponentTypeHttpServerService  = "HttpServerService"
	ComponentTypeHttpClientService  = "HttpClientService"
	ComponentTypeBunDatabase        = "BunDatabase"
	ComponentTypeMemoryRepository   = "MemoryRepository"
	ComponentTypeSqlite3Repository  = "Sqlite3Repository"
	ComponentTypeIssueRepository    = "IssueRepository"
	ComponentTypeIssueService       = "IssueService"
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
