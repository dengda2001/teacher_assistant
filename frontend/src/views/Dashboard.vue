<template>
  <div class="dashboard">
    <!-- Welcome Card -->
    <el-card class="welcome-card" shadow="never">
      <div class="welcome-content">
        <div class="welcome-left">
          <h1 class="welcome-title">{{ greeting }}，张老师</h1>
          <p class="welcome-date">{{ currentDate }}</p>
        </div>
        <div class="welcome-right">
          <div class="todo-badge">
            <el-icon><Bell /></el-icon>
            <span>今日待办事项：{{ todoCount }} 项</span>
          </div>
        </div>
      </div>
    </el-card>

    <!-- Stats Grid -->
    <el-row :gutter="16" class="stats-row">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card" shadow="never">
          <div class="stat-icon-wrapper blue">
            <el-icon><User /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.studentCount }}</div>
            <div class="stat-label">学生总数</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card" shadow="never">
          <div class="stat-icon-wrapper green">
            <el-icon><DocumentChecked /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.performanceCount }}</div>
            <div class="stat-label">表现记录</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card" shadow="never">
          <div class="stat-icon-wrapper orange">
            <el-icon><ChatDotRound /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.pendingComments }}</div>
            <div class="stat-label">待生成评语</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card" shadow="never">
          <div class="stat-icon-wrapper purple">
            <el-icon><TrendCharts /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.classAverage }}</div>
            <div class="stat-label">班级平均分</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Quick Actions -->
    <el-row :gutter="16" class="quick-actions-row">
      <el-col :span="24">
        <el-card class="quick-actions-card" shadow="never">
          <template #header>
            <div class="card-header">快捷入口</div>
          </template>
          <div class="quick-actions-grid">
            <div class="quick-action-item" @click="$router.push('/students')">
              <div class="quick-action-icon blue">
                <el-icon><User /></el-icon>
              </div>
              <span class="quick-action-label">学生管理</span>
            </div>
            <div class="quick-action-item" @click="$router.push('/performance')">
              <div class="quick-action-icon green">
                <el-icon><EditPen /></el-icon>
              </div>
              <span class="quick-action-label">添加记录</span>
            </div>
            <div class="quick-action-item" @click="$router.push('/comments')">
              <div class="quick-action-icon orange">
                <el-icon><ChatDotRound /></el-icon>
              </div>
              <span class="quick-action-label">生成评语</span>
            </div>
            <div class="quick-action-item" @click="$router.push('/scores')">
              <div class="quick-action-icon purple">
                <el-icon><TrendCharts /></el-icon>
              </div>
              <span class="quick-action-label">成绩管理</span>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Recent Activity Timeline -->
    <el-row :gutter="16" class="activity-row">
      <el-col :span="24">
        <el-card class="activity-card" shadow="never">
          <template #header>
            <div class="card-header">
              <span>最近动态</span>
              <el-button link type="primary" size="small" @click="viewAllActivities">
                查看全部
              </el-button>
            </div>
          </template>
          <div class="timeline-list">
            <div
              v-for="(activity, index) in recentActivities"
              :key="activity.id"
              class="timeline-item"
              :class="{ 'last': index === recentActivities.length - 1 }"
            >
              <div class="timeline-marker" :class="activity.type"></div>
              <div class="timeline-content">
                <div class="timeline-title">{{ activity.title }}</div>
                <div class="timeline-desc">{{ activity.description }}</div>
                <div class="timeline-time">{{ activity.time }}</div>
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
  User,
  DocumentChecked,
  ChatDotRound,
  TrendCharts,
  EditPen,
  Bell
} from '@element-plus/icons-vue'

// Greeting based on time of day
const greeting = computed(() => {
  const hour = new Date().getHours()
  if (hour < 12) return '早上好'
  if (hour < 18) return '下午好'
  return '晚上好'
})

// Current date in Chinese format
const currentDate = computed(() => {
  const date = new Date()
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const weekdays = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六']
  const weekday = weekdays[date.getDay()]
  return `${year}年${month}月${day}日 ${weekday}`
})

// Todo count
const todoCount = ref(3)

// Stats data
const stats = ref({
  studentCount: 48,
  performanceCount: 156,
  pendingComments: 12,
  classAverage: 82.5
})

// Recent activities data
const recentActivities = ref([
  {
    id: 1,
    type: 'success',
    title: '导入期中考试成绩',
    description: '成功导入数学、语文、英语三科成绩，共48条记录',
    time: '2小时前'
  },
  {
    id: 2,
    type: 'primary',
    title: '生成学生评语',
    description: '为15名学生生成了期末评语',
    time: '5小时前'
  },
  {
    id: 3,
    type: 'warning',
    title: '添加学生表现记录',
    description: '记录了张小明同学的课堂表现情况',
    time: '1天前'
  },
  {
    id: 4,
    type: 'info',
    title: '导出成绩报告',
    description: '导出2024年春季学期期中成绩报告',
    time: '2天前'
  },
  {
    id: 5,
    type: 'success',
    title: '新增学生',
    description: '添加了3名新转入学生信息',
    time: '3天前'
  }
])

// View all activities handler
const viewAllActivities = () => {
  console.log('View all activities')
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.dashboard {
  max-width: 1400px;
}

// Welcome Card
.welcome-card {
  margin-bottom: 16px;
  background: linear-gradient(135deg, var(--primary) 0%, var(--primary-light) 100%);
  border: none;

  :deep(.el-card__body) {
    padding: 24px;
  }

  .welcome-content {
    display: flex;
    justify-content: space-between;
    align-items: center;

    @media (max-width: 768px) {
      flex-direction: column;
      gap: 16px;
      text-align: center;
    }
  }

  .welcome-title {
    font-size: 24px;
    font-weight: 600;
    color: white;
    margin: 0 0 8px 0;
  }

  .welcome-date {
    font-size: 14px;
    color: rgba(255, 255, 255, 0.8);
    margin: 0;
  }

  .todo-badge {
    display: flex;
    align-items: center;
    gap: 8px;
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    padding: 10px 16px;
    border-radius: 20px;
    font-size: 14px;
    color: white;

    .el-icon {
      font-size: 16px;
    }
  }
}

// Stats Row
.stats-row {
  margin-bottom: 16px;

  .el-col {
    margin-bottom: 16px;
  }
}

.stat-card {
  :deep(.el-card__body) {
    display: flex;
    align-items: center;
    gap: 16px;
    padding: 20px;
  }

  .stat-icon-wrapper {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;

    .el-icon {
      font-size: 24px;
      color: white;
    }

    &.blue {
      background: linear-gradient(135deg, #3b82f6 0%, #60a5fa 100%);
    }

    &.green {
      background: linear-gradient(135deg, #10b981 0%, #34d399 100%);
    }

    &.orange {
      background: linear-gradient(135deg, #f59e0b 0%, #fbbf24 100%);
    }

    &.purple {
      background: linear-gradient(135deg, #8b5cf6 0%, #a78bfa 100%);
    }
  }

  .stat-info {
    flex: 1;
  }

  .stat-value {
    font-size: 28px;
    font-weight: 700;
    color: var(--foreground);
    line-height: 1.2;
    margin-bottom: 4px;
  }

  .stat-label {
    font-size: 14px;
    color: var(--muted-foreground);
  }
}

// Quick Actions
.quick-actions-row {
  margin-bottom: 16px;
}

.quick-actions-card {
  :deep(.el-card__header) {
    padding: 16px 20px;
    border-bottom: 1px solid var(--border);
  }

  :deep(.el-card__body) {
    padding: 20px;
  }

  .card-header {
    font-size: 16px;
    font-weight: 600;
    color: var(--foreground);
  }

  .quick-actions-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 16px;

    @media (max-width: 768px) {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  .quick-action-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
    padding: 24px 16px;
    border-radius: 12px;
    cursor: pointer;
    transition: all 0.2s ease;
    background: var(--muted);

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }
  }

  .quick-action-icon {
    width: 56px;
    height: 56px;
    border-radius: 16px;
    display: flex;
    align-items: center;
    justify-content: center;

    .el-icon {
      font-size: 28px;
      color: white;
    }

    &.blue {
      background: linear-gradient(135deg, #3b82f6 0%, #60a5fa 100%);
    }

    &.green {
      background: linear-gradient(135deg, #10b981 0%, #34d399 100%);
    }

    &.orange {
      background: linear-gradient(135deg, #f59e0b 0%, #fbbf24 100%);
    }

    &.purple {
      background: linear-gradient(135deg, #8b5cf6 0%, #a78bfa 100%);
    }
  }

  .quick-action-label {
    font-size: 14px;
    font-weight: 500;
    color: var(--foreground);
  }
}

// Activity Timeline
.activity-card {
  :deep(.el-card__header) {
    padding: 16px 20px;
    border-bottom: 1px solid var(--border);
  }

  :deep(.el-card__body) {
    padding: 20px;
  }

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 16px;
    font-weight: 600;
    color: var(--foreground);
  }
}

.timeline-list {
  position: relative;
}

.timeline-item {
  position: relative;
  padding-left: 28px;
  padding-bottom: 24px;

  &.last {
    padding-bottom: 0;
  }

  &::before {
    content: '';
    position: absolute;
    left: 7px;
    top: 12px;
    bottom: 0;
    width: 2px;
    background: var(--border);
  }

  &.last::before {
    display: none;
  }
}

.timeline-marker {
  position: absolute;
  left: 0;
  top: 4px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  border: 3px solid white;
  box-shadow: 0 0 0 2px var(--border);

  &.success {
    background: #10b981;
    box-shadow: 0 0 0 2px #10b981;
  }

  &.primary {
    background: #3b82f6;
    box-shadow: 0 0 0 2px #3b82f6;
  }

  &.warning {
    background: #f59e0b;
    box-shadow: 0 0 0 2px #f59e0b;
  }

  &.info {
    background: #6b7280;
    box-shadow: 0 0 0 2px #6b7280;
  }
}

.timeline-content {
  .timeline-title {
    font-size: 14px;
    font-weight: 600;
    color: var(--foreground);
    margin-bottom: 4px;
  }

  .timeline-desc {
    font-size: 13px;
    color: var(--muted-foreground);
    margin-bottom: 4px;
    line-height: 1.5;
  }

  .timeline-time {
    font-size: 12px;
    color: var(--muted-foreground);
  }
}
</style>
