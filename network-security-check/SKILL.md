---
name: network-security-check
description: PassWall/sing-box diagnosis, safe changes, and rollback.
metadata:
  author: FeiNiaoBF
  version: 2.0.0
---

# Network Security Check

这是一个**网络故障与变更流程 skill**，不是单次命令检查器。它把 PassWall、sing-box、OpenWrt/ImmortalWrt、VPS 和 Windows 客户端问题，从现象推进到可验证结论：

```text
接收症状 → 建立基线 → 缩小假设 → 最小实验 → 变更审批 → 备份变更 → 回归验证 → 关闭或回滚
```

目标是保护当前可用网络，避免为了修复一个域名而破坏全局流量。

## 何时使用

使用本 skill 处理以下完整任务：

- DNS、IP、IPv6 泄漏或解析异常的诊断
- PassWall 分流、节点、透明代理或重启变更
- sing-box VLESS Reality 节点的连通性验证
- Claude、Cloudflare 等服务的网络故障定位
- 已知配置变更后的验证、回滚和 incident closeout

仅问一个简单的命令或概念时，不启动完整流程；直接回答即可。

## 安全门

在进入变更阶段前，必须满足：

1. 明确症状、影响范围、目标和当前是否仍可联网。
2. 先完成只读基线；不要用猜测替代输出证据。
3. 对每个候选变更写明影响键、备份方式、成功标准和回滚命令。
4. 一次只做一个配置域的变更；上传、安装、校验、清理必须顺序执行。
5. 全局节点切换、DNS 模式切换、PassWall restart/reload 都属于高风险变更，必须说明可能的短暂断网并取得用户同意。
6. Skill 内只使用占位符，不保存真实 IP、节点 ID、局域网地址、主机名、用户名或凭据。

如果用户报告“国内外都断了”，暂停优化，直接进入 [Recovery path](#recovery-path)。

## Procedure

### 1. Intake：建立问题契约

先记录：

- **症状**：哪个域名、客户端或协议失败，错误是什么。
- **范围**：单域名、单节点、单设备，还是全局网络。
- **时间线**：最后一次已知正常状态，以及最近的变更。
- **目标**：恢复可用性、验证节点、修复分流，还是准备一次配置变更。
- **限制**：是否允许短暂断网、是否能 SSH、是否只能读配置。

输出一句可验证的问题定义，例如：

> 在不改变全局 TCP 节点的前提下，判断 `api.anthropic.com` 失败是 DNS、路由、Reality 握手还是浏览器挑战。

完成条件：范围、目标和变更权限均明确；仍有歧义时只问一个会改变流程的问题。

### 2. Baseline：采集只读证据

客户端基线：

```powershell
curl.exe -4 --max-time 15 https://api.ipify.org
curl.exe -6 --max-time 15 https://api64.ipify.org
nslookup baidu.com <ROUTER_IP>
nslookup api.ipify.org <ROUTER_IP>
curl.exe -4 -I --max-time 15 https://api.anthropic.com
```

路由器只读基线：

```sh
uci show passwall.@global[0]
uci show passwall.myshunt
uci show passwall.AIGC
ps w
netstat -lntup
free
swapon -s
```

记录时间、命令、退出码和关键输出。不要把完整配置、订阅内容、私钥或 token 粘贴进报告。

完成条件：已经区分“解析失败、连接失败、代理失败、服务端拒绝和浏览器挑战”；否则继续读取证据，不进入变更。

### 3. Hypothesis：建立故障树

按证据而不是直觉分类：

| 假设 | 先看什么 | 典型下一步 |
|---|---|---|
| DNS / 负缓存 | `nslookup`、生成的 DNS 配置、监听端口 | 查上游、端口和 Windows 缓存 |
| 路由 / 分流 | Direct list、shunt 规则、目标节点 IP | 用临时测试流量验证，不改全局 |
| Reality / TLS | `server_name`、公私钥方向、`short_id`、时间差 | 用真实 sing-box 客户端复测 |
| 浏览器挑战 | API HTTP 状态、浏览器与 curl 差异 | 不把 403 challenge 直接判为线路坏 |
| 客户端状态 | DNS 负缓存、旧进程、代理环境变量 | 清理并重复同一基线 |

详细命令和判读见 [PassWall/VPS runbook](references/passwall-vps-lessons.md)。

完成条件：保留一个或多个可被实验区分的假设，并明确每个实验不会改变什么。

### 4. Minimal test：先验证，不动主流量

优先使用临时本地 socks/mixed inbound 或单域名测试：

```sh
curl -4 -sS --socks5-hostname 127.0.0.1:<LOCAL_SOCKS_PORT> \
  --connect-timeout 10 --max-time 20 https://api.ipify.org
curl -4 -sS -I --socks5-hostname 127.0.0.1:<LOCAL_SOCKS_PORT> \
  --connect-timeout 10 --max-time 20 https://api.anthropic.com
```

实验纪律：

- 不为了单域名测试修改 `passwall.@global[0].tcp_node`。
- 不把临时成功误判为全局配置安全。
- 每次实验只改变一个变量，并保存前后输出。

完成条件：实验结果能排除至少一个假设，或明确说明为什么需要用户批准下一步高风险动作。

### 5. Change plan：写出可回滚方案

在任何写操作前，输出：

```text
目标：
将要改变的配置域：
当前值摘要：
备份位置/内容：
预期影响：
成功标准：
回滚命令：
验证命令：
```

只读证据不足时，不提出“顺手优化”或多个配置一起改。

完成条件：用户能看懂影响范围，并且失败后可以恢复到已知稳定状态。

### 6. Snapshot and mutation：备份后单点变更

- 保存相关 UCI 配置和生成文件的最小快照。
- 一次只改一个配置域。
- 写入、重载、校验按顺序执行。
- 原配置、原视频或其他源资产不覆盖、不删除。
- 高风险 restart/reload 前再次确认短暂断网风险。

完成条件：变更命令、退出码和实际写入值均已记录。

### 7. Validate：重复基线并验证目标

至少执行：

1. 与变更前相同的客户端和路由器基线。
2. 目标域名/节点的最小测试。
3. DNS、IPv4/IPv6、代理出口和 HTTP 状态检查。
4. PassWall 进程、监听端口和关键配置检查。
5. 若涉及浏览器服务，区分 API 可达性与浏览器挑战。

完成条件：成功标准全部有实际输出支持；否则进入回滚或继续诊断，不能报告“已修复”。

### 8. Recovery path：全局故障优先恢复

当国内外流量都失败时：

1. 停止新增优化。
2. 对照最近一次已知稳定快照恢复全局节点和 DNS 模式。
3. 顺序执行 stop → 检查残留进程/端口 → start。
4. 清理 Windows DNS 负缓存。
5. 重跑最小基线。
6. 记录恢复前后差异，再决定是否单独诊断原始问题。

恢复命令模板和成功标准见 [PassWall/VPS runbook](references/passwall-vps-lessons.md)。

完成条件：基础 DNS、国内域名、目标 API 和代理状态均已重新验证，或明确记录仍未恢复的范围。

## 输出格式

每次流程结束时报告：

```markdown
## 结论
- 根因 / 当前最佳假设：
- 影响范围：

## 证据
- 变更前：
- 实验结果：
- 变更后：

## 已执行变更
- 配置域：
- 备份：
- 回滚方式：

## 验证
- 通过：
- 未覆盖：
- 用户需要决定：
```

## 不属于本 skill

- 通用应用代码安全审计
- 渗透测试或攻击性操作
- 自动申请云凭据、生成私钥或写入订阅 token
- 没有症状和目标的无差别“优化网络”
