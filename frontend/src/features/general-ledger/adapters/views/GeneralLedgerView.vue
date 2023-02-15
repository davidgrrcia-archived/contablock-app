<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import GeneralLedgerTable from '@/features/general-ledger/adapters/views/GeneralLedgerTable.vue';

const configStore = useConfigStore();
configStore.activateCurrent('/general-ledger');
</script>

<template>
  <main>
    <suspense><GeneralLedgerTable></GeneralLedgerTable></suspense>
  </main>
</template>
