<template>
  <div class="comments-page">
    <!-- 页面标题区 -->
    <div class="page-header">
      <div class="header-title">
        <h1 class="page-title">评语生成</h1>
        <p class="page-subtitle">基于学生成绩和表现记录，智能生成个性化评语</p>
      </div>
      <div class="header-actions">
        <el-button @click="showSettings = true">
          <el-icon><Setting /></el-icon>
          生成设置
        </el-button>
        <el-button type="primary" @click="generateComments" :loading="generating">
          <el-icon><MagicStick /></el-icon>
          AI生成评语
        </el-button>
      </div>
    </div>

    <!-- 统计栏 -->
    <div class="stats-bar">
      <div class="stat-item">
        <div class="stat-value">{{ students.length }}</div>
        <div class="stat-label">总学生数</div>
      </div>
      <div class="stat-divider" />
      <div class="stat-item">
        <div class="stat-value generated">{{ generatedCount }}</div>
        <div class="stat-label">已生成</div>
      </div>
      <div class="stat-divider" />
      <div class="stat-item">
        <div class="stat-value pending">{{ pendingCount }}</div>
        <div class="stat-label">待生成</div>
      </div>
    </div>

    <!-- 学生表格 -->
    <el-card class="students-card">
      <el-table
        :data="students"
        v-loading="generating"
        stripe
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="学生" min-width="140">
          <template #default="{ row }">
            <div class="student-info">
              <el-avatar :size="36" :icon="UserFilled" class="student-avatar" />
              <span class="student-name">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="学科成绩" min-width="200">
          <template #default="{ row }">
            <div class="score-tags">
              <el-tag
                v-for="score in row.recentScores"
                :key="score.subject"
                size="small"
                :type="getSubjectTagType(score.subject)"
                class="subject-tag"
              >
                {{ score.subject }} {{ score.score }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="表现记录" min-width="100">
          <template #default="{ row }">
            <el-tag size="small" type="info" effect="plain">
              {{ row.recordCount }} 条记录
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="评语预览" min-width="280">
          <template #default="{ row }">
            <div v-if="row.comment" class="comment-preview-text" @click="openPreview(row)">
              {{ row.comment }}
            </div>
            <span v-else class="no-comment">-</span>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag
              :type="row.comment ? 'success' : 'warning'"
              size="small"
              effect="light"
            >
              <el-icon v-if="row.comment"><Check /></el-icon>
              <el-icon v-else><Clock /></el-icon>
              {{ row.comment ? '已生成' : '待生成' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button
              type="primary"
              link
              size="small"
              @click="openPreview(row)"
            >
              {{ row.comment ? '查看' : '生成' }}
            </el-button>
            <el-button
              v-if="row.comment"
              type="danger"
              link
              size="small"
              @click="clearComment(row)"
            >
              清除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 生成设置弹窗 -->
    <el-dialog
      v-model="showSettings"
      title="生成设置"
      width="500px"
      :close-on-click-modal="false"
    >
      <el-form :model="settings" label-width="100px" class="settings-form">
        <el-form-item label="评语风格">
          <el-radio-group v-model="settings.style">
            <el-radio-button label="encouragement">鼓励型</el-radio-button>
            <el-radio-button label="balanced">综合型</el-radio-button>
            <el-radio-button label="reminder">提醒型</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="评语长度">
          <el-slider
            v-model="settings.length"
            :min="50"
            :max="200"
            :step="10"
            show-stops
          />
          <div class="slider-hint">约 {{ settings.length }} 字</div>
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
        <el-button @click="showSettings = false">取消</el-button>
        <el-button type="primary" @click="saveSettings">保存设置</el-button>
      </template>
    </el-dialog>

    <!-- 评语预览/编辑弹窗 -->
    <el-dialog
      v-model="showPreview"
      :title="previewTitle"
      width="700px"
      :close-on-click-modal="false"
    >
      <div class="preview-content">
        <div class="preview-header">
          <el-avatar :size="48" :icon="UserFilled" />
          <div class="preview-student-info">
            <div class="preview-name">{{ currentStudent?.name }}</div>
            <div class="preview-meta">
              <el-tag
                v-for="score in currentStudent?.recentScores"
                :key="score.subject"
                size="small"
                :type="getSubjectTagType(score.subject)"
                class="meta-tag"
              >
                {{ score.subject }} {{ score.score }}
              </el-tag>
              <el-tag size="small" type="info" effect="plain">
                {{ currentStudent?.recordCount }} 条表现记录
              </el-tag>
            </div>
          </div>
        </div>

        <el-divider />

        <div class="preview-section">
          <div class="section-title">生成评语</div>
          <el-input
            v-model="editableComment"
            type="textarea"
            :rows="8"
            resize="none"
            placeholder="点击 AI生成评语 按钮生成评语，或在此手动输入"
          />
          <div class="char-count">{{ editableComment.length }} 字</div>
        </div>
      </div>

      <template #footer>
        <el-button @click="showPreview = false">关闭</el-button>
        <el-button @click="regenerateCurrent">
          <el-icon><Refresh /></el-icon>
          重新生成
        </el-button>
        <el-button type="primary" @click="saveCurrentComment">
          <el-icon><Check /></el-icon>
          保存评语
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Setting,
  MagicStick,
  UserFilled,
  Check,
  Clock,
  Refresh
} from '@element-plus/icons-vue'

interface Score {
  subject: string
  score: number
}

interface Student {
  id: string
  name: string
  class: string
  recentScores: Score[]
  recordCount: number
  comment?: string
}

// 状态
const generating = ref(false)
const showSettings = ref(false)
const showPreview = ref(false)
const currentStudent = ref<Student | null>(null)
const editableComment = ref('')
const selectedStudents = ref<Student[]>([])

// 设置
const settings = reactive({
  style: 'balanced',
  length: 100,
  include: ['score', 'performance', 'suggestion'],
  extra: ''
})

// 学生数据
const students = ref<Student[]>([
  {
    id: '1',
    name: '张三',
    class: '初二1班',
    recentScores: [
      { subject: '语文', score: 85 },
      { subject: '数学', score: 92 },
      { subject: '英语', score: 88 }
    ],
    recordCount: 5,
    comment: ''
  },
  {
    id: '2',
    name: '李四',
    class: '初二1班',
    recentScores: [
      { subject: '语文', score: 78 },
      { subject: '数学', score: 85 },
      { subject: '英语', score: 90 }
    ],
    recordCount: 3,
    comment: ''
  },
  {
    id: '3',
    name: '王五',
    class: '初二1班',
    recentScores: [
      { subject: '语文', score: 92 },
      { subject: '数学', score: 88 },
      { subject: '英语', score: 85 }
    ],
    recordCount: 4,
    comment: ''
  },
  {
    id: '4',
    name: '赵六',
    class: '初二1班',
    recentScores: [
      { subject: '语文', score: 88 },
      { subject: '数学', score: 76 },
      { subject: '英语', score: 92 }
    ],
    recordCount: 6,
    comment: ''
  },
  {
    id: '5',
    name: '钱七',
    class: '初二1班',
    recentScores: [
      { subject: '语文', score: 95 },
      { subject: '数学', score: 94 },
      { subject: '英语', score: 96 }
    ],
    recordCount: 2,
    comment: ''
  }
])

// 计算属性
const generatedCount = computed(() => students.value.filter(s => s.comment).length)
const pendingCount = computed(() => students.value.filter(s => !s.comment).length)
const previewTitle = computed(() => currentStudent.value ? `${currentStudent.value.name} - 评语详情` : '评语详情')

// 获取学科标签类型
const getSubjectTagType = (subject: string): any => {
  const typeMap: Record<string, any> = {
    '语文': 'danger',
    '数学': 'primary',
    '英语': 'success'
  }
  return typeMap[subject] || 'info'
}

// 处理选择变化
const handleSelectionChange = (selection: Student[]) => {
  selectedStudents.value = selection
}

// 打开预览弹窗
const openPreview = (student: Student) => {
  currentStudent.value = student
  editableComment.value = student.comment || ''
  showPreview.value = true
}

// 生成评语
const generateComments = async () => {
  const targets = selectedStudents.value.length > 0 ? selectedStudents.value : students.value

  if (targets.length === 0) {
    ElMessage.warning('请先选择学生')
    return
  }

  generating.value = true

  // 模拟AI生成
  setTimeout(() => {
    targets.forEach(student => {
      student.comment = generateMockComment(student)
    })
    generating.value = false
    ElMessage.success(`成功生成 ${targets.length} 条评语`)
  }, 2000)
}

// 模拟生成评语
const generateMockComment = (student: Student): string => {
  const styleTemplates: Record<string, string[]> = {
    encouragement: [
      `${student.name}同学本学期表现优异，学习态度积极向上。在${student.recentScores[0].subject}方面取得了${student.recentScores[0].score}分的成绩，展现出良好的学习潜力。希望你继续保持这种学习热情，在各个方面都能取得更大进步。`,
      `${student.name}是一个充满活力的学生，本学期在各项学习活动中都表现出色。${student.recentScores.map(s => s.subject + s.score + '分').join('、')}，各科成绩均衡发展。老师为你的进步感到骄傲，期待你下学期更精彩的表现！`
    ],
    balanced: [
      `${student.name}同学本学期表现良好，学习态度端正。本学期成绩：${student.recentScores.map(s => s.subject + s.score + '分').join('、')}。课堂表现积极，能够按时完成各项作业。希望在今后的学习中能够更加主动思考，提高学习效率。`,
      `${student.name}在本学期的学习中表现稳定。各科成绩较为均衡，其中${student.recentScores[0].subject}表现较为突出。日常表现记录显示该生遵守纪律，与同学相处融洽。建议在学习方法上多下功夫，争取更大突破。`
    ],
    reminder: [
      `${student.name}同学本学期在学习上还有提升空间。虽然${student.recentScores[0].subject}取得了${student.recentScores[0].score}分，但整体学习主动性有待加强。希望下学期能够更加专注于学习，提高课堂效率，争取更好的成绩。`,
      `${student.name}在本学期表现一般，需要更加努力。成绩方面：${student.recentScores.map(s => s.subject + s.score + '分').join('、')}。建议改进学习方法，增强学习自觉性，多向老师和同学请教，相信通过努力一定能取得进步。`
    ]
  }

  const templates = styleTemplates[settings.style] || styleTemplates.balanced
  let comment = templates[Math.floor(Math.random() * templates.length)]

  // 根据长度设置截断
  if (comment.length > settings.length) {
    comment = comment.substring(0, settings.length) + '...'
  }

  return comment
}

// 重新生成当前评语
const regenerateCurrent = () => {
  if (!currentStudent.value) return
  editableComment.value = generateMockComment(currentStudent.value)
  ElMessage.success('评语已重新生成')
}

// 保存当前评语
const saveCurrentComment = () => {
  if (currentStudent.value) {
    currentStudent.value.comment = editableComment.value
    ElMessage.success('评语已保存')
    showPreview.value = false
  }
}

// 清除评语
const clearComment = async (student: Student) => {
  try {
    await ElMessageBox.confirm(
      `确定要清除 ${student.name} 的评语吗？`,
      '确认清除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    student.comment = ''
    ElMessage.success('评语已清除')
  } catch {
    // 用户取消
  }
}

// 保存设置
const saveSettings = () => {
  ElMessage.success('设置已保存')
  showSettings.value = false
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.comments-page {
  padding: $spacing-lg;
}

// 页面头部
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: $spacing-xl;

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

  .header-actions {
    display: flex;
    gap: $spacing-md;

    .el-button {
      display: flex;
      align-items: center;
      gap: $spacing-xs;
    }
  }
}

// 统计栏
.stats-bar {
  display: flex;
  align-items: center;
  gap: $spacing-xl;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: $radius-lg;
  padding: $spacing-xl $spacing-2xl;
  margin-bottom: $spacing-xl;
  color: white;

  .stat-item {
    text-align: center;

    .stat-value {
      font-size: 32px;
      font-weight: 700;
      line-height: 1;
      margin-bottom: $spacing-xs;

      &.generated {
        color: #90EE90;
      }

      &.pending {
        color: #FFD700;
      }
    }

    .stat-label {
      font-size: 14px;
      opacity: 0.9;
    }
  }

  .stat-divider {
    width: 1px;
    height: 40px;
    background: rgba(255, 255, 255, 0.3);
  }
}

// 学生卡片
.students-card {
  .student-info {
    display: flex;
    align-items: center;
    gap: $spacing-md;

    .student-avatar {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    }

    .student-name {
      font-weight: 500;
      color: $text-primary;
    }
  }

  .score-tags {
    display: flex;
    flex-wrap: wrap;
    gap: $spacing-xs;

    .subject-tag {
      font-weight: 500;
    }
  }

  .comment-preview-text {
    color: $text-secondary;
    font-size: 13px;
    line-height: 1.6;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    cursor: pointer;
    padding: $spacing-sm;
    border-radius: $radius-md;
    transition: background-color 0.2s;

    &:hover {
      background-color: $bg-hover;
    }
  }

  .no-comment {
    color: $text-muted;
    font-style: italic;
  }
}

// 设置表单
.settings-form {
  .slider-hint {
    font-size: 12px;
    color: $text-muted;
    margin-top: $spacing-xs;
  }
}

// 预览弹窗
.preview-content {
  .preview-header {
    display: flex;
    align-items: center;
    gap: $spacing-lg;

    .preview-student-info {
      flex: 1;

      .preview-name {
        font-size: 18px;
        font-weight: 600;
        color: $text-primary;
        margin-bottom: $spacing-sm;
      }

      .preview-meta {
        display: flex;
        flex-wrap: wrap;
        gap: $spacing-xs;

        .meta-tag {
          font-weight: 500;
        }
      }
    }
  }

  .preview-section {
    margin-top: $spacing-lg;

    .section-title {
      font-size: 14px;
      font-weight: 600;
      color: $text-primary;
      margin-bottom: $spacing-md;
    }

    .char-count {
      text-align: right;
      font-size: 12px;
      color: $text-muted;
      margin-top: $spacing-sm;
    }
  }
}
</style>
