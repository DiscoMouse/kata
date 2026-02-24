package db

import (
	"testing"

	"github.com/DiscoMouse/kata/internal/lifecycle"
)

func TestDatabase_Lifecycle(t *testing.T) {
	// Initialize the component
	cfg := Config{URL: "postgres://localhost:5432/test"}
	d := New(cfg)

	// 1. Verify Initial State
	if d.Status() != lifecycle.Stopped {
		t.Errorf("expected initial state %v, got %v", lifecycle.Stopped, d.Status())
	}

	// 2. Test Start and check error
	if err := d.Start(); err != nil {
		t.Fatalf("failed to start database: %v", err)
	}

	// 3. Immediately defer cleanup so it stops even if the test fails later
	defer func() {
		if err := d.Stop(); err != nil {
			t.Logf("cleanup warning: %v", err)
		}
	}()

	// 4. Verify Running State
	if d.Status() != lifecycle.Running {
		t.Errorf("expected state %v, got %v", lifecycle.Running, d.Status())
	}
}
