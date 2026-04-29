<template>
  <div>
    <!-- 页面标题区域 -->
    <div class="page-header">
      <div>
        <h1 class="page-title">
          成绩管理
          <span class="subtitle">管理学生各科考试成绩</span>
        </h1>
      </div>
      <div class="actions">
        <el-button class="btn-secondary" @click="showImportDialog = true">
          <el-icon><Upload /></el-icon>
          <span>导入成绩</span>
        </el-button>
        <el-button type="primary" class="btn-primary" @click="showAddDialog = true">
          <el-icon><Plus /></el-icon>
          <span>录入成绩</span>
        </el-button>
      </div>
    </div>

    <!-- 主内容卡片 -->
    <el-card class="content-card" shadow="never">
      <!-- 过滤栏 -->
      <div class="filter-bar">
        <div class="filter-items">
          <el-select v-model="filterExam" placeholder="选择考试" clearable class="filter-select">
            <el-option
              v-for="exam in examOptions"
              :key="exam.value"
              :label="exam.label"
              :value="exam.value"
            />
          </el-select>
          <el-select v-model="filterSubject" placeholder="选择科目" clearable class="filter-select">
            <el-option
              v-for="subject in subjectOptions"
              :key="subject.value"
              :label="subject.label"
              :value="subject.value"
            />
          </el-select>
        </div>
        <span class="record-count">共 {{ filteredScores.length }} 条记录</span>
      </div>

      <!-- 成绩表格 -->
      <el-table :data="filteredScores" v-loading="loading" style="width: 100%" stripe>
        <!-- 学生信息列 -->
        <el-table-column label="学生" min-width="140">
          <template #default="{ row }">
            <div class="student-info-cell">
              <div class="avatar-circle">{{ row.studentName.charAt(0) }}</div>
              <span class="student-name">{{ row.studentName }}</span>
            </div>
          </template>
        </el-table-column>

        <!-- 语文成绩 -->
        <el-table-column label="语文" min-width="100">
          <template #default="{ row }">
            <span :class="['score-tag', getScoreClass(row.chinese)]">
              {{ row.chinese }}
            </span>
          </template>
        </el-table-column>

        <!-- 数学成绩 -->
        <el-table-column label="数学" min-width="100">
          <template #default="{ row }">
            <span :class="['score-tag', getScoreClass(row.math)]">
              {{ row.math }}
            </span>
          </template>
        </el-table-column>

        <!-- 英语成绩 -->
        <el-table-column label="英语" min-width="100">
          <template #default="{ row }">
            <span :class="['score-tag', getScoreClass(row.english)]">
              {{ row.english }}
            </span>
          </template>
        </el-table-column>

        <!-- 总分 -->
        <el-table-column label="总分" min-width="100" sortable>
          <template #default="{ row }">
            <span class="total-score">{{ row.total }}</span>
          </template>
        </el-table-column>

        <!-- 平均分 -->
        <el-table-column label="平均分" min-width="100" sortable>
          <template #default="{ row }">
            <span class="average-score">{{ row.average.toFixed(1) }}</span>
          </template>
        </el-table-column>

        <!-- 操作列 -->
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

    <!-- 导入成绩弹窗 -->
    <el-dialog
      v-model="showImportDialog"
      title="导入成绩"
      width="600px"
      class="custom-dialog"
      destroy-on-close
    >
      <p class="dialog-subtitle">支持 Excel 文件导入，请确保格式正确</p>
      <ExcelUploader @success="handleImportSuccess" />
    </el-dialog>

    <!-- 录入/编辑成绩弹窗 -->
    <el-dialog
      v-model="showAddDialog"
      :title="isEdit ? '编辑成绩' : '录入成绩'"
      width="500px"
      class="custom-dialog"
      destroy-on-close
    >
      <p class="dialog-subtitle">{{ isEdit ? '修改学生成绩信息' : '请输入学生的各科成绩' }}</p>
      <el-form :model="scoreForm" label-width="80px" :rules="rules" ref="formRef" class="custom-form">
        <el-form-item label="学生" prop="studentName">
          <el-input v-model="scoreForm.studentName" placeholder="请输入学生姓名" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="考试" prop="examName">
          <el-select v-model="scoreForm.examName" placeholder="请选择考试" style="width: 100%">
            <el-option
              v-for="exam in examOptions"
              :key="exam.value"
              :label="exam.label"
              :value="exam.value"
            />
          </el-select>
        </el-form-item>
        <el-row :gutter="16">
          <el-col :span="8">
            <el-form-item label="语文" prop="chinese">
              <el-input-number v-model="scoreForm.chinese" :min="0" :max="150" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="数学" prop="math">
              <el-input-number v-model="scoreForm.math" :min="0" :max="150" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="英语" prop="english">
              <el-input-number v-model="scoreForm.english" :min="0" :max="150" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showAddDialog = false" class="btn-secondary">取消</el-button>
          <el-button type="primary" @click="saveScore" class="btn-primary">
            {{ isEdit ? '保存' : '确认录入' }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Upload, Edit, Delete } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'
import ExcelUploader from '@/components/ExcelUploader.vue'

interface ScoreRecord {
  id: string
  studentName: string
  examName: string
  chinese: number
  math: number
  english: number
  total: number
  average: number
}

const loading = ref(false)
const showImportDialog = ref(false)
const showAddDialog = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

// 过滤条件
const filterExam = ref('')
const filterSubject = ref('')

// 选项数据
const examOptions = [
  { label: '期中考试', value: '期中考试' },
  { label: '期末考试', value: '期末考试' },
  { label: '月考一', value: '月考一' },
  { label: '月考二', value: '月考二' }
]

const subjectOptions = [
  { label: '语文', value: 'chinese' },
  { label: '数学', value: 'math' },
  { label: '英语', value: 'english' }
]

// 模拟成绩数据
const scores = ref<ScoreRecord[]>([
  { id: '1', studentName: '王小明', examName: '期中考试', chinese: 95, math: 88, english: 92, total: 275, average: 91.7 },
  { id: '2', studentName: '李小红', examName: '期中考试', chinese: 88, math: 95, english: 85, total: 268, average: 89.3 },
  { id: '3', studentName: '张三', examName: '期中考试', chinese: 78, math: 82, english: 76, total: 236, average: 78.7 },
  { id: '4', studentName: '赵四', examName: '期中考试', chinese: 65, math: 70, english: 68, total: 203, average: 67.7 },
  { id: '5', studentName: '刘五', examName: '期中考试', chinese: 92, math: 85, english: 90, total: 267, average: 89.0 },
  { id: '6', studentName: '陈六', examName: '期中考试', chinese: 55, math: 45, english: 58, total: 158, average: 52.7 },
  { id: '7', studentName: '周七', examName: '期中考试', chinese: 82, math: 78, english: 80, total: 240, average: 80.0 },
  { id: '8', studentName: '吴八', examName: '期中考试', chinese: 76, math: 88, english: 72, total: 236, average: 78.7 }
])

// 表单数据
const scoreForm = reactive({
  id: '',
  studentName: '',
  examName: '',
  chinese: 0,
  math: 0,
  english: 0
})

const rules: FormRules = {
  studentName: [{ required: true, message: '请输入学生姓名', trigger: 'blur' }],
  examName: [{ required: true, message: '请选择考试', trigger: 'change' }],
  chinese: [{ required: true, message: '请输入语文成绩', trigger: 'blur' }],
  math: [{ required: true, message: '请输入数学成绩', trigger: 'blur' }],
  english: [{ required: true, message: '请输入英语成绩', trigger: 'blur' }]
}

// 过滤后的成绩数据
const filteredScores = computed(() => {
  let result = scores.value
  if (filterExam.value) {
    result = result.filter(s => s.examName === filterExam.value)
  }
  return result
})

// 获取成绩颜色样式类
const getScoreClass = (score: number): string => {
  if (score >= 90) return 'score-excellent'    // >90 绿色
  if (score >= 75) return 'score-good'         // >75 蓝色
  if (score >= 60) return 'score-pass'         // >60 橙色
  return 'score-fail'                          // <60 红色
}

// 保存成绩
const saveScore = async () => {
  if (!formRef.value) return
  await formRef.value.validate((valid) => {
    if (valid) {
      const total = scoreForm.chinese + scoreForm.math + scoreForm.english
      const average = total / 3

      if (isEdit.value) {
        const index = scores.value.findIndex(s => s.id === scoreForm.id)
        if (index !== -1) {
          scores.value[index] = {
            ...scoreForm,
            total,
            average
          }
        }
        ElMessage.success('成绩更新成功')
      } else {
        scores.value.push({
          ...scoreForm,
          id: Date.now().toString(),
          total,
          average
        })
        ElMessage.success('成绩录入成功')
      }
      showAddDialog.value = false
      resetForm()
    }
  })
}

// 编辑成绩
const editScore = (record: ScoreRecord) => {
  isEdit.value = true
  Object.assign(scoreForm, {
    id: record.id,
    studentName: record.studentName,
    examName: record.examName,
    chinese: record.chinese,
    math: record.math,
    english: record.english
  })
  showAddDialog.value = true
}

// 删除成绩
const deleteScore = (record: ScoreRecord) => {
  ElMessageBox.confirm(
    `确定要删除 "${record.studentName}" 的成绩记录吗？`,
    '确认删除',
    { type: 'warning' }
  ).then(() => {
    scores.value = scores.value.filter(s => s.id !== record.id)
    ElMessage.success('删除成功')
  })
}

// 导入成功回调
const handleImportSuccess = (data: any) => {
  ElMessage.success(`成功导入 ${data.length} 条成绩记录`)
  showImportDialog.value = false
}

// 重置表单
const resetForm = () => {
  isEdit.value = false
  Object.assign(scoreForm, {
    id: '',
    studentName: '',
    examName: '',
    chinese: 0,
    math: 0,
    english: 0
  })
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

// 页面标题样式
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
      font-weight: 400;
      color: $text-secondary;
      margin-left: $spacing-md;
    }
  }

  .actions {
    display: flex;
    gap: $spacing-md;
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

// 内容卡片
.content-card {
  border-radius: $radius-lg;
  border: 1px solid $border;

  :deep(.el-card__body) {
    padding: $spacing-xl;
  }
}

// 过滤栏
.filter-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-lg;
  padding-bottom: $spacing-lg;
  border-bottom: 1px solid $border-light;

  .filter-items {
    display: flex;
    gap: $spacing-md;
  }

  .filter-select {
    width: 160px;

    :deep(.el-input__wrapper) {
      border-radius: $radius-md;
    }
  }

  .record-count {
    font-size: 14px;
    color: $text-secondary;
  }
}

// 学生信息单元格
.student-info-cell {
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

  .student-name {
    font-size: 14px;
    color: $text-primary;
    font-weight: 500;
  }
}

// 成绩标签样式
.score-tag {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 4px 12px;
  border-radius: $radius-md;
  font-size: 14px;
  font-weight: 600;
  min-width: 48px;

  &.score-excellent {
    background: $success-light;
    color: $success-text;
  }

  &.score-good {
    background: $info-light;
    color: $info-text;
  }

  &.score-pass {
    background: $warning-light;
    color: $warning-text;
  }

  &.score-fail {
    background: $danger-light;
    color: $danger-text;
  }
}

// 总分和平均分样式
.total-score {
  font-size: 14px;
  font-weight: 600;
  color: $text-primary;
}

.average-score {
  font-size: 14px;
  font-weight: 500;
  color: $primary;
}

// 操作按钮
.action-btns {
  display: flex;
  gap: $spacing-xs;

  .el-button {
    font-size: 16px;
  }
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
  .el-select .el-input__wrapper,
  .el-input-number .el-input__wrapper {
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
