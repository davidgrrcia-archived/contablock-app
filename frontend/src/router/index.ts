import { createRouter, createWebHistory } from 'vue-router';
import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router';
import DailyLedgerView from '@/features/daily-ledger/adapters/views/DailyLedgerView.vue';
import GeneralLedgerView from '@/features/general-ledger/adapters/views/GeneralLedgerView.vue';
import BalancesView from '@/features/balances/adapters/views/BalancesView.vue';
import NotFound from '@/components/NotFound.vue';
import LoginView from '@/features/users/adapters/views/LoginView.vue';
import { useConfigStore } from '@/stores/config';

//TODO: apply clean architecture

const requireAuth = (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
  const configStore = useConfigStore();

  if (typeof to.params.project !== 'string') {
    return;
  }

  if (!configStore.state.isAuthenticated && to.name !== 'Login') {
    next({
      name: 'login',
      params: { project: to.params.project },
      query: { redirect: to.path },
    });
  } else {
    next();
  }
};

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/projects/:project/',
      name: 'dailyLedger',
      beforeEnter: requireAuth,
      component: DailyLedgerView,
    },
    {
      path: '/projects/:project/general-ledger',
      name: 'generalLedger',
      beforeEnter: requireAuth,
      component: GeneralLedgerView,
    },
    {
      path: '/projects/:project/balances',
      name: 'balances',
      beforeEnter: requireAuth,
      component: BalancesView,
    },
    {
      path: '/projects/:project/login',
      name: 'login',
      component: LoginView,
    },
    { path: '/:pathMatch(.*)*', component: NotFound },
  ],
});

export default router;
