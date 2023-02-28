<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import HeaderNav from '@/components/HeaderNav.vue';
import BalanceTable from '@/features/balances/adapters/views/BalancesTable.vue';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const configStore = useConfigStore();
const route = useRoute();
const project = computed(() => route.params.project);

if (typeof project.value === 'string') {
  configStore.activateCurrent(project.value, '/projects/' + project.value + '/balances');
}
</script>

<template>
  <HeaderNav />
  <main>
    <suspense><BalanceTable></BalanceTable></suspense>
  </main>
</template>
