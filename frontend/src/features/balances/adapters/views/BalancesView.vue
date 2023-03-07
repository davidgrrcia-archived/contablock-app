<script setup lang="ts">
import HeaderNav from '@/components/HeaderNav.vue';
import BalanceTable from '@/features/balances/adapters/views/BalancesTable.vue';
import CreateEntryCTA from '@/components/CreateEntryCTA.vue';
</script>

<template>
  <HeaderNav />
  <main>
    <CreateEntryCTA
      :props="{
        title: 'Balance de situación',
        description: 'Estado financiero de recursos y obligaciones de la empresa en este momento',
      }"
    ></CreateEntryCTA>
    <suspense><BalanceTable></BalanceTable></suspense>
  </main>
</template>
