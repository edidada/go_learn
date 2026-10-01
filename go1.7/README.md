# Go 1.7 语言特性

Go 1.7 对“终止语句”（terminating statement）的判定作了一项小的规范澄清：判定
一个语句列表是否以终止语句结束时，应看 **最后一条非空语句**。这与此前 gc 和
gccgo 编译器的实际行为一致，不影响已有正确程序。

`termination.go` 用一个含尾随空语句的函数演示该规则：两个分支都 `return` 的 `if`
是最后的非空语句，因此函数已知会返回；尾随的第二个 `;` 不能让它变成“缺少
return”的函数。该双分号是刻意保留的语法演示，不能用 gofmt 改写该文件。

```powershell
cd go1.7
go test .
go run .
```

该目录从 `feature/go1.6` 演进而来；更早版本示例仍位于同级目录。局部 `go.work`
避免未来合并根 Go 1.18+ workspace 时出现配置冲突。

参考：[Go 1.7 Release Notes](https://go.dev/doc/go1.7)。
