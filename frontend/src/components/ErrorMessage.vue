<script setup lang="ts">
import { useMessagesStore, Severities } from '@/stores/messages';
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
import { ExclamationTriangleIcon, CheckCircleIcon } from '@heroicons/vue/24/outline';

const errorMessagesStore = useMessagesStore();

function mapMessageTypesToColors(severity: Severities) {
  switch (severity) {
    case Severities.error:
    case Severities.dangerousConfirm:
      return 'bg-red-100 border-red-400 text-red-700';
    case Severities.success:
      return 'bg-green-100 border-green-400 text-green-700';
    case Severities.info:
      return 'bg-blue-100 border-blue-400 text-blue-700';
    case Severities.warning:
      return 'bg-yellow-100 border-yellow-400 text-yellow-700';
    case Severities.confirm:
      return 'bg-indigo-100 border-indigo-400 text-indigo-700';
    default:
      return 'bg-red-100 border-red-400 text-red-700';
  }
}
</script>

<template>
  <div v-for="msg of errorMessagesStore.state.messages" :key="msg.id">
    <TransitionRoot as="template" :show="msg.show">
      <Dialog as="div" class="relative z-10" @close="msg.show = false">
        <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </TransitionChild>

        <div class="fixed inset-0 z-10 overflow-y-auto">
          <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                  <div class="sm:flex sm:items-start">
                    <div v-if="msg.severity == Severities.success" class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-green-500 sm:mx-0 sm:h-10 sm:w-10">
                      <CheckCircleIcon class="h-6 w-6 bg-green-500 text-white" aria-hidden="true" />
                    </div>
                    <div v-else :class="[mapMessageTypesToColors(msg.severity)]" class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full sm:mx-0 sm:h-10 sm:w-10">
                      <ExclamationTriangleIcon :class="[mapMessageTypesToColors(msg.severity)]" class="h-6 w-6" aria-hidden="true" />
                    </div>
                    <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                      <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">{{ msg.title }}</DialogTitle>
                      <div class="mt-2">
                        <p class="text-sm text-gray-700">{{ msg.description }}</p>
                      </div>
                    </div>
                  </div>
                </div>
                <div v-if="msg.severity == Severities.confirm || msg.severity == Severities.dangerousConfirm" class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                  <button type="button" :class="[mapMessageTypesToColors(msg.severity)]" class="inline-flex w-full justify-center rounded-md border border-transparent px-4 py-2 text-base font-medium shadow-sm sm:ml-3 sm:w-auto sm:text-sm" @click="msg.actionCallback">Aceptar</button>
                  <button type="button" class="inline-flex w-full justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-base font-medium text-gray-700 shadow-sm hover:bg-gray-50 sm:ml-3 sm:w-auto sm:text-sm" @click="errorMessagesStore.clearMessages()">Cancelar</button>
                </div>
                <div v-else class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                  <button type="button" :class="[mapMessageTypesToColors(msg.severity)]" class="inline-flex w-full justify-center rounded-md border border-transparent px-4 py-2 text-base font-medium shadow-sm sm:ml-3 sm:w-auto sm:text-sm" @click="errorMessagesStore.clearMessages()">Aceptar</button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>
