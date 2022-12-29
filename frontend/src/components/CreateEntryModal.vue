<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';
import CreateEntryForm from '@/features/entries/adapters/views/CreateEntryForm.vue';

const modalState = useCreateEntryModalStore();

function hideCreateEntryModal() {
  modalState.hideModal();
}
</script>

<template>
  <DialogPrimeVue
    header="Registrar nuevo asiento contable"
    v-model:visible="modalState.isOpen"
    :breakpoints="{ '960px': '75vw', '640px': '90vw' }"
    :style="{ width: '60vw' }"
    :modal="true"
  >
    <suspense><CreateEntryForm></CreateEntryForm></suspense>
    <template #footer>
      <ButtonPrimeVue label="No" icon="pi pi-times" @click="hideCreateEntryModal" class="p-button-text" />
      <ButtonPrimeVue label="Yes" icon="pi pi-check" @click="hideCreateEntryModal" autofocus />
    </template>
  </DialogPrimeVue>
</template>

<style scoped></style>
