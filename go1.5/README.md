# Go 1.5 及此前版本的语言特性

这个目录是 Go 1.5 及此前版本语言特性的学习入口。版本演进说明见
[`../docs/go1.5-and-earlier-language-features.md`](../docs/go1.5-and-earlier-language-features.md)，
覆盖 Go 1.0 至 Go 1.5 的语言新增与规则修订；它不罗列标准库、运行时和 `go` 命令的变化。

本目录的可执行代码聚焦 Go 1.5 相对于 Go 1.4 的新增语法。Go 1.5 的官方发布说明在
“Changes to the language”中只列出一项：**map
复合字面量的键可以省略元素类型**。`mapkeys.go` 和 `mapkeys_test.go` 因而完整
覆盖了该版本新增的语言变化：既给出 Go 1.4 的完整写法，也给出 Go 1.5 新增的
省略写法，并断言二者构造出的 map 相同。

例如：

```go
// Go 1.4 及以前。
map[Point]string{Point{1, 2}: "A"}

// Go 1.5 起合法。
map[Point]string{{1, 2}: "A"}
```

## 验证

在本目录执行：

```powershell
go test .
go run .
```

`go.mod` 的 `go 1.5` 语言版本会让现代工具链按 Go 1.5 的语言边界检查源码；
示例本身不使用模块、泛型、`any` 或其他 Go 1.5 之后的语法。历史 Go 1.5 工具链
尚未支持 modules，因此若直接用该工具链，请将源码置于 GOPATH 中再运行测试。

## 与后续分支的边界

本分支只新增 `go1.5/`。其中的 `go.work` 仅服务于现代工具链：它使在该目录运行
命令时优先使用本模块，不依赖或修改仓库根目录将在 Go 1.18 分支引入的 workspace。
后续 Go 1.6 示例应新建同级的 `go1.6/`，不要改动本目录或根 workspace，以保持
合并低冲突。

参考：[Go 1.5 Release Notes](https://go.dev/doc/go1.5)。
