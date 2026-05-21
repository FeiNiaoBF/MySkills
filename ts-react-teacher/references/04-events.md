# 阶段 4：事件处理与表单类型安全

## React 合成事件类型速查表

```tsx
// 鼠标事件
onClick: (e: React.MouseEvent<HTMLButtonElement>) => void
onMouseEnter: (e: React.MouseEvent<HTMLDivElement>) => void

// 输入事件
onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void
onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void

// 键盘事件
onKeyDown: (e: React.KeyboardEvent<HTMLInputElement>) => void

// 表单事件
onSubmit: (e: React.FormEvent<HTMLFormElement>) => void

// 焦点事件
onFocus: (e: React.FocusEvent<HTMLInputElement>) => void

// 拖放事件
onDrop: (e: React.DragEvent<HTMLDivElement>) => void
```

## 完整表单示例

```tsx
import { useState, FormEvent, ChangeEvent } from 'react';

interface FormData {
  username: string;
  email: string;
  age: number;
}

const LoginForm = () => {
  const [form, setForm] = useState<FormData>({
    username: '',
    email: '',
    age: 0,
  });

  // 通用字段更新处理器（泛型 keyof 保证字段名安全）
  const handleChange = <K extends keyof FormData>(
    field: K,
    value: FormData[K]
  ) => {
    setForm((prev) => ({ ...prev, [field]: value }));
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // form 的类型在这里是完整的 FormData
    console.log(form.username); // ✅ 类型安全
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        value={form.username}
        onChange={(e: ChangeEvent<HTMLInputElement>) =>
          handleChange('username', e.target.value)
        }
      />
      <input
        type="number"
        value={form.age}
        onChange={(e) => handleChange('age', Number(e.target.value))}
      />
      <button type="submit">提交</button>
    </form>
  );
};
```

## 常见陷阱

### e.target vs e.currentTarget
```tsx
// e.target：触发事件的元素（可能是子元素）
// e.currentTarget：绑定事件的元素（更安全）

// 如果 button 里有 span，点击 span 时：
// e.target = span
// e.currentTarget = button（你绑定 onClick 的元素）

const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
  // ✅ 用 currentTarget 更安全
  console.log(e.currentTarget.id);
};
```
