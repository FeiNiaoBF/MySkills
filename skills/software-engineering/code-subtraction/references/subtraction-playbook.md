# 减法场景手册

SKILL.md 工作流程的配套参考：各层次减法速查与输出报告格式。

## 代码层面

**删除死代码**
- 注释掉的大段代码 → 直接删（有 Git 不怕找不回来）
- 未使用的 import → 工具自动检测（ESLint `no-unused-vars`、Pyright、`go vet`）
- 未使用的函数/变量 → IDE "Find Usages" 确认后删除

**合并重复逻辑**

```python
# ❌ 减法前：重复
def process_user(user): validate(user); log("user"); save(user)
def process_admin(admin): validate(admin); log("admin"); save(admin)

# ✅ 减法后：参数化
def process_entity(entity, kind): validate(entity); log(kind); save(entity)
```

⚠️ 警惕过度参数化：为消除两行重复引入几十行配置，得不偿失。

**简化条件（卫语句）**

```javascript
// ❌ 减法前：嵌套
function process(data) {
  if (data) {
    if (data.valid) {
      if (data.items.length > 0) {
        return doWork(data);
      }
    }
  }
}

// ✅ 减法后：提前返回
function process(data) {
  if (!data || !data.valid || data.items.length === 0) return;
  return doWork(data);
}
```

**减少不必要的抽象**
- 接口只有一个实现且不会扩展 → 删除接口，用具体类
- 只调用一次的函数且不表达独立概念 → 内联
- `service → manager → handler → util` 四层 → 考虑合并为两层

## 依赖层面

```bash
npm:   npx depcheck
maven: mvn dependency:analyze
go:    go mod why <package>
pip:   pip-extra-reqs / deptry
```

优先用语言内置 API 替代外部库（如 `fetch` 替代 `axios`，`java.net.http` 替代 Apache HttpClient）。

## 模块/服务层面

合并信号：
- 两个模块**频繁一起修改**且被**相同调用方**使用 → 合并
- 一个"独立服务"只被**一个调用方**使用 → 降级为内部函数
- 接口**半年内无调用** → 先消融确认无真实依赖，再下线

## 减法分析报告格式

```
## 减法分析报告

### 🔍 扫描结果（共发现 N 处）

| 编号 | 层次 | 位置 | 问题 | 风险 |
|------|------|------|------|------|
| 1    | 句法 | ...  | ...  | 🟢   |

### 📋 逐项建议

#### #1：[简短标题]
**问题**：...
**建议**：删除 / 合并 / 内联
**消融实验**：基线 <测试/冒烟结果> → 移除后 <结果> → 判定 <可删/需保留>
**改后代码**：
```[语言]
...
```
**风险**：低 / 中 / 高，[验证方法]

### ✅ 安全删除顺序
1. 先删 🟢 低风险项（...）
2. 再处理 🟡 中等风险项（先标记废弃）
3. 🔴 高风险项补测试 + 消融对照后再操作
```
