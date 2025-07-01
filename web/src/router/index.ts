import StocksView from '@/views/StocksView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [{ path: '/', name: 'Home', component: StocksView }]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
