package application

import (
	"testing"
)

func TestApplication(t *testing.T) {
	app := New()
	result := app.Log("works")
	if result != "Application works" {
		t.Error("Expected Application to append 'works'")
	}
}
