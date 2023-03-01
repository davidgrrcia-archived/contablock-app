<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import GeneralLedgerTable from '@/features/general-ledger/adapters/views/GeneralLedgerTable.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <suspense><GeneralLedgerTable></GeneralLedgerTable></suspense>
  </main>
</template>
