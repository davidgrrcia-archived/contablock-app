<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import DailyLedgerTable from '@/features/daily-ledger/adapters/views/DailyLedgerTable.vue';
import CreateEntryModal from '@/features/entries/adapters/views/CreateEntryModal.vue';
import ModifyEntryModal from '@/features/entries/adapters/views/ModifyEntryModal.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <CreateEntryModal></CreateEntryModal>
    <ModifyEntryModal></ModifyEntryModal>
    <suspense><DailyLedgerTable></DailyLedgerTable></suspense>
  </main>
</template>
