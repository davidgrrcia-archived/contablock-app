import { createRouter, createWebHistory } from 'vue-router';
import DailyLedgerView from '@/features/daily-ledger/adapters/views/DailyLedgerView.vue';
import GeneralLedgerView from '@/features/general-ledger/adapters/views/GeneralLedgerView.vue';
import BalancesView from '@/features/balances/adapters/views/BalancesView.vue';

//TODO: apply clean architecture

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/projects/:project/',
      name: 'dailyLedger',
      component: DailyLedgerView,
    },
    {
      path: '/projects/:project/general-ledger',
      name: 'generalLedger',
      component: GeneralLedgerView,
    },
    {
      path: '/projects/:project/balances',
      name: 'balances',
      component: BalancesView,
    },
  ],
});

export default router;
