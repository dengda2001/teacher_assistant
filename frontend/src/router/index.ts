import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/views/Dashboard.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
      meta: { title: '仪表盘', icon: 'HomeFilled' }
    },
    {
      path: '/students',
      name: 'students',
      component: () => import('@/views/students/Index.vue'),
      meta: { title: '学生管理', icon: 'UserFilled' }
    },
    {
      path: '/scores',
      name: 'scores',
      component: () => import('@/views/scores/Index.vue'),
      meta: { title: '成绩管理', icon: 'TrendCharts' }
    },
    {
      path: '/scores/analysis',
      name: 'score-analysis',
      component: () => import('@/views/scores/Analysis.vue'),
      meta: { title: '成绩分析', icon: 'DataAnalysis' }
    },
    {
      path: '/performance',
      name: 'performance',
      component: () => import('@/views/performance/Index.vue'),
      meta: { title: '表现记录', icon: 'EditPen' }
    },
    {
      path: '/comments',
      name: 'comments',
      component: () => import('@/views/comments/Index.vue'),
      meta: { title: '评语生成', icon: 'ChatDotRound' }
    }
  ]
})

export default router