<script setup lang="ts">
import { useDailyLedgerStore } from '@/stores/dailyLedger';

const dailyLedgerStore = useDailyLedgerStore();
await dailyLedgerStore.loadDailyLedger();

const dailyLedger = dailyLedgerStore.state.storedDailyLedger
</script>

<template>

</template>

<style scoped>

</style>
