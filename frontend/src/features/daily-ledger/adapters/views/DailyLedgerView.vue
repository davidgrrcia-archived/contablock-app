<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import DailyLedgerTable from '@/features/daily-ledger/adapters/views/DailyLedgerTable.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
