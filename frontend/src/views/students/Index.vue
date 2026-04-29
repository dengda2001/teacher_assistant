<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">
          学生管理
          <span class="subtitle">管理班级学生信息</span>
        </h1>
      </div>
      <el-button type="primary" class="btn-primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        <span>添加学生</span>
      </el-button>
    </div>

    <el-card class="content-card" shadow="never">
      <div class="table-toolbar">
        <div class="search-box-table">
          <el-icon class="search-icon"><Search /></el-icon>
          <el-input
            v-model="searchQuery"
            placeholder="搜索学生姓名或学号..."
            class="search-input-clean"
            clearable
          />
        </div>
        <span class="student-count">共 {{ filteredStudents.length }} 名学生</span>
      </div>

      <el-table :data="filteredStudents" v-loading="loading" style="width: 100%">
        <el-table-column min-width="120">
          <template #default="{ row }">
            <div class="student-name-cell">
              <div class="avatar-circle">{{ row.name.charAt(0) }}</div>
              <span class="name-text">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="grade" label="年级" min-width="100">
          <template #default="{ row }">
            初中{{ row.grade }}
          </template>
        </el-table-column>
        <el-table-column label="性别" min-width="100">
          <template #default="{ row }">
            <span :class="['gender-tag', row.gender === '男' ? 'male' : 'female']">
              {{ row.gender }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <div class="action-btns">
              <el-button type="primary" link @click="editStudent(row)">
                <el-icon><Edit /></el-icon>
              </el-button>
              <el-button type="danger" link @click="deleteStudent(row)">
                <el-icon><Delete /></el-icon>
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑学生弹窗 -->
    <el-dialog
      v-model="showAddDialog"
      :title="isEdit ? '编辑学生' : '添加新学生'"
      width="500px"
      class="custom-dialog"
      destroy-on-close
    >
      <p class="dialog-subtitle">请填写学生的基本信息</p>
      <el-form :model="studentForm" label-width="80px" :rules="rules" ref="formRef" class="custom-form">
        <el-form-item label="姓名" prop="name">
          <el-input v-model="studentForm.name" placeholder="请输入学生姓名" />
        </el-form-item>
        <el-form-item label="学号" prop="studentNo">
          <el-input v-model="studentForm.studentNo" placeholder="请输入学号" />
        </el-form-item>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="班级" prop="class">
              <el-input v-model="studentForm.class" placeholder="例如：1班" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="年级" prop="grade">
              <el-select v-model="studentForm.grade" placeholder="请选择" style="width: 100%">
                <el-option label="初一" :value="7" />
                <el-option label="初二" :value="8" />
                <el-option label="初三" :value="9" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="性别" prop="gender">
          <el-radio-group v-model="studentForm.gender">
            <el-radio label="男">男</el-radio>
            <el-radio label="女">女</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showAddDialog = false" class="btn-secondary">取消</el-button>
          <el-button type="primary" @click="saveStudent" class="btn-primary">
            {{ isEdit ? '保存' : '确认添加' }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Edit, Delete } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'

interface Student {
  id?: string
  name: string
  studentNo: string
  grade: number
  class: string
  gender: string
}

const loading = ref(false)
const searchQuery = ref('')
const showAddDialog = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

const studentForm = reactive<Student>({
  name: '',
  studentNo: '',
  grade: 7,
  class: '',
  gender: '男'
})

const students = ref<Student[]>([
  { id: '1', name: '王小明', studentNo: '2024001', grade: 9, class: '1班', gender: '男' },
  { id: '2', name: '李小红', studentNo: '2024002', grade: 9, class: '1班', gender: '女' },
  { id: '3', name: '张三', studentNo: '2024003', grade: 9, class: '1班', gender: '男' },
  { id: '4', name: '赵四', studentNo: '2024004', grade: 9, class: '1班', gender: '男' },
  { id: '5', name: '刘五', studentNo: '2024005', grade: 9, class: '1班', gender: '女' },
  { id: '6', name: '陈六', studentNo: '2024006', grade: 9, class: '1班', gender: '男' },
  { id: '7', name: '周七', studentNo: '2024007', grade: 9, class: '1班', gender: '女' },
  { id: '8', name: '吴八', studentNo: '2024008', grade: 9, class: '1班', gender: '男' }
])

const filteredStudents = computed(() => {
  if (!searchQuery.value) return students.value
  const query = searchQuery.value.toLowerCase()
  return students.value.filter(s =>
    s.name.toLowerCase().includes(query) ||
    s.studentNo.includes(query)
  )
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  studentNo: [{ required: true, message: '请输入学号', trigger: 'blur' }],
  grade: [{ required: true, message: '请选择年级', trigger: 'change' }],
  class: [{ required: true, message: '请输入班级', trigger: 'blur' }],
  gender: [{ required: true, message: '请选择性别', trigger: 'change' }]
}

const saveStudent = async () => {
  if (!formRef.value) return
  await formRef.value.validate((valid) => {
    if (valid) {
      if (isEdit.value) {
        const index = students.value.findIndex(s => s.id === studentForm.id)
        if (index !== -1) {
          students.value[index] = { ...studentForm }
        }
        ElMessage.success('更新成功')
      } else {
        students.value.push({
          ...studentForm,
          id: Date.now().toString()
        })
        ElMessage.success('添加成功')
      }
      showAddDialog.value = false
      resetForm()
    }
  })
}

const editStudent = (student: Student) => {
  isEdit.value = true
  Object.assign(studentForm, student)
  showAddDialog.value = true
}

const deleteStudent = (student: Student) => {
  ElMessageBox.confirm(
    `确定要删除学生 "${student.name}" 吗？`,
    '确认删除',
    { type: 'warning' }
  ).then(() => {
    students.value = students.value.filter(s => s.id !== student.id)
    ElMessage.success('删除成功')
  })
}

const resetForm = () => {
  isEdit.value = false
  Object.assign(studentForm, {
    id: undefined,
    name: '',
    studentNo: '',
    grade: 7,
    class: '',
    gender: '男'
  })
}
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.content-card {
  border-radius: $radius-lg;
  border: 1px solid $border;

  :deep(.el-card__body) {
    padding: $spacing-xl;
  }
}

.table-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-lg;
}

.search-box-table {
  position: relative;
  width: 280px;

  .search-icon {
    position: absolute;
    left: $spacing-md;
    top: 50%;
    transform: translateY(-50%);
    color: $text-muted;
    font-size: 16px;
    z-index: 1;
  }

  .search-input-clean {
    :deep(.el-input__wrapper) {
      background: $bg-input;
      border-radius: $radius-md;
      box-shadow: none;
      padding-left: 36px;

      .el-input__inner {
        height: 36px;
        font-size: 14px;

        &::placeholder {
          color: $text-muted;
        }
      }

      &.is-focus {
        box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
      }
    }
  }
}

.student-count {
  font-size: 14px;
  color: $text-secondary;
}

.student-name-cell {
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

  .name-text {
    font-size: 14px;
    color: $text-primary;
    font-weight: 500;
  }
}

.action-btns {
  display: flex;
  gap: $spacing-xs;

  .el-button {
    font-size: 16px;
  }
}

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
  .el-select .el-input__wrapper {
    border-radius: $radius-md;
    box-shadow: 0 0 0 1px $border;

    &.is-focus {
      box-shadow: 0 0 0 1px $primary;
    }
  }

  .el-radio__input.is-checked .el-radio__inner {
    background: $primary;
    border-color: $primary;
  }

  .el-radio__input.is-checked + .el-radio__label {
    color: $primary;
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: $spacing-md;
  padding: $spacing-md 0 0;
}
</style>
