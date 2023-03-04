<script setup lang="ts">
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue';
import { ChevronDownIcon } from '@heroicons/vue/20/solid';
import { useDailyLedgerStore } from '@/features/daily-ledger/application/daily-ledger-store';
import { getState, getQuarter, formatAmount, formatDate, formatId, getTotalBalanceColor } from '../../../utils/utils';
import { useEntriesStore } from '@/features/entries/application/entries-store';
import { useMessagesStore } from '@/stores/messages';

const entriesStore = useEntriesStore();
const dailyLedgerStore = useDailyLedgerStore();
const messagesStore = useMessagesStore();
await dailyLedgerStore.loadDailyLedger();
const dailyLedger = dailyLedgerStore.state.storedDailyLedger.reverse();

function publishEntries(entries: any) {
  messagesStore.addConfirmMessage('¿Estás seguro de publicar los asientos seleccionados?', 'Esta es una acción que no se puede revertir, asegúrate de querer publicar este asiento', () => {
    entriesStore.publishEntries(entries);
  });
}
function removeEntries(entries: any) {
  messagesStore.addDangerousConfirmMessage('¿Estás seguro de eliminar los asientos seleccionados?', 'Esta es una acción que no se puede revertir, asegúrate de querer eliminar este asiento', () => {
    entriesStore.removeEntries(entries);
  });
}
function modifyEntries(entries: any) {
  entriesStore.modifyEntries(entries);
}
</script>

<template>
  <div class="px-8">
    <div class="flex items-center">
      <div class="flex-auto">
        <h1 class="text-3xl font-semibold text-gray-900">Libro Diario</h1>
        <p class="mt-2 text-sm text-gray-700">Registro de transacciones financieras diarias ordenadas cronológicamente.</p>
      </div>
      <div class="w-1/2 text-right">
        <button type="button" @click="entriesStore.showCreationModal()" class="btn my-3 mr-24 bg-blue-500 hover:bg-white hover:text-blue-500">Nuevo asiento</button>
      </div>
    </div>
    <div class="mt-8 flex flex-col">
      <div class="-my-2 -mx-8">
        <div class="inline-block min-w-full py-2 px-8 align-middle">
          <div class="shadow ring-1 ring-black ring-opacity-5">
            <table class="min-w-full">
              <thead class="bg-white">
                <tr>
                  <th scope="col" class="sticky top-14 z-10 rounded-tl-lg bg-gray-800 py-3.5 pr-3 pl-6 text-left text-sm font-semibold text-white">#</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Cuenta</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Divisa</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Fecha</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Trimestre</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Concepto</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Debe</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Haber</th>
                  <th scope="col" class="sticky top-14 z-10 rounded-tr-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Estado</th>
                </tr>
              </thead>
              <tbody class="bg-white">
                <template v-for="(record, index) in dailyLedger" :key="index">
                  <tr class="sticky top-24 z-10 bg-gray-800">
                    <th colspan="6" scope="colgroup" class="px-6 text-left text-sm font-semibold text-white">Asiento: {{ formatId(record.id) }}</th>
                    <th scope="colgroup" class="px-6 text-left text-sm font-bold" :class="[getTotalBalanceColor(record.totalDebitAmount, record.totalCreditAmount)]">{{ formatAmount(record.totalDebitAmount, 'EUR') }}</th>
                    <th scope="colgroup" class="px-6 text-left text-sm font-bold" :class="[getTotalBalanceColor(record.totalCreditAmount, record.totalDebitAmount)]">{{ formatAmount(record.totalCreditAmount, 'EUR') }}</th>
                    <th>
                      <Menu as="div" class="relative inline-block p-2 text-left">
                        <div>
                          <MenuButton class="inline-flex w-full justify-center gap-x-1.5 rounded-md bg-blue-500 px-3 py-2 text-sm font-semibold text-white">
                            Acciones
                            <ChevronDownIcon class="-mr-1 h-5 w-5 text-white" aria-hidden="true" />
                          </MenuButton>
                        </div>

                        <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95" enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75" leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
                          <MenuItems class="absolute right-0 z-20 mt-2 w-56 origin-top-right rounded-md bg-gray-800 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                            <div class="py-1">
                              <MenuItem v-slot="{ active }">
                                <a v-on:click="publishEntries(record.entriesExpanded)" :class="[active ? 'bg-gray-700 text-blue-500' : 'text-white', 'block px-4 py-2 text-sm']">Publicar</a>
                              </MenuItem>
                              <MenuItem v-slot="{ active }">
                                <a v-on:click="modifyEntries(record.entriesExpanded)" :class="[active ? 'bg-gray-700 text-blue-500' : 'text-white', 'block px-4 py-2 text-sm']">Modificar</a>
                              </MenuItem>
                              <MenuItem v-slot="{ active }">
                                <a v-on:click="removeEntries(record.entriesExpanded)" :class="[active ? 'bg-gray-700 text-blue-500' : 'text-white', 'block px-4 py-2 text-sm']">Borrar</a>
                              </MenuItem>
                            </div>
                          </MenuItems>
                        </transition>
                      </Menu>
                    </th>
                  </tr>
                  <tr v-for="(entryExpanded, j) in dailyLedger[index].entriesExpanded" :key="j" :class="j % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                    <td class="whitespace-nowrap px-3 py-4 text-xs text-gray-800">
                      {{ formatId(entryExpanded.entry.id) }}
                    </td>
                    <td class="px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.account.displayName }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.currency }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ formatDate(entryExpanded.entry.date) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ getQuarter(entryExpanded.entry.quarter) }}
                    </td>
                    <td class="px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.concept }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm font-medium text-gray-800">
                      {{ formatAmount(entryExpanded.entry.debitAmount, entryExpanded.entry.currency) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm font-medium text-gray-800">
                      {{ formatAmount(entryExpanded.entry.creditAmount, entryExpanded.entry.currency) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      <span :class="[getState(entryExpanded.entry.state) === 'Creado' ? 'bg-yellow-100 text-yellow-800' : getState(entryExpanded.entry.state) === 'Publicado' ? 'bg-green-100 text-green-800' : 'bg-blue-500']" class="inline-flex rounded-full px-2 text-xs font-semibold leading-5">{{ getState(entryExpanded.entry.state) }}</span>
                    </td>
                    <td></td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
