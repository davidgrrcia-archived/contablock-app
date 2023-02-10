<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';
import CreateEntryModal from '@/components/CreateEntryModal.vue';
import { useConfigStore } from '@/stores/config';
import {
  BriefcaseIcon,
  LinkIcon,
  MapPinIcon,
} from '@heroicons/vue/20/solid'

const configStore = useConfigStore();
const modalState = useCreateEntryModalStore();

function showCreateEntryModal() {
  modalState.showModal();
}
</script>

<template>
  <header class="fixed top-0 left-0 flex w-full items-center bg-gray-800">
    <div class="flex w-1/2 items-center">
      <img alt="Vue logo" class="ml-5 object-cover" src="@/assets/logo.png" />
      <RouterLink class="text-lg" v-for="item in configStore.state.navigation" :key="item.name" :to="item.href" :class="['rounded-md p-2 font-normal text-white ' + (item.current ? 'underline' : 'hover:bg-white hover:text-black')]">{{ item.name }}</RouterLink>
    </div>
    <div class="w-1/2 text-right">
      <button type="button" @click="showCreateEntryModal()" class="btn my-3 mr-24 bg-[var(--blue)] hover:bg-white hover:text-[var(--blue)]">Nuevo asiento</button>
    </div>
  </header>
  <CreateEntryModal></CreateEntryModal>

  <div class="mt-18 mb-5 flex items-center justify-between">
    <div class="min-w-0 flex-1 bg-gray-100 pt-24 pb-10 text-center">
      <h2 class="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">Zara España, SA</h2>
      <div class="grid justify-items-center ">
        <div class="flex mt-0  space-x-6">
          <div class="mt-2 flex items-center text-sm text-gray-500">
            <BriefcaseIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
            A15022510
          </div>
          <div class="mt-2 flex items-center text-sm text-gray-500">
            <MapPinIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
            Arteixo
          </div>
          <div class="mt-2 flex items-center text-sm text-gray-500">
            <LinkIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
            0xb794f5ea0ba39494ce839613fffba74279579268
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
