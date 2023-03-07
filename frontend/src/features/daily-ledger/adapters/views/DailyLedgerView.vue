<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import DailyLedgerTable from '@/features/daily-ledger/adapters/views/DailyLedgerTable.vue';
import CreateEntryCTA from '@/components/CreateEntryCTA.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <CreateEntryCTA
      :props="{
        title: 'Libro Diario',
        description: 'Registro de transacciones financieras diarias ordenadas cronológicamente.',
      }"
    ></CreateEntryCTA>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
