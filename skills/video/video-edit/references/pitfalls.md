# 实测坑

从 [video-edit](../SKILL.md) 拆出的按需参考；遇到异常或复现旧问题时读。

1. auto-editor v29 大改：无 `--silent-threshold`，改用 `--edit "audio:threshold=-35dB"`；`--export v1` 的 JSON 可直接当输入渲染。
2. faster-whisper GPU 要 `cublas64_12.dll`（venv 默认无）；CPU int8 足够（small 9.3x，medium 2.2x，14 线程）。
3. whisper 词级 token 中文 = 单字；small 不带标点、medium 带标点。切分必须用段落文本逐字对齐词序列（标点挂前词），直接按字数切会切碎词。
4. medium 段落常连到 30s+，SRT 必须重切；段边界词时间戳有交叠，全局单调化钳制必须有。
5. 数学口播必须 `initial_prompt` 术语表，否则“导数→岛数”“公式→公司”。残留错字模式：相迁／学历／解（减）／SOPX／区域（趋于）／delta 三种切法／简繁漂移。
6. `merge_fillers` 曾把剪除区间留成 GAP（不渲染 = 时长对，但 remap 漏算 → 字幕漂移 8.3s）。已修：交叠必须显式 99999 块零 GAP；`audit_v1.py` 就是防复发用的。
7. 渲染输出勿直接写同步目录（锁文件报“被占用”），先落本地再移动。
8. 手机竖屏视频常为 VFR：帧率不一致时秒→帧换算会与 auto-editor 内部帧号错位；`audit_v1` 偏差 > 0.5s 先查这个。
9. 删静音后响度感知更不稳定：用 `loudnorm=I=-14:LRA=11:TP=-1.5` 归一到流媒体标准。
10. 错字不可完全预测：同音频重跑 medium 错字会漂移；转写 JSON 是一次性资产勿重跑。同音错字 `lp` 正常，`logprob` 只能筛幻觉/漏听。
11. LLM 降耗分工（实测）：同主题第 2 个视频起，规则库 + 摘要预筛 + 残留驱动 = 审校类 token 降 60–80%；首视频无法省。
12. ffmpeg 多输入抽帧的流选择坑：`-ss t1 -i a -ss t2 -i b -frames:v 1 x y` 全部取自同一流 → 逐帧单独跑 ffmpeg 再 `xstack`/`hstack` 拼图；抽帧自检必须逐格看（曾发现 10 帧全同帧）。
13. auto-editor 渲染有 seek 风暴：v1 决策 chunk 密度高（>1000）时耗时随素材时长非线性爆炸。渲染改用 `v1_to_ffmpeg.py` 直剪，auto-editor 只保留 `--export v1` 做决策分析。
14. 后台长渲染进程易被外部 -15 杀（机器睡眠/杀软/OOM），特征：日志卡在 "Creating audio"，输出文件停止增长。长渲染分段跑（每段 < 2min）规避；检查退出码是否为 -15。
15. 原片时间轴 ≠ 剪辑时间轴：朝向/内容分界点必须用 `map_time()` 映射到剪辑轴再切（曾发生 `-ss` 超出 cut 片末尾的空文件事故）。
16. `concat`/滤镜 SAR 必须一致：`crop`/`scale` 后 SAR 会带尾数（1215:1216），`concat` 前两侧都 `setsar=1`，否则报 "SAR do not match"。
17. 字幕烧录（`subtitles`）路径：用绝对路径且转义盘符冒号，文件必须真实存在（曾写错路径连烧 3 次失败）；srt 先复制到本地 `$VIDEO_TMP` 再烧，别直接读同步目录。
18. 手机横持拍摄源（HEVC 常带旋转元数据）：ffmpeg 解码自动转成竖帧，内容侧躺；`transpose=1` 顺时针转正；`cropdetect` 在解码后帧上跑。

另有若干通用判断坑：

- 不从单次失败猜测同步盘、睡眠、显存或 seek 是根因；记录进程、日志和退出码后再下结论。
- 音频 true peak 超标说明输出存在峰值风险，不足以单独证明录音时 ADC 削波。
- 两个文件时长相近不足以证明内容重复；需比较时间线或抽样内容。
- 复用旧脚本前先做短样本端到端测试；固定 fps、滤镜链类型错误、脆弱断点续跑和提前删除分片均是阻断项。
