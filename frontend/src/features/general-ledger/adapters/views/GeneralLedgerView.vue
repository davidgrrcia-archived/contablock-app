<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import GeneralLedgerTable from '@/features/general-ledger/adapters/views/GeneralLedgerTable.vue';
import CreateEntryCTA from '@/components/CreateEntryCTA.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <CreateEntryCTA
      :props="{
        title: 'Libro Mayor',
        description: 'Registro de débitos y créditos de una empresa clasificados por cuentas contables.',
      }"
    ></CreateEntryCTA>
    <suspense><GeneralLedgerTable></GeneralLedgerTable></suspense>
  </main>
</template>
