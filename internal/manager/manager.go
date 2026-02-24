package manager

import "github.com/DiscoMouse/kata/internal/lifecycle"

type Manager struct {
	components []lifecycle.Component
}

func (m *Manager) StartAll() error {
	for _, comp := range m.components {
		if err := comp.Start(); err != nil {
			return err // Or handle partial failure
		}
	}
	return nil
}
