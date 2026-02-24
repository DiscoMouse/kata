package auth

import (
	"fmt"

	"github.com/DiscoMouse/kata/internal/db"
	"github.com/DiscoMouse/kata/internal/lifecycle"
)

type Service struct {
	database *db.Database
	state    lifecycle.State
}

func New(d *db.Database) *Service {
	return &Service{
		database: d,
		state:    lifecycle.Stopped,
	}
}

func (a *Service) Start() error {
	// Logic: Auth cannot start if DB isn't running
	if a.database.Status() != lifecycle.Running {
		return fmt.Errorf("auth cannot start: database is not running")
	}

	fmt.Println("Auth: Initializing security layers...")
	a.state = lifecycle.Running
	return nil
}

func (a *Service) Stop() error {
	fmt.Println("Auth: Revoking active sessions...")
	a.state = lifecycle.Stopped
	return nil
}

func (a *Service) Status() lifecycle.State {
	return a.state
}

func (a *Service) Name() string {
	return "AuthService"
}
