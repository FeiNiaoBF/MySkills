# Go 风格示例

这里的完整前后对照覆盖容易改错的场景。规范规则见 [SKILL.md](../SKILL.md)；本文件只补充可直接对照的代码。所有版本标记都以 `go.mod` 为准。

## 错误链与日志边界

```go
// 改前：丢失上下文，同一错误在两层重复记录。
func load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("read failed: %v", err)
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		log.Printf("unmarshal failed: %v", err)
		return nil, err
	}
	return &c, nil
}

// 改后：补充步骤，保留原因，只在外层记录一次。
func load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}
	return &c, nil
}
```

调用方现在可以通过包装后的错误匹配 `errors.Is(err, fs.ErrNotExist)`。这里的 `%q` 只打印输入路径；如果路径本身属于敏感信息，就不要把它写进错误串。

写路径的关闭错误需要具名返回值：

```go
func writeFile(path string, data []byte) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { err = errors.Join(err, f.Close()) }()
	_, err = f.Write(data)
	return err
}
```

`errors.Join` 需要 `[1.20+]`。更低版本可以在 `err == nil` 时返回 `f.Close()`，或分别报告写错误和关闭错误。

## 接口归属与 typed nil

```go
// 生产方 storage 预先定义了宽接口。
package storage

type Store interface {
	Get(context.Context, string) (*User, error)
	Put(context.Context, *User) error
	Delete(context.Context, string) error
	List(context.Context) ([]*User, error)
	Close() error
}

// 消费方 report 只用到其中一个方法。
package report

// 更好：在使用方定义窄接口。
type userGetter interface {
	Get(context.Context, string) (*User, error)
}

func Build(ctx context.Context, users userGetter) (*Report, error) { /* ... */ }
```

生产方仍然可以返回 `*PostgresStore`；需要完整能力的调用方继续使用具体类型。

```go
// 改前：nil 的 *bytes.Buffer 装进 io.Reader 后，接口本身不为 nil。
func newReader(ok bool) (io.Reader, error) {
	var buf *bytes.Buffer
	if ok {
		buf = &bytes.Buffer{}
	}
	return buf, nil
}

r, err := newReader(false)
if err != nil {
	return err
}
_, _ = r.Read(nil) // r != nil，随后 panic: nil pointer dereference
```

修复方式是返回具体类型 `(*bytes.Buffer, error)`，或者在「不存在」时显式返回 `nil` 接口。错误返回值同理：`var e *MyError; return e` 会得到一个非 nil 的 `error`。

## 切片别名、底层数组保留与 JSON 形状

```go
base := []int{1, 2, 3, 4}
a := base[:2]      // len(a) == 2, cap(a) == 4
b := append(a, 99) // 写入 base[2]
// b == [1 2 99]，base == [1 2 99 4]
```

按意图有两种正确写法：

```go
b := append(slices.Clone(a), 99) // 完全脱离原数组；[1.21+]
b := append(a[:2:2], 99)         // 限制 cap，append 时分配新数组
```

子切片会保留整个底层数组：

```go
func firstRecord(buf []byte) []byte {
	return buf[:1024] // 调用方会一直持有整个 buf
}

func firstRecord(buf []byte) []byte {
	return slices.Clone(buf[:1024]) // 只保留这条记录
}
```

`nil` 切片和空切片在 JSON 边界上可观察地不同：

```go
type Response struct {
	Items []Item `json:"items"`
}

json.Marshal(Response{})                // {"items":null}
json.Marshal(Response{Items: []Item{}}) // {"items":[]}
```

内部默认保留 `nil`；只有外部形状要求 `[]` 时，才在出口构造非 nil 切片，并用测试固定这个行为。

## map 顺序与键存在性

```go
// 改前：每次运行的输出顺序都可能不同。
var out []string
for k, v := range m {
	out = append(out, k+"="+v)
}

// 改后：顺序确定。Go 1.21 之前可用 sort.Strings。
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
slices.Sort(keys) // [1.21+]
for _, k := range keys {
	out = append(out, k+"="+m[k])
}
```

区分「不存在」和「存在且为零值」：

```go
if v, ok := m[k]; ok {
	use(v) // m[k] 本身可能就是 ""
}
```

遍历时删除当前 key 是安全的；遍历时新增的条目可能被产出，也可能不被产出，因此不要依赖遍历期间的写入。

## goroutine 生命周期、并发上限与 timer

```go
// 改前：无上限、无归属、忽略取消和错误。
func process(ctx context.Context, items []Item) {
	for _, it := range items {
		go func() { handle(it) }()
	}
}
```

```go
// 改后：有界、可取消、错误可见。
func process(ctx context.Context, items []Item) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(8)
	for _, it := range items {
		it := it // 仅 Go 1.22 之前需要
		g.Go(func() error { return handle(ctx, it) })
	}
	return g.Wait()
}
```

`errgroup` 来自 `golang.org/x/sync/errgroup`；项目已有该依赖时使用。没有时可以用 `context.WithCancel`、信号量 channel 和 `WaitGroup` 组合出同样语义。

只用标准库的有界写法：

```go
sem := make(chan struct{}, 8)
var wg sync.WaitGroup
loop:
for _, it := range items {
	it := it
	select {
	case sem <- struct{}{}:
	case <-ctx.Done():
		break loop
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { <-sem }()
		handle(ctx, it)
	}()
}
wg.Wait()
```

这个变体要自行收集 `handle` 的错误（例如加锁的切片或 error channel）；`errgroup` 已经代劳。

长循环中的 timer 复用：

```go
// 改前：模块低于 Go 1.23 时，每个未触发的 timer 都会持续可达。
for {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Second):
		tick()
	}
}
```

```go
// 改后：只用一个 timer，显式重置。
t := time.NewTimer(time.Second)
defer t.Stop()
for {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		tick()
		t.Reset(time.Second)
	}
}
```

`[1.23+]` 起，不再被引用的 timer 会立即被回收，timer channel 也变为无缓冲，因此 `time.After` 写法可以接受；但显式 timer 对周期的表达仍更清楚。Go 1.25 增加了 `wg.Go(f)`，它在启动 goroutine 前完成 `Add`。

## 表驱动测试、helper 与并行

```go
func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    int
		wantErr error
	}{
		{"empty", "", 0, ErrEmpty},
		{"number", "42", 42, nil},
	}

	for _, tc := range tests {
		tc := tc // 仅 Go 1.22 之前需要
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := Parse(tc.in)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Parse(%q) error = %v, want %v", tc.in, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Parse(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
```

```go
func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	t.Cleanup(srv.Close)
	return srv
}
```

没有 `t.Helper()` 时，共享初始化函数内部的失败会指向 helper 自己的行号，而不是调用点。`t.Cleanup` 在子测试失败时也会执行，而只注册在父测试里的清理不会按子测试粒度执行。

异步副作用要用同步信号等待，而不是 sleep：

```go
done := make(chan struct{})
go func() {
	defer close(done)
	worker()
}()

select {
case <-done:
case <-time.After(5 * time.Second):
	t.Fatal("worker did not finish within 5s")
}
```
