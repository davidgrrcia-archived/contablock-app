<template>
  <div class="relative">
    <div v-if="modalState.isOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div
        class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0"
      >
        <div
          class="relative rounded-lg shadow-xl w-11/12 sm:w-3/4 lg:w-1/2 xl:w-1/3"
        >
          <button
            @click="hideCreateEntryModal()"
            class="text-3xl leading-none absolute top-1/2 right-0 m-4 text-gray-500 hover:text-red-500"
          >
            &times;
          </button>
          <div class="bg-white rounded-lg py-4 px-6 mb-4">
            <h3 class="text-2xl font-bold mb-2 text-gray-800">Modal Title</h3>
            <p class="text-gray-600 mb-8">Modal content goes here</p>
            <button
              @click="hideCreateEntryModal()"
              class="px-4 py-2 font-bold text-white bg-blue-500 rounded-full hover:bg-blue-700 focus:outline-none focus:shadow-outline-blue active:bg-blue-800"
            >
              Close
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';

const modalState = useCreateEntryModalStore();

function hideCreateEntryModal() {
  modalState.hideModal();
}
</script>
