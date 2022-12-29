<script setup lang="ts">
import { RouterView } from 'vue-router';
import HeaderNav from '@/components/HeaderNav.vue';
import ErrorModal from '@/components/ErrorMessage.vue';
</script>

<template>
  <header></header>
  <ErrorModal></ErrorModal>
  <HeaderNav />
  <RouterView />
</template>
