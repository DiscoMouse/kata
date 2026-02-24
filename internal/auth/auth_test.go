package auth

import (
	"testing"

	"github.com/DiscoMouse/kata/internal/db"
	"github.com/DiscoMouse/kata/internal/lifecycle"
)

func TestAuth_StartRequiresDB(t *testing.T) {
	// Setup
	database := db.New(db.Config{URL: "memskip"})
	authSvc := New(database)

	// 1. Attempt to start Auth while DB is Stopped
	// This should fail according to our logic in auth.go
	err := authSvc.Start()
	if err == nil {
		t.Error("Expected error when starting Auth without a running DB, got nil")
	}

	// 2. Start DB and check the error
	if err := database.Start(); err != nil {
		t.Fatalf("Setup failed: could not start database for test: %v", err)
	}

	// 3. Try to start Auth again
	err = authSvc.Start()
	if err != nil {
		t.Errorf("Expected Auth to start successfully once DB is running, got: %v", err)
	}

	// 4. Final state verification
	if authSvc.Status() != lifecycle.Running {
		t.Errorf("Expected Auth state to be Running, got %v", authSvc.Status())
	}
}
