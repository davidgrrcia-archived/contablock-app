<script setup lang="ts">
import { useGeneralLedgerStore } from '@/features/general-ledger/application/general-ledger-store';
import { getState, getQuarter, formatAmount, formatDate, getEntryStateColor } from '../../../utils/utils';
import { ref, watch, onMounted } from 'vue';

const generalLedgerStore = useGeneralLedgerStore();
await generalLedgerStore.loadGeneralLedger();

const generalLedger = ref();

watch(
  () => generalLedgerStore.state.storedGeneralLedger,
  newStoredGeneralLedger => {
    generalLedger.value = newStoredGeneralLedger;
  },
);

onMounted(async () => {
  await generalLedgerStore.loadGeneralLedger();
  generalLedger.value = generalLedgerStore.state.storedGeneralLedger;
});

function getDebitTotalAmount(totalDebitAmount: any, totalCreditAmount: any) {
  totalDebitAmount = totalDebitAmount || 0;
  totalCreditAmount = totalCreditAmount || 0;

  return totalDebitAmount > totalCreditAmount ? formatAmount(totalDebitAmount - totalCreditAmount, 'EUR') : '';
}

function getCreditTotalAmount(totalDebitAmount: any, totalCreditAmount: any) {
  totalDebitAmount = totalDebitAmount || 0;
  totalCreditAmount = totalCreditAmount || 0;

  return totalCreditAmount > totalDebitAmount ? formatAmount(totalCreditAmount - totalDebitAmount, 'EUR') : '';
}
</script>

<template>
  <div class="m-8 flex flex-col">
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
                <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Saldo deudor</th>
                <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Saldo acreedor</th>
                <th scope="col" class="sticky top-14 z-10 rounded-tr-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Estado</th>
              </tr>
            </thead>
            <tbody class="bg-white">
              <template v-for="(record, index) in generalLedger" :key="index">
                <tr class="sticky top-24 z-10 bg-gray-700">
                  <th colspan="4" scope="colgroup" class="px-3 py-4 text-left text-sm font-semibold text-white">{{ record.account.displayName }}</th>
                  <th scope="colgroup" class="px-3 py-4 text-left text-sm font-semibold text-white">{{ formatAmount(record.totalDebitAmount, 'EUR') }}</th>
                  <th scope="colgroup" class="px-3 py-4 text-left text-sm font-semibold text-white">{{ formatAmount(record.totalCreditAmount, 'EUR') }}</th>
                  <th scope="colgroup" class="px-3 py-4 text-left text-sm font-semibold text-white">{{ getDebitTotalAmount(record.totalDebitAmount, record.totalCreditAmount) }}</th>
                  <th scope="colgroup" class="px-3 py-4 text-left text-sm font-semibold text-white">{{ getCreditTotalAmount(record.totalDebitAmount, record.totalCreditAmount) }}</th>
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
                  <td></td>
                  <td></td>
                  <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                    <span :class="[getEntryStateColor(entryExpanded.state)]" class="inline-flex rounded-full px-2 text-xs font-semibold leading-5">{{ getState(entryExpanded.state) }}</span>
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
</template>
