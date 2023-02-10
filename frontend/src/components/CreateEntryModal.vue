<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';
import CreateEntryForm from '@/features/entries/adapters/views/CreateEntryForm.vue';
import { useEntriesStore } from '@/stores/entries';
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';

const modalState = useCreateEntryModalStore();
const entriesStore = useEntriesStore();

function hideCreateEntryModal() {
  entriesStore.clearEntries();
  modalState.hideModal();
}

function createEntries() {
  entriesStore.createEntries();
  modalState.hideModal();
}
</script>

<template>
  <TransitionRoot as="template" :show="modalState.isOpen">
    <Dialog as="div" class="relative z-10" @close="modalState.isOpen = false">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
      </TransitionChild>
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="flex min-h-full items-end items-center justify-center p-4 p-0 text-center">
          <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 translate-y-0 scale-95" enter-to="opacity-100 translate-y-0 scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 scale-100" leave-to="opacity-0 translate-y-4 translate-y-0 scale-95">
            <DialogPanel class="relative my-8 w-full max-w-fit transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all">
              <div class="bg-white p-6 px-4 pt-5 pb-4 pb-4">
                <div class="flex items-start">
                  <div class="mt-3 mt-0 ml-4 text-left text-center">
                    <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">Registrar nuevo asiento contable</DialogTitle>
                    <div class="mt-2">
                      <suspense><CreateEntryForm></CreateEntryForm></suspense>
                    </div>
                  </div>
                </div>
              </div>
              <div class="flex flex-row-reverse bg-gray-50 px-4 py-3 px-6">
                <button type="button" class="ml-3 inline-flex w-full w-auto justify-center rounded-md border border-transparent bg-red-600 px-4 py-2 text-base text-sm font-medium text-white shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2" @click="createEntries">Crear Asiento</button>
                <button type="button" class="mt-3 mt-0 ml-3 inline-flex w-full w-auto justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-base text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" @click="hideCreateEntryModal" ref="cancelButtonRef">Cancelar</button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
