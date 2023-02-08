<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import DailyLedgerTable from '@/components/DailyLedgerTable.vue';

const configStore = useConfigStore();
configStore.activateCurrent('/');
</script>

<template>
  <main>
    <h1>Libro Diario</h1>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
