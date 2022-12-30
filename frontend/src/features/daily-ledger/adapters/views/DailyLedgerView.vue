<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import EntriesTable from '@/components/EntriesTable.vue';

const configStore = useConfigStore();
configStore.activateCurrent('/');
</script>

<template>
  <main>
    <suspense><EntriesTable></EntriesTable></suspense>
  </main>
</template>
