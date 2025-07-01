import { createRouter, createWebHistory } from 'vue-router'
import InfiniteScroll from '../views/InfiniteScroll.vue'

const routes = [{ path: '/', name: 'Home', component: InfiniteScroll }]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
