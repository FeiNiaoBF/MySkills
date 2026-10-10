---
name: linear-algebra
description: >-
  基于 Gilbert Strang（MIT 18.06）体系的线性代数笔记 skill。
  用于在 Obsidian 知识库中新建和更新线性代数知识笔记——写对、写好、边写边自检。
  核心原则：纯数学、Strang 为依据、自然双链、三层骨架。
  涉及"线代"、"线性代数"、"Strang"、"向量"、"矩阵"、"特征值"等主题时触发。
metadata:
  origin: local skill migrated from ~/.agents/skills
  version: "1.0.0"
compatibility:
  - Obsidian Flavored Markdown ($...$ LaTeX, > [!tip] callout, [[wikilinks]])
  - Gilbert Strang "Introduction to Linear Algebra" 6th ed.
  - MIT OCW 18.06 / 2020 Vision / Five Factorizations
  - Vault Meta 规范（Library YAML 属性规范、Tag 命名空间规范等）
---

# Linear Algebra Note Skill

基于 Gilbert Strang（MIT 18.06）体系的线性代数笔记 skill。服务于 Obsidian 知识库「Codex Vitae」的 `03.Library/Notes/02.Mathematics/Math/` 目录。

---

## 核心原则

### 1. 纯数学

笔记内容为纯数学知识，不包含：
- 考试考点提示（教资、考研等——如有需求另开笔记，不混入）
- 编程/框架/库的具体应用（如有需求另开笔记，不混入）
- 个人学习计划、进度追踪等元信息

但允许：
- 用编程/物理/日常场景做类比辅助理解（类比不是应用笔记）
- 标注"这个概念的典型应用领域"作为知识背景

### 2. Strang 为依据

教材依据按优先级从高到低：

| 优先级 | 来源 | 覆盖范围 | 使用场景 |
|--------|------|---------|---------|
| 1 | MIT OCW 18.06 课堂录像 + 讲稿 | Lecture 1–34 | 核实 Strang 的具体说法、授课顺序 |
| 2 | Strang, *Intro to Linear Algebra* 6th ed. | 全书 10 章 | 章节结构、定义、练习题型 |
| 3 | 2020 Vision / Five Factorizations (2023) | 补充视频 | 2020 年后的教学观点更新 |

- 核心概念笔记 → 搜 Lecture 讲稿 + 教材对应章节大纲
- 小修补/确认某个引用 → 只搜争议关键词
- **不从记忆里编 Strang 的原话。**

### 3. 自然双链

[[Wikilinks]] **只能出现在正文叙述中**。不出现任何独立链接列表：
- ❌ "关键关系"、"See Also"、"相关笔记"、"前驱概念"、"后续概念"等独立列表
- ❌ 正文末尾的链接汇总
- ✅ "矩阵乘法本质上就是[[线性变换|变换的复合]]……"——在叙述中自然引出

> 唯一例外：笔记末尾的"依据" callout 可以包含指向 SignalSource 或外部资源的链接，那些不是概念双链，是溯源。

### 4. 三层骨架

每篇笔记的正文由至多三个自然段落组成。不强制全有，但成熟的笔记应完整：

```
┌─────────────────────────────────────────────┐
│ ①　直觉入口（第一句话或 callout）              │
│     一句话/一个类比/一个反常识问题              │
│     目的：拉读者进来，让大脑知道"要接什么"       │
│     如果 reader 只看这一句，ta 应该能大致         │
│     知道这篇在讲什么                           │
├─────────────────────────────────────────────┤
│ ②　核心叙述（正文主体）                        │
│     按知识本身的内在逻辑展开，不拘一格            │
│     允许：分节、推导、对话、对比、例子            │
│     [[双链]] 只出现在这里的行文中                │
│     全文唯一展示双链的地方                      │
├─────────────────────────────────────────────┤
│ ③　溯源锚点                                    │
│     这篇知识从哪来的                           │
│     格式：```> [!cite] 依据                    │
│     一行即可，列教材/课程/论文/实践来源         │
│     放在文末，不混入正文                       │
└─────────────────────────────────────────────┘
```

### 5. 跨类型通用

三层骨架适用于所有 type，但 ② 的展开方式因 type 而异：

| type | ② 核心叙述的倾向 |
|------|---------------|
| `concept` | 定义 → 直觉拆解 → 关键例子 |
| `method` | 问题 → 操作步骤 → 适用范围与局限 |
| `case` | 背景 → 过程 → 结果与提取 |
| `question` | 什么被质疑 → 几种立场 → 结论 |
| `summary` | 要汇总什么 → 关键提炼 |
| `moc` | 按三层骨架走不通——moc 的写法由 MOC 模板单独规范 |

### 6. 边写边自检

写作过程中同步做以下检查，不在写完后另起 QA 步骤：

- **溯源检查**：刚写下的引述/结论，是否需要搜原始材料确认？→ 按两级搜索深度执行
- **双链检查**：刚写的这段叙述，是否存在一个现有笔记可以用 `[[wikilink]]` 自然嵌入？→ 优先给概念名加链，不强行
- **规范检查**：frontmatter 是否完整？type/tags 是否符合 Library YAML 规范？
- **冗余检查**：刚写的内容，是否已有其他笔记覆盖过？→ 更新已有笔记的对应段落，不重复写新笔记

---

## 前置条件

在执行任何笔记操作前，确认以下规范已读：

| 规范文件 | 原因 |
|---------|------|
| Library YAML 属性规范 | type、status、subject、moc、tags 等字段规则 |
| Library 笔记模板规范 | Note.md / Knowledge.md 的初始格式参考 |
| Tag 命名空间规范 | tags 的格式、命名空间、禁止规则 |
| Tag 治理巡检规范 | 写完后是否引入了治理问题 |
| 03.Library 规范 | Library 层级的整体边界 |

MOC 相关笔记同时参考 MOC 模板（顶层 MOC.md / 普通主题 MOC.md）。

---

## 触发方式

### 主动触发

```
/linear-algebra <指令>
```

指令包括：

| 指令 | 行为 |
|------|------|
| `写笔记：<标题>` | 按三层骨架 + 搜索规范新建一篇笔记 |
| `更新：<标题>` | 更新已有笔记内容，按修改幅度选择搜索深度 |
| `补充：<标题>` | 在已有笔记中新增一个知识点段落 |
| `检查：<标题>` | 对已有笔记做一遍六项自检，报告结果不修改 |
| `自检` | 对该目录下所有 `status: reviewing` 的笔记做批量自检 |

### 被动触发

当对话中涉及以下主题且需要写/更新笔记时，自动激活：
- "线代"、"线性代数"、"Strang"、"18.06"
- 向量、矩阵、线性变换、行列式、特征值、特征向量、SVD
- 四个子空间、列空间、零空间、行空间、左零空间
- 线性无关、基、维数、秩、正交、投影
- 高斯消元、LU 分解、QR 分解、Gram-Schmidt
- 对角化、二次型、正定矩阵

### 不触发

当上下文明显不是笔记场景时（如闲聊、学习过程提问、考试练习），不激活 skill。

---

## 输出格式

### 新建笔记

```yaml
# frontmatter 符合现有 Library YAML 规范
---
type: concept    # 按内容实际选择
status: reviewing
created: <当日>
updated: <当日>
subject: Mathematics
moc: "[[MOC｜线性代数]]"
tags:
  - topic/mathematics
---
```

正文按三层骨架组织。

### 更新笔记

不做全文重写，只做段落级 edit。保留已有前言的合理内容，只修正/补充/替换需要变的部分。

---

## 外部参考

- MIT OCW: https://ocw.mit.edu/courses/18-06-linear-algebra-spring-2010/
- MIT 2020 Vision: https://ocw.mit.edu/courses/res-18-010-a-2020-vision-of-linear-algebra-spring-2020/
- Strang ILA 6th ed. outline: https://math.mit.edu/~gs/linearalgebra/ila6/ila6outline.pdf
