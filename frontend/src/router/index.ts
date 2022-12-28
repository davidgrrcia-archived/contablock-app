import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/features/home/adapters/views/HomeView.vue';

//TODO: apply clean architecture

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
  ],
});

export default router;
