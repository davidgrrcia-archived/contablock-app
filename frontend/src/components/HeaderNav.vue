<template>
  <div class="columns-3 bg-gray-800">
    <div class="w-full p-5">
      <img
        class="h-8 w-auto object-right"
        src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
        alt="Your Company"
      />
    </div>
    <div class="w-full p-5 space-x-4">
      <a
        v-for="item in navigation"
        :key="item.name"
        :href="item.href"
        :class="[
          item.current
            ? 'text-amber-400 underline decoration-2 underline-offset-8'
            : 'text-white hover:bg-amber-400 hover:text-black',
          'px-3 py-2 rounded-md text-sm font-medium',
        ]"
        :aria-current="item.current ? 'page' : undefined"
        >{{ item.name }}</a
      >
    </div>
    <div class="w-full p-4">
      <button
        class="px-3 py-2 text-sm font-medium rounded-md bg-amber-400 text-gray-800"
        type="button"
        @click="showCreateEntryModal()"
      >
        New Entry
      </button>
    </div>
  </div>
  <CreateEntryModal></CreateEntryModal>
</template>

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
