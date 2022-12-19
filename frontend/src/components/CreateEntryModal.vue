<template>
  <div class="relative">
    <div
      v-if="modalState.isOpen"
      class="fixed inset-0 z-50 overflow-y-auto bg-gray-800/25"
    >
      <div class="p-28 text-center">
        <div class="bg-white rounded-md p-4">
          <h3 class="text-2xl font-bold mb-2 text-gray-800">
            Create New Entry
          </h3>

          <button
            @click="hideCreateEntryModal()"
            class="p-4 text-sm font-bold text-white bg-red-500 rounded-md hover:bg-gray-800"
          >
            Cancel
          </button>
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
