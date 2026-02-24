package db

import (
	"testing"

	"github.com/DiscoMouse/kata/internal/lifecycle"
)

func TestDatabase_Lifecycle(t *testing.T) {
	cfg := Config{URL: "postgres://localhost:5432/test"}
	d := New(cfg)

	// 1. Check initial state
	if d.Status() != lifecycle.Stopped {
		t.Errorf("expected initial state Stopped, got %v", d.Status())
	}

	// 2. Test Start
	err := d.Start()
	if err != nil {
		t.Fatalf("failed to start database: %v", err)
	}

	if d.Status() != lifecycle.Running {
		t.Errorf("expected state Running after Start, got %v", d.Status())
	}

	// 3. Test Stop
	err = d.Stop()
	if err != nil {
		t.Fatalf("failed to stop database: %v", err)
	}

	if d.Status() != lifecycle.Stopped {
		t.Errorf("expected state Stopped after Stop, got %v", d.Status())
	}
}

func TestDatabase_Concurrency(t *testing.T) {
	// Because our Manager might check status while the DB is starting/stopping,
	// we should ensure our Status() call is thread-safe and doesn't race.
	d := New(Config{URL: "mem"})

	go func() {
		for i := 0; i < 100; i++ {
			_ = d.Start()
			_ = d.Stop()
		}
	}()

	for i := 0; i < 100; i++ {
		_ = d.Status()
	}
}
