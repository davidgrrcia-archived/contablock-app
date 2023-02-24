<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import DailyLedgerTable from '@/features/daily-ledger/adapters/views/DailyLedgerTable.vue';
import { useRoute } from 'vue-router';
import { ref } from 'vue';

const configStore = useConfigStore();
configStore.activateCurrent('/');

const route = useRoute();
const project = ref(route.params.project);

console.log(project);
</script>

<template>
  <main>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
