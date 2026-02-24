package manager

import (
	"context"
	"fmt"
	"sync"

	"github.com/DiscoMouse/kata/internal/lifecycle"
)

type Manager struct {
	components []lifecycle.Component
	mu         sync.RWMutex
}

// New creates a manager with a defined order of components.
// Order matters! The first component starts first and stops last.
func New(comps ...lifecycle.Component) *Manager {
	return &Manager{
		components: comps,
	}
}

// StartAll moves every component from Stopped -> Starting -> Running.
// If any component fails, it halts the entire boot sequence.
func (m *Manager) StartAll() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, comp := range m.components {
		fmt.Printf("Manager: Starting %s...\n", comp.Name())
		if err := comp.Start(); err != nil {
			return fmt.Errorf("component %s failed to start: %w", comp.Name(), err)
		}
	}
	return nil
}

// StopAll moves components from Running -> Stopping -> Stopped.
// Best Practice: We iterate in REVERSE order to ensure dependencies
// (like a DB) are the last things to go down.
func (m *Manager) StopAll(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var lastErr error

	// Iterate backwards
	for i := len(m.components) - 1; i >= 0; i-- {
		comp := m.components[i]

		// Check if the context (timeout) has expired before stopping next
		select {
		case <-ctx.Done():
			return fmt.Errorf("shutdown timed out: %w", ctx.Err())
		default:
			fmt.Printf("Manager: Stopping %s...\n", comp.Name())
			if err := comp.Stop(); err != nil {
				fmt.Printf("Error stopping %s: %v\n", comp.Name(), err)
				lastErr = err
			}
		}
	}
	return lastErr
}
