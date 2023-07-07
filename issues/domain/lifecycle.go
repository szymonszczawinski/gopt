package domain

const (
	LIFECYCLE_STATE_NEW          = "New"
	LIFECYCLE_STATE_OPEN         = "Open"
	LIFECYCLE_STATE_ANALISYS     = "Analisys"
	LIFECYCLE_STATE_DESIGN       = "Design"
	LIFECYCLE_STATE_DEVELOPMENT  = "Development"
	LIFECYCLE_STATE_FIXED        = "Fixed"
	LIFECYCLE_STATE_CLOSED       = "Closed"
	LIFECYCLE_STATE_VERIFICATION = "Verification"
	LIFECYCLE_STATE_INTEGRATION  = "Integration"
	LIFECYCLE_STATE_RETEST       = "Retest"
	LIFECYCLE_STATE_REJECTED     = "Rejected"
	LIFECYCLE_STATE_DRAFT        = "Draft"
)

type LifecycleState struct {
	Entity
	name string
}

func NewLifecycleState(id int, name string) LifecycleState {
	state := LifecycleState{
		Entity: Entity{id: id},
		name:   name,
	}
	return state
}
func (self LifecycleState) GetValue() string {
	return self.name
}

type Lifecycle struct {
	Entity
	name        string
	startState  LifecycleState
	transitions map[LifecycleState][]LifecycleState
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
		Entity:      Entity{id: lcb.id},
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
	lifecycle Lifecycle
	state     LifecycleState
}

func (self LivecycleManaged) GetLifecycle() Lifecycle {
	return self.lifecycle
}

func (self LivecycleManaged) GetState() LifecycleState {
	return self.state
}
