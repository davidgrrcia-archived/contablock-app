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
  <header class="fixed top-0 left-0 flex w-full items-center bg-black">
    <div class="w-1/3">
      <img alt="Vue logo" class="object-cover" src="@/assets/logo.png" />
    </div>
    <nav class="ml-10 w-1/3">
      <RouterLink v-for="item in configStore.state.navigation" :key="item.name" :to="item.href" :class="['rounded-md p-2 font-normal text-white ' + (item.current ? 'underline' : 'hover:bg-white hover:text-black')]">{{ item.name }}</RouterLink>
    </nav>
    <div class="w-1/3">
      <button type="button" @click="showCreateEntryModal()" class="btn mt-3 mb-3 bg-[var(--blue)] hover:bg-indigo-700 focus:ring-indigo-500">Nuevo asiento</button>
    </div>
  </header>
  <CreateEntryModal></CreateEntryModal>
</template>
