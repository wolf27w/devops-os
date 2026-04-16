<template>
  <div class="db-container">
    <el-card>
      <template #header>
        <div class="header">
          <h3>Database Connections</h3>
          <el-button type="primary" :icon="Plus">New Connection</el-button>
        </div>
      </template>
      
      <el-empty description="Database module placeholder" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Plus } from '@element-plus/icons-vue'
</script>

<style scoped>
.db-container {
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