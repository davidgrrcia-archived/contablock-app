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
    <img alt="Vue logo" class="object-cover" src="@/assets/logo.png" />
    <nav class="ml-10">
      <RouterLink v-for="item in configStore.state.navigation" :key="item.name" :to="item.href" :class="['rounded-md p-2 font-normal text-white ' + (item.current ? 'underline' : 'hover:bg-white hover:text-black')]">{{ item.name }}</RouterLink>
    </nav>
    <button type="button" @click="showCreateEntryModal()" class="btn mt-3 mb-3 bg-[var(--blue)] hover:bg-indigo-700 focus:ring-indigo-500">Nuevo asiento</button>
  </header>
  <CreateEntryModal></CreateEntryModal>
</template>
