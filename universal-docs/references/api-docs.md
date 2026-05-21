# API 文档生成规范

## 扫描策略（按语言）

### TypeScript / JavaScript
```bash
# 扫描导出的公开接口
grep -r "^export" src/ --include="*.ts" --include="*.tsx" -l
# 读取类型定义
find src/ -name "*.ts" | xargs grep -l "export interface\|export type\|export function\|export class"
```

### Python
```bash
# 扫描 __all__ 或 public 函数/类
grep -r "^def \|^class \|^async def " src/ --include="*.py" | grep -v "_"
```

### Go
```bash
find . -name "*.go" | xargs grep -l "^func [A-Z]"  # 大写 = 公开
```

### HTTP API（路由扫描）
```bash
# Express/Koa
grep -r "router\.\(get\|post\|put\|delete\|patch\)" src/ --include="*.ts"
# FastAPI/Flask
grep -r "@app\.\|@router\." . --include="*.py"
# Go (gin/echo)
grep -r "\.GET\|\.POST\|\.PUT\|\.DELETE" . --include="*.go"
```

## HTTP API 文档结构

```markdown
# API Reference

## 认证（Authentication）
[描述认证方式：Bearer Token / API Key / OAuth2 等]

## 基础 URL
`https://api.example.com/v1`

## 错误码

| 状态码 | 含义 |
|--------|------|
| 400 | 请求参数错误 |
| 401 | 未认证 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 端点列表

### POST /users
创建新用户。

**请求体**
\```json
{
  "name": "Alice",
  "email": "alice@example.com",
  "role": "user"
}
\```

**响应 201**
\```json
{
  "id": "usr_abc123",
  "name": "Alice",
  "email": "alice@example.com",
  "createdAt": "2024-01-01T00:00:00Z"
}
\```

**错误响应 400**
\```json
{ "error": "VALIDATION_ERROR", "message": "email is required" }
\```
```

## 函数库 API 文档结构

```markdown
## `functionName(param1, param2?)`

简短描述：做什么，解决什么问题。

**参数**

| 参数 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| param1 | string | ✅ | — | 说明 |
| param2 | number | ❌ | 0 | 说明 |

**返回值**：`Promise<Result>` — 描述返回内容

**异常**：`ValidationError` — 当 param1 为空时抛出

**示例**
\```typescript
const result = await functionName('value', 42);
\```
```

## 质量检查清单

生成 API 文档后自查：
- [ ] 每个端点/函数都有真实的请求和响应示例
- [ ] 错误场景有对应的错误码和示例
- [ ] 认证方式清晰说明，并有示例 header
- [ ] 可选参数标注默认值
- [ ] 示例中的 URL、字段名与代码中实际一致（已通过扫描确认）
