<template>
  <div>
    <div class="page-header">
      <h1 class="page-title">成绩分析</h1>
      <el-radio-group v-model="analysisType">
        <el-radio-button label="overview">总览</el-radio-button>
        <el-radio-button label="trend">趋势</el-radio-button>
        <el-radio-button label="subject">学科对比</el-radio-button>
      </el-radio-group>
    </div>

    <el-row :gutter="24">
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-title">平均分</div>
          <div class="stat-value">{{ stats.average }}</div>
          <div class="stat-change positive">
            <el-icon><ArrowUp /></el-icon>
            <span>比上次 +2.5</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-title">最高分</div>
          <div class="stat-value">{{ stats.max }}</div>
          <div class="stat-label">张同学</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-title">及格率</div>
          <div class="stat-value">{{ stats.passRate }}%</div>
          <div class="stat-change positive">
            <el-icon><ArrowUp /></el-icon>
            <span>比上次 +5%</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stat-card">
          <div class="stat-title">优秀率</div>
          <div class="stat-value">{{ stats.excellentRate }}%</div>
          <div class="stat-change negative">
            <el-icon><ArrowDown /></el-icon>
            <span>比上次 -2%</span>
          </div>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="24" class="mt-6">
      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>分数段分布</span>
            </div>
          </template>
          <div ref="scoreDistributionChart" class="chart-container"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>学科对比雷达图</span>
            </div>
          </template>
          <div ref="radarChart" class="chart-container"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="24" class="mt-6">
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>成绩趋势</span>
            </div>
          </template>
          <div ref="trendChart" class="chart-container"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { ArrowUp, ArrowDown } from '@element-plus/icons-vue'
import * as echarts from 'echarts'

const analysisType = ref('overview')
const scoreDistributionChart = ref<HTMLElement>()
const radarChart = ref<HTMLElement>()
const trendChart = ref<HTMLElement>()

const stats = ref({
  average: '82.5',
  max: '98',
  passRate: '92',
  excellentRate: '35'
})

let chartInstances: echarts.ECharts[] = []

onMounted(() => {
  initScoreDistributionChart()
  initRadarChart()
  initTrendChart()

  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  chartInstances.forEach(chart => chart.dispose())
})

const handleResize = () => {
  chartInstances.forEach(chart => chart.resize())
}

const initScoreDistributionChart = () => {
  if (!scoreDistributionChart.value) return
  const chart = echarts.init(scoreDistributionChart.value)
  chartInstances.push(chart)

  const option = {
    tooltip: { trigger: 'axis' },
    xAxis: {
      type: 'category',
      data: ['<60', '60-70', '70-80', '80-90', '90-100']
    },
    yAxis: { type: 'value' },
    series: [{
      data: [3, 8, 15, 20, 4],
      type: 'bar',
      itemStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          { offset: 0, color: '#6366f1' },
          { offset: 1, color: '#8b5cf6' }
        ])
      },
      borderRadius: [8, 8, 0, 0]
    }]
  }
  chart.setOption(option)
}

const initRadarChart = () => {
  if (!radarChart.value) return
  const chart = echarts.init(radarChart.value)
  chartInstances.push(chart)

  const option = {
    tooltip: {},
    radar: {
      indicator: [
        { name: '语文', max: 100 },
        { name: '数学', max: 100 },
        { name: '英语', max: 100 },
        { name: '物理', max: 100 },
        { name: '化学', max: 100 }
      ]
    },
    series: [{
      type: 'radar',
      data: [
        {
          value: [85, 90, 78, 88, 82],
          name: '班级平均',
          areaStyle: {
            color: 'rgba(99, 102, 241, 0.2)'
          },
          lineStyle: {
            color: '#6366f1'
          }
        }
      ]
    }]
  }
  chart.setOption(option)
}

const initTrendChart = () => {
  if (!trendChart.value) return
  const chart = echarts.init(trendChart.value)
  chartInstances.push(chart)

  const option = {
    tooltip: { trigger: 'axis' },
    legend: { data: ['语文', '数学', '英语'] },
    xAxis: {
      type: 'category',
      data: ['第一次月考', '期中考试', '第二次月考', '期末考试']
    },
    yAxis: { type: 'value', min: 60, max: 100 },
    series: [
      {
        name: '语文',
        type: 'line',
        smooth: true,
        data: [82, 85, 83, 88],
        lineStyle: { color: '#6366f1' }
      },
      {
        name: '数学',
        type: 'line',
        smooth: true,
        data: [78, 82, 85, 90],
        lineStyle: { color: '#10b981' }
      },
      {
        name: '英语',
        type: 'line',
        smooth: true,
        data: [80, 78, 82, 85],
        lineStyle: { color: '#f59e0b' }
      }
    ]
  }
  chart.setOption(option)
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-xl;
}

.stat-card {
  background: $bg-card;
  border-radius: $radius-lg;
  padding: $spacing-xl;
  box-shadow: $shadow-sm;

  .stat-title {
    font-size: 13px;
    color: $text-secondary;
    margin-bottom: $spacing-sm;
  }

  .stat-value {
    font-size: 32px;
    font-weight: 600;
    color: $text-primary;
    margin-bottom: $spacing-sm;
  }

  .stat-label {
    font-size: 13px;
    color: $text-muted;
  }

  .stat-change {
    font-size: 13px;
    display: flex;
    align-items: center;
    gap: $spacing-xs;

    &.positive {
      color: $success;
    }

    &.negative {
      color: $danger;
    }
  }
}

.mt-6 {
  margin-top: $spacing-2xl;
}

.card-header {
  font-weight: 600;
}

.chart-container {
  height: 300px;
}
</style>
