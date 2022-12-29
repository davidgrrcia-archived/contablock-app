<script setup lang="ts">
import { useConfigStore } from '@/stores/config';

const configStore = useConfigStore();
configStore.activateCurrent('/');
</script>

<template>
  <main>
    <h1>Libro Diario</h1>
  </main>
</template>
