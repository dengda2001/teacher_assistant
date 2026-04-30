# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

教师助手 (Teacher Assistant) - An intelligent teaching aid tool for middle school teachers to manage grades, record student performance, and generate AI-powered comments.

## Tech Stack

- **Backend**: Go 1.21+ + Gin + GORM + SQLite (default) / MySQL
- **Frontend**: Vue 3 + TypeScript + Element Plus + Pinia + Vite
- **Dashboard**: Next.js 15 (experimental admin dashboard in `/dashboard`)
- **Design**: SCSS with warm color palette (#faf9f7 background, #6366f1 primary)

## Development Commands

### Backend (Go)
```bash
cd backend
go mod download          # Install dependencies
go run cmd/main.go       # Start dev server on :8080
go run cmd/simple.go     # Alternative simple server
```

### Frontend (Vue)
```bash
cd frontend
npm install              # Install dependencies
npm run dev              # Start dev server on :5173
npm run build            # Production build
npm run type-check       # TypeScript checking
```

### Dashboard (Next.js)
```bash
cd dashboard
npm install
npm run dev              # Start on :3000
npm run build
npm run lint
```

## Architecture

### Backend Structure (Layered Architecture)
```
backend/
├── cmd/
│   ├── main.go           # Production entry (SQLite default)
│   └── simple.go         # Minimal test server
├── internal/
│   ├── handler/          # HTTP handlers (Gin handlers)
│   ├── service/          # Business logic layer
│   ├── repository/       # DB initialization
│   └── model/            # GORM models + migration
```

**Key Models** (`internal/model/`):
- `Student` - Student info with associations
- `ExamScore` - Grade records per subject
- `PerformanceRecord` - Behavioral/performance notes
- `Comment` - Generated teacher comments

**API Routes** (`/api/v1/`):
- `/students` - CRUD for students
- `/scores` - Grades + import/export/analysis
- `/records` - Performance records
- `/comments` - AI generation + batch export (Word/Excel)

### Frontend Structure
```
frontend/src/
├── views/                # Page components
│   ├── Dashboard.vue
│   ├── students/Index.vue
│   ├── scores/Index.vue & Analysis.vue
│   ├── performance/Index.vue
│   └── comments/Index.vue
├── components/           # Shared components
│   ├── Layout.vue        # Sidebar + header layout
│   ├── ExcelUploader.vue
│   └── CommentPreview.vue
├── router/index.ts       # Vue Router config
├── styles/               # SCSS variables + global
└── api/                  # API client modules
```

**Important Config**:
- Vite proxy: `/api` → `http://localhost:8080`
- Path alias: `@/` maps to `src/`
- Port: 5173 (frontend), 8080 (backend)

## Database

**Default**: SQLite (`teacher_assistant.db` in backend dir)
- Zero-config for development
- Auto-migration on startup via `model.AutoMigrate()`

**MySQL**: Edit `main.go` and switch to `repository.InitMySQL()`
```sql
CREATE DATABASE teacher_assistant CHARACTER SET utf8mb4;
```

## Key Design Patterns

1. **Backend Handlers**: Follow `NewXxxHandler(service) -> handler.List/Create/Update/Delete` pattern
2. **Frontend Views**: Use Element Plus components + SCSS with CSS variables
3. **API Integration**: Frontend uses proxy (vite.config.ts) to reach backend
4. **File Upload**: Excel import via `ExcelUploader.vue` → `/scores/import`

## Design System

From `frontend/src/styles/variables.scss`:
- Background: `#faf9f7` (warm white)
- Primary: `#6366f1` (indigo)
- Card radius: `16px`
- Use `text-wrap: pretty` for CJK text

## Notes

- `/design-demos/` contains HTML prototype pages for UI design exploration (not production code)
- `/pictures/` contains UI mockup images
- Backend uses GORM DeletedAt for soft deletes
- Frontend icons from `@element-plus/icons-vue`
