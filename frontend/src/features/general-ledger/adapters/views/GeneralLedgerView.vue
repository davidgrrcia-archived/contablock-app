<script setup lang="ts">
import { useConfigStore } from '@/stores/config';

const configStore = useConfigStore();
configStore.activateCurrent('/general-ledger');
</script>

<template>
  <main>
    <h1>Libro Mayor</h1>
  </main>
</template>
