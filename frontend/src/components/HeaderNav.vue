<script setup lang="ts">
import CreateEntryModal from '@/features/entries/adapters/views/CreateEntryModal.vue';
import { useConfigStore } from '@/stores/config';
import BusinessHeading from '@/features/business/adapters/views/BusinessHeading.vue';

const configStore = useConfigStore();

function logOut() {
  localStorage.setItem(configStore.state.project + 'Authentication', JSON.stringify(false));
  window.location.reload();
}
</script>

<template>
  <header class="fixed top-0 left-0 flex w-full items-center bg-gray-800">
    <div class="flex w-1/2 items-center">
      <img alt="Vue logo" class="ml-5 object-cover" src="@/assets/logo.png" />
      <RouterLink class="text-lg" v-for="item in configStore.state.navigation" :key="item.name" :to="item.href" :class="['rounded-md p-2 font-normal text-white underline-offset-4 ' + (item.current ? 'underline' : 'hover:underline')]">{{ item.name }}</RouterLink>
    </div>
    <div class="w-1/2 text-right">
      <button type="button" @click="logOut()" class="btn my-3 mr-24 bg-blue-500 hover:bg-white hover:text-blue-500">Cerrar sesión</button>
    </div>
  </header>
  <CreateEntryModal></CreateEntryModal>
  <suspense><BusinessHeading></BusinessHeading></suspense>
</template>
