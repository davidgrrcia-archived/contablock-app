<script setup lang="ts">
import { useDailyLedgerStore } from '@/features/daily-ledger/application/daily-ledger-store';
import { getState, getQuarter, formatAmount, formatDate, formatId, getTotalBalanceColor } from '../../../utils/utils';
import { useEntriesStore } from '@/features/entries/application/entries-store';
import { useMessagesStore } from '@/stores/messages';
import { useCreateEntryModalStore } from '@/features/entries/application/create-entry-modal-store';

const entriesStore = useEntriesStore();
const dailyLedgerStore = useDailyLedgerStore();
const messagesStore = useMessagesStore();
await dailyLedgerStore.loadDailyLedger();
const dailyLedger = dailyLedgerStore.state.storedDailyLedger.reverse();
const modalState = useCreateEntryModalStore();

function showCreateEntryModal() {
  modalState.showModal();
}

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
</script>

<template>
  <div class="px-8">
    <div class="flex items-center">
      <div class="flex-auto">
        <h1 class="text-3xl font-semibold text-gray-900">Libro Diario</h1>
        <p class="mt-2 text-sm text-gray-700">Registro de transacciones financieras diarias ordenadas cronológicamente.</p>
      </div>
      <div class="w-1/2 text-right">
        <button type="button" @click="showCreateEntryModal()" class="btn my-3 mr-24 bg-blue-500 hover:bg-white hover:text-blue-500">Nuevo asiento</button>
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
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Estado</th>
                  <th scope="col" class="sticky top-14 z-10 rounded-tr-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Acciones</th>
                </tr>
              </thead>
              <tbody class="bg-white">
                <template v-for="(record, index) in dailyLedger" :key="index">
                  <tr class="sticky top-24 z-10 bg-gray-600">
                    <th colspan="6" scope="colgroup" class="px-6 text-left text-sm font-semibold text-white">Asiento: {{ formatId(record.id) }}</th>
                    <th scope="colgroup" class="px-6 text-left text-sm font-semibold" :class="[getTotalBalanceColor(record.totalDebitAmount, record.totalCreditAmount)]">{{ formatAmount(record.totalDebitAmount, 'EUR') }}</th>
                    <th scope="colgroup" class="px-6 text-left text-sm font-semibold" :class="[getTotalBalanceColor(record.totalDebitAmount, record.totalCreditAmount)]">{{ formatAmount(record.totalCreditAmount, 'EUR') }}</th>
                    <th></th>
                    <th>
                      <div v-if="record.entriesExpanded[0].entry.state == 'CREATED'" class="flex space-x-4 py-2 px-2">
                        <button type="button" class="btn bg-blue-500" v-on:click="publishEntries(record.entriesExpanded)">Publicar</button>
                        <button type="button" class="btn bg-red-500" v-on:click="removeEntries(record.entriesExpanded)">Eliminar</button>
                      </div>
                      <div v-else class="flex space-x-4 py-2 px-2">
                        <span class="p-5"></span>
                      </div>
                    </th>
                  </tr>
                  <tr v-for="(entryExpanded, j) in dailyLedger[index].entriesExpanded" :key="j" :class="j % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                    <td class="whitespace-nowrap px-3 py-4 text-xs text-gray-800">
                      {{ formatId(entryExpanded.entry.id) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
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
