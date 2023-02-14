<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import DailyLedgerTable from '@/features/daily-ledger/adapters/views/DailyLedgerTable.vue';

const configStore = useConfigStore();
configStore.activateCurrent('/');
</script>

<template>
  <main>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
