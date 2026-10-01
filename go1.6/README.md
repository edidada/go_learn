# Go 1.6 语言特性

Go 1.6 的官方发行说明明确指出：**该版本没有语言规范变化**。因此本目录不伪造
“Go 1.6 新语法”示例；`LanguageChangeCount` 及其测试将这个结论作为可执行的分支
检查。

Go 1.6 仍支持 Go 1.5 及更早版本的语言特性。Go 1.5 的新增 map 复合字面量键类型
省略示例保留在上游分支已包含的 [`../go1.5`](../go1.5/) 目录中。

```powershell
cd go1.6
go test .
go run .
```

本目录的 `go.work` 仅隔离现代工具链的模块发现，不修改仓库根目录的 Go 1.18+
workspace；后续版本应继续新建同级目录。

参考：[Go 1.6 Release Notes](https://go.dev/doc/go1.6)。
