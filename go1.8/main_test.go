package main

import "testing"

func TestStructConversionIgnoresTags(t *testing.T) {
	if got := convertIgnoringTags(jsonBar{8}); got.X != 8 {
		t.Fatal(got)
	}
}
