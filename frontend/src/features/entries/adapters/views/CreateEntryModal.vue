<script setup lang="ts">
import CreateEntryForm from '@/features/entries/adapters/views/CreateEntryForm.vue';
import { useEntriesStore } from '@/features/entries/application/entries-store.js';
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';

const entriesStore = useEntriesStore();
</script>

<template>
  <TransitionRoot as="template" :show="entriesStore.state.creationModalOpen">
    <Dialog as="div" class="relative z-10" @close="entriesStore.state.creationModalOpen = false">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
      </TransitionChild>
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="flex min-h-full items-end justify-center p-4 text-center">
          <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 translate-y-0 scale-95" enter-to="opacity-100 translate-y-0 scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 scale-100" leave-to="opacity-0 translate-y-4 translate-y-0 scale-95">
            <DialogPanel class="relative my-8 w-full max-w-full transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all">
              <div class="bg-white p-6 px-4 pt-5 pb-4">
                <div class="mt-3 ml-4 text-left">
                  <div class="grid">
                    <div>
                      <DialogTitle as="h3" class="text-2xl font-medium leading-6 text-gray-700">Registrar nuevo asiento contable</DialogTitle>
                    </div>
                    <div class="-mt-10 text-right">
                      <button type="button" class="mt-3 ml-3 inline-flex w-auto justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-base font-medium text-gray-700 shadow-sm hover:bg-gray-50" @click="entriesStore.clearEntriesAndHideModals()" ref="cancelButtonRef">Cancelar</button>
                      <button type="button" class="ml-3 inline-flex w-auto justify-center rounded-md border border-transparent bg-blue-500 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-gray-700" @click="entriesStore.createEntries()">Crear Asiento</button>
                    </div>
                  </div>
                  <div class="mt-2">
                    <suspense><CreateEntryForm></CreateEntryForm></suspense>
                  </div>
                </div>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
