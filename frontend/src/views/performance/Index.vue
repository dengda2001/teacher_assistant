<template>
  <div class="performance-page">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-title">
        <h1 class="page-title">表现记录</h1>
        <p class="page-subtitle">记录学生日常表现，生成个性化评语</p>
      </div>
      <el-button type="primary" size="large" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        <span>添加记录</span>
      </el-button>
    </div>

    <el-row :gutter="24">
      <!-- 左侧：记录卡片列表 -->
      <el-col :span="16">
        <div class="records-list">
          <div
            v-for="record in filteredRecords"
            :key="record.id"
            class="record-card"
          >
            <div class="record-main">
              <!-- 头像 -->
              <div
                class="record-avatar"
                :class="{ 'avatar-success': record.type === 'success', 'avatar-warning': record.type === 'warning' }"
              >
                {{ record.studentName.charAt(0) }}
              </div>

              <!-- 内容区 -->
              <div class="record-content-wrapper">
                <div class="record-header">
                  <div class="record-meta">
                    <span class="student-name">{{ record.studentName }}</span>
                    <el-tag
                      :type="record.type === 'success' ? 'success' : record.type === 'warning' ? 'warning' : 'primary'"
                      size="small"
                      class="type-tag"
                    >
                      {{ record.type === 'success' ? '表扬' : record.type === 'warning' ? '提醒' : '记录' }}
                    </el-tag>
                    <span class="record-time">{{ formatDate(record.recordDate) }}</span>
                  </div>
                  <div class="record-actions">
                    <el-button type="primary" link @click="editRecord(record)">
                      <el-icon><Edit /></el-icon>
                    </el-button>
                    <el-button type="danger" link @click="deleteRecord(record)">
                      <el-icon><Delete /></el-icon>
                    </el-button>
                  </div>
                </div>

                <div class="record-content">{{ record.content }}</div>

                <div class="record-tags">
                  <el-tag
                    v-for="tag in record.tags"
                    :key="tag"
                    size="small"
                    effect="light"
                    class="record-tag"
                  >
                    {{ tag }}
                  </el-tag>
                </div>
              </div>
            </div>
          </div>

          <el-empty v-if="filteredRecords.length === 0" description="暂无记录" />
        </div>
      </el-col>

      <!-- 右侧：快捷面板 -->
      <el-col :span="8">
        <!-- 快捷标签 -->
        <div class="side-panel">
          <div class="panel-header">
            <el-icon><CollectionTag /></el-icon>
            <span>快捷标签</span>
          </div>
          <div class="quick-tags">
            <el-tag
              v-for="tag in quickTags"
              :key="tag"
              :effect="selectedTag === tag ? 'dark' : 'plain'"
              class="quick-tag"
              @click="toggleTagFilter(tag)"
            >
              {{ tag }}
            </el-tag>
            <el-tag
              v-if="selectedTag"
              effect="plain"
              class="quick-tag clear-tag"
              @click="clearTagFilter"
            >
              清除筛选
            </el-tag>
          </div>
        </div>

        <!-- 使用说明 -->
        <div class="side-panel">
          <div class="panel-header">
            <el-icon><InfoFilled /></el-icon>
            <span>使用说明</span>
          </div>
          <div class="usage-tips">
            <div class="tip-item">
              <span class="tip-num">1</span>
              <span>点击"添加记录"记录学生日常表现</span>
            </div>
            <div class="tip-item">
              <span class="tip-num">2</span>
              <span>可使用标签分类记录类型</span>
            </div>
            <div class="tip-item">
              <span class="tip-num">3</span>
              <span>表现记录将用于评语生成</span>
            </div>
            <div class="tip-item">
              <span class="tip-num">4</span>
              <span>点击标签可快速筛选相关记录</span>
            </div>
          </div>
        </div>

        <!-- 统计卡片 -->
        <div class="side-panel stats-panel">
          <div class="panel-header">
            <el-icon><TrendCharts /></el-icon>
            <span>数据统计</span>
          </div>
          <div class="stats-grid">
            <div class="stat-item">
              <div class="stat-value">{{ stats.total }}</div>
              <div class="stat-label">总记录</div>
            </div>
            <div class="stat-item">
              <div class="stat-value stat-success">{{ stats.positive }}</div>
              <div class="stat-label">表扬</div>
            </div>
            <div class="stat-item">
              <div class="stat-value stat-warning">{{ stats.warning }}</div>
              <div class="stat-label">提醒</div>
            </div>
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 添加/编辑记录弹窗 -->
    <el-dialog
      v-model="showAddDialog"
      :title="isEdit ? '编辑记录' : '添加表现记录'"
      width="600px"
      destroy-on-close
    >
      <el-form
        :model="recordForm"
        label-width="80px"
        :rules="rules"
        ref="formRef"
      >
        <el-form-item label="学生" prop="studentId">
          <el-select
            v-model="recordForm.studentId"
            placeholder="选择学生"
            style="width: 100%"
          >
            <el-option
              v-for="student in students"
              :key="student.id"
              :label="student.name"
              :value="student.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="recordForm.type">
            <el-radio label="success">表扬</el-radio>
            <el-radio label="warning">提醒</el-radio>
            <el-radio label="primary">普通记录</el-radio>
          </el-radio-group>
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
            <el-option
              v-for="tag in allTags"
              :key="tag"
              :label="tag"
              :value="tag"
            />
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
        <el-button @click="showAddDialog = false">取消</el-button>
        <el-button type="primary" @click="saveRecord">
          {{ isEdit ? '保存' : '添加' }}
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Edit, Delete, CollectionTag, InfoFilled, TrendCharts } from '@element-plus/icons-vue'
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
  type: 'primary' | 'success' | 'warning' | 'danger'
}

// 弹窗控制
const showAddDialog = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

// 标签筛选
const selectedTag = ref('')

// 学生列表
const students = ref<Student[]>([
  { id: '1', name: '张三' },
  { id: '2', name: '李四' },
  { id: '3', name: '王五' },
  { id: '4', name: '赵六' },
  { id: '5', name: '孙七' }
])

// 记录列表
const records = ref<PerformanceRecord[]>([
  {
    id: '1',
    studentId: '1',
    studentName: '张三',
    content: '今天课堂上积极发言，主动回答了老师的提问，表现出色。对数学问题的理解很到位，思路清晰。',
    tags: ['积极发言', '课堂活跃'],
    recordDate: '2026-04-28',
    type: 'success'
  },
  {
    id: '2',
    studentId: '2',
    studentName: '李四',
    content: '作业完成质量有所下降，需要加强课后复习。建议家长督促孩子按时完成作业。',
    tags: ['作业延迟', '需要关注'],
    recordDate: '2026-04-27',
    type: 'warning'
  },
  {
    id: '3',
    studentId: '3',
    studentName: '王五',
    content: '帮助同学解答数学难题，耐心细致，展现了良好的互助精神。',
    tags: ['助人为乐', '数学'],
    recordDate: '2026-04-26',
    type: 'success'
  }
])

// 快捷标签
const quickTags = ['积极发言', '课堂活跃', '作业优秀', '作业延迟', '助人为乐', '遵守纪律', '需要关注', '进步明显']
const allTags = [...quickTags, '态度认真', '思维活跃', '团结同学', '热爱劳动']

// 表单数据
const recordForm = reactive({
  id: '',
  studentId: '',
  type: 'primary' as 'primary' | 'success' | 'warning',
  content: '',
  tags: [] as string[],
  recordDate: new Date().toISOString().split('T')[0]
})

// 表单校验规则
const rules: FormRules = {
  studentId: [{ required: true, message: '请选择学生', trigger: 'change' }],
  type: [{ required: true, message: '请选择记录类型', trigger: 'change' }],
  recordDate: [{ required: true, message: '请选择日期', trigger: 'change' }],
  content: [{ required: true, message: '请输入记录内容', trigger: 'blur' }]
}

// 筛选后的记录
const filteredRecords = computed(() => {
  if (!selectedTag.value) return records.value
  return records.value.filter(record => record.tags.includes(selectedTag.value))
})

// 统计数据
const stats = computed(() => {
  const total = records.value.length
  const positive = records.value.filter(r => r.type === 'success').length
  const warning = records.value.filter(r => r.type === 'warning').length
  return { total, positive, warning }
})

// 格式化日期
const formatDate = (date: string) => {
  const d = new Date(date)
  const today = new Date()
  const yesterday = new Date(today)
  yesterday.setDate(yesterday.getDate() - 1)

  if (date === today.toISOString().split('T')[0]) return '今天'
  if (date === yesterday.toISOString().split('T')[0]) return '昨天'
  return `${d.getMonth() + 1}月${d.getDate()}日`
}

// 切换标签筛选
const toggleTagFilter = (tag: string) => {
  if (selectedTag.value === tag) {
    selectedTag.value = ''
  } else {
    selectedTag.value = tag
  }
}

// 清除标签筛选
const clearTagFilter = () => {
  selectedTag.value = ''
}

// 保存记录
const saveRecord = async () => {
  if (!formRef.value) return
  await formRef.value.validate((valid) => {
    if (valid) {
      const student = students.value.find(s => s.id === recordForm.studentId)
      if (isEdit.value) {
        const index = records.value.findIndex(r => r.id === recordForm.id)
        if (index !== -1) {
          records.value[index] = {
            ...recordForm,
            studentName: student?.name || ''
          }
        }
        ElMessage.success('记录更新成功')
      } else {
        const record: PerformanceRecord = {
          ...recordForm,
          id: Date.now().toString(),
          studentName: student?.name || ''
        }
        records.value.unshift(record)
        ElMessage.success('记录添加成功')
      }
      showAddDialog.value = false
      resetForm()
    }
  })
}

// 编辑记录
const editRecord = (record: PerformanceRecord) => {
  isEdit.value = true
  Object.assign(recordForm, {
    id: record.id,
    studentId: record.studentId,
    type: record.type,
    content: record.content,
    tags: [...record.tags],
    recordDate: record.recordDate
  })
  showAddDialog.value = true
}

// 删除记录
const deleteRecord = (record: PerformanceRecord) => {
  ElMessageBox.confirm(
    '确定要删除这条记录吗？',
    '确认删除',
    { type: 'warning' }
  ).then(() => {
    records.value = records.value.filter(r => r.id !== record.id)
    ElMessage.success('删除成功')
  })
}

// 重置表单
const resetForm = () => {
  isEdit.value = false
  recordForm.id = ''
  recordForm.studentId = ''
  recordForm.type = 'primary'
  recordForm.content = ''
  recordForm.tags = []
  recordForm.recordDate = new Date().toISOString().split('T')[0]
  formRef.value?.resetFields()
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.performance-page {
  padding: $spacing-xl;
}

// 页面头部
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-2xl;

  .header-title {
    .page-title {
      font-size: 24px;
      font-weight: 600;
      color: $text-primary;
      margin: 0 0 $spacing-sm 0;
    }

    .page-subtitle {
      font-size: 14px;
      color: $text-secondary;
      margin: 0;
    }
  }
}

// 记录列表
.records-list {
  display: flex;
  flex-direction: column;
  gap: $spacing-lg;
}

// 记录卡片
.record-card {
  background: $bg-card;
  border-radius: $radius-lg;
  border: 1px solid $border;
  padding: $spacing-lg;
  transition: all $transition-base;

  &:hover {
    box-shadow: $shadow-lg;
    border-color: $primary-light;
  }

  .record-main {
    display: flex;
    gap: $spacing-lg;
  }

  // 头像
  .record-avatar {
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
    flex-shrink: 0;

    &.avatar-success {
      background: $success-light;
      color: $success-text;
    }

    &.avatar-warning {
      background: $warning-light;
      color: $warning-text;
    }
  }

  // 内容区
  .record-content-wrapper {
    flex: 1;
    min-width: 0;
  }

  .record-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: $spacing-md;

    .record-meta {
      display: flex;
      align-items: center;
      gap: $spacing-sm;

      .student-name {
        font-weight: 600;
        font-size: 15px;
        color: $text-primary;
      }

      .type-tag {
        font-size: 12px;
      }

      .record-time {
        font-size: 13px;
        color: $text-muted;
      }
    }

    .record-actions {
      display: flex;
      gap: $spacing-xs;

      .el-button {
        padding: $spacing-xs;
      }
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

    .record-tag {
      font-size: 12px;
    }
  }
}

// 右侧面板
.side-panel {
  background: $bg-card;
  border-radius: $radius-lg;
  border: 1px solid $border;
  padding: $spacing-lg;
  margin-bottom: $spacing-lg;

  .panel-header {
    display: flex;
    align-items: center;
    gap: $spacing-sm;
    font-weight: 600;
    font-size: 15px;
    color: $text-primary;
    margin-bottom: $spacing-lg;

    .el-icon {
      color: $primary;
    }
  }
}

// 快捷标签
.quick-tags {
  display: flex;
  flex-wrap: wrap;
  gap: $spacing-sm;

  .quick-tag {
    cursor: pointer;
    transition: all $transition-fast;

    &:hover {
      transform: translateY(-1px);
    }

    &.clear-tag {
      color: $text-muted;
      border-style: dashed;
    }
  }
}

// 使用说明
.usage-tips {
  display: flex;
  flex-direction: column;
  gap: $spacing-md;

  .tip-item {
    display: flex;
    align-items: flex-start;
    gap: $spacing-sm;
    font-size: 13px;
    color: $text-secondary;
    line-height: 1.5;

    .tip-num {
      width: 20px;
      height: 20px;
      border-radius: $radius-full;
      background: $primary-lighter;
      color: $primary;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 12px;
      font-weight: 600;
      flex-shrink: 0;
    }
  }
}

// 统计面板
.stats-panel {
  .stats-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: $spacing-md;
  }

  .stat-item {
    text-align: center;
    padding: $spacing-md;
    background: $bg-primary;
    border-radius: $radius-md;

    .stat-value {
      font-size: 24px;
      font-weight: 700;
      color: $text-primary;
      line-height: 1.2;

      &.stat-success {
        color: $success;
      }

      &.stat-warning {
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
</style>
