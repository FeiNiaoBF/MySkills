# 规则出处与边界

本文件记录 [SKILL.md](../SKILL.md) 中规则的来源和失效边界。规则有争议、涉及版本差异或准备修改规则时再读；日常改代码不需要它，正文已经写明可直接执行的判断。

核对上游链接的日期：2026-11-13。

## 规则如何取舍

- **通用的 Go 惯例**（格式化、错误串、接口大小、goroutine 归属、测试）直接写成规则，因为它们在任意 Go 模块中都成立，且由 Go 项目自身推荐或强制。
- **Google Go Style** 提供判断顺序（清晰、简单、简洁、可维护、一致）以及部分命名和文档决策。它的 Decisions 页只对 Google 代码库具有规范效力；本 Skill 只在某条 Google 规则对任意项目也站得住时才写入。
- **有意省略**：Google 内部的评审政策、内部工具、仓库布局规则，以及「不设固定行宽」这类 Google 特有决定，都不作为通用要求。本 Skill 也不要求单独产出风格报告。
- **有取舍的规则**明确写出条件，而不是一律强制：`%w` 对 `%v`、`Mutex` 对 `RWMutex`、nil 切片对空切片、struct 选项对函数式选项，以及必须先测量才能做的性能改动。

## 核心来源

| 来源 | 用途与限制 |
| --- | --- |
| [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) | Go Wiki 的实践指导，覆盖格式化、命名、接口、错误、context 和 goroutine 生命周期。它补充 Effective Go，但要放回原文语境使用。 |
| [Google Go Style 总览](https://google.github.io/styleguide/go/index)与 [Guide](https://google.github.io/styleguide/go/guide) | Guide 对 **Google Go 代码**是权威且规范的；它规定清晰、简单、简洁、可维护、一致的优先顺序。它不设固定行宽是有意的风格决定。 |
| [Google Go Style Decisions](https://google.github.io/styleguide/go/decisions) | 对 Google 评审者是规范，但不是普遍适用的 Go 要求；应查具体条目，而不是把每条 Google 规则都当成通用规则。 |
| [Google Go Style Best Practices](https://google.github.io/styleguide/go/best-practices) | 既非规范也非权威，提供依赖上下文的示例和取舍。 |
| [Go Test Comments](https://go.dev/wiki/TestComments) | 测试失败诊断与结构；共享检查逻辑适合表驱动，断言逻辑不同则适合拆开。 |
| [Go Doc Comments](https://go.dev/doc/comment) | 包、导出声明和示例的文档注释格式。 |
| [Effective Go](https://go.dev/doc/effective_go) | 基础惯例；其前言说明文本没有跟上 modules、泛型等新特性。 |

## 版本门槛

以 `go.mod` 的 `go` 指令为准，而不是本机安装的工具链。自 Go 1.23 起，`go vet` 的 `stdversion` 分析器会报告超出模块有效版本的符号；循环变量作用域等语言语义和 timer 行为等运行时行为也由同一个指令控制。

| 特性 | 最低版本 | 对规则的影响 |
| --- | --- | --- |
| `any` 别名 | 1.18 | 只在真正动态的边界使用 `any`。 |
| `errors.Join`、多个 `%w` | 1.20 | 把关闭错误和操作错误合并返回。 |
| `slices`、`maps`、`min`、`max` | 1.21 | 替代手写循环、排序和比较。 |
| 每次迭代独立的循环变量 | 1.22 | 1.22+ 不需要 `v := v`，`go vet` 会提示多余复制。 |
| `for i := range n` | 1.22 | 需要索引时使用；不需要索引时仍优先 `range` 容器。 |
| timer 可被回收、timer channel 无缓冲 | 1.23 | 1.23+ 长循环里用 `time.After` 安全；更低版本复用 `time.NewTimer`。 |
| `maps.Keys`/`maps.Values` 返回迭代器、`slices.Sorted` | 1.23 | 不要假设 1.21 的返回切片形式。 |
| `sync.WaitGroup.Go` | 1.25 | 替代 `Add` + `go` + `defer Done`；旧模块不要依赖。 |
| `atomic.Int64` 等类型化原子值 | 1.19 | 优先于旧的函数式原子操作。 |
| `os.ReadFile`、`os.WriteFile`、`io.ReadAll` | 1.16 | 优先于 `ioutil` 等价物。 |
| `t.Setenv` | 1.17 | 测试中的环境变量隔离。 |
| `t.TempDir` | 1.15 | 临时目录自动清理。 |
| `t.Cleanup` | 1.14 | 与子测试配合的清理顺序。 |

参考：[Go 1.22 发布说明](https://go.dev/doc/go1.22)、[Go 1.23 发布说明](https://go.dev/doc/go1.23)（timer 与迭代器变更）、[循环变量试验](https://go.dev/wiki/LoopvarExperiment)、[`sync.WaitGroup`](https://pkg.go.dev/sync#WaitGroup.Go)、[`slices`](https://pkg.go.dev/slices)、[`maps`](https://pkg.go.dev/maps)。

## 错误行为

- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)：用 `%w` 包装会把底层错误通过 `errors.Is` 和 `errors.As` 暴露给调用方，可能成为 API 契约的一部分。包装是设计决定，不是条件反射。
- [Go 规范：错误](https://go.dev/ref/spec#Errors)：`error` 是普通接口，把 typed nil 装进去后结果非 nil。
- [Go 规范：map 类型](https://go.dev/ref/spec#Map_types)：遍历顺序未指定；nil map 可读但不可写。
- [Go 规范：append](https://go.dev/ref/spec#Appending_and_copying_slices)：append 可能复用目标底层数组。

## 工具与格式化

- [gofmt](https://pkg.go.dev/cmd/gofmt) 是标准格式化工具。[goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) 还会调整 import；项目已经依赖它时沿用，不要只为风格新增依赖。
- [go vet](https://pkg.go.dev/cmd/vet) 能发现这些规则依赖的复合字面量、锁复制、循环变量和 stdversion 问题。
- 竞态检测（`go test -race`）是并发结论的证据；只靠阅读评审不算验证。
