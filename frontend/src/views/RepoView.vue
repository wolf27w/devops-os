<template>
  <div class="repo-container">
    <el-card>
      <template #header>
        <div class="header">
          <h3>Repositories</h3>
          <el-button type="primary" :icon="Plus">New Repository</el-button>
        </div>
      </template>
      
      <el-empty description="Repository module placeholder" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Plus } from '@element-plus/icons-vue'
</script>

<style scoped>
.repo-container {
  height: 100%;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header h3 {
  margin: 0;
}
</style>