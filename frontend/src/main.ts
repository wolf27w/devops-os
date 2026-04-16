import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'
import { permissionDirective, roleDirective } from './utils/permission'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus)

// 注册权限指令
app.directive('permission', permissionDirective)
app.directive('role', roleDirective)

app.mount('#app')