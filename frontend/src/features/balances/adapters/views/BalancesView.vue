<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import BalanceTable from '@/features/balances/adapters/views/BalancesTable.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <suspense><BalanceTable></BalanceTable></suspense>
  </main>
</template>
