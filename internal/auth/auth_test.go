package auth

import (
	"testing"

	"github.com/DiscoMouse/kata/internal/db"
	"github.com/DiscoMouse/kata/internal/lifecycle"
)

func TestAuth_StartRequiresDB(t *testing.T) {
	// Setup components
	database := db.New(db.Config{URL: "mem"})
	authSvc := New(database)

	// 1. Test: Auth should fail to start if DB is stopped
	if err := authSvc.Start(); err == nil {
		t.Error("expected error when starting Auth without a running DB, got nil")
	}

	// 2. Test: Start DB correctly
	if err := database.Start(); err != nil {
		t.Fatalf("setup failed: could not start database: %v", err)
	}

	// Defer DB cleanup
	defer func() {
		if err := database.Stop(); err != nil {
			t.Logf("db cleanup warning: %v", err)
		}
	}()

	// 3. Test: Auth should now start successfully
	if err := authSvc.Start(); err != nil {
		t.Errorf("expected Auth to start, got error: %v", err)
	}

	// Defer Auth cleanup
	defer func() {
		if err := authSvc.Stop(); err != nil {
			t.Logf("auth cleanup warning: %v", err)
		}
	}()

	// 4. Verify final state using the lifecycle import
	if authSvc.Status() != lifecycle.Running {
		t.Errorf("expected auth state %v, got %v", lifecycle.Running, authSvc.Status())
	}
}
