package main

import "testing"

func TestMapKeyElision(t *testing.T) {
	want := map[Point]string{Point{1, 2}: "a"}
	got := map[Point]string{{1, 2}: "a"} // 1.5 键省略写法
	if got[Point{1, 2}] != want[Point{1, 2}] {
		t.Fatal("两种写法结果应一致")
	}
}
