# 教师助手Agent系统设计文档

**日期**: 2026-04-29  
**版本**: v1.0  
**状态**: 待实现

---

## 1. 项目概述

教师助手Agent是一款面向初中教师的智能教学辅助工具，通过AI能力简化教师的日常重复工作，包括成绩管理、学生表现记录、评语生成和课件制作。

### 1.1 核心目标
- 减少教师在成绩统计、评语撰写上的重复劳动
- 提供智能化的教学资源生成能力
- 建立可溯源的专属题库系统
- 生成专业美观的PPT课件

### 1.2 使用场景
- **场景1**: 期中考试后，上传Excel成绩表，系统自动分析班级情况，生成学生个人成绩单和评语
- **场景2**: 日常教学中，快速记录学生表现，期末批量生成包含成绩和表现的综合评语
- **场景3**: 备课时，输入课题和年级，系统自动生成教案并从题库选题，输出精美PPT

---

## 2. 目标用户

**主要用户**: 初中教师（7-9年级）
**用户画像**:
- 年轻教师群体，对科技产品接受度高
- 希望提高工作效率，减少重复劳动
- 注重教学资源的质量和美观度

**部署模式**: 
- **第一阶段**: 个人本地版（单机部署，一位教师使用）
- **后续阶段**: 支持多用户和云部署

---

## 3. 核心功能模块

### 3.1 成绩管理模块

#### 功能列表
| 功能 | 优先级 | 描述 |
|------|--------|------|
| Excel智能导入 | P0 | 自动识别Excel列结构，支持用户确认映射关系 |
| 成绩查询 | P0 | 按学生、科目、考试时间筛选查看 |
| 成绩分析 | P1 | 趋势分析、偏科分析（雷达图）、稳定性分析 |
| 排名统计 | P1 | 班级排名、年级排名（预留）、排名变化轨迹 |
| 学习小组分组 | P1 | 智能算法分组，确保每组成绩分布均衡 |
| 数据导出 | P1 | Excel（原始数据）、PDF（分析报告） |

#### Excel导入流程
```
用户上传Excel文件
    ↓
系统自动识别表头和数据列
    ↓
展示列映射预览（姓名列→姓名，列A→语文成绩...）
    ↓
用户确认或调整映射
    ↓
数据验证（检查重复、异常值）
    ↓
导入成功，展示统计摘要
```

#### 学习小组分组算法
- 输入：学生名单及最近一次综合成绩
- 目标：将班级分为N组，每组平均成绩接近，组内成绩有梯度（便于互帮互助）
- 算法思路：先排序，再蛇形分配，最后微调平衡

---

### 3.2 学生表现记录模块

#### 功能列表
| 功能 | 优先级 | 描述 |
|------|--------|------|
| 文本记录 | P0 | 时间、标签、内容、关联学生 |
| 标签管理 | P0 | 自定义标签（积极发言、作业延迟、助人为乐等） |
| 语音输入（网页） | P1 | 浏览器MediaRecorder API录音转文字 |
| 语音输入（小程序） | P2 | 微信小程序录音（后续实现） |
| 记录查询 | P1 | 按学生、标签、时间段筛选 |

#### 记录数据结构
```json
{
  "id": "uuid",
  "student_id": "学生ID",
  "student_name": "学生姓名",
  "content": "记录内容",
  "tags": ["积极发言", "课堂活跃"],
  "record_date": "2026-04-29",
  "created_at": "timestamp",
  "source": "text|voice"
}
```

---

### 3.3 评语生成模块

#### 功能列表
| 功能 | 优先级 | 描述 |
|------|--------|------|
| 数据源整合 | P0 | 自动拉取学生成绩+表现记录 |
| AI评语生成 | P0 | 基于大模型生成个性化评语 |
| 模板管理 | P1 | 预设多种风格（鼓励型、提醒型、综合型） |
| 批量生成 | P1 | 一键生成全班评语 |
| 导出Word | P1 | 表格形式：学生姓名+评语 |
| 导出Excel | P1 | 批量导出，便于复制到成绩系统 |

#### 生成Prompt设计
```
角色：你是一位经验丰富的初中班主任，善于用温暖而专业的语言评价学生。

背景信息：
- 学生：{姓名}
- 年级：初中{年级}
- 近期成绩：{各科成绩及排名变化}
- 日常表现：{表现记录列表}

要求：
1. 生成一段{风格}的期末评语，100-150字
2. 结合具体成绩表现和日常行为
3. 既有肯定也有期望
4. 语言亲切自然，避免套话

风格选项：鼓励型/提醒型/综合型
```

---

### 3.4 题库与课件制作模块

#### 3.4.1 题库管理

**功能列表**
| 功能 | 优先级 | 描述 |
|------|--------|------|
| 题目录入 | P0 | 支持单题录入和批量导入 |
| AI辅助提取 | P1 | 粘贴题目文本，AI自动提取结构化数据 |
| 来源标注 | P1 | 记录题目来源URL/文档，标记可信度 |
| 题目审核 | P2 | 人工抽检确认题目准确性 |
| 题目检索 | P1 | 按知识点、难度、题型筛选 |

**题目数据结构**
```json
{
  "id": "uuid",
  "subject": "数学",
  "grade": 8,
  "type": "选择题",
  "knowledge_point": "二次函数",
  "difficulty": 3,
  "content": "题目内容",
  "options": ["A", "B", "C", "D"],
  "answer": "B",
  "analysis": "解析",
  "source": {
    "type": "website|document|manual",
    "url": "来源链接",
    "reliability": "high|medium|low"
  },
  "status": "pending|verified",
  "created_at": "timestamp"
}
```

#### 3.4.2 课件制作（核心亮点）

**功能列表**
| 功能 | 优先级 | 描述 |
|------|--------|------|
| AI教案生成 | P0 | 根据年级、科目、主题生成完整教学大纲 |
| 智能选题 | P1 | 从题库匹配知识点，自动选择题目 |
| PPT生成 | P0 | 生成专业美观的PPT文件 |
| 主题适配 | P1 | 根据学科自动匹配风格（数学-科技蓝/语文-古典雅） |
| 教学动画 | P1 | 答案延迟显示、逐步揭示解题步骤 |
| AI配图 | P2 | 调用nanobanana等接口生成配图 |

**课件生成流程**
```
用户输入：年级 + 科目 + 主题 + 特殊要求
    ↓
AI教案Agent生成教学大纲
    ↓
用户确认/修改大纲
    ↓
智能选题（匹配知识点，从题库选择）
    ↓
PPT生成引擎渲染
    - 主题风格选择
    - 题目排版
    - 动画设置
    - AI配图（如有）
    ↓
输出.pptx文件供下载
```

**PPT技术方案**
- 引擎：`python-pptx`（Go后端调用Python服务）
- 风格：专业精品风 + 暖调背景
- 动画：支持淡入、延迟显示、逐步揭示等教学动画
- 生图接口：预留多协议兼容层（nanobanana优先）

---

## 4. 技术架构

### 4.1 整体架构

```
┌─────────────────────────────────────────────────────────┐
│                    前端层 (Vue3)                         │
│     专业精品风UI + 暖调背景 + 精致交互                    │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐        │
│  │ 成绩管理 │ │ 表现记录 │ │ 评语生成 │ │ 课件制作 │        │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘        │
└─────────────────────────────────────────────────────────┘
                           │
┌─────────────────────────────────────────────────────────┐
│              API层 (Go + Gin)                           │
│  • RESTful API  • 统一错误处理  • 请求日志              │
└─────────────────────────────────────────────────────────┘
                           │
┌─────────────────────────────────────────────────────────┐
│                   业务服务层 (Go)                        │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐   │
│  │ 成绩服务  │ │ 记录服务  │ │ 评语服务  │ │ 课件服务  │   │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘   │
│  ┌──────────┐ ┌──────────┐                              │
│  │ 题库服务  │ │ AI服务    │                              │
│  └──────────┘ └──────────┘                              │
└─────────────────────────────────────────────────────────┘
                           │
┌─────────────────────────────────────────────────────────┐
│                   基础设施层                              │
│  MySQL ─┬─ 抽象数据库接口（支持后续云数据库）             │
│  Redis ─┴─ 缓存                                          │
│  本地文件系统 ── Excel/PPT/图片存储                       │
└─────────────────────────────────────────────────────────┘
```

### 4.2 技术栈选型

| 层级 | 选型 | 理由 |
|------|------|------|
| 后端框架 | Go + Gin | 高性能、并发强、编译单二进制文件 |
| 数据库 | MySQL + GORM | 关系型数据为主，GORM提供抽象层 |
| 缓存 | Redis | 会话、热点数据缓存 |
| 前端框架 | Vue3 + Composition API | 响应式系统适合交互丰富的后台 |
| UI组件库 | Element Plus（深度定制） | 组件丰富，可定制为专业精品风格 |
| 状态管理 | Pinia | 比Vuex更简洁 |
| 图表库 | ECharts | 成绩分析可视化 |
| Excel处理 | xlsx.js(前端) + excelize(Go) | 前后端协同 |
| PPT生成 | python-pptx | 功能完整，动画支持好 |
| AI大模型 | DeepSeek/通义千问 | 中文能力强，价格优，国内稳定 |
| AI生图 | nanobanana接口 | 本地部署友好，预留多协议兼容 |
| 语音识别 | 浏览器Web Speech API | 免费，满足基础需求 |

### 4.3 抽象接口设计

为支持后续扩展，核心业务接口需抽象：

```go
// 数据库驱动接口
 type DatabaseDriver interface {
     Connect(config DBConfig) error
     Close() error
     Query(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
     Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
     // ...
 }
 
 // 图片生成接口
 type ImageGenerator interface {
     Generate(prompt string, options GenOptions) (ImageResult, error)
     Supports() []string // 返回支持的协议列表
 }
 
 // LLM提供商接口
 type LLMProvider interface {
     Complete(ctx context.Context, messages []Message, options LLMOptions) (CompletionResult, error)
     StreamComplete(ctx context.Context, messages []Message, options LLMOptions) (<-chan StreamChunk, error)
 }
```

---

## 5. 数据模型设计

### 5.1 核心实体

```sql
-- 学生表
CREATE TABLE students (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    student_no VARCHAR(20) UNIQUE,
    class VARCHAR(20),
    grade TINYINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 考试成绩表
CREATE TABLE exam_scores (
    id VARCHAR(36) PRIMARY KEY,
    student_id VARCHAR(36) NOT NULL,
    exam_name VARCHAR(100) NOT NULL,
    exam_date DATE NOT NULL,
    subject VARCHAR(50) NOT NULL,
    score DECIMAL(5,2),
    full_score DECIMAL(5,2) DEFAULT 100,
    class_rank INT,
    grade_rank INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_student (student_id),
    INDEX idx_exam (exam_name, exam_date)
);

-- 表现记录表
CREATE TABLE performance_records (
    id VARCHAR(36) PRIMARY KEY,
    student_id VARCHAR(36) NOT NULL,
    content TEXT NOT NULL,
    tags JSON,
    record_date DATE NOT NULL,
    source ENUM('text', 'voice') DEFAULT 'text',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_student_date (student_id, record_date)
);

-- 题库表
CREATE TABLE questions (
    id VARCHAR(36) PRIMARY KEY,
    subject VARCHAR(50) NOT NULL,
    grade TINYINT,
    type ENUM('choice', 'fill', 'answer') NOT NULL,
    knowledge_point VARCHAR(100),
    difficulty TINYINT DEFAULT 3,
    content TEXT NOT NULL,
    options JSON,
    answer TEXT,
    analysis TEXT,
    source_type VARCHAR(20),
    source_url VARCHAR(500),
    reliability ENUM('high', 'medium', 'low') DEFAULT 'medium',
    status ENUM('pending', 'verified') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_subject_kp (subject, knowledge_point)
);

-- 评语记录表
CREATE TABLE comments (
    id VARCHAR(36) PRIMARY KEY,
    student_id VARCHAR(36) NOT NULL,
    semester VARCHAR(20),
    style VARCHAR(20) DEFAULT 'comprehensive',
    content TEXT NOT NULL,
    generated_by ENUM('ai', 'manual') DEFAULT 'ai',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 6. UI/UX设计规范

### 6.1 设计理念
**专业精品风 + 暖调背景**
- 精致排版、微妙层次感、高级感
- 暖调背景色，护眼温馨
- 像使用Linear、Notion般的精品体验

### 6.2 色彩系统

```css
/* 主色调 */
--primary: #6366f1;           /* 柔和紫蓝 */
--primary-light: #818cf8;     /* 浅紫 */
--primary-dark: #4f46e5;      /* 深紫 */

/* 背景色（暖调） */
--bg-primary: #faf9f7;        /* 暖白/米白 */
--bg-secondary: #f5f4f2;      /* 浅暖灰 */
--bg-card: #ffffff;           /* 纯白卡片 */

/* 文字色 */
--text-primary: #1a1a2e;      /* 深色文字 */
--text-secondary: #6b7280;    /* 次要文字 */
--text-muted: #9ca3af;        /* 辅助文字 */

/* 边框 */
--border: rgba(0,0,0,0.06);   /* 微妙边框 */
--border-light: rgba(0,0,0,0.04);

/* 强调色 */
--success: #10b981;           /* 成功绿 */
--warning: #f59e0b;           /* 警告橙 */
--error: #ef4444;             /* 错误红 */
```

### 6.3 布局规范

- **卡片圆角**: `16px`（大卡片）、`12px`（小卡片）、`8px`（标签按钮）
- **阴影层次**:
  - 基础: `0 1px 3px rgba(0,0,0,0.1)`
  - 悬浮: `0 10px 40px rgba(0,0,0,0.08)`
  - 弹窗: `0 20px 60px rgba(0,0,0,0.12)`
- **间距系统**: 4px基准，使用8、12、16、24、32、48px
- **字体**: Inter/PingFang SC，字号层次清晰

### 6.4 动效规范

- **过渡时长**: 200ms（快速）、300ms（标准）、500ms（舒缓）
- **缓动函数**: `cubic-bezier(0.4, 0, 0.2, 1)`
- **悬浮效果**: 轻微上浮 + 阴影加深
- **页面切换**: 淡入 + 轻微上滑

---

## 7. 开发阶段规划

### 7.1 第一阶段（MVP - 核心功能）

**目标**: 完成基础成绩管理和评语生成，验证核心价值

| 模块 | 功能 | 工期 |
|------|------|------|
| 基础框架 | Go后端搭建、Vue3前端框架、数据库设计 | 3天 |
| 成绩管理 | Excel导入、成绩查询、基础分析 | 5天 |
| 表现记录 | 文本记录、标签管理 | 3天 |
| 评语生成 | AI集成、批量生成、导出Word/Excel | 5天 |
| UI打磨 | 专业精品风UI实现、暖调色彩调整 | 4天 |

**里程碑1**: 教师可以导入成绩、记录表现、批量生成评语并导出

### 7.2 第二阶段（增强功能）

**目标**: 完善题库和课件制作，形成完整闭环

| 模块 | 功能 | 工期 |
|------|------|------|
| 成绩模块 | 学习小组分组、PDF报告导出、高级分析图表 | 4天 |
| 表现模块 | 网页语音输入、时间轴视图 | 3天 |
| 题库模块 | AI辅助提取、题目管理、来源标注 | 5天 |
| 课件模块 | AI教案生成、PPT生成引擎、基础动画 | 7天 |
| 集成优化 | 教案-选题-PPT全流程打通 | 3天 |

**里程碑2**: 可以生成完整的AI课件（教案+PPT）

### 7.3 第三阶段（体验升级）

**目标**: 小程序接入、多用户支持、云部署准备

| 模块 | 功能 | 工期 |
|------|------|------|
| 小程序 | 语音输入、快捷记录 | 5天 |
| 账号体系 | 用户注册登录、数据隔离 | 4天 |
| 云部署 | Docker容器化、配置中心 | 3天 |
| 高级功能 | AI配图、PPT主题多样化 | 5天 |

**里程碑3**: 支持多教师使用，可云部署

---

## 8. 接口设计要点

### 8.1 成绩管理接口

```
POST   /api/v1/scores/import          Excel导入成绩
GET    /api/v1/scores                  查询成绩列表
GET    /api/v1/scores/:studentId       查询学生成绩详情
POST   /api/v1/scores/analyze          成绩分析
POST   /api/v1/scores/grouping         学习小组分组
GET    /api/v1/scores/export           导出成绩数据
```

### 8.2 表现记录接口

```
POST   /api/v1/records                 创建表现记录
GET    /api/v1/records                 查询记录列表
PUT    /api/v1/records/:id             更新记录
DELETE /api/v1/records/:id             删除记录
GET    /api/v1/records/tags            获取标签列表
```

### 8.3 评语生成接口

```
POST   /api/v1/comments/generate       AI生成评语
POST   /api/v1/comments/batch          批量生成评语
GET    /api/v1/comments                查询评语列表
GET    /api/v1/comments/export/word    导出Word
GET    /api/v1/comments/export/excel   导出Excel
```

### 8.4 课件制作接口

```
POST   /api/v1/lessons/generate        生成教案大纲
POST   /api/v1/lessons/ppt             生成PPT
GET    /api/v1/lessons/:id/ppt/download 下载PPT文件
```

### 8.5 题库接口

```
POST   /api/v1/questions               创建题目
POST   /api/v1/questions/extract       AI提取题目
GET    /api/v1/questions               查询题目列表
PUT    /api/v1/questions/:id/verify    审核题目
```

---

## 9. 风险评估与应对

| 风险 | 影响 | 应对措施 |
|------|------|---------|
| AI生成内容不准确 | 高 | 关键数据（评语、教案）必须人工确认后方可使用 |
| Excel格式复杂多样 | 中 | 提供模板示例，智能识别后必须用户确认 |
| PPT生成效果不达预期 | 中 | 预留模板定制空间，允许用户上传自定义模板 |
| 题库版权问题 | 高 | 仅支持用户自行录入或明确来源的题目，系统不主动爬取 |
| 后期扩展架构瓶颈 | 中 | 数据库、图片生成、LLM均使用抽象接口，便于切换 |

---

## 10. 附录

### 10.1 命名规范

- **Go**: 包名小写单数，变量驼峰，常量全大写下划线
- **Vue**: 组件PascalCase，组合式函数useXxx，文件kebab-case
- **数据库**: 表名复数小写下划线，字段小写下划线

### 10.2 项目结构

```
teacher-assistant/
├── backend/
│   ├── cmd/                    # 启动入口
│   ├── internal/
│   │   ├── domain/            # 领域模型
│   │   ├── service/           # 业务逻辑
│   │   ├── repository/        # 数据访问（抽象接口）
│   │   ├── handler/           # HTTP处理器
│   │   └── pkg/               # 内部工具包
│   │       ├── ai/            # AI服务封装
│   │       ├── excel/         # Excel处理
│   │       └── ppt/           # PPT生成
│   ├── pkg/                   # 公共包
│   └── configs/               # 配置文件
├── frontend/
│   ├── src/
│   │   ├── views/             # 页面视图
│   │   ├── components/        # 组件
│   │   ├── api/               # API接口
│   │   ├── stores/            # Pinia状态
│   │   ├── styles/            # 全局样式
│   │   └── utils/             # 工具函数
│   └── public/
├── ai-agents/                 # AI Agent定义
│   ├── lesson-agent/
│   └── comment-agent/
└── docs/
```

---

**文档编制**: Claude Code  
**审核状态**: 待用户审核
