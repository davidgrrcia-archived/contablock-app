<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';
import CreateEntryForm from '@/features/entries/adapters/views/CreateEntryForm.vue';
import { useEntriesStore } from '@/stores/entries';
import { useMessagesStore } from '@/stores/messages';

const modalState = useCreateEntryModalStore();
const entriesStore = useEntriesStore();
const messagesStore = useMessagesStore();

function hideCreateEntryModal() {
  modalState.hideModal();
}

function createEntries() {
  if (entriesStore.state.totalBalance !== 0) {
    messagesStore.addErrorMessage(
      'El principio de partida doble es obligatorio, el balance actual es: ' + entriesStore.state.totalBalance,
    );
  }
  if (entriesStore.state.entriesCreation.length === 0) {
    messagesStore.addErrorMessage('No hay asientos creados');
    return;
  }

  entriesStore.createEntries();
  modalState.hideModal();
}
</script>

<template>
  <DialogPrimeVue
    header="Registrar nuevo asiento contable"
    v-model:visible="modalState.isOpen"
    :breakpoints="{ '960px': '75vw', '640px': '90vw' }"
    :style="{ width: '90vw' }"
    :modal="true"
  >
    <suspense><CreateEntryForm></CreateEntryForm></suspense>
    <template #footer>
      <ButtonPrimeVue label="Cancelar" icon="pi pi-times" @click="hideCreateEntryModal" class="btn-cancel" />
      <ButtonPrimeVue label="Crear Asiento" icon="pi pi-check" @click="createEntries" class="btn-create" />
    </template>
  </DialogPrimeVue>
</template>

<style scoped>
.btn-cancel {
  color: var(--red);
  background-color: var(--white) !important;
}
.btn-cancel:hover {
  color: var(--white);
  background-color: var(--red) !important;
}
.btn-create {
  background-color: var(--yellow) !important;
}
</style>
