<template>
  <div class="comment-preview">
    <div class="student-header">
      <el-avatar :size="48" :icon="UserFilled" />
      <div class="student-info">
        <div class="name">{{ student?.name }}</div>
        <div class="class">{{ student?.class }}</div>
      </div>
    </div>

    <el-divider />

    <div class="preview-section">
      <h4>数据源</h4>
      <div class="data-source">
        <div class="source-item">
          <span class="label">近期成绩:</span>
          <el-tag v-for="score in student?.recentScores" :key="score.subject" size="small" class="score-tag">
            {{ score.subject }}: {{ score.score }}
          </el-tag>
        </div>
        <div class="source-item">
          <span class="label">表现记录:</span>
          <span>{{ student?.recordCount }} 条</span>
        </div>
      </div>
    </div>

    <div class="preview-section">
      <h4>生成评语</h4>
      <el-input
        v-model="editableComment"
        type="textarea"
        :rows="6"
        resize="none"
      />
      <div class="char-count">{{ editableComment.length }} 字</div>
    </div>

    <div class="preview-actions">
      <el-button @click="$emit('regenerate')">
        <el-icon><Refresh /></el-icon>
        重新生成
      </el-button>
      <el-button type="primary" @click="handleSave">
        <el-icon><Check /></el-icon>
        保存评语
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { UserFilled, Refresh, Check } from '@element-plus/icons-vue'

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
}

const props = defineProps<{
  student: Student | null
  comment: string
}>()

const emit = defineEmits(['regenerate', 'save'])

const editableComment = ref('')

watch(() => props.comment, (newVal) => {
  editableComment.value = newVal
}, { immediate: true })

const handleSave = () => {
  emit('save', editableComment.value)
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.comment-preview {
  padding: $spacing-lg;
}

.student-header {
  display: flex;
  align-items: center;
  gap: $spacing-lg;

  .student-info {
    .name {
      font-size: 18px;
      font-weight: 600;
      color: $text-primary;
    }

    .class {
      font-size: 13px;
      color: $text-secondary;
      margin-top: $spacing-xs;
    }
  }
}

.preview-section {
  margin: $spacing-xl 0;

  h4 {
    font-size: 14px;
    font-weight: 600;
    color: $text-primary;
    margin-bottom: $spacing-md;
  }
}

.data-source {
  background: $bg-secondary;
  border-radius: $radius-md;
  padding: $spacing-lg;

  .source-item {
    display: flex;
    align-items: center;
    gap: $spacing-md;
    margin-bottom: $spacing-md;

    &:last-child {
      margin-bottom: 0;
    }

    .label {
      font-weight: 500;
      color: $text-secondary;
      min-width: 80px;
    }
  }

  .score-tag {
    margin-right: $spacing-xs;
  }
}

.char-count {
  text-align: right;
  font-size: 12px;
  color: $text-muted;
  margin-top: $spacing-sm;
}

.preview-actions {
  display: flex;
  justify-content: flex-end;
  gap: $spacing-md;
  margin-top: $spacing-xl;
}
</style>
