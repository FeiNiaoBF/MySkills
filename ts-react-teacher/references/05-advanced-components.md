# 阶段 5：高级组件模式

## 泛型组件

### 为什么需要泛型组件

假设你要写一个通用列表组件。没有泛型时：

```tsx
// ❌ 使用 any - 完全失去类型安全
interface ListProps {
  items: any[];
  renderItem: (item: any) => React.ReactNode;
}
```

有了泛型：

```tsx
// ✅ 类型安全的泛型列表
interface ListProps<T> {
  items: T[];
  renderItem: (item: T, index: number) => React.ReactNode;
  keyExtractor: (item: T) => string;
}

// 注意 .tsx 文件中的 <T,> 逗号！防止 TS 把 <T> 解析为 JSX 标签
function List<T,>({ items, renderItem, keyExtractor }: ListProps<T>) {
  return (
    <ul>
      {items.map((item, index) => (
        <li key={keyExtractor(item)}>{renderItem(item, index)}</li>
      ))}
    </ul>
  );
}

// 使用时 TS 自动推断 T = User
interface User { id: number; name: string }
const users: User[] = [{ id: 1, name: 'Alice' }];

<List
  items={users}
  keyExtractor={(user) => String(user.id)}  // user 自动是 User 类型
  renderItem={(user) => <span>{user.name}</span>}
/>
```

## ComponentProps 继承原生属性

```tsx
import { ComponentProps } from 'react';

// 继承所有原生 button 属性（type、disabled、aria-*…）
// 然后只扩展你自己的 props
type ButtonProps = ComponentProps<'button'> & {
  variant?: 'primary' | 'secondary';
  loading?: boolean;
};

const Button = ({ variant = 'primary', loading, children, ...rest }: ButtonProps) => {
  return (
    // ...rest 会把所有原生 button 属性（onClick、type 等）透传
    <button {...rest} disabled={loading || rest.disabled}>
      {loading ? '加载中...' : children}
    </button>
  );
};

// 使用时，所有原生 button 属性都有类型提示！
<Button variant="primary" type="submit" aria-label="提交表单">
  提交
</Button>
```

## 多态组件（Polymorphic Components）

多态组件允许调用方指定渲染为哪个 HTML 元素或组件：

```tsx
// 步骤 1：定义多态 Props 类型
type AsComponent = keyof JSX.IntrinsicElements | React.ComponentType<any>;

type PolymorphicProps<C extends AsComponent> = {
  as?: C;
  children?: React.ReactNode;
} & React.ComponentPropsWithoutRef<C>;

// 步骤 2：实现多态组件
function Text<C extends AsComponent = 'p'>({
  as,
  children,
  ...rest
}: PolymorphicProps<C>) {
  const Component = (as ?? 'p') as AsComponent;
  return <Component {...rest}>{children}</Component>;
}

// 步骤 3：使用 - 类型会随 as 变化！
<Text>默认是 p 标签</Text>
<Text as="h1" id="title">变成 h1 了</Text>
<Text as="a" href="/about">变成 a 标签，自动要求 href</Text>
// ✅ href 是 a 标签特有的，TS 会自动提示且校验
```
