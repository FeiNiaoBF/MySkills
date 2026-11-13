---
name: video-edit
description: 当用户要处理本地长视频的转写、口播精简、字幕对齐或成片校验时使用；负责媒体与时间轴处理，不负责数学内容审校或 Manim 动画。
metadata:
  author: FeiNiaoBF
  version: 4.0.0
disable-model-invocation: true
---

# Video Edit

将手机长视频处理为无特效的讲解成片。本 skill 负责媒体与时间轴，不判断数学内容，也不制作 Manim。

2026-09-05 用 P1.mp4（41.9min 数学口播，1080×1920）全流程验证：28.2min 成品（-32%）+ 747 条对齐字幕。

## 输入与输出

输入：源视频、目标画幅、剪辑策略、术语提示和 episode 状态文件。

输出：

- 原片轴转写 JSON/SRT/TXT
- 可复现的剪辑决策及参数清单
- 剪辑轴字幕
- 预览成片与最终成片
- 机器校验结果和人工抽查清单

## 路径约定

本技能正文不写机器绝对路径；命令使用下列符号，具体值按机器设置（换机器只改这一处）。

不要用 `TMP`／`TEMP` 作变量名：它们是操作系统自带的临时目录变量（Git Bash 里 `$TMP=/tmp`），会与流水线的中间产物目录冲突。

| 符号 | 含义 |
|---|---|
| `$AI_ROOT` | 本地 AI 工具根：流水线脚本、规则库、ASR 模型、venv |
| `$VIDEO_TMP` | 决策文件与渲染中间产物目录（不进同步目录） |
| `$VIDEO_ROOT` | 视频工程根，每集一个 `P{n}/` 子目录 |

| 资产 | 相对位置 |
|---|---|
| 流水线脚本 | `$AI_ROOT/scripts/{video2srt,fillers_from_json,merge_fillers,remap_srt,review_filter,v1_to_ffmpeg}.py` |
| 修复引擎 + 规则库 | `$AI_ROOT/scripts/fix_srt_text.py` + `srt_rules.json` |
| 校验脚本 | `$AI_ROOT/scripts/{verify_srt,audit_v1}.py` |
| faster-whisper | `$AI_ROOT/models/faster-whisper-{medium,small}` |
| venv | `$AI_ROOT/venv`（faster-whisper + auto-editor 29.x） |
| 每视频决策文件 | `$VIDEO_TMP/${STEM}_{base,final,fillers}.json`（保留，重渲全靠它） |

## Procedure

### 1. 探测源文件

```bash
V="<视频路径>"; STEM="<名>"
ffprobe -v error -show_entries "format=duration:stream=codec_name,width,height" \
        -of default=noprint_wrappers=1 "$V"
# VFR 检查（手机视频常为 VFR）
ffprobe -v error -select_streams v -show_entries stream=r_frame_rate,avg_frame_rate \
        -of default=noprint_wrappers=1 "$V"
```

记录容器时长、视频/音频流、真实帧率、旋转元数据、色彩信息和音频采样率。逐段抽帧判断内容方向，旋转元数据不能代替画面验证。

`r_frame_rate != avg_frame_rate` 即 VFR，必须先转 CFR，否则帧号区间会错位：

```bash
ffmpeg -y -i "$V" -fps_mode cfr -c:v libx264 -preset fast -crf 18 -c:a copy "$VIDEO_TMP/${STEM}_cfr.mp4"
```

新版 FFmpeg 已移除 `-vsync`，必须用输出级 `-fps_mode cfr`。同步目录（OneDrive）云占位文件会转写失败，先确认文件已落地。

完成条件：状态文件含源文件指纹和完整探测摘要。

### 2. 本地转写

```bash
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/video2srt.py" "$V"
```

优先本地 ASR，把课程术语作为**提示**而不是强制替换。换主题先改脚本内 `initial_prompt` 术语表，否则数学口播会出系统性错字。

完成条件：三种转写产物可解析，末条时间不超过媒体时长，原始转写有只读备份（脚本自动存 `.bak`，已存在则跳过）。

### 3. 生成粗剪提案

```bash
# 语气词区间：只剪孤立的（前后停顿 ≥ 0.25s）
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/fillers_from_json.py" "<视频目录>/$STEM.json" "$VIDEO_TMP/${STEM}_fillers.json"
# 静音候选（约 40s/小时素材）
"$AI_ROOT/venv/Scripts/auto-editor.exe" "$V" --edit "audio:threshold=-35dB" --margin 0.2s \
    --export v1 -o "$VIDEO_TMP/${STEM}_base.json" --no-open --progress none
# 合并（末参 = 源视频帧率）
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/merge_fillers.py" \
    "$VIDEO_TMP/${STEM}_base.json" "$VIDEO_TMP/${STEM}_fillers.json" "$VIDEO_TMP/${STEM}_final.json" 30
```

静音检测只提出候选删除区间；安全边距保留自然呼吸和句首句尾。孤立语气词可列为候选，嵌在连续语流中的词不自动删除。先输出低成本预览或区间表给用户确认。

完成条件：删除区间无重叠、无未定义空洞，并能计算预期成片时长。

### 4. 建立时间映射

```bash
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/remap_srt.py" \
    "<视频目录>/$STEM.json" "$VIDEO_TMP/${STEM}_final.json" "<视频目录>/${STEM}_cut.srt"
```

以时间戳区间建立原片轴→剪辑轴映射；**字幕和画面方向分界都使用同一映射**。原片时间轴 ≠ 剪辑时间轴：朝向/内容分界点必须用决策文件的 `map_time()` 映射到剪辑轴再切，直接按原片秒数切会切错甚至超出剪辑片末尾。

映射后检查字幕单调、非负、无越界。段边界词时间戳有交叠，全局单调化钳制必须有。

完成条件：映射预测时长与预览实测时长在一帧量级内；否则停止渲染并定位差异。

### 5. 处理画幅

竖屏源要出横版 16:9 时用。手机横持拍摄内容侧躺（字迹转 90°），`transpose=1`（顺时针）转正——用户持机习惯固定（P1 段2 与 P2 全片同向，已验证）。

```bash
# 1. 全片朝向扫描（cropdetect 每秒采样 → 有效画幅变化点）
ffmpeg -hide_banner -i "$V" -vf "fps=1,cropdetect=limit=24:round=2:reset=0" -f null - 2> "$VIDEO_TMP/${STEM}_crop.log"
```

2. 解析日志归并连续段（逐秒取 `crop=w:h` 众数，按 `(w,h,朝向)` 归并）→ 朝向段列表。`crop` 值全程唯一 = 单一朝向；多变 = 混合朝向。
3. 抽帧 A/B 实测旋转方向：抽一帧 `transpose=1` 与 `transpose=2` 各转一次，判哪张字迹水平。
4. 混合朝向：分界点在原片轴，必须映射到剪辑轴再切（P1 实测：原片 1831s → 剪辑版 1230.17s）。
5. 每段各自转码（NVENC 8M）：

```bash
# 侧躺段（零画质损失，原生横版）
-vf "transpose=1"
# 竖条段（裁黑边放大，有代价）
-vf "crop=W:H:X:Y,scale=1920:1080:flags=lanczos,unsharp=5:5:0.4:5:5:0.0"
```

6. 分段转码后 `concat -c copy`；两侧都要 `setsar=1`。

`crop` 的 `W:H:X:Y` 以 cropdetect 日志实测值为准（P1 段1 是 `crop=1080:608:0:656`）。单一朝向的源（如 P2）直接全片 `transpose=1` 一步到位。

完成条件：每种方向至少有代表帧通过人工确认，拼接流参数一致。

### 6. 渲染

```bash
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/v1_to_ffmpeg.py" \
    "$VIDEO_TMP/${STEM}_final.json" "$VIDEO_TMP/${STEM}_cut.mp4" --rotate --ebu
```

用 `ffmpeg` 直剪（`trim`+`concat` filter，分批，速度线性），**不要**用 auto-editor 渲染。输出先落本地 `$VIDEO_TMP` 再移动进成品目录。响度归一放在音频滤镜链并单独验证（`loudnorm=I=-14:LRA=11:TP=-1.5`）；视频滤镜链只处理画面。长任务保存可恢复清单，清单绑定源指纹、决策哈希、参数和已完成分片。

完成条件：命令退出码为 0，输出能从头到尾完整解码。

### 7. 字幕与交付

```bash
# 渲染后校验（必做）
ffprobe -v error -show_entries format=duration -of default=noprint_wrappers=1:nokey=1 "${STEM}_cut.mp4"
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/audit_v1.py" "$VIDEO_TMP/${STEM}_final.json" <实测秒> 30   # GAP/OVERLAP 必须 0，偏差 <0.5s
"$AI_ROOT/venv/Scripts/python.exe" "$AI_ROOT/scripts/verify_srt.py" "<视频目录>/${STEM}_cut.srt" <实测秒>  # 格式/空文本/单调/超时长
```

先交外挂 SRT，再按状态文件决定是否烧录。烧录前抽查字体、边距、换行和公式字符；发布版使用 `faststart`。不加转场、滤镜包装或 BGM。再抽听成品 3 处（开头/语气词处/结尾）确认无剪崩。

字幕样式预设（`force_style`）：

```text
# 讲解版（默认，已验证）：白字黑边底部
FontName=Microsoft YaHei,FontSize=12,PrimaryColour=&H00FFFFFF,OutlineColour=&H00000000,BorderStyle=1,Outline=1.2,Shadow=0.4,MarginV=12
# 短视频高亮版：大字加粗黄字
FontName=Microsoft YaHei,FontSize=16,Bold=1,PrimaryColour=&H0000FFFF,OutlineColour=&H00000000,BorderStyle=1,Outline=1.5,Shadow=0.6,MarginV=20
```

烧录双代编码（剪辑渲染 + 字幕烧录）在 8M 码率下画质损失可忽略，勿降到 5M 以下。

完成条件：最终视频、外挂字幕、参数记录和验证报告齐全。

## 按需参考

- 字幕修复与参数：[text-repair.md](references/text-repair.md)。
- 内容审校：[content-review.md](references/content-review.md)。
- 实测坑：[pitfalls.md](references/pitfalls.md)。
- 选型来源：[sources.md](references/sources.md)。

## 工程目录与命名约定

```text
$VIDEO_ROOT/            # 同步目录，顶层放系列 README.md 索引
  P{n}/                 # 每支视频独立子目录
    README.md           # 本篇索引
    P{n}_内容勘误.md    # 内容审查产出（数学/事实核对 + 口误对照）
    原片/P{n}.mp4       # 原片（勿改勿删）
    转写/               # P{n}.{srt,json,txt,json.bak} 原片轴全量转写（json = 母版，词级时间戳）
    字幕/P{n}_cut.srt   # 剪辑轴字幕（已修错字 + 简繁 + 口误标注，母本）
    成品/
      竖版/  P{n}_cut / P{n}_final / P{n}_v2（各配同名 .srt）
      横版/  P{n}_cut_16x9 / P{n}_final_16x9 / P{n}_v2_16x9（各配同名 .srt）
$VIDEO_TMP/                   # 决策文件 + 渲染中间产物（不进同步目录）
```

转写产物先落原片目录再移进 `转写/`；成品旁的 srt 用字幕母本复制，保证播放器外挂同名自动加载。

## 验证

至少完成：

1. `ffprobe` 核对目标流规格。
2. `ffmpeg -f null` 完整解码视频与音频。
3. 决策时长、成片时长和字幕末尾交叉核对。
4. 首、中、尾及每个剪辑/旋转边界抽帧。
5. 抽听开头、密集剪辑段和结尾。
6. 确认源文件与最终产物均未被覆盖。
7. 删除类操作需用户实时批准，批量删除前列清单等确认。
