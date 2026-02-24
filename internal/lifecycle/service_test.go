package lifecycle

import "testing"

func TestState_Values(t *testing.T) {
	// Verify that iota is working and states are unique
	states := map[State]string{
		Stopped:  "Stopped",
		Starting: "Starting",
		Running:  "Running",
		Stopping: "Stopping",
	}

	if len(states) != 4 {
		t.Errorf("Expected 4 unique states, got %d", len(states))
	}

	// Verify specific order/values
	if Stopped != 0 {
		t.Errorf("Expected Stopped to be 0, got %d", Stopped)
	}
	if Running != 2 {
		t.Errorf("Expected Running to be 2, got %d", Running)
	}
}

// TestComponentInterface ensures that we haven't broken the contract.
// This is a "static" test that will fail to compile if the interface changes.
func TestComponentInterface(t *testing.T) {
	var _ Component = (*mockComponent)(nil)
}

type mockComponent struct{}

func (m *mockComponent) Start() error  { return nil }
func (m *mockComponent) Stop() error   { return nil }
func (m *mockComponent) Status() State { return Stopped }
func (m *mockComponent) Name() string  { return "mock" }
