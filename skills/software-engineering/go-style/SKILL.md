---
name: go-style
description: 编写、修改、测试或审查 Go 代码时使用：按内置的具体规则直接优化命名、错误处理、接口与依赖、切片与 map、并发和测试，说明依据与例外；明确请求时可只评估不修改。
license: MIT
metadata:
  version: 0.2.0
---

# Go Style

在获得授权的 Go 改动中直接按本文件的规则改写代码，不需要再查外部风格指南。每条规则给出「现状 → 改成 → 依据」，并注明不适用条件。先保证行为、兼容性和项目约定，再谈风格。

## 适用范围

- 用于新增或修改 `.go` 文件、测试、公开 API 和 Go 示例。
- 用户要求「评估/审查 Go 风格」但未要求修改时，只列出具体位置、改法和依据，不动代码。
- 仅阅读或解释 Go 代码不构成修改授权。不把范围内的改动扩大成全仓库改名、全仓库格式化或无关重构。

## 先确认约束

1. 读 `go.mod` 的 `go` 指令和 build 约束，确定模块语言版本；规则里的 `[1.21+]`、`[1.22+]` 指这个版本，不是本机工具链版本。低于门槛时用等价写法，不引入新符号。
2. 读目标文件附近的代码、项目 formatter/lint 配置和公开 API 调用方。改导出标识符、签名或错误链前先查引用；不要只凭搜索文本判断。
3. 冲突时按此优先级取舍：可观察行为与兼容性 > 项目既有约定和工具链 > 本文件规则 > 个人偏好。若某条改写会改变 JSON 形状、错误链契约、公开签名或并发时序，先确认；无法确认就保留原行为并说明。
4. 生成代码只改生成器和模板，不手改产物。

## 命名

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| `user_id`、`User_id`、`userID` 混用 | 未导出 `userID`，导出 `UserID` | Go 用 MixedCaps，不用下划线 |
| `Url`、`HttpClient`、`Id` | `URL`、`HTTPClient`、`ID` | 缩写是一个词：全大写或全小写，不拆成 `Url` |
| 包名 `utils`、`common`、`helpers`、`models` | 按职责命名：`retry`、`billing`、`auth`、`user` | 包名出现在每个调用点，泛名不提供信息 |
| `chubby.ChubbyFile`、`http.HTTPServer` | `chubby.File`、`http.Server` | 调用点已经带包名，标识符不要重复（stutter） |
| 接收者写 `this`、`self` | 类型首字母（`c`、`s`），同一类型所有方法保持一致 | Go 没有对象指针概念；一致才可读 |
| `GetUser()` 只读字段 | `User()`；昂贵操作才用 `FetchUser()`、`LoadUser()` | `Get` 前缀在 Go 中冗余；动词应表达代价 |
| 大作用域里叫 `x`、`d` | `timeout`、`deadline`、`retryBudget` | 名字长度随作用域增大；短名只适合小作用域 |
| 任意值叫 `c` | 索引用 `i`、`j`，通道用 `ch`，错误用 `err`，上下文用 `ctx` | 约定俗成的单字母含义，降低阅读成本 |
| 布尔变量叫 `flag`、`status` | `ok`、`found`、`valid`、`enabled` | 名字直接表达真假含义，调用处可读 |
| 导出声明注释写 `// returns the user` | `// User returns ...` | 文档注释以声明名开头，godoc 才会正确归属 |
| 测试名 `Test1`、`TestParseCase` | `TestParse`、`TestParse_EmptyInput`，子测试名描述场景 | 失败输出能定位行为和输入 |

不适用：生成的 protobuf/数据库字段名、既有公开 API 改名会破坏兼容、外部规范强制的名称；存在对应的 `SetUser` 时保留 `GetUser` 也符合惯例。只有获得授权并迁移全部调用方时才改名。

## 错误处理

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| `if err != nil { return err }` 裸返回 | `return fmt.Errorf("load config: %w", err)` | 包装说明「哪一步失败」，`%w` 保留 `errors.Is/As` |
| 只因怕暴露内部实现就把所有错误改成 `%v` | 调用方需要判断错误类型时用 `%w`；确实不应暴露底层错误时才用 `%v` | `%w` 把底层错误纳入 API；是否暴露要按契约决定 |
| `errors.New(fmt.Sprintf(...))` | `fmt.Errorf(...)`，需要包裹时用 `%w` | 前者无法保留错误链，还多一次拼接 |
| `err.Error() == "not found"`、`err.(*MyError)` | `errors.Is(err, ErrNotFound)`、`errors.As(err, &target)` | 包装后字符串和动态类型都不再可靠 |
| `fmt.Errorf("Failed to load config: %v", err)` | `fmt.Errorf("load config: %w", err)` | 错误串小写、无结尾标点、动词开头；不要只写 failed |
| 每一层都 `log.Printf` 后再返回 | 在最外层记录一次，或返回给调用方处理 | 重复日志掩盖真正的边界，调用方也无法处理 |
| `_ = f.Close()` 静默丢弃写错误 | 写路径：`if err := f.Close(); err != nil { return err }`；读路径确需忽略时注明原因 | 写关闭失败意味着数据可能没有落盘 |
| `defer f.Close()` 后不检查错误 | 具名返回 + `defer func() { err = errors.Join(err, f.Close()) }()`，或显式关闭 | `defer` 保证执行，不保证错误被处理 |
| 循环里 `defer resp.Body.Close()` | 把循环体抽成函数，或每轮显式关闭 | `defer` 在函数返回时才执行，循环会堆积连接 |
| 忽略 `rows.Err()`、`scanner.Err()` | 迭代结束后检查 | 遍历过程中的错误只在这里暴露 |
| `panic("file missing")` 处理普通失败 | 返回 `error` | panic 只用于程序员错误或不可能状态 |
| `if err == nil { ... } else { ... }` 深层嵌套 | `if err != nil { return ... }` 提前返回 | 减少嵌套，主路径保持在左侧 |

不适用：底层库不记录日志；公共错误文本已被外部依赖匹配时保留原文并加测试。`recover` 只用于 goroutine 或请求边界，且要转成 error 或重新 panic。`errors.Join` 需要 `[1.20+]`。

## 接口与依赖

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| 生产者包预先定义 `UserStore` 接口 | 生产者按需返回具体类型；消费方定义自己用到的最小接口 | 接口属于使用方，有真实需求才抽象 |
| 为 mock 给每个类型建接口 | 只在外部边界（数据库、HTTP、时钟）定义 1–2 个方法的接口；简单类型直接测 | 接口爆炸增加跳转和实现维护成本 |
| 5 个方法的接口只用到 1 个 | 拆成 `Reader`/`Writer` 式窄接口，参数接受窄接口 | 接口越大抽象越弱，窄接口更易替换 |
| 函数返回接口类型 | 返回具体类型（如 `*Client`），除非抽象本身是契约 | 调用方需要额外方法或字段时不被接口卡住 |
| 全局 `var db *sql.DB`、`init()` 里建连接 | 构造函数显式注入，在 `main` 或装配层创建 | 隐式依赖难测、难并行，生命周期不清 |
| `init()` 做业务初始化 | 显式 `New`/`Setup`；只有驱动、flag 等注册场景用 `init` | `init` 顺序隐式，且无法返回错误 |
| 用 `any` 传结构化数据 | 定义 struct；只有真正任意的 JSON 边界才用 `any` | 字段不可发现，错误推迟到运行时 |
| 单次使用的类型也加泛型 | 先写具体类型；≥2 处实例化，或标准库泛型已覆盖时再泛型化 | 泛型增加阅读成本，收益必须具体 |
| 选项 struct 和函数式选项混用 | 选项集合稳定用 struct；需要向后兼容地增长用函数式 | 两套机制让调用方猜哪套生效 |
| `ctx` 存进 struct 或当可选参数 | `ctx context.Context` 放首参，沿调用链传递 | 取消和截止时间需要跨边界传播 |
| 用 context 传业务参数 | 用显式参数 | context 只承载请求域数据，如 trace、deadline |
| 构造失败时返回 `(*T)(nil)` 当 error | 返回真正的 `nil` error；检查 typed nil | 非 nil 接口装 nil 指针时 `err != nil` 为真 |

不适用：既有公开接口不能改；第三方库要求实现其接口；稳定且被广泛实现的抽象（如 `io.Reader`）可以在生产者侧定义。API 已发布后为新增参数使用函数式选项是合理的。

## 切片与 map

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| `s := make([]T, 0)` 表示空 | `var s []T` | nil 切片可直接 `len`/`range`/`append`，少一次分配 |
| 已知要 append n 个元素 | `make([]T, 0, n)` | 预分配容量，避免多次扩容复制 |
| `make([]T, n)` 后又 append | 用 `make([]T, n)` 直接索引赋值，或 `make([]T, 0, n)` | 前者长度已是 n，再 append 会得到 2n 个元素 |
| `b := append(a, x)` 后继续用 `a` | `b := append(slices.Clone(a), x)`，或确认 `a` 不再使用 | 容量足够时 append 复用底层数组，会改写 `a` |
| 长期保存 `s[0:2]` | `slices.Clone(s[0:2])` | 子切片保留整个底层数组，阻止大对象回收 |
| 依赖 map 遍历顺序输出 | 收集 key 后 `slices.Sort`，或改用有序切片 | map 遍历顺序未指定，每次可能不同 |
| `v := m[k]; if v != zero` 判断存在 | `v, ok := m[k]` | 零值无法区分「不存在」和「存在且为零值」 |
| 对 nil map 直接写入 | 先 `make`；只有读路径才允许 nil map | nil map 读安全，写会 panic |
| 用 `map[string]any` 当动态对象 | 定义 struct；只有真正任意的 JSON 边界才用 `any` | 字段不可发现，错误推迟到运行时 |
| 并发读写普通 map | 加锁；只有「写一次读多次」或键集不重叠时才用 `sync.Map` | 并发 map 写是数据竞争，运行时会直接 fatal |
| 手写循环做包含、排序、比较 | `[1.21+]` `slices.Contains`、`slices.Sort`、`maps.Equal` | 标准库更短，边界更少出错 |
| 用 `==` 比较两个 map | `maps.Equal` `[1.21+]`，或 `reflect.DeepEqual` | map 不可比较，只能和 nil 比较 |

不适用：JSON/API 明确要求 `[]` 而不是 `null` 时返回空切片并加测试；protobuf repeated 字段按其语义；需要复用缓冲区的热路径（`s[:0]`）以测量为准；模块语言版本低于 1.21 时不能使用 `slices`/`maps` 的泛型函数。

## 并发

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| `go f()` 后无人等待或取消 | 明确 owner：项目已有 `errgroup` 时用 `errgroup.Group`，否则 `WaitGroup` + `ctx` | 没有停止条件的 goroutine 会泄漏，测试和进程退出都受影响 |
| `wg.Add(1)` 写在 goroutine 内 | `wg.Add(1)` 放在 `go` 之前 | Add 和 Wait 竞争会 panic 或提前返回 |
| 接收方 `close(ch)` | 只由发送方关闭；接收方用 `v, ok := <-ch` 判断结束 | 接收方关闭会向发送方 panic |
| 用关闭 channel 表示「所有接收者完成」 | 用 `WaitGroup` 计数，或用独立的 `done` channel 广播 | 关闭是发送方语义，不是完成计数 |
| 大缓冲 channel 当队列堆积请求 | 有界 worker pool 或信号量 `chan struct{}` | 缓冲不解决生产快于消费，只推迟内存问题 |
| `select { case ...: default: }` 空转 | 阻塞 select，或加 ticker、退避 | default 让循环变成忙等，吃满 CPU |
| 循环里 `case <-time.After(d):` | 复用 `time.NewTimer` + `Reset`/`Stop`；`[1.23+]` 可直接用 | 旧版本未触发的 timer 不会被回收，长循环持续泄漏 |
| 生产循环用 `time.Sleep` 做周期任务 | `time.NewTicker` + `defer ticker.Stop()`，并在 `select` 里响应 `ctx.Done()` | Sleep 无法取消，错误累积漂移，也无法及时退出 |
| 所有状态有的加锁有的不加 | 同一状态的所有读写走同一把锁 | 漏掉的访问仍是数据竞争，`-race` 可发现 |
| 复制含 `sync.Mutex` 或 `sync.WaitGroup` 的 struct | 传指针，方法用指针接收者 | 复制同步原语会让副本保护不同状态，`go vet` 也会报 copylocks |
| 默认用 `sync.RWMutex` 替换 `Mutex` | 先用 `Mutex`；读写比实测悬殊再用 `RWMutex` | `RWMutex` 有额外开销，短临界区可能更慢 |
| `[<1.22]` 循环里启动 goroutine 捕获循环变量 | 循环内 `v := v` 再捕获，或升级模块版本 | 旧语义所有迭代共享同一变量 |
| 并发改动从不跑竞态检测 | 至少跑一次 `go test -race ./...` | 竞争往往只在调度变化时暴露 |

不适用：有意长期运行的后台 worker 必须在注释和生命周期里说明如何停止；项目统一使用 `x/sync/errgroup` 时沿用；`[1.23+]` 的 timer 行为以 `go.mod` 为准；`[1.25+]` 可用 `wg.Go(f)`，它把 Add 放在启动前。

## 测试

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| 同一检查逻辑复制成多个测试函数 | 表驱动 + `t.Run(name, ...)` | 用例增删只改数据，失败信息带用例名 |
| 每个用例断言逻辑都不同还硬套表 | 拆成独立测试函数 | 共享检查逻辑才是表驱动的条件 |
| 辅助函数失败时指向辅助函数行 | 辅助函数首行 `t.Helper()` | 失败位置回到调用点 |
| 手动 `defer os.RemoveAll` | `t.TempDir()`、`t.Cleanup(fn)` | 自动清理，子测试也正确 |
| `t.Errorf("got %v", got)` | `t.Errorf("Parse(%q) = %v, want %v", in, got, want)` | 失败信息包含输入和期望，现场可诊断 |
| `time.Sleep` 等异步结果 | channel/`WaitGroup` 同步，或带超时的轮询辅助 | sleep 既慢又不稳定 |
| `t.Parallel()` 共享循环变量 | `[<1.22]` 用 `tc := tc`；或升级模块版本 | 并行子测试共享变量会串数据 |
| 比较错误字符串 | `errors.Is`/`errors.As` | 错误文本不是契约，包装后会变 |
| 只测私有实现细节 | 测公开行为，外部测试包 `foo_test` | 重构不应导致测试重写 |
| 为格式化或重命名补测试 | 不加 | 没有行为变化，测试只增加维护量 |

不适用：集成/时序测试必须真实等待时，用宽松超时并在失败信息里打印实际耗时；项目已使用 golden file 或 `cmp.Diff` 时沿用；测试第三方行为时以对方文档为准。`t.Fatal` 只能从运行测试函数的那个 goroutine 调用，子 goroutine 用 `t.Error` + return 或把错误传回。

## 可读性与 API 表面

| 现状 | 改成 | 依据 |
| --- | --- | --- |
| 包级可变全局状态 | 用依赖注入或显式参数传递 | 全局状态让测试、并行和初始化顺序互相干扰 |
| 多字段或跨包 struct 用无字段名 literal | `T{Name: x, Size: y}` | 增加字段时不会静默错位；`go vet` 也会检查跨包用法 |
| 循环里 `s += x` 拼接字符串 | `strings.Builder` | 每次 `+=` 都复制整个字符串 |
| `fmt.Sprintf("%d", n)` 做简单转换 | `strconv.Itoa(n)` | 少一层反射和格式化开销 |
| 长函数里多层嵌套的 `if` | 提取带名字的小函数，再配合提前返回 | 缩小每段逻辑的作用域，调用点直接表达意图 |
| 注释只复述代码 | 说明原因、不变量、副作用和取舍 | 代码已说明「做什么」，注释应说明「为什么」 |
| 改了行为但不更新文档注释 | 同步更新导出声明的文档和示例 | 文档注释是公开 API 的一部分 |

不适用：性能只在测量证明是热路径后才值得牺牲可读性；项目已有日志、错误或配置框架时沿用其模式，不另建一套。

## 常见误判

- 「nil 切片一律改 `var`」不成立：JSON/API 需要空数组时保留非 nil。
- 「所有错误都 `%w`」不成立：会把底层实现细节变成公开契约；按调用方是否需要判断错误决定。
- 「所有接口都删掉」不成立：公开契约和稳定抽象仍需接口。
- 「map 一律改有序结构」不成立：确实不关心顺序时用 map 更直接。
- 「用新特性更现代」不成立：以 `go.mod` 为门槛；`go vet` `[1.23+]` 的 stdversion 分析会标记超出模块版本的符号。
- 「顺手全仓库 gofmt/rename」不成立：只格式化本次改动的文件，只迁移本次需要的调用方。

## 验证与交付

1. 只对改动的 Go 文件运行 `gofmt -w`；不要格式化无关文件。
2. 运行 `go build ./...`、`go vet ./...` 和覆盖改动的窄测试；并发改动加 `-race`，缓存干扰时加 `-count=1`。可能下载依赖或访问外部服务的命令先确认。
3. 检查 diff：没有无关改动、没有遗留调试输出、公开 API 变更已迁移调用方。
4. 汇报实际运行的检查和结果，并注明未验证的部分；不要声称未执行的检查已通过。普通编码任务不额外产出风格报告。

## 参考

- [references/sources.md](references/sources.md)：规则出处与边界，区分通用 Go 惯例、Google 特有约定和版本敏感行为。
- [references/examples.md](references/examples.md)：易错场景的完整前后对照，包括错误链、typed nil、切片别名、map 顺序、goroutine 生命周期和并行测试。
