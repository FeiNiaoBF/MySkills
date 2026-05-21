# 阶段 6：状态管理类型

## Zustand 最佳实践

```tsx
import { create } from 'zustand';

interface CounterStore {
  count: number;
  increment: () => void;
  setCount: (n: number) => void;
}

// ✅ 用接口明确 store 的完整类型
const useCounterStore = create<CounterStore>((set) => ({
  count: 0,
  increment: () => set((state) => ({ count: state.count + 1 })),
  setCount: (n) => set({ count: n }),
}));

// 组件中使用（只订阅需要的部分，避免不必要重渲染）
const count = useCounterStore((state) => state.count);
```

## Zod 校验 API 响应

为什么不用 `as` 断言？因为**运行时数据不可信**，类型断言在编译期通过，但无法阻止服务端返回意外字段或类型。

```tsx
import { z } from 'zod';

// 定义 Schema（Zod 自动推导 TS 类型）
const UserSchema = z.object({
  id: z.number(),
  name: z.string(),
  email: z.string().email(),
  role: z.enum(['admin', 'user']),
});

type User = z.infer<typeof UserSchema>; // 自动推断，不用手写接口

// 在数据获取时验证
async function fetchUser(id: number): Promise<User> {
  const res = await fetch(`/api/users/${id}`);
  const raw = await res.json();
  
  // ✅ 运行时校验，不符合则抛错
  return UserSchema.parse(raw);
  
  // ❌ 不要这样 - 没有任何运行时保护
  // return raw as User;
}
```

---

# 阶段 7：React Router 路由类型

```tsx
import { useParams, useSearchParams } from 'react-router-dom';

// useParams - 路由参数
// 路由定义：/users/:userId/posts/:postId
const { userId, postId } = useParams<{ userId: string; postId: string }>();
// 注意：params 始终是 string，需要自行转换
const id = Number(userId); // 如果需要数字

// useSearchParams - 查询参数（?page=1&filter=active）
const [searchParams, setSearchParams] = useSearchParams();
const page = searchParams.get('page'); // string | null
const pageNum = page ? parseInt(page, 10) : 1;
```

---

# 阶段 8：工程实践

## 自定义 Hooks 泛型设计

```tsx
// 通用 useFetch Hook
interface FetchState<T> {
  data: T | null;
  loading: boolean;
  error: Error | null;
}

function useFetch<T>(url: string): FetchState<T> {
  const [state, setState] = useState<FetchState<T>>({
    data: null,
    loading: true,
    error: null,
  });

  useEffect(() => {
    fetch(url)
      .then((res) => res.json())
      .then((data: T) => setState({ data, loading: false, error: null }))
      .catch((error) => setState({ data: null, loading: false, error }));
  }, [url]);

  return state;
}

// 使用：T 由调用方决定
const { data, loading } = useFetch<User[]>('/api/users');
// data 的类型是 User[] | null ✅
```

---

# 阶段 9：代码质量规则速查

## 关键 tsconfig.json 严格选项

```json
{
  "compilerOptions": {
    "strict": true,          // 启用所有严格检查
    "noImplicitAny": true,   // 禁止隐式 any
    "strictNullChecks": true, // null/undefined 必须显式处理
    "noUnusedLocals": true,  // 未使用变量报错
    "noUnusedParameters": true // 未使用参数报错
  }
}
```

## 常见反模式 vs 正确写法

```tsx
// ❌ 类型断言掩盖问题
const input = document.getElementById('search') as HTMLInputElement;
input.value = 'test'; // 如果元素不存在，运行时报错

// ✅ 安全的类型收窄
const el = document.getElementById('search');
if (el instanceof HTMLInputElement) {
  el.value = 'test';
}

// ❌ 不必要的类型断言
const name = (user as User).name; // 如果 user 已经是 User 类型

// ✅ 直接使用
const name = user.name;
```
