package main

import "testing"

func TestTrailingEmptyStatementDoesNotHideTermination(t *testing.T) {
	if got := returnBeforeTrailingEmptyStatement(); got != 7 {
		t.Fatalf("returnBeforeTrailingEmptyStatement() = %d, want 7", got)
	}
}
