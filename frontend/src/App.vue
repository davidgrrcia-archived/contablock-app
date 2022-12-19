<script setup lang="ts">
import { RouterView } from 'vue-router';
import HeaderNav from '@/components/HeaderNav.vue';
import CreateEntryModal from '@/components/CreateEntryModal.vue';
</script>

<template>
  <header></header>
  <HeaderNav />
  <CreateEntryModal></CreateEntryModal>
  <RouterView />
</template>
