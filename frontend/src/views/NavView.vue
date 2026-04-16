<template>
  <div class="nav-container">
    <el-card>
      <template #header>
        <div class="header">
          <h3>Navigation Items</h3>
          <el-button type="primary" :icon="Plus">New Item</el-button>
        </div>
      </template>
      
      <el-empty description="Navigation module placeholder" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Plus } from '@element-plus/icons-vue'
</script>

<style scoped>
.nav-container {
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