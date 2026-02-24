package auth

import (
	"testing"

	"github.com/DiscoMouse/kata/internal/db"
	"github.com/DiscoMouse/kata/internal/lifecycle"
)

func TestAuth_StartRequiresDB(t *testing.T) {
	database := db.New(db.Config{URL: "memskip"})
	authSvc := New(database)

	// Attempt to start Auth while DB is Stopped
	err := authSvc.Start()

	if err == nil {
		t.Error("Expected error when starting Auth without a running DB, got nil")
	}

	// Now start DB and try again
	database.Start()
	err = authSvc.Start()

	if err != nil {
		t.Errorf("Expected Auth to start successfully once DB is running, got: %v", err)
	}

	if authSvc.Status() != lifecycle.Running {
		t.Errorf("Expected Auth state to be Running, got %v", authSvc.Status())
	}
}
