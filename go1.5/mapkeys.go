package main

import "fmt"

// Point 作为 map 键：演示 1.5 的复合字面量键类型省略。
type Point struct{ Lat, Lon float64 }

func demoMapKeys() {
	// 1.5 之前的写法：map 键必须重复写完整类型
	before := map[Point]string{
		Point{29.935523, 52.891566}:   "Persepolis",
		Point{-25.352594, 131.034361}: "Uluru",
		Point{37.422455, -122.084306}: "Googleplex",
	}

	// Go 1.5 新语言特性（本版本唯一的语言改动）：
	// 复合字面量的“键”也可以省略类型 —— 此前只有切片/map 的值能省略。
	// 下面的 {29.935523, ...} 不带 Point 前缀，1.4 编译器会报错。
	after := map[Point]string{
		{29.935523, 52.891566}:   "Persepolis",
		{-25.352594, 131.034361}: "Uluru",
		{37.422455, -122.084306}: "Googleplex",
	}

	fmt.Printf("省略键类型写法 == 完整写法? %v\n", before[Point{29.935523, 52.891566}] == after[Point{29.935523, 52.891566}])
	fmt.Printf("查表: (37.422455,-122.084306) -> %s\n", lookupGoogle(after))
}

// 值省略早就有（1.0），键省略是 1.5 补上的对称规则
func lookupGoogle(m map[Point]string) string {
	return m[Point{37.422455, -122.084306}]
}
