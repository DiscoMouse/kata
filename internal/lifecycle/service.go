package lifecycle

type State int

const (
	Stopped State = iota
	Starting
	Running
	Stopping
)

type Component interface {
	Start() error
	Stop() error
	Status() State
	Name() string
}
