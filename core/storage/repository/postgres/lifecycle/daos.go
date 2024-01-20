package lifecycle

type LifecycleStateRow struct {
	Name string
	Id   int
}

type LifecycleRow struct {
	Name         string
	Id           int
	StartStateId int
}

type StateTransition struct {
	LifecycleId int
	FromStateId int
	ToStateId   int
}
