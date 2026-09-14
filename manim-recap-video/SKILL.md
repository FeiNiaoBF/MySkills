---
name: manim-recap-video
description: 将获批的数学概念制作为无旁白 Manim 重温段。
metadata:
  author: FeiNiaoBF
  version: 3.0.0
---

# Manim Recap Video

把用户批准的候选概念制作成教学视频末尾的无旁白动画重温段。本 skill 不选择候选、不剪手写正片，也不添加 BGM。

## 使用边界

**仅用于用户明确要求的片尾独立重温段**（手写讲解在前、动画在后）。

讲解中途替换画面、原声连续的任务走 `math-manim-insertion`（场景设计）+ `manim-video-insertion`（装配），**不要**加载这里的补静音轨与片尾拼接流程。两者的装配契约相反：这里是追加拼接，那里是画面替换且总时长不变。

结构定式：手写讲解在前（真实感）+ Manim 动画在后（直观重温）。动画段无旁白，纯内嵌字幕，BGM 留口子但不主动加。符号系统与手写一致（`x_0`／`Δx`／`f'(x_0)`）。

## 输入与输出

输入：获批候选、已核对公式、目标画幅、正片流规格、课程工作区和 episode。

输出：可阅读源码、低清预览、正式动画、装配版视频及验证记录。

## 路径约定

本技能正文不写机器绝对路径；具体值按机器设置。

| 符号 | 含义 |
|---|---|
| `<COURSE_WORKSPACE>` | 课程仓库工作区 |
| `<MANIM_ROOT>` | Manim 工程根，当前为 `<COURSE_WORKSPACE>/animations/math/` |
| `$TEXLIVE_ROOT` | TeX Live 安装根（MathTex 依赖，用前确认 `xelatex` 可用） |

沿用课程仓库现有 `animations/math` 环境（`.venv` + `uv.lock`，Python 3.12，`manim>=0.21`）；除非依赖冲突被实测确认，不为每集复制 venv。历史工程曾用独立目录 `<name>_manim`，仅作参考，不作为新集的做法。

## Procedure

### 1. 建立场景契约

为每个获批候选写：教学问题、初始状态、变化过程、最终结论、公式来源和预计时长。符号与已审校字幕一致。

完成条件：用户确认场景顺序和每场的教学目的。

### 2. 实现低清预览

```bash
# 预览（输出目录名 = 高度 + fps）
uv run manim render -r 960,540 --fps 15 <场景模块> <场景类>...
# 竖版（若需要）
uv run manim render -r 540,960 --fps 15 <场景模块> <场景类>...
```

用状态驱动的 Manim 结构表达变化关系；几何标注从对象坐标推导，避免独立悬浮装饰。字幕先退出再进入，防止交叠。横版布局按实际画幅重排，不把竖版坐标机械复用。

完成条件：所有场景能在共享环境中低清渲染，且无异常退出。

### 3. 预览验收

逐场完整观看，并抽取开始、关键变化、结束三类帧。检查公式、中文字体、遮挡、运动方向、停留时间和视觉因果。

完成条件：用户明确批准预览或列出修改项；未批准不进入正式渲染。

### 4. 正式渲染与装配

```bash
uv run manim render -r 1920,1080 --fps 30 <场景模块> <场景类>...
```

按正片目标规格正式渲染。动画无旁白时必须补与正片兼容的静音音轨；若编码参数不一致，先统一转码再拼接：

```bash
# 1. 各段 concat -c copy → anim.mp4
# 2. 补静音轨（动画无音轨）
ffmpeg -y -i anim.mp4 -f lavfi -i "anullsrc=r=44100:cl=stereo" -shortest \
    -c:v h264_nvenc -b:v 8M -pix_fmt yuv420p -c:a aac -b:a 192k -ar 44100 anim_pad.mp4
# 3. 与剪辑成品 concat -c copy → P{n}_v2（音频参数必须一致）
```

装配结果保留独立动画文件，便于学习和重做。

完成条件：动画和正片拼接成功，接缝前后音画正常，完整解码通过。

## 场景代码模板

- Base 类含 `set_cap()` 底部字幕助手：`Text`（Microsoft YaHei）+ FadeOut 旧 / FadeIn 新。
- 中文 `Text` 用 Microsoft YaHei（否则方块）；公式用 `MathTex`。
- 动态用 `ValueTracker` + `always_redraw` 驱动。

P1 场景清单（可复用模板）：

| 场景 | 内容 |
|---|---|
| S1 Increment | 微分棍缩短 → `dx` |
| S2 Secant to Tangent | 割线直角阶梯（`Δx` 横腿 + `Δy` 竖腿） |
| S3 Derivative Definition | 定义式用 `TransformMatchingShapes` 成形 |
| S4 Area Invariant | 面积恒为 2（`always_redraw` Polygon + DecimalNumber 不变量展示） |
| S5 First Order Term | 结论快闪 + 收束语 |

时长以数学节奏为准，勿注水（P1 五景共 87s）。

## 学习资产

每集工程至少包含：

```text
<MANIM_ROOT>/
├── README.md        # 讲什么、对应哪集、如何预览
├── src/             # 场景源码（按集分包子目录）
├── docs/            # 每段动画的数学内容与进出点设计
├── previews/        # 低清验收产物（抽帧 + 视频）
└── media/           # 正式渲染产物
```

README 记录课程映射，但视频分集编号与 MIT lecture/unit 编号是两套身份。课程版本与对应 lecture 只有在官方来源确认后才填写。

## Verification

- 共享 Manim 环境的版本与命令已记录。
- 每个场景低清预览通过且获用户批准。
- 正式动画分辨率、帧率、SAR、像素格式和音频流与装配策略一致。
- 动画公式与 `math-video-review` 的确认结果一致。
- 完整解码和接缝抽查通过。

## Pitfalls

实测坑，勿再踩：

1. **几何元素必须“长在几何上”**：辅助线（`Δx`/`Δy` 直角阶梯）从 P/Q 坐标推出（`corner=[Q.x, P.y]`），禁画悬空独立小棍——S2 首版悬空被用户点名返工。Q→P 重合时按 tracker 淡出，否则残留脏元素。
2. **concat 音频参数必须一致**：Manim 输出无音轨，先补 `anullsrc` 静音轨并显式对齐编码（aac／44100／stereo）。
3. **竖改横不是只改 `-r`**：内容坐标（如 S4 面积数字 `DOWN*1.15`）在横版会压到图形上，需重排到空白区（`DOWN*2.6 + LEFT*4.2`）；改完预览抽帧确认。
4. **`set_cap` 交叉淡化瞬态**：FadeOut 旧 + FadeIn 新 同时进行时，新旧字幕宽度不同会短暂交叠约 0.5s；在意的话改成先后淡出。
5. **ffmpeg 多输入抽帧的流选择坑**：`-ss t1 -i a -ss t2 -i b -frames:v 1 x y` 全部取自同一流 → 逐帧单独跑 ffmpeg 再拼图。

另有：

- `always_redraw` 的回调必须返回 Mobject；标量关系使用普通函数。
- 动态辅助线应从被标注对象计算，重合时显式淡出。
- Manim 默认无音轨；直接与含音轨正片拼接可能失败或出现异常。
- 输出分辨率变化会改变可用布局，不只是一项渲染参数变化。
- 无视觉检查能力时把预览交给用户验收，不宣称画面正确；静帧不能验证动态过程，涉及运动/收敛/转场时抽一小段序列看。
