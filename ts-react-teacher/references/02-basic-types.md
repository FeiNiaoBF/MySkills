# 阶段 2：React 基础类型

## 函数组件类型

### React.FC 与直接类型注解
`React.FC`（别名 `React.FunctionComponent`）是历史遗留写法，现代 React + TS 项目更推荐**直接注解返回类型**，原因如下：

- `React.FC` 在 React 18 之前隐式包含 `children?: ReactNode`，导致类型不精确
- 直接注解更清晰，IDE 提示更准确

```tsx
import React from 'react';

// ❌ 旧写法（React 17 及以前常见）
const OldButton: React.FC<{ label: string }> = ({ label }) => {
  return <button>{label}</button>;
};

// ✅ 现代推荐写法
interface ButtonProps {
  label: string;
  onClick?: () => void;
}

const Button = ({ label, onClick }: ButtonProps): React.ReactElement => {
  return <button onClick={onClick}>{label}</button>;
};
```

## Props 类型定义

### 基础 Props
```tsx
interface UserCardProps {
  name: string;           // 必填
  age?: number;           // 可选（number | undefined）
  email: string | null;   // 必填但可为 null（与可选不同！）
}
```

### 带 children 的 Props
```tsx
import { ReactNode, PropsWithChildren } from 'react';

// 方式 1：手动声明
interface CardProps {
  title: string;
  children: ReactNode; // 推荐，比 JSX.Element 更宽泛
}

// 方式 2：用工具类型
type CardProps2 = PropsWithChildren<{ title: string }>;

const Card = ({ title, children }: CardProps) => (
  <div>
    <h2>{title}</h2>
    {children}
  </div>
);
```

## 核心类型区别

| 类型 | 含义 | 使用场景 |
|------|------|----------|
| `ReactNode` | 任何可渲染内容（string、number、JSX、null、数组…） | `children` prop 的类型 |
| `ReactElement` | 必须是 JSX 元素（不含 null/string） | 组件返回类型（更精确） |
| `JSX.Element` | 等同于 `ReactElement<any, any>`，较宽松 | 少用，优先 ReactElement |
| `CSSProperties` | React 的内联样式对象类型 | `style` prop |

```tsx
import { CSSProperties } from 'react';

interface BoxProps {
  style?: CSSProperties; // 比 object 精确，比手写所有属性省事
  className?: string;
}
```
