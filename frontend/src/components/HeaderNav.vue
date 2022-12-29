<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';
import CreateEntryModal from '@/components/CreateEntryModal.vue';
import { useConfigStore } from '@/stores/config';

const configStore = useConfigStore();
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
        v-for="item in configStore.state.navigation"
        :key="item.name"
        :to="item.href"
        :class="[item.current ? 'current' : 'not-current', 'menu']"
        >{{ item.name }}</RouterLink
      >
    </nav>

    <button class="btn" type="button" @click="showCreateEntryModal()">New Entry</button>
  </header>
  <CreateEntryModal></CreateEntryModal>
</template>

<style scoped>
header {
  position: fixed;
  display: flex;
  top: 0;
  left: 0;
  height: auto;
  width: 100%;
  background-color: var(--primary);
  place-items: center;
}

.logo {
  display: block;
  margin: 20px;
}

nav {
  width: 100%;
  text-align: center;
}

.menu {
  font-size: 18px;
  border-radius: 5px;
  margin-left: 20px;
  text-decoration: none;
  padding: 10px;
}

.not-current {
  color: var(--quaternary);
}

.not-current:hover {
  background-color: var(--quaternary);
  color: #222222;
}

.current {
  background-color: var(--quaternary);
  color: #222222;
}

.btn {
  border: none;
  padding: 10px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 18px;
  border-radius: 5px;
  background-color: var(--enphasis);
  color: #222222;
  margin-right: 50px;
  white-space: nowrap;
}
</style>
