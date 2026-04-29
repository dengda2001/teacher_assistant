<template>
  <el-container class="layout">
    <!-- 侧边栏 -->
    <el-aside width="240px" class="sidebar">
      <div class="logo">
        <div class="logo-icon">
          <el-icon size="24" color="#6366f1"><School /></el-icon>
        </div>
        <div class="logo-text">
          <span class="title">教师助手</span>
          <span class="subtitle">智能教学管理</span>
        </div>
      </div>

      <el-menu
        :default-active="$route.path"
        router
        class="nav-menu"
      >
        <el-menu-item v-for="route in routes" :key="route.path" :index="route.path">
          <el-icon>
            <component :is="route.meta?.icon" />
          </el-icon>
          <span>{{ route.meta?.title }}</span>
        </el-menu-item>
      </el-menu>

      <div class="sidebar-footer">
        <el-menu-item index="/settings" class="settings-item">
          <el-icon><Setting /></el-icon>
          <span>设置</span>
        </el-menu-item>
      </div>
    </el-aside>

    <!-- 主内容区 -->
    <el-container class="main-container">
      <el-header class="header">
        <div class="search-box">
          <el-icon class="search-icon"><Search /></el-icon>
          <input type="text" placeholder="搜索学生、记录..." class="search-input-header" />
        </div>
        <div class="header-actions">
          <el-button class="icon-btn" text>
            <el-icon size="18"><Bell /></el-icon>
          </el-button>
          <div class="user-info">
            <el-avatar :size="36" class="user-avatar">张</el-avatar>
            <div class="user-meta">
              <span class="user-name">张老师</span>
              <span class="user-role">高三(1)班班主任</span>
            </div>
          </div>
        </div>
      </el-header>

      <el-main class="main-content">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import router from '@/router'
import { School, Setting, Search, Bell } from '@element-plus/icons-vue'

const $route = useRoute()

const routes = computed(() => {
  return router.getRoutes().filter(route => route.meta?.title && route.path !== '/settings')
})
</script>

<style scoped lang="scss">
@use '@/styles/variables.scss' as *;

.layout {
  height: 100vh;
  background: $bg-primary;
}

.sidebar {
  background: $bg-card;
  border-right: 1px solid $border;
  display: flex;
  flex-direction: column;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.04);
}

.logo {
  height: 80px;
  display: flex;
  align-items: center;
  padding: 0 $spacing-xl;
  gap: $spacing-md;
  border-bottom: 1px solid $border-light;

  .logo-icon {
    width: 40px;
    height: 40px;
    background: $primary-lighter;
    border-radius: $radius-md;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .logo-text {
    display: flex;
    flex-direction: column;

    .title {
      font-size: 18px;
      font-weight: 600;
      color: $text-primary;
      letter-spacing: -0.3px;
      line-height: 1.3;
    }

    .subtitle {
      font-size: 12px;
      color: $text-muted;
      line-height: 1.3;
    }
  }
}

.nav-menu {
  flex: 1;
  padding: $spacing-md 0;
  border-right: none;

  :deep(.el-menu-item) {
    height: 48px;
    line-height: 48px;
    margin: $spacing-xs $spacing-lg;
    border-radius: $radius-md;
    color: $text-secondary;
    font-size: 14px;

    .el-icon {
      margin-right: $spacing-md;
      font-size: 18px;
    }

    &:hover {
      background: $bg-hover;
      color: $text-primary;
    }

    &.is-active {
      background: $primary-lighter;
      color: $primary;
      font-weight: 500;

      &::before {
        content: '';
        position: absolute;
        left: -$spacing-lg;
        top: 50%;
        transform: translateY(-50%);
        width: 3px;
        height: 20px;
        background: $primary;
        border-radius: 0 2px 2px 0;
      }
    }
  }
}

.sidebar-footer {
  padding: $spacing-md 0;
  border-top: 1px solid $border-light;

  .settings-item {
    height: 48px;
    line-height: 48px;
    margin: 0 $spacing-lg;
    border-radius: $radius-md;
    color: $text-secondary;

    &:hover {
      background: $bg-hover;
      color: $text-primary;
    }
  }
}

.main-container {
  background: $bg-primary;
}

.header {
  height: 64px;
  background: $bg-card;
  border-bottom: 1px solid $border;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 $spacing-2xl;
}

.search-box {
  position: relative;
  width: 320px;

  .search-icon {
    position: absolute;
    left: $spacing-md;
    top: 50%;
    transform: translateY(-50%);
    color: $text-muted;
    font-size: 16px;
    z-index: 1;
  }

  .search-input-header {
    width: 100%;
    height: 40px;
    background: $bg-input;
    border: none;
    border-radius: 24px;
    padding: 0 $spacing-lg 0 40px;
    font-size: 14px;
    color: $text-primary;
    outline: none;
    transition: all $transition-fast;

    &::placeholder {
      color: $text-muted;
    }

    &:focus {
      background: $bg-card;
      box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
    }
  }
}

.header-actions {
  display: flex;
  align-items: center;
  gap: $spacing-lg;

  .icon-btn {
    width: 36px;
    height: 36px;
    border-radius: $radius-md;
    color: $text-secondary;

    &:hover {
      background: $bg-hover;
      color: $text-primary;
    }
  }
}

.user-info {
  display: flex;
  align-items: center;
  gap: $spacing-md;

  .user-avatar {
    background: linear-gradient(135deg, $primary 0%, $primary-light 100%);
    color: white;
    font-weight: 600;
  }

  .user-meta {
    display: flex;
    flex-direction: column;

    .user-name {
      font-size: 14px;
      font-weight: 600;
      color: $text-primary;
      line-height: 1.3;
    }

    .user-role {
      font-size: 12px;
      color: $text-muted;
      line-height: 1.3;
    }
  }
}

.main-content {
  padding: $spacing-2xl;
  overflow-y: auto;
}

// 过渡动画
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
