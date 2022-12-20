<script setup lang="ts">
import type { Navigation } from '@/interfaces/Navigation';
import { useCreateEntryModalStore } from '@/stores/createEntryModal';
import CreateEntryModal from '@/components/CreateEntryModal.vue';

const navigation: Navigation[] = [
  { name: 'Dashboard', href: '#', current: true },
  { name: 'Team', href: '#', current: false },
  { name: 'Projects', href: '#', current: false },
  { name: 'Calendar', href: '#', current: false },
];

const modalState = useCreateEntryModalStore();

function showCreateEntryModal() {
  modalState.showModal();
}
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="50" height="50" />

      <nav>
        <RouterLink
          v-for="item in navigation"
          :key="item.name"
          :to="item.href"
          :class="[
            item.current
            ? 'current'
            : 'not-current','menu'
          ]"

        >{{item.name}}</RouterLink>
      </nav>

    <button
      class="button"
      type="button"
      @click="showCreateEntryModal()"
    >
      New Entry
    </button>
  </header>
</template>

<style scoped>
header {
  position: fixed;
  display: flex;
  top: 0;
  left: 0;
  height: auto;
  width: 100%;
  background-color: #222222;
  place-items: center;
  padding: 10px;
}

.logo {
  display: block;
  margin: 0 2rem 0 0;
}

nav {
  width: 100%;
  margin-top: 1rem;
  margin-left: -1rem;
  padding: 1rem 0;
  text-align: center;
}

.menu{
  font-size: 18px;
  border-radius: 5px;
  margin-left: 20px;
  text-decoration: none;
  padding: 10px;
}

.not-current{
  color: #fbbf24;
}

.not-current:hover{
  background-color: #fbbf24;
  color: #222222;
}

.current{
  background-color: #fbbf24;
  color: #222222;
}

.button{
    display: block;
    margin: 0 2rem 0 0;
}

</style>
