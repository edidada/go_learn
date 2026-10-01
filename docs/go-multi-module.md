# Go 多 Module 组织方式（类比 Maven 多模块）

## 概念对应

| Maven | Go |
|---|---|
| parent pom | `go.work`（本地联合编译，无"父工程"概念） |
| `<modules>` | `go.work` 中的 `use` 指令 / 各子目录独立 `go.mod` |
| 模块间依赖 | `require` + `replace`（或靠 `go.work` 解析） |
| reactor 统一构建 | `go build ./...`（workspace 模式下） |
| 版本 tag `artifactId-1.0.0` | 仓库 tag 带模块路径前缀 `services/api/v1.0.0` |

核心区别：Go 没有 parent module，每个 module 是独立的模块图根，各自维护自己的依赖。

## 本仓库结构

```text
go_learn/
├── go.work                  # workspace：联合所有 module（对应 Maven reactor）
├── .gitignore
├── go.mod                   # 根 module：go_learn（入口/示例代码）
├── main.go
├── docs/                    # 文档（提交，不 ignore）
├── libs/
│   └── common/
│       ├── go.mod           # module go_learn/libs/common
│       └── common.go        # 被其他 module 引用的公共包
├── services/
│   ├── api/
│   │   ├── go.mod           # module go_learn/services/api
│   │   └── main.go
│   └── worker/
│       ├── go.mod           # module go_learn/services/worker
│       └── main.go
```

## 关键文件

### go.work（仓库根，提交到 git）

```go
go 1.26

use (
	.
	./libs/common
	./services/api
	./services/worker
)
```

存在 `go.work` 时，`go build` / `go vet` / `go run` 在 workspace 模式下运行，跨 module 引用直接解析到本地目录，无需版本号。

### 子 module 的 go.mod（以 services/api 为例）

```go
module go_learn/services/api

go 1.26

require go_learn/libs/common v0.0.0

replace go_learn/libs/common => ../../libs/common
```

- `require` 声明依赖，`replace` 指到本地路径：脱离 workspace 也能编译（CI 单独构建子目录时有用）。
- 正式发布后可去掉 `replace`，改为依赖打好的 tag 版本。

## 常用命令

```bash
# 新建子 module
mkdir -p services/foo && cd services/foo
go mod init go_learn/services/foo

# 把新 module 加入 workspace（仓库根执行）
go work use ./services/foo

# workspace 模式下构建/测试全部 module
go build ./...
go test ./...

# 单独进入某个 module 操作
cd services/api && go mod tidy
```

## 发布与版本

- 每个子 module 独立打 tag，tag 必须带模块路径尾部：`git tag services/api/v1.2.0 && git push --tags`。
- 其他仓库引用：`require go_learn/services/api v1.2.0`（前提是模块路径有可访问的 VCS 路径；纯本地学习仓库用 `replace` 即可）。
- `go.work` / `go.work.sum` 一般提交；纯个人本地试验可加入 `.gitignore`。

## 什么时候用多 module

| 场景 | 建议 |
|---|---|
| 学习 / 小项目 | **单 module + 分包**（`internal/`、`pkg/`），最简单 |
| 各模块需独立版本发布、依赖隔离 | 多 module |
| 模块依赖方向复杂、循环依赖风险 | 多 module 强制边界 |
| 团队按模块独立演进、独立 CI | 多 module |

单 module 时，Maven 的"module"对应 Go 的 package：

```text
go_learn/
├── go.mod
├── cmd/server/
├── internal/service/
├── internal/repo/
└── pkg/util/
```
