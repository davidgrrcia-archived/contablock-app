<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import BalanceTable from '@/features/balances/adapters/views/BalancesTable.vue';

const configStore = useConfigStore();
configStore.activateCurrent('/balances');
</script>

<template>
  <main>
    <suspense><BalanceTable></BalanceTable></suspense>
  </main>
</template>
