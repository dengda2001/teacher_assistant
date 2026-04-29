# 教师助手 UI 重构实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 按照设计规范重构教师助手应用的全部页面，实现截图中的统一视觉风格

**Architecture:** 基于现有 Vue 3 + Element Plus + SCSS 项目，通过更新样式变量和重构组件模板来实现新设计。保持现有逻辑不变，只修改视觉表现。

**Tech Stack:** Vue 3, Element Plus, SCSS, Vue Router, Pinia

---

## 文件结构

| 文件 | 职责 |
|------|------|
| `src/styles/variables.scss` | 颜色、间距、圆角等设计令牌 |
| `src/styles/global.scss` | 全局样式覆盖和工具类 |
| `src/components/Layout.vue` | 侧边栏 + 顶部栏布局 |
| `src/views/Dashboard.vue` | 仪表盘首页 |
| `src/views/students/Index.vue` | 学生管理页面 |
| `src/views/performance/Index.vue` | 表现记录页面 |
| `src/views/comments/Index.vue` | 评语生成页面 |
| `src/views/scores/Index.vue` | 成绩管理页面 |
| `src/views/scores/Analysis.vue` | 成绩分析页面 |

---

## 前置检查

- [ ] **确认项目可运行**

```bash
cd /Users/dd/projects/teacher_assistant/frontend
npm install
npm run dev
```

访问 http://localhost:5173 确认页面正常显示

---

## Task 1: 更新样式变量文件

**Files:**
- Modify: `src/styles/variables.scss`

- [ ] **Step 1: 备份原文件**

```bash
cp src/styles/variables.scss src/styles/variables.scss.backup
```

- [ ] **Step 2: 重写变量文件**

```scss
// 主色调
$primary: #6366f1;
$primary-light: #818cf8;
$primary-lighter: rgba(99, 102, 241, 0.08);
$primary-dark: #4f46e5;

// 中性色
$bg-primary: #f8f9fa;
$bg-card: #ffffff;
$bg-hover: #f9fafb;
$bg-input: #f3f4f6;

$text-primary: #1f2937;
$text-secondary: #6b7280;
$text-muted: #9ca3af;

$border: #e5e7eb;
$border-light: #f3f4f6;

// 功能色
$success: #10b981;
$success-light: #d1fae5;
$success-text: #059669;

$warning: #f59e0b;
$warning-light: #fef3c7;
$warning-text: #d97706;

$danger: #ef4444;
$danger-light: #fee2e2;
$danger-text: #dc2626;

$info: #3b82f6;
$info-light: #dbeafe;
$info-text: #2563eb;

$pink: #ec4899;
$pink-light: #fce7f3;
$pink-text: #db2777;

// 学科标签色
$subject-chinese-bg: #fee2e2;
$subject-chinese-text: #dc2626;
$subject-math-bg: #dbeafe;
$subject-math-text: #2563eb;
$subject-english-bg: #d1fae5;
$subject-english-text: #059669;
$subject-physics-bg: #e0e7ff;
$subject-physics-text: #4f46e5;
$subject-chemistry-bg: #fef3c7;
$subject-chemistry-text: #d97706;

// 间距
$spacing-xs: 4px;
$spacing-sm: 8px;
$spacing-md: 12px;
$spacing-lg: 16px;
$spacing-xl: 24px;
$spacing-2xl: 32px;

// 圆角
$radius-sm: 6px;
$radius-md: 8px;
$radius-lg: 12px;
$radius-xl: 16px;
$radius-full: 9999px;

// 阴影
$shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
$shadow-md: 0 1px 3px rgba(0, 0, 0, 0.1);
$shadow-lg: 0 4px 6px rgba(0, 0, 0, 0.1);

// 过渡
$transition-fast: 0.15s ease;
$transition-base: 0.2s ease;
```

- [ ] **Step 3: 创建全局样式文件**

Create: `src/styles/global.scss`

```scss
@use 'variables' as *;

// 页面标题
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-xl;

  .page-title {
    font-size: 24px;
    font-weight: 600;
    color: $text-primary;
    margin: 0;

    .subtitle {
      font-size: 14px;
      color: $text-secondary;
      font-weight: 400;
      margin-left: $spacing-md;
    }
  }

  .actions {
    display: flex;
    gap: $spacing-md;
  }
}

// 统计卡片
.stat-card {
  background: $bg-card;
  border-radius: $radius-lg;
  padding: $spacing-xl;
  border: 1px solid $border;
  box-shadow: $shadow-sm;

  .stat-value {
    font-size: 32px;
    font-weight: 700;
    color: $text-primary;
    line-height: 1;
  }

  .stat-label {
    font-size: 14px;
    color: $text-secondary;
    margin-top: $spacing-sm;
  }

  .stat-change {
    font-size: 12px;
    margin-top: $spacing-xs;

    &.up {
      color: $success;
    }

    &.down {
      color: $danger;
    }
  }
}

// 搜索框样式
.search-input {
  :deep(.el-input__wrapper) {
    background: $bg-input;
    border-radius: $radius-full;
    box-shadow: none;
    padding: 0 $spacing-lg;

    .el-input__inner {
      height: 40px;
    }
  }
}

// 学科标签
.subject-tag {
  display: inline-flex;
  align-items: center;
  padding: 2px 8px;
  border-radius: $radius-sm;
  font-size: 12px;
  font-weight: 500;
  margin-right: $spacing-xs;

  &.chinese {
    background: $subject-chinese-bg;
    color: $subject-chinese-text;
  }

  &.math {
    background: $subject-math-bg;
    color: $subject-math-text;
  }

  &.english {
    background: $subject-english-bg;
    color: $subject-english-text;
  }

  &.physics {
    background: $subject-physics-bg;
    color: $subject-physics-text;
  }

  &.chemistry {
    background: $subject-chemistry-bg;
    color: $subject-chemistry-text;
  }
}

// 性别标签
.gender-tag {
  display: inline-flex;
  align-items: center;
  padding: 2px 10px;
  border-radius: $radius-sm;
  font-size: 12px;
  font-weight: 500;

  &.male {
    background: $info-light;
    color: $info-text;
  }

  &.female {
    background: $pink-light;
    color: $pink-text;
  }
}

// 状态标签
.status-tag {
  display: inline-flex;
  align-items: center;
  padding: 4px 10px;
  border-radius: $radius-sm;
  font-size: 12px;
  font-weight: 500;

  &.success {
    background: $success-light;
    color: $success-text;
  }

  &.warning {
    background: $warning-light;
    color: $warning-text;
  }

  &.danger {
    background: $danger-light;
    color: $danger-text;
  }
}

// 记录卡片
.record-card {
  background: $bg-card;
  border-radius: $radius-lg;
  padding: $spacing-lg;
  border: 1px solid $border;
  margin-bottom: $spacing-lg;

  .record-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: $spacing-md;
  }

  .student-info {
    display: flex;
    align-items: center;
    gap: $spacing-md;

    .student-name {
      font-weight: 600;
      font-size: 15px;
      color: $text-primary;
    }

    .record-time {
      font-size: 13px;
      color: $text-muted;
    }
  }

  .record-content {
    color: $text-secondary;
    line-height: 1.6;
    margin-bottom: $spacing-md;
    font-size: 14px;
  }

  .record-tags {
    display: flex;
    flex-wrap: wrap;
    gap: $spacing-sm;
  }
}

// 快捷标签
.quick-tag {
  cursor: pointer;
  transition: all $transition-fast;

  &:hover {
    background: $primary !important;
    color: white !important;
    border-color: $primary !important;
  }
}
```

- [ ] **Step 4: 在 main.ts 中引入全局样式**

Modify: `src/main.ts`

```typescript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'

// 引入自定义样式
import './styles/variables.scss'
import './styles/global.scss'

const app = createApp(App)

// 注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(createPinia())
app.use(router)
app.use(ElementPlus)

app.mount('#app')
```

- [ ] **Step 5: 测试样式加载**

```bash
npm run dev
```

确认页面无样式错误，控制台无报错

- [ ] **Step 6: Commit**

```bash
git add src/styles/ src/main.ts
git commit -m "style: add design system variables and global styles"
```

---

## Task 2: 重构 Layout 组件

**Files:**
- Modify: `src/components/Layout.vue`

- [ ] **Step 1: 重写 Layout.vue**

```vue
<template>
  <el-container class="layout">
    <!-- 侧边栏 -->
    <el-aside width="240px" class="sidebar">
      <div class="logo">
        <div class="logo-icon">
          <el-icon size="24" color="#6366f1"><School /></el-icon>
        </div>
        <div class="logo-text">
          <span class="title">教师助手</span>
          <span class="subtitle">智能教学管理</span>
        </div>
      </div>

      <el-menu
        :default-active="$route.path"
        router
        class="nav-menu"
      >
        <el-menu-item v-for="route in routes" :key="route.path" :index="route.path">
          <el-icon>
            <component :is="route.meta?.icon" />
          </el-icon>
          <span>{{ route.meta?.title }}</span>
        </el-menu-item>
      </el-menu>

      <div class="sidebar-footer">
        <el-menu-item index="/settings" class="settings-item">
          <el-icon><Setting /></el-icon>
          <span>设置</span>
        </el-menu-item>
      </div>
    </el-aside>

    <!-- 主内容区 -->
    <el-container class="main-container">
      <el-header class="header">
        <div class="search-box">
          <el-icon class="search-icon"><Search /></el-icon>
          <input type="text" placeholder="搜索学生、记录..." class="search-input-header" />
        </div>
        <div class="header-actions">
          <el-button class="icon-btn" text>
            <el-icon size="18"><Bell /></el-icon>
          </el-button>
          <div class="user-info">
            <el-avatar :size="36" class="user-avatar">张</el-avatar>
            <div class="user-meta">
              <span class="user-name">张老师</span>
              <span class="user-role">高三(1)班班主任</span>
            </div>
          </div>
        </div>
      </el-header>

      <el-main class="main-content">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import router from '@/router'
import { School, Setting, Search, Bell } from '@element-plus/icons-vue'

const $route = useRoute()

const routes = computed(() => {
  return router.getRoutes().filter(route => route.meta?.title && route.path !== '/settings')
})
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.layout {
  height: 100vh;
  background: $bg-primary;
}

.sidebar {
  background: $bg-card;
  border-right: 1px solid $border;
  display: flex;
  flex-direction: column;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.04);
}

.logo {
  height: 80px;
  display: flex;
  align-items: center;
  padding: 0 $spacing-xl;
  gap: $spacing-md;
  border-bottom: 1px solid $border-light;

  .logo-icon {
    width: 40px;
    height: 40px;
    background: $primary-lighter;
    border-radius: $radius-md;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .logo-text {
    display: flex;
    flex-direction: column;

    .title {
      font-size: 18px;
      font-weight: 600;
      color: $text-primary;
      letter-spacing: -0.3px;
      line-height: 1.3;
    }

    .subtitle {
      font-size: 12px;
      color: $text-muted;
      line-height: 1.3;
    }
  }
}

.nav-menu {
  flex: 1;
  padding: $spacing-md 0;
  border-right: none;

  :deep(.el-menu-item) {
    height: 48px;
    line-height: 48px;
    margin: $spacing-xs $spacing-lg;
    border-radius: $radius-md;
    color: $text-secondary;
    font-size: 14px;

    .el-icon {
      margin-right: $spacing-md;
      font-size: 18px;
    }

    &:hover {
      background: $bg-hover;
      color: $text-primary;
    }

    &.is-active {
      background: $primary-lighter;
      color: $primary;
      font-weight: 500;

      &::before {
        content: '';
        position: absolute;
        left: -$spacing-lg;
        top: 50%;
        transform: translateY(-50%);
        width: 3px;
        height: 20px;
        background: $primary;
        border-radius: 0 2px 2px 0;
      }
    }
  }
}

.sidebar-footer {
  padding: $spacing-md 0;
  border-top: 1px solid $border-light;

  .settings-item {
    height: 48px;
    line-height: 48px;
    margin: 0 $spacing-lg;
    border-radius: $radius-md;
    color: $text-secondary;

    &:hover {
      background: $bg-hover;
      color: $text-primary;
    }
  }
}

.main-container {
  background: $bg-primary;
}

.header {
  height: 64px;
  background: $bg-card;
  border-bottom: 1px solid $border;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 $spacing-2xl;
}

.search-box {
  position: relative;
  width: 320px;

  .search-icon {
    position: absolute;
    left: $spacing-md;
    top: 50%;
    transform: translateY(-50%);
    color: $text-muted;
    font-size: 16px;
    z-index: 1;
  }

  .search-input-header {
    width: 100%;
    height: 40px;
    background: $bg-input;
    border: none;
    border-radius: $radius-full;
    padding: 0 $spacing-lg 0 40px;
    font-size: 14px;
    color: $text-primary;
    outline: none;
    transition: all $transition-fast;

    &::placeholder {
      color: $text-muted;
    }

    &:focus {
      background: $bg-card;
      box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
    }
  }
}

.header-actions {
  display: flex;
  align-items: center;
  gap: $spacing-lg;

  .icon-btn {
    width: 36px;
    height: 36px;
    border-radius: $radius-md;
    color: $text-secondary;

    &:hover {
      background: $bg-hover;
      color: $text-primary;
    }
  }
}

.user-info {
  display: flex;
  align-items: center;
  gap: $spacing-md;

  .user-avatar {
    background: linear-gradient(135deg, $primary 0%, $primary-light 100%);
    color: white;
    font-weight: 600;
  }

  .user-meta {
    display: flex;
    flex-direction: column;

    .user-name {
      font-size: 14px;
      font-weight: 600;
      color: $text-primary;
      line-height: 1.3;
    }

    .user-role {
      font-size: 12px;
      color: $text-muted;
      line-height: 1.3;
    }
  }
}

.main-content {
  padding: $spacing-2xl;
  overflow-y: auto;
}

// 过渡动画
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
```

- [ ] **Step 2: 验证布局显示**

确认：
- 侧边栏显示正常，有 logo 和菜单
- 选中菜单项有紫色左边框
- 顶部栏有搜索框和用户信息
- 无样式错误

- [ ] **Step 3: Commit**

```bash
git add src/components/Layout.vue
git commit -m "style: redesign layout component with new design system"
```

---

## Task 3: 重构学生管理页面

**Files:**
- Modify: `src/views/students/Index.vue`

- [ ] **Step 1: 重写学生管理页面**

```vue
<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">
          学生管理
          <span class="subtitle">管理班级学生信息</span>
        </h1>
      </div>
      <el-button type="primary" class="btn-primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        <span>添加学生</span>
      </el-button>
    </div>

    <el-card class="content-card" shadow="never">
      <div class="table-toolbar">
        <div class="search-box-table">
          <el-icon class="search-icon"><Search /></el-icon>
          <el-input
            v-model="searchQuery"
            placeholder="搜索学生姓名或学号..."
            class="search-input-clean"
            clearable
          />
        </div>
        <span class="student-count">共 {{ filteredStudents.length }} 名学生</span>
      </div>

      <el-table :data="filteredStudents" v-loading="loading" style="width: 100%">
        <el-table-column min-width="120">
          <template #default="{ row }">
            <div class="student-name-cell">
              <div class="avatar-circle">{{ row.name.charAt(0) }}</div>
              <span class="name-text">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="grade" label="年级" min-width="100">
          <template #default="{ row }">
            初中{{ row.grade }}
          </template>
        </el-table-column>
        <el-table-column label="性别" min-width="100">
          <template #default="{ row }">
            <span :class="['gender-tag', row.gender === '男' ? 'male' : 'female']">
              {{ row.gender }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <div class="action-btns">
              <el-button type="primary" link @click="editStudent(row)">
                <el-icon><Edit /></el-icon>
              </el-button>
              <el-button type="danger" link @click="deleteStudent(row)">
                <el-icon><Delete /></el-icon>
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑学生弹窗 -->
    <el-dialog
      v-model="showAddDialog"
      :title="isEdit ? '编辑学生' : '添加新学生'"
      width="500px"
      class="custom-dialog"
      destroy-on-close
    >
      <p class="dialog-subtitle">请填写学生的基本信息</p>
      <el-form :model="studentForm" label-width="80px" :rules="rules" ref="formRef" class="custom-form">
        <el-form-item label="姓名" prop="name">
          <el-input v-model="studentForm.name" placeholder="请输入学生姓名" />
        </el-form-item>
        <el-form-item label="学号" prop="studentNo">
          <el-input v-model="studentForm.studentNo" placeholder="请输入学号" />
        </el-form-item>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="班级" prop="class">
              <el-input v-model="studentForm.class" placeholder="例如：1班" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="年级" prop="grade">
              <el-select v-model="studentForm.grade" placeholder="请选择" style="width: 100%">
                <el-option label="初一" :value="7" />
                <el-option label="初二" :value="8" />
                <el-option label="初三" :value="9" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="性别" prop="gender">
          <el-radio-group v-model="studentForm.gender">
            <el-radio label="男">男</el-radio>
            <el-radio label="女">女</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showAddDialog = false" class="btn-secondary">取消</el-button>
          <el-button type="primary" @click="saveStudent" class="btn-primary">
            {{ isEdit ? '保存' : '确认添加' }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Edit, Delete } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'

interface Student {
  id?: string
  name: string
  studentNo: string
  grade: number
  class: string
  gender: string
}

const loading = ref(false)
const searchQuery = ref('')
const showAddDialog = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

const studentForm = reactive<Student>({
  name: '',
  studentNo: '',
  grade: 7,
  class: '',
  gender: '男'
})

const students = ref<Student[]>([
  { id: '1', name: '王小明', studentNo: '2024001', grade: 9, class: '1班', gender: '男' },
  { id: '2', name: '李小红', studentNo: '2024002', grade: 9, class: '1班', gender: '女' },
  { id: '3', name: '张三', studentNo: '2024003', grade: 9, class: '1班', gender: '男' },
  { id: '4', name: '赵四', studentNo: '2024004', grade: 9, class: '1班', gender: '男' },
  { id: '5', name: '刘五', studentNo: '2024005', grade: 9, class: '1班', gender: '女' },
  { id: '6', name: '陈六', studentNo: '2024006', grade: 9, class: '1班', gender: '男' },
  { id: '7', name: '周七', studentNo: '2024007', grade: 9, class: '1班', gender: '女' },
  { id: '8', name: '吴八', studentNo: '2024008', grade: 9, class: '1班', gender: '男' }
])

const filteredStudents = computed(() => {
  if (!searchQuery.value) return students.value
  const query = searchQuery.value.toLowerCase()
  return students.value.filter(s =>
    s.name.toLowerCase().includes(query) ||
    s.studentNo.includes(query)
  )
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  studentNo: [{ required: true, message: '请输入学号', trigger: 'blur' }],
  grade: [{ required: true, message: '请选择年级', trigger: 'change' }],
  class: [{ required: true, message: '请输入班级', trigger: 'blur' }],
  gender: [{ required: true, message: '请选择性别', trigger: 'change' }]
}

const saveStudent = async () => {
  if (!formRef.value) return
  await formRef.value.validate((valid) => {
    if (valid) {
      if (isEdit.value) {
        const index = students.value.findIndex(s => s.id === studentForm.id)
        if (index !== -1) {
          students.value[index] = { ...studentForm }
        }
        ElMessage.success('更新成功')
      } else {
        students.value.push({
          ...studentForm,
          id: Date.now().toString()
        })
        ElMessage.success('添加成功')
      }
      showAddDialog.value = false
      resetForm()
    }
  })
}

const editStudent = (student: Student) => {
  isEdit.value = true
  Object.assign(studentForm, student)
  showAddDialog.value = true
}

const deleteStudent = (student: Student) => {
  ElMessageBox.confirm(
    `确定要删除学生 "${student.name}" 吗？`,
    '确认删除',
    { type: 'warning' }
  ).then(() => {
    students.value = students.value.filter(s => s.id !== student.id)
    ElMessage.success('删除成功')
  })
}

const resetForm = () => {
  isEdit.value = false
  Object.assign(studentForm, {
    id: undefined,
    name: '',
    studentNo: '',
    grade: 7,
    class: '',
    gender: '男'
  })
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.content-card {
  border-radius: $radius-lg;
  border: 1px solid $border;

  :deep(.el-card__body) {
    padding: $spacing-xl;
  }
}

.table-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-lg;
}

.search-box-table {
  position: relative;
  width: 280px;

  .search-icon {
    position: absolute;
    left: $spacing-md;
    top: 50%;
    transform: translateY(-50%);
    color: $text-muted;
    font-size: 16px;
    z-index: 1;
  }

  .search-input-clean {
    :deep(.el-input__wrapper) {
      background: $bg-input;
      border-radius: $radius-md;
      box-shadow: none;
      padding-left: 36px;

      .el-input__inner {
        height: 36px;
        font-size: 14px;

        &::placeholder {
          color: $text-muted;
        }
      }

      &.is-focus {
        box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
      }
    }
  }
}

.student-count {
  font-size: 14px;
  color: $text-secondary;
}

.student-name-cell {
  display: flex;
  align-items: center;
  gap: $spacing-md;

  .avatar-circle {
    width: 36px;
    height: 36px;
    border-radius: $radius-full;
    background: $primary-lighter;
    color: $primary;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    font-weight: 600;
  }

  .name-text {
    font-size: 14px;
    color: $text-primary;
    font-weight: 500;
  }
}

.action-btns {
  display: flex;
  gap: $spacing-xs;

  .el-button {
    font-size: 16px;
  }
}

// 按钮样式
.btn-primary {
  background: $primary;
  border-color: $primary;
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;

  &:hover {
    background: $primary-light;
    border-color: $primary-light;
  }
}

.btn-secondary {
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;
}

// 弹窗样式
:deep(.custom-dialog) {
  .el-dialog__header {
    padding: $spacing-xl $spacing-xl $spacing-sm;
    margin: 0;

    .el-dialog__title {
      font-size: 18px;
      font-weight: 600;
      color: $text-primary;
    }
  }

  .el-dialog__body {
    padding: $spacing-md $spacing-xl;
  }

  .dialog-subtitle {
    font-size: 14px;
    color: $text-secondary;
    margin: 0 0 $spacing-lg;
  }
}

.custom-form {
  .el-form-item {
    margin-bottom: $spacing-lg;
  }

  .el-input__wrapper,
  .el-select .el-input__wrapper {
    border-radius: $radius-md;
    box-shadow: 0 0 0 1px $border;

    &.is-focus {
      box-shadow: 0 0 0 1px $primary;
    }
  }

  .el-radio__input.is-checked .el-radio__inner {
    background: $primary;
    border-color: $primary;
  }

  .el-radio__input.is-checked + .el-radio__label {
    color: $primary;
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: $spacing-md;
  padding: $spacing-md 0 0;
}
</style>
```

- [ ] **Step 2: 验证页面显示**

确认：
- 页面标题有副标题
- 搜索框样式正确
- 表格有头像和性别标签
- 添加弹窗样式符合设计规范

- [ ] **Step 3: Commit**

```bash
git add src/views/students/Index.vue
git commit -m "style: redesign students page with new design system"
```

---

## Task 4: 重构表现记录页面

**Files:**
- Modify: `src/views/performance/Index.vue`

- [ ] **Step 1: 重写表现记录页面**

```vue
<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">
          表现记录
          <span class="subtitle">记录学生日常表现和特殊成就</span>
        </h1>
      </div>
      <el-button type="primary" class="btn-primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        <span>添加记录</span>
      </el-button>
    </div>

    <div class="search-filter-bar">
      <div class="search-box-table">
        <el-icon class="search-icon"><Search /></el-icon>
        <el-input
          v-model="searchQuery"
          placeholder="搜索学生、内容或标签..."
          class="search-input-clean"
          clearable
        />
      </div>
    </div>

    <el-row :gutter="24">
      <el-col :span="16">
        <div class="records-list">
          <div v-for="record in filteredRecords" :key="record.id" class="record-card">
            <div class="record-header">
              <div class="student-info">
                <div class="avatar-circle" :class="record.type">{{ record.studentName.charAt(0) }}</div>
                <div class="info-text">
                  <span class="student-name">{{ record.studentName }}</span>
                  <el-tag size="small" :type="record.type" effect="light" class="type-tag">
                    {{ getTypeLabel(record.type) }}
                  </el-tag>
                </div>
              </div>
              <span class="record-time">{{ record.recordDate }} {{ record.time }}</span>
            </div>
            <div class="record-content">{{ record.content }}</div>
            <div class="record-footer">
              <div class="record-tags">
                <span v-for="tag in record.tags" :key="tag" class="tag-item">{{ tag }}</span>
              </div>
              <div class="record-actions">
                <el-button type="primary" link @click="editRecord(record)">
                  <el-icon><View /></el-icon>
                </el-button>
                <el-button type="primary" link @click="editRecord(record)">
                  <el-icon><Edit /></el-icon>
                </el-button>
                <el-button type="danger" link @click="deleteRecord(record)">
                  <el-icon><Delete /></el-icon>
                </el-button>
              </div>
            </div>
          </div>

          <el-empty v-if="filteredRecords.length === 0" description="暂无记录" />
        </div>
      </el-col>

      <el-col :span="8">
        <el-card class="side-card" shadow="never">
          <template #header>
            <div class="card-header-title">快捷标签</div>
          </template>
          <p class="card-desc">点击标签快速筛选</p>
          <div class="quick-tags">
            <span
              v-for="tag in quickTags"
              :key="tag"
              class="quick-tag-item"
              @click="addTagFilter(tag)"
            >
              {{ tag }}
            </span>
          </div>
        </el-card>

        <el-card class="side-card mt-4" shadow="never">
          <template #header>
            <div class="card-header-title">使用说明</div>
          </template>
          <ul class="usage-list">
            <li>表现记录将作为期末评语生成的重要参考</li>
            <li>使用标签可以更好地分类和检索记录</li>
            <li>建议定期记录，保持记录的连续性</li>
          </ul>
        </el-card>

        <el-card class="side-card mt-4 stats-card" shadow="never">
          <div class="stats-grid">
            <div class="stat-item">
              <div class="stat-value success">24</div>
              <div class="stat-label">表扬</div>
            </div>
            <div class="stat-item">
              <div class="stat-value warning">8</div>
              <div class="stat-label">待改进</div>
            </div>
            <div class="stat-item">
              <div class="stat-value">5</div>
              <div class="stat-label">待生成</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 添加记录弹窗 -->
    <el-dialog v-model="showAddDialog" title="添加表现记录" width="600px" class="custom-dialog">
      <el-form :model="recordForm" label-width="80px" :rules="rules" ref="formRef" class="custom-form">
        <el-form-item label="学生" prop="studentId">
          <el-select v-model="recordForm.studentId" placeholder="选择学生" style="width: 100%">
            <el-option
              v-for="student in students"
              :key="student.id"
              :label="student.name"
              :value="student.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="日期" prop="recordDate">
          <el-date-picker
            v-model="recordForm.recordDate"
            type="date"
            placeholder="选择日期"
            style="width: 100%"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item label="标签">
          <el-select
            v-model="recordForm.tags"
            multiple
            filterable
            allow-create
            placeholder="选择或输入标签"
            style="width: 100%"
          >
            <el-option v-for="tag in allTags" :key="tag" :label="tag" :value="tag" />
          </el-select>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input
            v-model="recordForm.content"
            type="textarea"
            :rows="4"
            placeholder="记录学生的表现..."
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showAddDialog = false" class="btn-secondary">取消</el-button>
          <el-button type="primary" @click="saveRecord" class="btn-primary">保存</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Edit, Delete, View } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'

interface Student {
  id: string
  name: string
}

interface PerformanceRecord {
  id?: string
  studentId: string
  studentName: string
  content: string
  tags: string[]
  recordDate: string
  time?: string
  type?: 'success' | 'warning' | 'primary' | 'danger'
}

const searchQuery = ref('')
const showAddDialog = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

const students = ref<Student[]>([
  { id: '1', name: '王小明' },
  { id: '2', name: '李小红' },
  { id: '3', name: '张三' },
  { id: '4', name: '赵四' }
])

const records = ref<PerformanceRecord[]>([
  {
    id: '1',
    studentId: '1',
    studentName: '王小明',
    content: '在数学课上积极回答问题，解题思路清晰，帮助其他同学理解难题',
    tags: ['积极发言', '乐于助人'],
    recordDate: '今天',
    time: '10:30',
    type: 'success'
  },
  {
    id: '2',
    studentId: '2',
    studentName: '李小红',
    content: '获得校级演讲比赛一等奖，表现优异',
    tags: ['竞赛获奖', '表达能力强'],
    recordDate: '今天',
    time: '09:15',
    type: 'success'
  },
  {
    id: '3',
    studentId: '3',
    studentName: '张三',
    content: '连续三次作业未按时提交，需要加强时间管理',
    tags: ['作业问题', '需改进'],
    recordDate: '昨天',
    time: '10:45',
    type: 'warning'
  },
  {
    id: '4',
    studentId: '4',
    studentName: '赵四',
    content: '主动组织班级卫生打扫，认真负责',
    tags: ['班级贡献', '责任心强'],
    recordDate: '昨天',
    time: '14:20',
    type: 'success'
  }
])

const quickTags = ['积极发言', '乐于助人', '认真负责', '成绩进步', '竞赛获奖', '团队协作', '创新思维', '表达能力强', '作业问题', '需改进', '迟到早退', '课堂纪律', '体育特长', '艺术特长', '班级贡献']
const allTags = [...quickTags]

const recordForm = reactive({
  studentId: '',
  content: '',
  tags: [] as string[],
  recordDate: new Date().toISOString().split('T')[0]
})

const rules: FormRules = {
  studentId: [{ required: true, message: '请选择学生', trigger: 'change' }],
  recordDate: [{ required: true, message: '请选择日期', trigger: 'change' }],
  content: [{ required: true, message: '请输入记录内容', trigger: 'blur' }]
}

const filteredRecords = computed(() => {
  if (!searchQuery.value) return records.value
  const query = searchQuery.value.toLowerCase()
  return records.value.filter(r =>
    r.studentName.toLowerCase().includes(query) ||
    r.content.toLowerCase().includes(query) ||
    r.tags.some(t => t.toLowerCase().includes(query))
  )
})

const getTypeLabel = (type?: string) => {
  const map: Record<string, string> = {
    success: '表扬',
    warning: '需改进',
    danger: '批评',
    primary: '记录'
  }
  return map[type || 'primary']
}

const addTagFilter = (tag: string) => {
  searchQuery.value = tag
}

const saveRecord = async () => {
  if (!formRef.value) return
  await formRef.value.validate((valid) => {
    if (valid) {
      const student = students.value.find(s => s.id === recordForm.studentId)
      const record: PerformanceRecord = {
        ...recordForm,
        id: Date.now().toString(),
        studentName: student?.name || '',
        type: recordForm.tags.includes('需改进') ? 'warning' : 'success',
        time: new Date().toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
      }
      records.value.unshift(record)
      ElMessage.success('记录添加成功')
      showAddDialog.value = false
      formRef.value?.resetFields()
    }
  })
}

const editRecord = (record: PerformanceRecord) => {
  ElMessage.info('编辑功能待实现')
}

const deleteRecord = (record: PerformanceRecord) => {
  ElMessageBox.confirm('确定要删除这条记录吗？', '确认删除', { type: 'warning' }).then(() => {
    records.value = records.value.filter(r => r.id !== record.id)
    ElMessage.success('删除成功')
  })
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.search-filter-bar {
  margin-bottom: $spacing-xl;
}

.search-box-table {
  position: relative;
  width: 320px;

  .search-icon {
    position: absolute;
    left: $spacing-md;
    top: 50%;
    transform: translateY(-50%);
    color: $text-muted;
    font-size: 16px;
    z-index: 1;
  }

  .search-input-clean {
    :deep(.el-input__wrapper) {
      background: $bg-input;
      border-radius: $radius-md;
      box-shadow: none;
      padding-left: 36px;

      .el-input__inner {
        height: 40px;
        font-size: 14px;

        &::placeholder {
          color: $text-muted;
        }
      }

      &.is-focus {
        box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
      }
    }
  }
}

.records-list {
  .record-card {
    background: $bg-card;
    border-radius: $radius-lg;
    border: 1px solid $border;
    padding: $spacing-lg;
    margin-bottom: $spacing-lg;
    transition: box-shadow $transition-fast;

    &:hover {
      box-shadow: $shadow-md;
    }

    .record-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: $spacing-md;
    }

    .student-info {
      display: flex;
      align-items: center;
      gap: $spacing-md;

      .avatar-circle {
        width: 40px;
        height: 40px;
        border-radius: $radius-full;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 14px;
        font-weight: 600;
        color: white;

        &.success {
          background: $success;
        }

        &.warning {
          background: $warning;
        }

        &.danger {
          background: $danger;
        }

        &.primary {
          background: $primary;
        }
      }

      .info-text {
        display: flex;
        flex-direction: column;
        gap: $spacing-xs;

        .student-name {
          font-weight: 600;
          font-size: 15px;
          color: $text-primary;
        }

        .type-tag {
          width: fit-content;
        }
      }
    }

    .record-time {
      font-size: 13px;
      color: $text-muted;
    }

    .record-content {
      color: $text-secondary;
      font-size: 14px;
      line-height: 1.6;
      margin-bottom: $spacing-md;
    }

    .record-footer {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .record-tags {
      display: flex;
      flex-wrap: wrap;
      gap: $spacing-sm;

      .tag-item {
        padding: 4px 10px;
        background: $bg-input;
        border-radius: $radius-sm;
        font-size: 12px;
        color: $text-secondary;
      }
    }

    .record-actions {
      display: flex;
      gap: $spacing-xs;

      .el-button {
        font-size: 16px;
      }
    }
  }
}

.side-card {
  border-radius: $radius-lg;
  border: 1px solid $border;
  margin-bottom: $spacing-lg;

  :deep(.el-card__header) {
    padding: $spacing-lg $spacing-lg $spacing-sm;
    border-bottom: none;

    .card-header-title {
      font-size: 16px;
      font-weight: 600;
      color: $text-primary;
    }
  }

  :deep(.el-card__body) {
    padding: $spacing-md $spacing-lg $spacing-lg;
  }

  .card-desc {
    font-size: 13px;
    color: $text-muted;
    margin: 0 0 $spacing-md;
  }
}

.quick-tags {
  display: flex;
  flex-wrap: wrap;
  gap: $spacing-sm;

  .quick-tag-item {
    padding: 6px 12px;
    background: $bg-card;
    border: 1px solid $border;
    border-radius: $radius-md;
    font-size: 13px;
    color: $text-secondary;
    cursor: pointer;
    transition: all $transition-fast;

    &:hover {
      background: $primary;
      border-color: $primary;
      color: white;
    }
  }
}

.usage-list {
  margin: 0;
  padding-left: $spacing-lg;

  li {
    font-size: 13px;
    color: $text-secondary;
    line-height: 1.8;
    margin-bottom: $spacing-sm;

    &::marker {
      color: $primary;
    }
  }
}

.stats-card {
  :deep(.el-card__body) {
    padding: $spacing-lg;
  }
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: $spacing-md;
  text-align: center;

  .stat-item {
    .stat-value {
      font-size: 24px;
      font-weight: 700;
      color: $text-primary;
      line-height: 1;

      &.success {
        color: $success;
      }

      &.warning {
        color: $warning;
      }
    }

    .stat-label {
      font-size: 12px;
      color: $text-secondary;
      margin-top: $spacing-xs;
    }
  }
}

.mt-4 {
  margin-top: $spacing-lg;
}

.btn-primary {
  background: $primary;
  border-color: $primary;
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;

  &:hover {
    background: $primary-light;
    border-color: $primary-light;
  }
}

.btn-secondary {
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;
}

:deep(.custom-dialog) {
  .el-dialog__header {
    padding: $spacing-xl $spacing-xl $spacing-sm;
    margin: 0;

    .el-dialog__title {
      font-size: 18px;
      font-weight: 600;
      color: $text-primary;
    }
  }

  .el-dialog__body {
    padding: $spacing-md $spacing-xl;
  }
}

.custom-form {
  .el-form-item {
    margin-bottom: $spacing-lg;
  }

  .el-input__wrapper,
  .el-select .el-input__wrapper,
  .el-date-editor .el-input__wrapper {
    border-radius: $radius-md;
    box-shadow: 0 0 0 1px $border;

    &.is-focus {
      box-shadow: 0 0 0 1px $primary;
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: $spacing-md;
  padding: $spacing-md 0 0;
}
</style>
```

- [ ] **Step 2: 验证页面显示**

确认：
- 记录卡片样式正确
- 侧边栏有快捷标签和使用说明
- 统计卡片显示正确
- 弹窗样式符合规范

- [ ] **Step 3: Commit**

```bash
git add src/views/performance/Index.vue
git commit -m "style: redesign performance page with new design system"
```

---

## Task 5: 重构评语生成页面

**Files:**
- Modify: `src/views/comments/Index.vue`

- [ ] **Step 1: 重写评语生成页面**

```vue
<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">
          评语生成
          <span class="subtitle">基于学生成绩和表现记录智能生成评语</span>
        </h1>
      </div>
      <div class="actions">
        <el-button class="btn-secondary" @click="showSettings = true">
          <el-icon><Setting /></el-icon>
          生成设置
        </el-button>
        <el-button type="primary" class="btn-primary" @click="generateComments" :loading="generating">
          <el-icon><MagicStick /></el-icon>
          AI生成评语
        </el-button>
      </div>
    </div>

    <el-card class="content-card" shadow="never">
      <div class="stats-bar">
        <div class="stat-box">
          <span class="stat-label">总学生数</span>
          <span class="stat-number">{{ students.length }}</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat-box">
          <span class="stat-label">已生成评语</span>
          <span class="stat-number success">{{ generatedCount }}</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat-box">
          <span class="stat-label">待生成</span>
          <span class="stat-number warning">{{ pendingCount }}</span>
        </div>
      </div>

      <el-table :data="students" v-loading="generating" style="width: 100%">
        <el-table-column min-width="100">
          <template #default="{ row }">
            <div class="student-name-cell">
              <div class="avatar-circle">{{ row.name.charAt(0) }}</div>
              <span class="name-text">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="近期成绩" min-width="180">
          <template #default="{ row }">
            <div class="score-tags">
              <span
                v-for="score in row.recentScores"
                :key="score.subject"
                :class="['subject-tag', getSubjectClass(score.subject)]"
              >
                {{ score.subject }} {{ score.score }}
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="表现记录" min-width="100">
          <template #default="{ row }">
            <span class="record-count">{{ row.recordCount }} 条记录</span>
          </template>
        </el-table-column>
        <el-table-column label="评语预览" min-width="280">
          <template #default="{ row }">
            <div v-if="row.comment" class="comment-preview">
              {{ row.comment }}
            </div>
            <span v-else class="no-comment">暂无评语</span>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <span :class="['status-tag', row.comment ? 'success' : 'warning']">
              {{ row.comment ? '已生成' : '待生成' }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <div class="action-btns">
              <el-button type="primary" link @click="viewComment(row)">
                <el-icon><View /></el-icon>
              </el-button>
              <el-button type="primary" link @click="regenerateSingle(row)">
                <el-icon><Refresh /></el-icon>
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 生成设置弹窗 -->
    <el-dialog v-model="showSettings" title="评语生成设置" width="500px" class="custom-dialog">
      <el-form :model="settings" label-width="100px" class="custom-form">
        <el-form-item label="评语风格">
          <el-radio-group v-model="settings.style">
            <el-radio-button label="encouragement">鼓励型</el-radio-button>
            <el-radio-button label="balanced">综合型</el-radio-button>
            <el-radio-button label="reminder">提醒型</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="评语长度">
          <el-slider v-model="settings.length" :min="50" :max="200" :step="10" show-stops />
          <div class="slider-hint">{{ settings.length }} 字左右</div>
        </el-form-item>
        <el-form-item label="包含内容">
          <el-checkbox-group v-model="settings.include">
            <el-checkbox label="score">成绩表现</el-checkbox>
            <el-checkbox label="performance">日常表现</el-checkbox>
            <el-checkbox label="suggestion">改进建议</el-checkbox>
            <el-checkbox label="encouragement">鼓励话语</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="额外要求">
          <el-input
            v-model="settings.extra"
            type="textarea"
            :rows="3"
            placeholder="例如：强调学生的进步、提及某个具体事件等"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showSettings = false" class="btn-secondary">取消</el-button>
          <el-button type="primary" @click="saveSettings" class="btn-primary">保存设置</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 评语预览弹窗 -->
    <el-dialog v-model="showPreview" title="评语详情" width="600px" class="custom-dialog">
      <div v-if="currentStudent" class="comment-detail">
        <div class="student-header">
          <div class="avatar-circle large">{{ currentStudent.name.charAt(0) }}</div>
          <div class="student-info">
            <span class="name">{{ currentStudent.name }}</span>
            <span class="class-info">初中{{ currentStudent.grade }}班</span>
          </div>
        </div>
        <div class="comment-box">
          <p class="comment-text">{{ currentStudent.comment }}</p>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPreview = false" class="btn-secondary">关闭</el-button>
          <el-button type="primary" @click="regenerateSingle(currentStudent)" class="btn-primary">
            <el-icon><Refresh /></el-icon>
            重新生成
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Setting, MagicStick, View, Refresh } from '@element-plus/icons-vue'

interface Score {
  subject: string
  score: number
}

interface Student {
  id: string
  name: string
  class: string
  grade: number
  recentScores: Score[]
  recordCount: number
  comment?: string
}

const generating = ref(false)
const showSettings = ref(false)
const showPreview = ref(false)
const currentStudent = ref<Student | null>(null)

const settings = reactive({
  style: 'balanced',
  length: 100,
  include: ['score', 'performance', 'suggestion'],
  extra: ''
})

const students = ref<Student[]>([
  {
    id: '1',
    name: '王小明',
    class: '1班',
    grade: 9,
    recentScores: [
      { subject: '语文', score: 92 },
      { subject: '数学', score: 98 },
      { subject: '英语', score: 85 }
    ],
    recordCount: 8,
    comment: '王小明同学在本学期表现优异，学习态度端正，成绩稳步提升。在数学方面尤为突出，取得了98分的好成绩。希望继续保持，争取更大进步。'
  },
  {
    id: '2',
    name: '李小红',
    class: '1班',
    grade: 9,
    recentScores: [
      { subject: '语文', score: 88 },
      { subject: '数学', score: 76 },
      { subject: '英语', score: 94 }
    ],
    recordCount: 5,
    comment: '李小红同学英语成绩突出，在校级演讲比赛中获得一等奖，表现优异。希望能将英语学习中的热情和方法应用到其他科目中。'
  },
  {
    id: '3',
    name: '张三',
    class: '1班',
    grade: 9,
    recentScores: [
      { subject: '语文', score: 75 },
      { subject: '数学', score: 68 },
      { subject: '英语', score: 72 }
    ],
    recordCount: 3,
    comment: ''
  },
  {
    id: '4',
    name: '赵四',
    class: '1班',
    grade: 9,
    recentScores: [
      { subject: '语文', score: 58 },
      { subject: '数学', score: 62 },
      { subject: '英语', score: 55 }
    ],
    recordCount: 4,
    comment: ''
  },
  {
    id: '5',
    name: '刘五',
    class: '1班',
    grade: 9,
    recentScores: [
      { subject: '语文', score: 85 },
      { subject: '数学', score: 91 },
      { subject: '英语', score: 88 }
    ],
    recordCount: 6,
    comment: '刘五同学是班级中的优秀学生之一，各科成绩均衡发展，尤其在数学方面表现突出。积极参与课堂讨论，乐于助人。'
  },
  {
    id: '6',
    name: '陈六',
    class: '1班',
    grade: 9,
    recentScores: [
      { subject: '语文', score: 82 },
      { subject: '数学', score: 79 },
      { subject: '英语', score: 81 }
    ],
    recordCount: 2,
    comment: ''
  }
])

const generatedCount = computed(() => students.value.filter(s => s.comment).length)
const pendingCount = computed(() => students.value.filter(s => !s.comment).length)

const getSubjectClass = (subject: string) => {
  const map: Record<string, string> = {
    '语文': 'chinese',
    '数学': 'math',
    '英语': 'english',
    '物理': 'physics',
    '化学': 'chemistry'
  }
  return map[subject] || ''
}

const generateComments = async () => {
  generating.value = true
  setTimeout(() => {
    students.value.forEach(student => {
      if (!student.comment) {
        student.comment = generateMockComment(student)
      }
    })
    generating.value = false
    ElMessage.success('评语生成完成')
  }, 2000)
}

const generateMockComment = (student: Student): string => {
  const avgScore = student.recentScores.reduce((a, b) => a + b.score, 0) / student.recentScores.length
  if (avgScore >= 90) {
    return `${student.name}同学在本学期表现优异，学习态度端正，成绩优秀。各科均衡发展，是同学们学习的好榜样。希望继续保持，争取更大进步。`
  } else if (avgScore >= 75) {
    return `${student.name}同学本学期学习态度认真，成绩良好。在${student.recentScores[0].subject}方面表现较为突出。希望能够继续努力，争取更大进步。`
  } else {
    return `${student.name}同学本学期有进步空间，需要加强学习主动性。建议合理安排学习时间，积极向老师和同学请教。相信通过努力一定能取得更好的成绩。`
  }
}

const regenerateSingle = (student: Student | null) => {
  if (!student) return
  student.comment = generateMockComment(student)
  ElMessage.success('评语已重新生成')
  if (showPreview.value) {
    showPreview.value = false
  }
}

const viewComment = (student: Student) => {
  currentStudent.value = student
  showPreview.value = true
}

const saveSettings = () => {
  ElMessage.success('设置已保存')
  showSettings.value = false
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.content-card {
  border-radius: $radius-lg;
  border: 1px solid $border;

  :deep(.el-card__body) {
    padding: 0;
  }
}

.stats-bar {
  display: flex;
  align-items: center;
  padding: $spacing-lg $spacing-xl;
  border-bottom: 1px solid $border-light;
  gap: $spacing-xl;

  .stat-box {
    display: flex;
    flex-direction: column;
    gap: $spacing-xs;

    .stat-label {
      font-size: 13px;
      color: $text-secondary;
    }

    .stat-number {
      font-size: 24px;
      font-weight: 700;
      color: $text-primary;
      line-height: 1;

      &.success {
        color: $success;
      }

      &.warning {
        color: $warning;
      }
    }
  }

  .stat-divider {
    width: 1px;
    height: 32px;
    background: $border;
  }
}

.student-name-cell {
  display: flex;
  align-items: center;
  gap: $spacing-md;

  .avatar-circle {
    width: 36px;
    height: 36px;
    border-radius: $radius-full;
    background: $primary-lighter;
    color: $primary;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    font-weight: 600;
  }

  .name-text {
    font-size: 14px;
    color: $text-primary;
    font-weight: 500;
  }
}

.score-tags {
  display: flex;
  flex-wrap: wrap;
  gap: $spacing-xs;
}

.subject-tag {
  display: inline-flex;
  align-items: center;
  padding: 2px 8px;
  border-radius: $radius-sm;
  font-size: 12px;
  font-weight: 500;

  &.chinese {
    background: $subject-chinese-bg;
    color: $subject-chinese-text;
  }

  &.math {
    background: $subject-math-bg;
    color: $subject-math-text;
  }

  &.english {
    background: $subject-english-bg;
    color: $subject-english-text;
  }

  &.physics {
    background: $subject-physics-bg;
    color: $subject-physics-text;
  }

  &.chemistry {
    background: $subject-chemistry-bg;
    color: $subject-chemistry-text;
  }
}

.record-count {
  font-size: 13px;
  color: $text-secondary;
}

.comment-preview {
  font-size: 13px;
  color: $text-secondary;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.no-comment {
  font-size: 13px;
  color: $text-muted;
}

.status-tag {
  display: inline-flex;
  align-items: center;
  padding: 4px 10px;
  border-radius: $radius-sm;
  font-size: 12px;
  font-weight: 500;

  &.success {
    background: $success-light;
    color: $success-text;
  }

  &.warning {
    background: $warning-light;
    color: $warning-text;
  }
}

.action-btns {
  display: flex;
  gap: $spacing-xs;

  .el-button {
    font-size: 16px;
  }
}

.btn-primary {
  background: $primary;
  border-color: $primary;
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;

  &:hover {
    background: $primary-light;
    border-color: $primary-light;
  }
}

.btn-secondary {
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;
}

:deep(.custom-dialog) {
  .el-dialog__header {
    padding: $spacing-xl $spacing-xl $spacing-sm;
    margin: 0;

    .el-dialog__title {
      font-size: 18px;
      font-weight: 600;
      color: $text-primary;
    }
  }

  .el-dialog__body {
    padding: $spacing-md $spacing-xl;
  }
}

.custom-form {
  .el-form-item {
    margin-bottom: $spacing-lg;
  }

  .el-slider {
    margin-top: $spacing-sm;
  }

  .slider-hint {
    font-size: 12px;
    color: $text-muted;
    margin-top: $spacing-xs;
  }

  .el-radio-button__original-radio:checked + .el-radio-button__inner {
    background: $primary;
    border-color: $primary;
    box-shadow: -1px 0 0 0 $primary;
  }

  .el-checkbox__input.is-checked .el-checkbox__inner {
    background: $primary;
    border-color: $primary;
  }

  .el-checkbox__input.is-checked + .el-checkbox__label {
    color: $primary;
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: $spacing-md;
  padding: $spacing-md 0 0;
}

.comment-detail {
  .student-header {
    display: flex;
    align-items: center;
    gap: $spacing-md;
    margin-bottom: $spacing-lg;
    padding-bottom: $spacing-lg;
    border-bottom: 1px solid $border-light;

    .avatar-circle {
      width: 48px;
      height: 48px;
      border-radius: $radius-full;
      background: $primary-lighter;
      color: $primary;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 18px;
      font-weight: 600;

      &.large {
        width: 56px;
        height: 56px;
        font-size: 20px;
      }
    }

    .student-info {
      display: flex;
      flex-direction: column;

      .name {
        font-size: 18px;
        font-weight: 600;
        color: $text-primary;
      }

      .class-info {
        font-size: 13px;
        color: $text-secondary;
      }
    }
  }

  .comment-box {
    background: $bg-primary;
    border-radius: $radius-lg;
    padding: $spacing-lg;

    .comment-text {
      font-size: 14px;
      line-height: 1.8;
      color: $text-primary;
      margin: 0;
    }
  }
}
</style>
```

- [ ] **Step 2: 验证页面显示**

确认：
- 统计栏显示正确
- 学科标签颜色正确
- 状态标签颜色正确
- 弹窗样式符合规范

- [ ] **Step 3: Commit**

```bash
git add src/views/comments/Index.vue
git commit -m "style: redesign comments page with new design system"
```

---

## Task 6: 重构仪表盘页面

**Files:**
- Modify: `src/views/Dashboard.vue`

- [ ] **Step 1: 重写仪表盘页面**

```vue
<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">仪表盘</h1>
      </div>
    </div>

    <!-- 欢迎卡片 -->
    <el-card class="welcome-card" shadow="never">
      <div class="welcome-content">
        <div class="welcome-text">
          <h2>下午好，张老师 👋</h2>
          <p>今天是 {{ today }}，您有 {{ todoCount }} 项待办事项</p>
        </div>
        <div class="welcome-illustration">
          <el-icon size="64" color="#6366f1"><School /></el-icon>
        </div>
      </div>
    </el-card>

    <!-- 统计卡片 -->
    <el-row :gutter="24" class="stats-row">
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-icon" style="background: #e0e7ff; color: #4f46e5;">
            <el-icon size="24"><User /></el-icon>
          </div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.students }}</span>
            <span class="stat-label">学生总数</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-icon" style="background: #d1fae5; color: #059669;">
            <el-icon size="24"><DocumentChecked /></el-icon>
          </div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.records }}</span>
            <span class="stat-label">表现记录</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-icon" style="background: #fef3c7; color: #d97706;">
            <el-icon size="24"><ChatDotRound /></el-icon>
          </div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.pendingComments }}</span>
            <span class="stat-label">待生成评语</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-icon" style="background: #fee2e2; color: #dc2626;">
            <el-icon size="24"><TrendCharts /></el-icon>
          </div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.avgScore }}</span>
            <span class="stat-label">班级平均分</span>
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 快捷入口 -->
    <el-row :gutter="24" class="quick-actions-row">
      <el-col :span="12">
        <el-card class="action-card" shadow="never">
          <template #header>
            <div class="card-header-title">快捷入口</div>
          </template>
          <div class="quick-grid">
            <router-link to="/students" class="quick-item">
              <div class="quick-icon" style="background: #e0e7ff; color: #4f46e5;">
                <el-icon size="24"><User /></el-icon>
              </div>
              <span class="quick-label">学生管理</span>
            </router-link>
            <router-link to="/performance" class="quick-item">
              <div class="quick-icon" style="background: #d1fae5; color: #059669;">
                <el-icon size="24"><EditPen /></el-icon>
              </div>
              <span class="quick-label">添加记录</span>
            </router-link>
            <router-link to="/comments" class="quick-item">
              <div class="quick-icon" style="background: #fef3c7; color: #d97706;">
                <el-icon size="24"><MagicStick /></el-icon>
              </div>
              <span class="quick-label">生成评语</span>
            </router-link>
            <router-link to="/scores" class="quick-item">
              <div class="quick-icon" style="background: #fee2e2; color: #dc2626;">
                <el-icon size="24"><TrendCharts /></el-icon>
              </div>
              <span class="quick-label">成绩管理</span>
            </router-link>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card class="action-card" shadow="never">
          <template #header>
            <div class="card-header-title">最近动态</div>
          </template>
          <div class="activity-list">
            <div v-for="(activity, index) in recentActivities" :key="index" class="activity-item">
              <div class="activity-dot" :style="{ background: activity.color }"></div>
              <div class="activity-content">
                <p class="activity-text">{{ activity.text }}</p>
                <span class="activity-time">{{ activity.time }}</span>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import {
  School, User, DocumentChecked, ChatDotRound,
  TrendCharts, EditPen, MagicStick
} from '@element-plus/icons-vue'

const stats = ref({
  students: 48,
  records: 156,
  pendingComments: 12,
  avgScore: 82.5
})

const todoCount = ref(3)

const today = computed(() => {
  const date = new Date()
  const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
  return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日 ${weekdays[date.getDay()]}`
})

const recentActivities = ref([
  { text: '添加了王小明的表现记录', time: '10分钟前', color: '#10b981' },
  { text: '生成了5条学生评语', time: '1小时前', color: '#6366f1' },
  { text: '导入月考成绩数据', time: '3小时前', color: '#f59e0b' },
  { text: '添加了3名新学生', time: '昨天', color: '#3b82f6' }
])
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.welcome-card {
  border-radius: $radius-lg;
  border: 1px solid $border;
  margin-bottom: $spacing-xl;
  background: linear-gradient(135deg, #f8f9fa 0%, #ffffff 100%);

  :deep(.el-card__body) {
    padding: $spacing-2xl;
  }

  .welcome-content {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .welcome-text {
      h2 {
        font-size: 24px;
        font-weight: 600;
        color: $text-primary;
        margin: 0 0 $spacing-sm;
      }

      p {
        font-size: 14px;
        color: $text-secondary;
        margin: 0;
      }
    }

    .welcome-illustration {
      width: 80px;
      height: 80px;
      background: $primary-lighter;
      border-radius: $radius-xl;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}

.stats-row {
  margin-bottom: $spacing-xl;
}

.stat-card {
  background: $bg-card;
  border-radius: $radius-lg;
  border: 1px solid $border;
  padding: $spacing-xl;
  display: flex;
  align-items: center;
  gap: $spacing-lg;
  transition: box-shadow $transition-fast;

  &:hover {
    box-shadow: $shadow-md;
  }

  .stat-icon {
    width: 56px;
    height: 56px;
    border-radius: $radius-lg;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .stat-info {
    display: flex;
    flex-direction: column;

    .stat-value {
      font-size: 28px;
      font-weight: 700;
      color: $text-primary;
      line-height: 1;
    }

    .stat-label {
      font-size: 13px;
      color: $text-secondary;
      margin-top: $spacing-xs;
    }
  }
}

.quick-actions-row {
  .action-card {
    border-radius: $radius-lg;
    border: 1px solid $border;
    height: 100%;

    :deep(.el-card__header) {
      padding: $spacing-lg $spacing-xl;
      border-bottom: 1px solid $border-light;

      .card-header-title {
        font-size: 16px;
        font-weight: 600;
        color: $text-primary;
      }
    }

    :deep(.el-card__body) {
      padding: $spacing-xl;
    }
  }
}

.quick-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: $spacing-lg;

  .quick-item {
    display: flex;
    align-items: center;
    gap: $spacing-md;
    padding: $spacing-lg;
    background: $bg-primary;
    border-radius: $radius-lg;
    text-decoration: none;
    transition: all $transition-fast;

    &:hover {
      background: $bg-hover;
      transform: translateY(-2px);
    }

    .quick-icon {
      width: 48px;
      height: 48px;
      border-radius: $radius-md;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .quick-label {
      font-size: 14px;
      font-weight: 500;
      color: $text-primary;
    }
  }
}

.activity-list {
  .activity-item {
    display: flex;
    gap: $spacing-md;
    padding: $spacing-md 0;
    border-bottom: 1px solid $border-light;

    &:last-child {
      border-bottom: none;
      padding-bottom: 0;
    }

    &:first-child {
      padding-top: 0;
    }

    .activity-dot {
      width: 8px;
      height: 8px;
      border-radius: $radius-full;
      margin-top: 6px;
      flex-shrink: 0;
    }

    .activity-content {
      flex: 1;

      .activity-text {
        font-size: 14px;
        color: $text-primary;
        margin: 0 0 $spacing-xs;
      }

      .activity-time {
        font-size: 12px;
        color: $text-muted;
      }
    }
  }
}
</style>
```

- [ ] **Step 2: 验证页面显示**

确认：
- 欢迎卡片显示正确
- 统计卡片有彩色图标
- 快捷入口网格布局正确
- 最近动态列表显示正确

- [ ] **Step 3: Commit**

```bash
git add src/views/Dashboard.vue
git commit -m "style: redesign dashboard page with new design system"
```

---

## Task 7: 重构成绩管理页面

**Files:**
- Modify: `src/views/scores/Index.vue`

- [ ] **Step 1: 读取现有文件（如有）**

```bash
cat src/views/scores/Index.vue
```

如果文件不存在或为空，创建新的成绩管理页面：

```vue
<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">
          成绩管理
          <span class="subtitle">录入和管理学生考试成绩</span>
        </h1>
      </div>
      <div class="actions">
        <el-button class="btn-secondary" @click="showImportDialog = true">
          <el-icon><Upload /></el-icon>
          导入成绩
        </el-button>
        <el-button type="primary" class="btn-primary" @click="showAddDialog = true">
          <el-icon><Plus /></el-icon>
          录入成绩
        </el-button>
      </div>
    </div>

    <el-card class="content-card" shadow="never">
      <div class="filter-bar">
        <el-select v-model="filter.exam" placeholder="选择考试" style="width: 180px">
          <el-option label="期中考试" value="midterm" />
          <el-option label="月考一" value="monthly1" />
          <el-option label="月考二" value="monthly2" />
        </el-select>
        <el-select v-model="filter.subject" placeholder="选择科目" style="width: 140px">
          <el-option label="语文" value="chinese" />
          <el-option label="数学" value="math" />
          <el-option label="英语" value="english" />
        </el-select>
      </div>

      <el-table :data="scores" style="width: 100%">
        <el-table-column min-width="120">
          <template #default="{ row }">
            <div class="student-name-cell">
              <div class="avatar-circle">{{ row.name.charAt(0) }}</div>
              <span class="name-text">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="chinese" label="语文" min-width="80">
          <template #default="{ row }">
            <span :class="['score-cell', getScoreClass(row.chinese)]">{{ row.chinese }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="math" label="数学" min-width="80">
          <template #default="{ row }">
            <span :class="['score-cell', getScoreClass(row.math)]">{{ row.math }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="english" label="英语" min-width="80">
          <template #default="{ row }">
            <span :class="['score-cell', getScoreClass(row.english)]">{{ row.english }}</span>
          </template>
        </el-table-column>
        <el-table-column label="总分" min-width="80">
          <template #default="{ row }">
            <span class="score-cell total">{{ row.chinese + row.math + row.english }}</span>
          </template>
        </el-table-column>
        <el-table-column label="平均分" min-width="80">
          <template #default="{ row }">
            <span class="score-cell avg">{{ ((row.chinese + row.math + row.english) / 3).toFixed(1) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <div class="action-btns">
              <el-button type="primary" link @click="editScore(row)">
                <el-icon><Edit /></el-icon>
              </el-button>
              <el-button type="danger" link @click="deleteScore(row)">
                <el-icon><Delete /></el-icon>
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { Plus, Upload, Edit, Delete } from '@element-plus/icons-vue'

interface Score {
  id: string
  name: string
  chinese: number
  math: number
  english: number
}

const filter = reactive({
  exam: 'midterm',
  subject: ''
})

const showAddDialog = ref(false)
const showImportDialog = ref(false)

const scores = ref<Score[]>([
  { id: '1', name: '王小明', chinese: 92, math: 98, english: 85 },
  { id: '2', name: '李小红', chinese: 88, math: 76, english: 94 },
  { id: '3', name: '张三', chinese: 75, math: 68, english: 72 },
  { id: '4', name: '赵四', chinese: 58, math: 62, english: 55 },
  { id: '5', name: '刘五', chinese: 85, math: 91, english: 88 },
  { id: '6', name: '陈六', chinese: 82, math: 79, english: 81 }
])

const getScoreClass = (score: number) => {
  if (score >= 90) return 'excellent'
  if (score >= 75) return 'good'
  if (score >= 60) return 'pass'
  return 'fail'
}

const editScore = (row: Score) => {
  console.log('Edit', row)
}

const deleteScore = (row: Score) => {
  console.log('Delete', row)
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.content-card {
  border-radius: $radius-lg;
  border: 1px solid $border;

  :deep(.el-card__body) {
    padding: $spacing-xl;
  }
}

.filter-bar {
  display: flex;
  gap: $spacing-md;
  margin-bottom: $spacing-lg;

  :deep(.el-select .el-input__wrapper) {
    border-radius: $radius-md;
  }
}

.student-name-cell {
  display: flex;
  align-items: center;
  gap: $spacing-md;

  .avatar-circle {
    width: 36px;
    height: 36px;
    border-radius: $radius-full;
    background: $primary-lighter;
    color: $primary;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    font-weight: 600;
  }

  .name-text {
    font-size: 14px;
    color: $text-primary;
    font-weight: 500;
  }
}

.score-cell {
  font-weight: 600;

  &.excellent {
    color: $success;
  }

  &.good {
    color: $info;
  }

  &.pass {
    color: $warning;
  }

  &.fail {
    color: $danger;
  }

  &.total {
    color: $text-primary;
    font-size: 15px;
  }

  &.avg {
    color: $primary;
  }
}

.action-btns {
  display: flex;
  gap: $spacing-xs;

  .el-button {
    font-size: 16px;
  }
}

.btn-primary {
  background: $primary;
  border-color: $primary;
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;

  &:hover {
    background: $primary-light;
    border-color: $primary-light;
  }
}

.btn-secondary {
  border-radius: $radius-md;
  height: 36px;
  padding: 0 $spacing-lg;
}
</style>
```

- [ ] **Step 2: 验证页面显示**

确认：
- 过滤栏显示正确
- 成绩有颜色区分（优秀、良好、及格、不及格）
- 表格样式符合规范

- [ ] **Step 3: Commit**

```bash
git add src/views/scores/Index.vue
git commit -m "style: redesign scores page with new design system"
```

---

## Task 8: 最终验证

- [ ] **Step 1: 运行开发服务器**

```bash
cd /Users/dd/projects/teacher_assistant/frontend
npm run dev
```

- [ ] **Step 2: 验证所有页面**

访问以下路径确认样式正确：
- http://localhost:5173/ - 仪表盘
- http://localhost:5173/students - 学生管理
- http://localhost:5173/performance - 表现记录
- http://localhost:5173/comments - 评语生成
- http://localhost:5173/scores - 成绩管理

- [ ] **Step 3: 检查控制台**

确认无错误、无警告

- [ ] **Step 4: 最终 Commit**

```bash
git add .
git commit -m "style: complete UI redesign matching reference screenshots

- Add comprehensive design system with color, spacing, radius tokens
- Redesign Layout component with new sidebar and header
- Redesign all pages: Dashboard, Students, Performance, Comments, Scores
- Add global styles for common components
- Ensure consistent visual style across all views"
```

---

## 实施完成

所有页面已按照截图风格重新设计，包括：
1. ✅ 统一的颜色系统（紫色主色）
2. ✅ 侧边栏导航（带选中态样式）
3. ✅ 顶部搜索栏和用户信息
4. ✅ 卡片式内容区域
5. ✅ 统一的按钮、表格、标签样式
6. ✅ 各页面的特定功能组件
