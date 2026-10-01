package main

import "testing"

func TestGo16HasNoLanguageChanges(t *testing.T) {
	if got := LanguageChangeCount(); got != 0 {
		t.Fatalf("Go 1.6 language change count = %d, want 0", got)
	}
}
