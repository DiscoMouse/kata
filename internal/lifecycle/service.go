package lifecycle

// State represents the current lifecycle phase of a component.
type State int

const (
	// Stopped means the component is initialized but not active.
	Stopped State = iota
	// Starting means the component is in the middle of its boot sequence.
	Starting
	// Running means the component is fully operational.
	Running
	// Stopping means the component is currently shutting down.
	Stopping
)

// Component defines the interface for all "moving parts" in the system.
// Any struct that implements these four methods can be managed by your Manager.
type Component interface {
	Start() error
	Stop() error
	Status() State
	Name() string
}
