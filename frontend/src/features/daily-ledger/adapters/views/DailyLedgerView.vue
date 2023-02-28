<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import HeaderNav from '@/components/HeaderNav.vue';
import DailyLedgerTable from '@/features/daily-ledger/adapters/views/DailyLedgerTable.vue';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const configStore = useConfigStore();
const route = useRoute();
const project = computed(() => route.params.project);

if (typeof project.value === 'string') {
  configStore.activateCurrent(project.value, '/projects/' + project.value);
}
</script>

<template>
  <HeaderNav />
  <main>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
