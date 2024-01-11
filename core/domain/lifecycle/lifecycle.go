package lifecycle

import "gosi/core/domain/common/model"

type LifecycleState struct {
	name string
	model.Entity
}

func NewLifecycleState(id int, name string) LifecycleState {
	state := LifecycleState{
		Entity: model.Entity{Id: id},
		name:   name,
	}
	return state
}

func (ls LifecycleState) GetValue() string {
	return ls.name
}

type Lifecycle struct {
	name        string
	transitions map[LifecycleState][]LifecycleState
	startState  LifecycleState
	model.Entity
}

func (l Lifecycle) GetStartState() LifecycleState {
	return l.startState
}

type LifecycleBuilder struct {
	name        string
	transitions map[LifecycleState][]LifecycleState
	startState  LifecycleState
	id          int
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

// TODO: FIX
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

func (l Lifecycle) GetName() string {
	return l.name
}

type LivecycleManaged struct {
	State     LifecycleState
	Lifecycle Lifecycle
}

func (lm LivecycleManaged) GetLifecycle() Lifecycle {
	return lm.Lifecycle
}

func (lm LivecycleManaged) GetState() LifecycleState {
	return lm.State
}
