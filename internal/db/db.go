package db

import (
	"fmt"
	"sync"

	"github.com/DiscoMouse/kata/internal/lifecycle"
)

type Config struct {
	URL string
}

type Database struct {
	config Config
	state  lifecycle.State
	mu     sync.RWMutex
}

func New(cfg Config) *Database {
	return &Database{
		config: cfg,
		state:  lifecycle.Stopped,
	}
}

func (d *Database) Start() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Printf("DB: Connecting to %s...\n", d.config.URL)
	// Simulate connection logic
	d.state = lifecycle.Running
	return nil
}

func (d *Database) Stop() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Println("DB: Closing connections...")
	d.state = lifecycle.Stopped
	return nil
}

func (d *Database) Status() lifecycle.State {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.state
}

func (d *Database) Name() string {
	return "PostgresDB"
}
