# Skill 编辑清单

改 `skills/` 下的 `SKILL.md` 前读这一页；完整规则见 [CONTRIBUTING.md](../../CONTRIBUTING.md)。

## 入口

- `name` 与目录名一致，frontmatter 从首行 `---` 开始且是可解析 YAML。
- `description` 先说明触发条件，再说明能力范围；**前 57 个字符应说明何时使用**。
- 自动调用型 `description` 不超过 140 字符，手动触发型不超过 100 字符。近义词与触发短语放正文 `## When to use`，不塞进描述。
- `metadata` 必须为映射；作者、版本、日期等自定义字段放在其中。
- 本 Skill 内的相对链接必须指向真实文件；跨 Skill 用名称路由，不用跨目录相对链接。

## 内容

- Skill 必须能在另一台机器上独立使用：关键步骤、安全边界和验证条件不能依赖聊天记录或作者机器。
- 先写可执行步骤和完成条件，再提供按分支读取的参考入口；不要求每次通读所有材料。
- 同一规则只维护一处。不重复缓存配置、脚本帮助或可直接查到的结构；保留不能从环境推导出的理由、约束与经验。
- 按任务需要保留指导，不堆叠模型已能可靠完成的通用流程。用实际用例验证精简效果，不把更短等同于更好。
- 示例使用 `<VPS_IP>`、`<ROUTER_IP>`、`<SSH_HOST>`、`<your-key>.pem` 等占位符。公开文件不得包含真实用户的地址、主机名、用户名、节点 ID、凭据、私钥、订阅 token 或机器专属路径。
- 网络、路由器、云主机及媒体操作先做只读诊断；变更前说明影响、备份、回滚和验证方法。
- 不写入、不上传、不覆盖、不删除的约束保持原句，不压缩成概括。

## 验证

在仓库根目录执行，虚拟环境准备见 CONTRIBUTING：

```bash
python -B scripts/validate_skills.py
python -B -m unittest discover -s tests -v
git diff --check
python -B scripts/context_budget.py
```

<!-- ctx-canary: CTX-CANARY-7F3A9（加载验证标记；新会话应能通过 AGENTS 指针读到本行） -->
