# 教师助手设计规范

> 融合 Kinfolk 自然主义 + Notion 模块化 + Apple 毛玻璃效果

---

## 1. 设计哲学

### 核心原则
- **温暖不刺眼**：采用暖色调，长时间使用不疲劳
- **模块化布局**：可折叠、可重组的内容区块
- **轻盈通透**：毛玻璃效果营造层次感
- **年轻审美**：圆润、有机、有温度的视觉语言

### 参考风格
- **Kinfolk**：自然色系、衬线字体、留白美学
- **Notion**：模块化分区、可折叠块、侧边栏导航
- **Apple Design**：毛玻璃 (Glassmorphism)、柔和光晕、精致阴影

---

## 2. 色彩系统

### 主色调（自然暖色系）

| 名称 | 色值 | 用途 |
|------|------|------|
| **背景主色** | `#FDF8F3` | 页面主背景 |
| **背景次色** | `#F8F0E8` | 渐变背景、次要区域 |
| **卡片背景** | `rgba(255, 253, 250, 0.7)` | 玻璃态卡片 |
| **主强调色** | `#C4A882` | 按钮、选中态、图标 |
| **次强调色** | `#A8B89C` | 成功态、第二按钮 |
| ** Accent Coral** | `#E8A598` | 警告、高亮 |
| **文字主色** | `#2C2C2C` | 标题、正文 |
| **文字次色** | `#6B5B4F` | 次要文字、描述 |
| **文字辅助** | `#9B8B7B` | 占位符、时间戳 |

### 数据卡片配色

| 类型 | 色值 | 用途 |
|------|------|------|
| **学生数据** | `#E8C4A8` | 学生数量卡片 |
| **记录数据** | `#C5D4B8` | 表现记录卡片 |
| **评语数据** | `#F0D5C5` | 待评数量卡片 |
| **成绩数据** | `#E8D0E0` | 均分卡片 |

### 玻璃态配色

```css
--glass-bg: rgba(255, 253, 250, 0.65);
--glass-border: rgba(255, 255, 255, 0.6);
--glass-blur: blur(20px) saturate(180%);
```

### 光晕效果

```css
/* 暖色光晕 - 右上角 */
background: radial-gradient(circle, rgba(232, 196, 168, 0.25) 0%, transparent 70%);

/* 绿色光晕 - 左下角 */
background: radial-gradient(circle, rgba(197, 212, 184, 0.2) 0%, transparent 70%);

/* 珊瑚光晕 - 中右 */
background: radial-gradient(circle, rgba(240, 213, 197, 0.18) 0%, transparent 70%);
```

---

## 3. 字体系统

### 字体栈

```css
/* 标题字体 - 优雅衬线 */
font-family: 'Newsreader', Georgia, serif;

/* 正文字体 - 清晰无衬线 */
font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
```

### 字体层级

| 层级 | 字体 | 大小 | 字重 | 行高 | 用途 |
|------|------|------|------|------|------|
| **页面标题** | Newsreader | 32px | 400/500 | 1.3 | 页面主标题 |
| **区块标题** | Newsreader | 18px | 500 | 1.4 | 卡片标题 |
| **问候语** | Newsreader | 32px | 400 italic | 1.3 | "早上好" |
| **数据数值** | Newsreader | 32px | 600 | 1.2 | 统计数据 |
| **正文** | Inter | 14px | 400 | 1.6 | 描述文字 |
| **小字** | Inter | 12px | 400 | 1.5 | 辅助信息 |
| **标签** | Inter | 11-13px | 500 | 1.4 | 按钮、标签 |

### 字体样式

```css
/* 斜体问候语 */
.greeting {
  font-style: italic;
  font-family: 'Newsreader', serif;
}

/* 数字强调 */
.stat-value {
  font-family: 'Newsreader', serif;
  font-weight: 600;
  letter-spacing: -0.5px;
}
```

---

## 4. 间距系统

### 基础单位

| Token | 值 | 用途 |
|-------|-----|------|
| `xs` | 4px | 图标间距、紧凑内联 |
| `sm` | 8px | 小间隙、标签间距 |
| `md` | 12px | 卡片内边距、元素间距 |
| `lg` | 16px | 区块间距、卡片间隙 |
| `xl` | 20px | 大区块间距 |
| `2xl` | 24px | 页面边距 |
| `3xl` | 32px | 大区块边距 |

### 页面布局

```css
/* 侧边栏 */
.sidebar {
  width: 240px;
  padding: 24px 16px;
}

/* 主内容区 */
.main {
  margin-left: 240px;
  padding: 32px;
}

/* 卡片内边距 */
.card {
  padding: 20px;
  border-radius: 20px;
}
```

### 组件间距

| 组件 | 内边距 | 元素间距 |
|------|--------|----------|
| 数据卡片 | 20px 16px | - |
| 快捷操作按钮 | 16px 10px | 6px |
| 学生列表项 | 12px | 12px |
| 时间线条目 | 12px 0 | 4px |
| 待办项 | 10px | 12px |

---

## 5. 圆角系统

| Token | 值 | 用途 |
|-------|-----|------|
| `sm` | 8px | 小标签、复选框 |
| `md` | 12px | 按钮、输入框 |
| `lg` | 14px | 导航项、工具按钮 |
| `xl` | 16px | 中等卡片、头像 |
| `2xl` | 20px | 大卡片、内容区块 |
| `full` | 50% | 圆形头像、徽章 |

---

## 6. 阴影系统

### 卡片阴影

```css
/* 默认卡片 */
box-shadow: 0 8px 32px rgba(0, 0, 0, 0.05);

/* 悬停状态 */
box-shadow: 0 12px 40px rgba(0, 0, 0, 0.1);

/* 按钮阴影 */
box-shadow: 0 4px 16px rgba(196, 168, 130, 0.35);
```

---

## 7. 组件规范

### 7.1 玻璃态卡片 (Glass Card)

```css
.glass-card {
  background: rgba(255, 253, 250, 0.7);
  backdrop-filter: blur(20px) saturate(180%);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.6);
  padding: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.05);
}
```

### 7.2 主按钮

```css
.primary-button {
  padding: 12px 24px;
  border-radius: 14px;
  border: none;
  background: linear-gradient(135deg, #C4A882 0%, #A8B89C 100%);
  color: white;
  font-size: 14px;
  font-weight: 600;
  box-shadow: 0 4px 16px rgba(196, 168, 130, 0.35);
}
```

### 7.3 次按钮/玻璃按钮

```css
.glass-button {
  padding: 10px 18px;
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.6);
  background: rgba(255, 255, 255, 0.6);
  backdrop-filter: blur(10px);
  color: #2C2C2C;
  font-size: 14px;
}
```

### 7.4 数据卡片

```css
.stat-card {
  padding: 20px 16px;
  border-radius: 16px;
  text-align: center;
  background-color: /* 根据类型使用不同配色 */;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.12);
}
```

### 7.5 快捷操作卡片

```css
.action-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  padding: 16px 10px;
  border-radius: 14px;
  border: 2px solid /* 根据类型 */;
  background: rgba(255, 255, 255, 0.5);
}
```

### 7.6 侧边栏导航项

```css
.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
  border-radius: 12px;
  font-size: 14px;
  color: #2C2C2C;
}

.nav-item.active {
  background: rgba(196, 168, 130, 0.2);
}
```

### 7.7 学生头像

```css
.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #E8C4A8 0%, #C5D4B8 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 600;
}
```

### 7.8 标签

```css
.tag {
  font-size: 11-12px;
  padding: 2px 8px;
  background: rgba(196, 168, 130, 0.15);
  border-radius: 10px;
  color: #6B5B4F;
}
```

### 7.9 时间线

```css
.timeline-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #C4A882;
  box-shadow: 0 0 0 3px rgba(196, 168, 130, 0.2);
}

.timeline-connector {
  width: 2px;
  background: rgba(196, 168, 130, 0.2);
}
```

### 7.10 复选框

```css
.checkbox {
  width: 20px;
  height: 20px;
  border: 2px solid #D4C4B0;
  border-radius: 6px;
}

.checkbox.checked {
  background: #C4A882;
  border-color: #C4A882;
}
```

---

## 8. 布局模式

### 8.1 三栏布局（首页）

```
┌─────────────────────────────────────────────────────────┐
│  侧边栏  │     数据概览      │     最近动态     │   待办/日历  │
│  240px   │     (1fr)         │     (1fr)        │   320px      │
└─────────────────────────────────────────────────────────┘
```

### 8.2 两栏布局（学生管理、评语生成）

```
┌─────────────────────────────────────────────────────────┐
│  侧边栏  │           主内容区                             │
│  240px   │  列表/选择面板   │   详情/操作面板             │
└─────────────────────────────────────────────────────────┘
```

### 8.3 侧边栏布局

```
┌─────────┬───────────────────────────────────────────────┐
│ Logo    │                                               │
│ 导航项   │              主内容区                          │
│ 导航项   │                                               │
│ 导航项   │                                               │
│         │                                               │
│ 用户信息 │                                               │
└─────────┴───────────────────────────────────────────────┘
```

---

## 9. 动效规范

### 9.1 过渡时间

| 类型 | 时长 | 用途 |
|------|------|------|
| 快速 | 0.15s | 按钮悬停、小状态变化 |
| 正常 | 0.2s | 卡片悬停、展开折叠 |
| 缓慢 | 0.3s | 页面切换、大型过渡 |

### 9.2 缓动函数

```css
/* 默认 */
transition: all 0.2s ease;

/* 卡片悬停 - 弹性 */
transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
```

### 9.3 悬停效果

```css
/* 数据卡片上浮 */
.card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.12);
}

/* 按钮缩放 */
.button:hover {
  transform: scale(1.02);
}
```

---

## 10. 响应式断点

| 断点 | 宽度 | 调整 |
|------|------|------|
| Desktop | >1280px | 三栏布局 |
| Laptop | 1024-1280px | 两栏布局 |
| Tablet | 768-1024px | 侧边栏收起 |
| Mobile | <768px | 单栏堆叠 |

---

## 11. 图标使用

### 图标风格
- 使用 Emoji 作为主要图标（温暖、亲和）
- 或使用 Lucide Icons / Element Plus Icons（简洁线条）

### 常用图标映射

| 功能 | Emoji | 备选 |
|------|-------|------|
| 首页 | 🏠 | Home |
| 学生 | 👤 | User |
| 成绩 | 📊 | BarChart |
| 评语 | 📝 | FileText |
| 记录 | ⭐ | Star |
| 待办 | ✅ | CheckCircle |
| 日历 | 📅 | Calendar |
| 通知 | 🔔 | Bell |
| 添加 | + | Plus |
| 搜索 | 🔍 | Search |

---

## 12. 文件组织

```
docs/
├── design-system.md          # 本文件
└── assets/                   # 设计资源
    ├── color-palette.png
    ├── typography.png
    └── components.png
```

---

## 13. 实施检查清单

### 颜色
- [ ] 背景使用 `#FDF8F3`
- [ ] 所有卡片使用玻璃态效果
- [ ] 主按钮使用渐变色 `#C4A882` → `#A8B89C`
- [ ] 数据卡片使用对应配色

### 字体
- [ ] 标题使用 Newsreader
- [ ] 正文使用 Inter
- [ ] 问候语使用斜体

### 组件
- [ ] 卡片圆角 20px
- [ ] 按钮圆角 12-14px
- [ ] 头像使用渐变背景
- [ ] 时间线使用主色点缀

### 布局
- [ ] 侧边栏固定 240px
- [ ] 主内容区 margin-left: 240px
- [ ] 使用三栏/两栏布局避免右侧留白
- [ ] 添加背景光晕效果

---

*Version 1.0 | Created for 教师助手 Project*
