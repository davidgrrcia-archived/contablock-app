<script setup lang="ts">
import { RouterView } from 'vue-router';
import SiteNavbar from '@/components/SiteNavbar.vue';
</script>

<template>
  <header></header>
  <SiteNavbar />
  <RouterView />
</template>
