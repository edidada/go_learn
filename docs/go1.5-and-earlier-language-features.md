# Go 1.5 及以前的语言特性

本文按发行版本记录 **Go 1.0 到 Go 1.5 的语言增量与语言规则修订**。它不罗列
标准库、运行时和 `go` 命令的变化；例如 Go 1.5 的并发 GC 和 vendoring 都不属于
语言特性。Go 1.0 是稳定 Go 1 系列的基线，因而这里的“新增”是相对于 Go 1.0 前的
r60，而不是试图重写整份 Go 语言规范。

| 版本 | 语言新增 | 规则修订 / 不兼容点 |
| --- | --- | --- |
| Go 1.0 | `append([]byte, string...)`、指针元素复合字面量类型省略、`rune`、`error`、`delete`、结构体/数组相等比较 | 禁止关闭只接收 channel；map 遍历无序；明确多重赋值求值顺序；禁止函数与 map（除 `nil` 外）相等比较等 |
| Go 1.1 | 方法值 | 常量零作整数除数为编译错误；Unicode 字面量不允许 surrogate half |
| Go 1.2 | 三索引切片 | 空地址求值必须安全地 panic 或返回安全值 |
| Go 1.3 | 无 | 无语言变更（仅内存模型澄清） |
| Go 1.4 | 无变量 `for range` | 不再允许对 `**T` 隐式解引用后调用 `T` 的方法 |
| Go 1.5 | map 复合字面量键类型省略 | 无其他语言变更 |

## Go 1.0

### 语言新增

- 可把 string 逐字节追加到 `[]byte`：`b = append(b, "text"...)`。
- 数组、切片、map 的元素类型为指针时，复合字面量可省略 `&T` 中的 `T`：
  `[]*Date{{"Feb", 14}}`。
- 初始化期间启动的 goroutine 可以在整个程序初始化完成前运行；`init` 中可等待它。
- `rune` 成为 `int32` 的别名，字符字面量的默认类型为 `rune`。
- 预声明接口类型 `error`（`interface { Error() string }`）成为语言的一部分。
- 预声明函数 `delete(m, key)` 替代旧的 `m[key] = value, false` 删除语法。
- 当全部字段或元素可比较时，结构体和数组支持 `==`、`!=`，也可作为 map 键。
- 包可复制其他包定义但含未导出字段的结构体值；仍不能直接访问该字段。

### 规则修订

- 不能对只接收 channel（`<-chan T`）调用 `close`。
- `range` 遍历 map 的顺序明确为不可预测，不能依赖顺序。
- 多重赋值先按从左至右的顺序计算左操作数，再从左至右赋值。例如：

  ```go
  s := []int{1, 2, 3}
  i := 0
  i, s[i] = 1, 2 // s[0] 为 2，而不是 s[1]
  ```

- 命名返回值被同名局部变量遮蔽时，裸 `return` 被编译器拒绝。
- 函数值与 map 只能和 `nil` 比较；切片仍不能比较。结构体/数组只支持相等比较，
  不支持 `<`、`<=`、`>`、`>=`。

## Go 1.1

- **方法值**：`f := w.Write` 绑定接收者 `w`，之后可像普通函数一样调用 `f(p)`；
  它不同于方法表达式 `(*Writer).Write`，后者仍要求显式传入接收者。
- `x / 0` 中的常量零整数除数改为编译错误，而不是运行时 panic。
- string 与 rune 字面量不允许使用 Unicode surrogate half（`U+D800` 至 `U+DFFF`）。

## Go 1.2

- **三索引切片**可同时约束长度和容量：`part := array[2:4:7]`。结果的
  `len(part)` 是 `2`，`cap(part)` 是 `5`；第三个索引可防止 append 覆盖该边界之后
  的底层数组元素。
- 任何显式或隐式需要计算 nil 地址的表达式，必须 panic 或返回正确、安全的非 nil
  结果，不能错误访问低地址内存。

## Go 1.3

没有语言变化。发行说明只澄清了内存模型中带缓冲 channel 的 send/receive
同步关系；它不是新的语法或类型系统功能。

## Go 1.4

- 无需索引和值时，可写 `for range values { ... }`；此前必须写
  `for _ = range values { ... }`。
- 规范只允许为方法调用自动插入一次解引用。因此 `var p **T; p.M()` 不再被接受，
  应显式写作 `(*p).M()`。

## Go 1.5

map 复合字面量的键与值一样，若可由 map 类型推断，可省略元素类型：

```go
type Point struct{ X, Y int }

// Go 1.4 及以前。
oldStyle := map[Point]string{Point{1, 2}: "A"}

// Go 1.5 起。
go15Style := map[Point]string{{1, 2}: "A"}
```

本分支的 [`go1.5`](../go1.5/) 模块用可运行测试覆盖这项 Go 1.5 新语法。

## 官方资料

- [Go 1.0 Release Notes](https://go.dev/doc/go1)
- [Go 1.1 Release Notes](https://go.dev/doc/go1.1)
- [Go 1.2 Release Notes](https://go.dev/doc/go1.2)
- [Go 1.3 Release Notes](https://go.dev/doc/go1.3)
- [Go 1.4 Release Notes](https://go.dev/doc/go1.4)
- [Go 1.5 Release Notes](https://go.dev/doc/go1.5)
