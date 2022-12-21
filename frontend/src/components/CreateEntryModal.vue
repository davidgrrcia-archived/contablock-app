<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';

const modalState = useCreateEntryModalStore();

function hideCreateEntryModal() {
  modalState.hideModal();
}
</script>

<template>
  <Dialog
    header="Registrar nuevo asiento contable"
    v-model:visible="modalState.isOpen"
    :breakpoints="{ '960px': '75vw', '640px': '90vw' }"
    :style="{ width: '50vw' }"
    :modal="true"
  >
    <template #footer>
      <Button label="No" icon="pi pi-times" @click="hideCreateEntryModal" class="p-button-text" />
      <Button label="Yes" icon="pi pi-check" @click="hideCreateEntryModal" autofocus />
    </template>
  </Dialog>
</template>
