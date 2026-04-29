<template>
  <div class="excel-uploader">
    <!-- 第一步：上传文件 -->
    <div v-if="step === 'upload'" class="upload-area">
      <el-upload
        drag
        action="#"
        :auto-upload="false"
        :on-change="handleFileChange"
        accept=".xlsx,.xls"
        class="upload-dragger"
      >
        <el-icon class="upload-icon"><Upload /></el-icon>
        <div class="upload-text">
          拖拽Excel文件到此处，或<em>点击上传</em>
        </div>
        <template #tip>
          <div class="upload-tip">
            支持 .xlsx, .xls 格式，文件大小不超过 10MB
          </div>
        </template>
      </el-upload>

      <div class="template-tip">
        <el-alert type="info" :closable="false">
          <template #title>
            Excel格式说明
          </template>
          <div class="format-example">
            <p>系统会智能识别列，但建议包含以下列：</p>
            <ul>
              <li>姓名/学生姓名</li>
              <li>学号（可选）</li>
              <li>各科成绩列（如：语文、数学、英语）</li>
            </ul>
          </div>
        </el-alert>
      </div>
    </div>

    <!-- 第二步：列映射确认 -->
    <div v-if="step === 'mapping'" class="mapping-area">
      <h3>请确认列映射</h3>
      <p class="mapping-desc">系统已自动识别Excel列，请检查并调整映射关系：</p>

      <el-table :data="mappingData" border style="width: 100%">
        <el-table-column prop="excelColumn" label="Excel列名" min-width="200" />
        <el-table-column prop="sample" label="示例数据" min-width="150" />
        <el-table-column label="映射为" min-width="200">
          <template #default="{ row }">
            <el-select v-model="row.mapping" placeholder="选择映射" clearable style="width: 100%">
              <el-option label="学生姓名" value="name" />
              <el-option label="学号" value="studentNo" />
              <el-option label="语文" value="chinese" />
              <el-option label="数学" value="math" />
              <el-option label="英语" value="english" />
              <el-option label="物理" value="physics" />
              <el-option label="化学" value="chemistry" />
              <el-option label="班级" value="class" />
              <el-option label="忽略此列" value="ignore" />
            </el-select>
          </template>
        </el-table-column>
      </el-table>

      <div class="mapping-actions">
        <el-button @click="step = 'upload'">重新上传</el-button>
        <el-button type="primary" @click="handlePreview">下一步：预览数据</el-button>
      </div>
    </div>

    <!-- 第三步：数据预览 -->
    <div v-if="step === 'preview'" class="preview-area">
      <h3>数据预览</h3>
      <p class="preview-desc">共识别 {{ previewData.length }} 条记录，请确认数据无误：</p>

      <el-table :data="previewData" border height="300" style="width: 100%">
        <el-table-column type="index" width="60" label="序号" />
        <el-table-column
          v-for="col in previewColumns"
          :key="col.prop"
          :prop="col.prop"
          :label="col.label"
          min-width="120"
        />
      </el-table>

      <div class="preview-actions">
        <el-button @click="step = 'mapping'">上一步</el-button>
        <el-button type="primary" :loading="importing" @click="handleImport">确认导入</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Upload } from '@element-plus/icons-vue'
import * as XLSX from 'xlsx'

const emit = defineEmits(['success'])

interface MappingItem {
  excelColumn: string
  sample: string
  mapping: string
}

type Step = 'upload' | 'mapping' | 'preview'

const step = ref<Step>('upload')
const mappingData = ref<MappingItem[]>([])
const previewData = ref<any[]>([])
const previewColumns = ref<{ prop: string; label: string }[]>([])
const importing = ref(false)

const handleFileChange = (file: any) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    try {
      const data = new Uint8Array(e.target?.result as ArrayBuffer)
      const workbook = XLSX.read(data, { type: 'array' })
      const firstSheet = workbook.Sheets[workbook.SheetNames[0]]
      const jsonData = XLSX.utils.sheet_to_json(firstSheet, { header: 1 }) as any[][]

      if (jsonData.length < 2) {
        ElMessage.error('Excel文件数据不完整')
        return
      }

      // 提取列名和样例数据
      const headers = jsonData[0]
      const sampleRow = jsonData[1]

      mappingData.value = headers.map((header: string, index: number) => ({
        excelColumn: header,
        sample: sampleRow[index] || '',
        mapping: guessMapping(header)
      }))

      step.value = 'mapping'
    } catch (error) {
      ElMessage.error('解析Excel失败，请检查文件格式')
    }
  }
  reader.readAsArrayBuffer(file.raw)
}

const guessMapping = (columnName: string): string => {
  const name = columnName.toLowerCase()
  if (name.includes('姓名') || name.includes('名字') || name.includes('学生')) return 'name'
  if (name.includes('学号') || name.includes('编号')) return 'studentNo'
  if (name.includes('语文')) return 'chinese'
  if (name.includes('数学')) return 'math'
  if (name.includes('英语')) return 'english'
  if (name.includes('物理')) return 'physics'
  if (name.includes('化学')) return 'chemistry'
  if (name.includes('班级') || name.includes('班')) return 'class'
  return ''
}

const handlePreview = () => {
  // 检查是否至少映射了姓名列
  const hasName = mappingData.value.some(m => m.mapping === 'name')
  if (!hasName) {
    ElMessage.warning('请至少映射"学生姓名"列')
    return
  }

  // 生成预览数据（模拟）
  previewColumns.value = mappingData.value
    .filter(m => m.mapping && m.mapping !== 'ignore')
    .map(m => ({
      prop: m.mapping,
      label: m.excelColumn
    }))

  previewData.value = [
    { name: '张三', studentNo: '2024001', chinese: 85, math: 92, english: 78 },
    { name: '李四', studentNo: '2024002', chinese: 90, math: 88, english: 85 },
    { name: '王五', studentNo: '2024003', chinese: 76, math: 95, english: 82 }
  ]

  step.value = 'preview'
}

const handleImport = () => {
  importing.value = true
  setTimeout(() => {
    importing.value = false
    ElMessage.success('成绩导入成功')
    emit('success', previewData.value)
  }, 1500)
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.excel-uploader {
  padding: $spacing-lg;
}

.upload-area {
  .upload-dragger {
    width: 100%;
  }

  .upload-icon {
    font-size: 48px;
    color: $primary;
    margin-bottom: $spacing-md;
  }

  .upload-text {
    font-size: 16px;
    color: $text-secondary;

    em {
      color: $primary;
      font-style: normal;
      font-weight: 500;
    }
  }

  .upload-tip {
    margin-top: $spacing-md;
    font-size: 13px;
    color: $text-muted;
  }
}

.template-tip {
  margin-top: $spacing-xl;

  .format-example {
    font-size: 13px;

    p {
      margin-bottom: $spacing-sm;
    }

    ul {
      padding-left: $spacing-xl;
      margin: 0;

      li {
        margin: $spacing-xs 0;
        color: $text-secondary;
      }
    }
  }
}

.mapping-area,
.preview-area {
  h3 {
    font-size: 18px;
    font-weight: 600;
    margin-bottom: $spacing-md;
    color: $text-primary;
  }

  .mapping-desc,
  .preview-desc {
    color: $text-secondary;
    margin-bottom: $spacing-lg;
  }
}

.mapping-actions,
.preview-actions {
  margin-top: $spacing-xl;
  display: flex;
  justify-content: flex-end;
  gap: $spacing-md;
}
</style>