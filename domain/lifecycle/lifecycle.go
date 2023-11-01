package lifecycle

import "gosi/domain/common/model"

const ()

type LifecycleState struct {
	model.Entity
	name string
}

func NewLifecycleState(id int, name string) LifecycleState {
	state := LifecycleState{
		Entity: model.Entity{Id: id},
		name:   name,
	}
	return state
}
func (self LifecycleState) GetValue() string {
	return self.name
}

type Lifecycle struct {
	model.Entity
	name        string
	startState  LifecycleState
	transitions map[LifecycleState][]LifecycleState
}

func (self Lifecycle) GetStartState() LifecycleState {
	return self.startState
}

type LifecycleBuilder struct {
	id          int
	name        string
	startState  LifecycleState
	transitions map[LifecycleState][]LifecycleState
}

func NewLifeCycleBuilder(id int, name string, start LifecycleState) *LifecycleBuilder {
	builder := LifecycleBuilder{
		id:          id,
		name:        name,
		startState:  start,
		transitions: map[LifecycleState][]LifecycleState{},
	}
	return &builder
}

func (lcb *LifecycleBuilder) AddTransition(from LifecycleState, to LifecycleState) *LifecycleBuilder {
	transitions, exists := lcb.transitions[from]
	if !exists {
		transitions := make([]LifecycleState, 0)
		lcb.transitions[from] = transitions
	}
	transitions = append(transitions, to)
	return lcb
}

func (lcb LifecycleBuilder) Build() Lifecycle {
	lifecycle := Lifecycle{
		Entity:      model.Entity{Id: lcb.id},
		name:        lcb.name,
		startState:  lcb.startState,
		transitions: lcb.transitions,
	}
	return lifecycle
}

func (self Lifecycle) GetName() string {
	return self.name
}

type LivecycleManaged struct {
	Lifecycle Lifecycle
	State     LifecycleState
}

func (self LivecycleManaged) GetLifecycle() Lifecycle {
	return self.Lifecycle
}

func (self LivecycleManaged) GetState() LifecycleState {
	return self.State
}
