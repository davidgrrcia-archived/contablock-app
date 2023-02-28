<script setup lang="ts">
import { RouterView } from 'vue-router';
import ErrorModal from '@/components/ErrorMessage.vue';
</script>

<template>
  <header></header>
  <ErrorModal></ErrorModal>
  <RouterView class="body" />
</template>
