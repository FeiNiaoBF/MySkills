# 阶段 3：核心 Hooks 类型

## useState

### 自动推断（大多数情况够用）
```tsx
// TS 从初始值推断类型，无需显式声明
const [count, setCount] = useState(0);        // number
const [name, setName] = useState('');         // string
const [flag, setFlag] = useState(false);      // boolean
```

### 显式泛型（初始值为 null 或复杂类型时必须）
```tsx
interface User {
  id: number;
  name: string;
}

// ✅ 明确告诉 TS 这个值会是 User，但初始为 null
const [user, setUser] = useState<User | null>(null);

// ✅ 数组类型
const [items, setItems] = useState<string[]>([]);

// ❌ 这样 TS 推断为 never[]，之后 push 会报错
const [items2, setItems2] = useState([]);
```

## useRef

`useRef` 有两种完全不同的用途，类型也不同：

```tsx
import { useRef } from 'react';

// 用途 1：引用 DOM 元素（只读，React 管理）
// 初始值传 null，泛型传元素类型
const inputRef = useRef<HTMLInputElement>(null);
// inputRef.current 的类型是 HTMLInputElement | null
// 使用前需要判空：if (inputRef.current) { ... }

// 用途 2：保存可变值（不触发重渲染）
// 初始值非 null，类型自动包含 undefined
const timerRef = useRef<ReturnType<typeof setTimeout>>();
// 或者
const counterRef = useRef(0); // MutableRefObject<number>
```

## useContext

```tsx
import { createContext, useContext } from 'react';

interface ThemeContextType {
  theme: 'light' | 'dark';
  toggleTheme: () => void;
}

// 推荐：提供默认值而非 undefined，避免使用时判空
const ThemeContext = createContext<ThemeContextType>({
  theme: 'light',
  toggleTheme: () => {},
});

// 自定义 Hook 封装（最佳实践）
export const useTheme = () => {
  const ctx = useContext(ThemeContext);
  if (!ctx) throw new Error('useTheme 必须在 ThemeProvider 内使用');
  return ctx;
};
```

## useReducer

Action 联合类型设计是 `useReducer` 的精华：

```tsx
// 用"可辨识联合"(discriminated union) 设计 Action
type Action =
  | { type: 'INCREMENT' }
  | { type: 'DECREMENT' }
  | { type: 'SET'; payload: number }; // 只有 SET 有 payload

interface State {
  count: number;
  loading: boolean;
}

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case 'INCREMENT':
      return { ...state, count: state.count + 1 };
    case 'DECREMENT':
      return { ...state, count: state.count - 1 };
    case 'SET':
      // TS 在这里自动知道 action.payload 是 number
      return { ...state, count: action.payload };
  }
}

// 使用
const [state, dispatch] = useReducer(reducer, { count: 0, loading: false });
dispatch({ type: 'SET', payload: 10 }); // ✅ 类型安全
dispatch({ type: 'INCREMENT', payload: 1 }); // ❌ TS 报错：INCREMENT 没有 payload
```

## useCallback / useMemo 泛型推导

这两个 Hooks 通常能自动推断，但在复杂情况下可以显式标注：

```tsx
// useCallback - 通常自动推断
const handleChange = useCallback(
  (e: React.ChangeEvent<HTMLInputElement>) => {
    setValue(e.target.value);
  },
  [setValue] // 依赖数组！忘记会导致 stale closure
);

// useMemo - 自动推断返回类型
const filteredItems = useMemo(
  () => items.filter((item) => item.active),
  [items]
);
// filteredItems 的类型 = typeof items[number][] （自动推断）
```
