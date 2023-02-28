<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import HeaderNav from '@/components/HeaderNav.vue';
import GeneralLedgerTable from '@/features/general-ledger/adapters/views/GeneralLedgerTable.vue';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const configStore = useConfigStore();
const route = useRoute();
const project = computed(() => route.params.project);

if (typeof project.value === 'string') {
  configStore.activateCurrent(project.value, '/projects/' + project.value + '/general-ledger');
}
</script>

<template>
  <HeaderNav />
  <main>
    <suspense><GeneralLedgerTable></GeneralLedgerTable></suspense>
  </main>
</template>
