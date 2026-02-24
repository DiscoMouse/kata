package manager

import (
	"errors"
	"sync"
	"testing"

	"github.com/DiscoMouse/kata/internal/lifecycle"
)

// MockComponent allows us to simulate a "moving part"
type MockComponent struct {
	name       string
	state      lifecycle.State
	failStart  bool
	startCount int
	mu         sync.Mutex
}

func (m *MockComponent) Start() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.startCount++
	if m.failStart {
		return errors.New("failed to start")
	}
	m.state = lifecycle.Running
	return nil
}

func (m *MockComponent) Stop() error {
	m.state = lifecycle.Stopped
	return nil
}

func (m *MockComponent) Name() string            { return m.name }
func (m *MockComponent) Status() lifecycle.State { return m.state }
func TestManager_StartAll_Success(t *testing.T) {
	// Setup
	compA := &MockComponent{name: "DB"}
	compB := &MockComponent{name: "Auth"}

	mgr := New(compA, compB)

	// Execute
	err := mgr.StartAll()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if compA.state != lifecycle.Running || compB.state != lifecycle.Running {
		t.Errorf("Components should be Running. A: %v, B: %v", compA.state, compB.state)
	}
}
func TestManager_StartAll_HaltsOnError(t *testing.T) {
	// Setup: DB will fail, so Auth should never start
	db := &MockComponent{name: "DB", failStart: true}
	auth := &MockComponent{name: "Auth"}

	mgr := New(db, auth)

	// Execute
	err := mgr.StartAll()

	// Assert
	if err == nil {
		t.Error("Expected an error from Manager, but got nil")
	}

	if auth.startCount > 0 {
		t.Error("Auth service should not have attempted to start because DB failed")
	}
}
