<script setup lang="ts">
import { useGeneralLedgerStore } from '@/features/general-ledger/application/general-ledger-store';
import { getState, getQuarter, formatAmount, formatDate } from '../../../utils/utils';

const generalLedgerStore = useGeneralLedgerStore();
await generalLedgerStore.loadGeneralLedger();
const generalLedger = generalLedgerStore.state.storedGeneralLedger.reverse();
</script>

<template>
  <div class="px-8">
    <div class="flex items-center">
      <div class="flex-auto">
        <h1 class="text-3xl font-semibold text-gray-900">Libro Mayor</h1>
        <p class="mt-2 text-sm text-gray-700">Registro de débitos y créditos de una empresa clasificados por cuentas contables.</p>
      </div>
    </div>
    <div class="mt-8 flex flex-col">
      <div class="-my-2 -mx-8">
        <div class="inline-block min-w-full py-2 px-8 align-middle">
          <div class="shadow ring-1 ring-black ring-opacity-5">
            <table class="min-w-full">
              <thead class="bg-white">
                <tr>
                  <th scope="col" class="sticky top-14 z-10 rounded-tl-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Divisa</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Fecha</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Trimestre</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Concepto</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Debe</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Haber</th>
                  <th scope="col" class="sticky top-14 z-10 rounded-br-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Estado</th>
                </tr>
              </thead>
              <tbody class="bg-white">
                <template v-for="(record, index) in generalLedger" :key="index">
                  <tr class="sticky top-24 z-10 bg-gray-600">
                    <th colspan="4" scope="colgroup" class="px-6 py-5 text-left text-sm font-semibold text-white">{{ record.account.displayName }}</th>
                    <th scope="colgroup" class="px-6 py-5 text-left text-sm font-semibold text-white">{{ formatAmount(record.totalDebitAmount, 'EUR') }}</th>
                    <th scope="colgroup" class="px-6 py-5 text-left text-sm font-semibold text-white">{{ formatAmount(record.totalCreditAmount, 'EUR') }}</th>
                    <th colspan="2"></th>
                  </tr>
                  <tr v-for="(entryExpanded, j) in generalLedger[index].entries" :key="j" :class="j % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.currency }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ formatDate(entryExpanded.date) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ getQuarter(entryExpanded.quarter) }}
                    </td>
                    <td class="px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.concept }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm font-medium text-gray-800">
                      {{ formatAmount(entryExpanded.debitAmount, entryExpanded.currency) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm font-medium text-gray-800">
                      {{ formatAmount(entryExpanded.creditAmount, entryExpanded.currency) }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      <span :class="[getState(entryExpanded.state) === 'Creado' ? 'bg-yellow-100 text-yellow-800' : getState(entryExpanded.state) === 'Publicado' ? 'bg-green-100 text-green-800' : 'bg-blue-500']" class="inline-flex rounded-full px-2 text-xs font-semibold leading-5">{{ getState(entryExpanded.state) }}</span>
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
