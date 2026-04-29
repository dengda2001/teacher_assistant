# 教师助手

面向初中教师的智能教学辅助工具，帮助教师管理成绩、记录学生表现、AI生成评语。

## 技术栈

- **后端**: Go + Gin + GORM + MySQL
- **前端**: Vue3 + TypeScript + Element Plus + Pinia
- **AI**: DeepSeek/通义千问 API

## 功能特性

- **成绩管理**: Excel导入、成绩分析、数据导出
- **表现记录**: 文本记录、标签管理
- **评语生成**: AI智能生成、批量导出
- **课件制作**: 教案生成、PPT制作（开发中）

## 快速开始

### 1. 环境准备

- Go 1.21+
- Node.js 18+
- MySQL 8.0+

### 2. 数据库初始化

```sql
CREATE DATABASE teacher_assistant CHARACTER SET utf8mb4;
```

修改 `backend/cmd/main.go` 中的数据库连接字符串。

### 3. 启动后端

```bash
cd backend
go mod download
go run cmd/main.go
```

后端服务将在 http://localhost:8080 启动

### 4. 启动前端

```bash
cd frontend
npm install
npm run dev
```

前端服务将在 http://localhost:5173 启动

## 项目结构

```
teacher-assistant/
├── backend/                # Go后端
│   ├── cmd/               # 启动入口
│   ├── internal/
│   │   ├── handler/       # HTTP处理器
│   │   ├── service/       # 业务逻辑
│   │   ├── model/         # 数据模型
│   │   └── repository/    # 数据访问
│   └── go.mod
├── frontend/              # Vue3前端
│   ├── src/
│   │   ├── views/         # 页面视图
│   │   ├── components/    # 组件
│   │   ├── api/           # API接口
│   │   └── stores/        # 状态管理
│   └── package.json
└── docs/                  # 文档
```

## UI设计

采用**专业精品风 + 暖调背景**设计：
- 主色调: 柔和紫蓝 `#6366f1`
- 背景色: 暖白 `#faf9f7`
- 卡片圆角: 16px
- 精致排版和微妙阴影

## 开发计划

- [x] 项目架构搭建
- [x] 前端UI框架
- [x] 后端API框架
- [ ] Excel导入/导出
- [ ] AI评语生成集成
- [ ] 课件制作功能
- [ ] 小程序语音输入

## License

MIT