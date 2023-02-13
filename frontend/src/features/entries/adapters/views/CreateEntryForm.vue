<script setup lang="ts">
import { ref } from 'vue';
import { v4 as uuid } from 'uuid';
import type { Entry } from '@/features/entries/application/entries-model';
import { useEntriesStore } from '@/stores/entries';
import { useAccountsStore } from '@/stores/accounts';

//store
const entriesStore = useEntriesStore();
const accountsStore = useAccountsStore();
await accountsStore.loadAccounts();

const selectedCurrency = ref('EUR');
const currentDate = new Date();
const date = ref(currentDate);
const concept = ref();
const debitAmount = ref();
const creditAmount = ref();
const id = uuid();
const multipleEntryId = uuid();
const debitAccount = ref();
const creditAccount = ref();
const accounts = accountsStore.state.accounts;

if (entriesStore.state.multipleEntryId === '') {
  entriesStore.state.multipleEntryId = multipleEntryId;
}

function addDebitEntry() {
  const newEntry: Entry = {
    id: id,
    accountId: debitAccount.value,
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: selectedCurrency.value,
    date: date.value.toISOString(),
    displayDate: date.value.toLocaleDateString(),
    concept: concept.value,
    debitAmount: debitAmount.value,
    creditAmount: null,
  };

  entriesStore.addDebitEntry(newEntry);
}

function addCreditEntry() {
  const newEntry: Entry = {
    id: id,
    accountId: creditAccount.value,
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: selectedCurrency.value,
    date: date.value.toISOString(),
    displayDate: date.value.toLocaleDateString(),
    concept: concept.value,
    debitAmount: null,
    creditAmount: creditAmount.value,
  };

  entriesStore.addCreditEntry(newEntry);
}

const currencies = ref([
  {
    code: 'EUR',
    symbol: '€',
  },
  {
    code: 'GBP',
    symbol: '£',
  },
  {
    code: 'USD',
    symbol: '$',
  },
]);

function findCurrencyByCode(code: string): string {
  const c = currencies.value.find(currency => currency.code === code);

  return c ? c.symbol : '';
}
</script>

<template>
  <div class="mt-0 flex w-full space-x-10 overflow-hidden rounded-md p-6 text-left shadow">
    <div class="w-1/4">
      <label for="currency" class="block text-sm font-medium text-gray-700">Divisa</label>
      <select v-model="selectedCurrency" id="currency" name="currency" autocomplete="currency-name" class="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm sm:text-sm">
        <option v-for="(currency, i) in currencies" :key="i">{{ currency.code }}</option>
      </select>
    </div>
    <div class="w-1/4">
      <label for="concept" class="block text-sm font-medium text-gray-700">Concepto</label>
      <input type="text" v-model="concept" maxlength="64" name="concept" id="concept" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm sm:text-sm" />
    </div>
    <div class="w-1/4">
      <label for="calendar" class="block text-sm font-medium text-gray-700">Fecha del asiento</label>
      <CalendarPrimeVue v-model="date" :showIcon="true" dateFormat="dd/mm/yy" name="calendar" id="calendar" />
    </div>
    <div class="w-1/4">
      <label for="balance" class="block text-sm font-medium text-gray-700">Balance del asiento</label>
      <p id="balance" name="balance" class="text-xl">{{ entriesStore.state.totalBalance }}</p>
    </div>
  </div>

  <div class="mt-2 mb-10 flex space-x-4 text-left">
    <div class="w-1/2">
      <div class="mt-0 overflow-hidden rounded-md shadow">
        <div class="bg-gray-700 py-3 px-6 text-white">
          <h4 class="text-2xl">Debe</h4>
        </div>
        <div class="space-y-4 bg-white p-6">
          <div>
            <label for="debit-account" class="block text-sm font-medium text-gray-700">Cuenta</label>
            <TreeSelectPrimeVue v-model="debitAccount" :options="accounts" placeholder="Selecciona una cuenta para añadirla al debe" id="account" name="account"></TreeSelectPrimeVue>
          </div>
          <div>
            <label for="debit-amount" class="block text-sm font-medium text-gray-700">Importe</label>
            <div class="relative mt-1 rounded-md shadow-sm">
              <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                <span class="text-gray-500 sm:text-sm">{{ findCurrencyByCode(selectedCurrency) }}</span>
              </div>
              <input v-model="debitAmount" type="text" name="debit-amount" id="debit-amount" class="block w-full rounded-md border-gray-300 pl-7 pr-12 focus:border-blue-500 focus:ring-blue-500 sm:text-sm" placeholder="0.00" />
              <div class="absolute inset-y-0 right-0 flex items-center">
                <label for="currency" class="sr-only">Divisa</label>
                <select v-model="selectedCurrency" id="currency" name="currency" class="h-full rounded-md border-transparent bg-transparent py-0 pl-2 pr-7 text-gray-500 focus:border-blue-500 focus:ring-blue-500 sm:text-sm">
                  <option v-for="(currency, i) in currencies" :key="i">{{ currency.code }}</option>
                </select>
              </div>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 py-3 px-6 text-right">
          <button type="submit" id="credit-amount" name="credit-amount" @click="addDebitEntry" class="inline-flex w-full justify-center rounded-md bg-blue-500 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-gray-700">Añadir al debe</button>
        </div>
      </div>
    </div>

    <div class="w-1/2">
      <div class="mt-0 overflow-hidden rounded-md shadow">
        <div class="bg-gray-700 py-3 px-6 text-white">
          <h4 class="text-right text-2xl">Haber</h4>
        </div>
        <div class="space-y-4 bg-white p-6">
          <div>
            <label for="debit-account" class="block text-sm font-medium text-gray-700">Cuenta</label>
            <TreeSelectPrimeVue v-model="creditAccount" :options="accounts" placeholder="Selecciona una cuenta para añadirla al haber" id="account" name="account"></TreeSelectPrimeVue>
          </div>
          <div>
            <label for="credit-amount" class="block text-sm font-medium text-gray-700">Importe</label>
            <div class="relative mt-1 rounded-md shadow-sm">
              <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                <span class="text-gray-500 sm:text-sm">{{ findCurrencyByCode(selectedCurrency) }}</span>
              </div>
              <input v-model="creditAmount" type="text" name="credit-amount" id="credit-amount" class="block w-full rounded-md border-gray-300 pl-7 pr-12 focus:border-blue-500 focus:ring-blue-500 sm:text-sm" placeholder="0.00" />
              <div class="absolute inset-y-0 right-0 flex items-center">
                <label for="currency" class="sr-only">Divisa</label>
                <select v-model="selectedCurrency" id="currency" name="currency" class="h-full rounded-md border-transparent bg-transparent py-0 pl-2 pr-7 text-gray-500 focus:border-blue-500 focus:ring-blue-500 sm:text-sm">
                  <option v-for="(currency, i) in currencies" :key="i">{{ currency.code }}</option>
                </select>
              </div>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 py-3 px-6 text-right">
          <button type="submit" id="credit-amount" name="credit-amount" @click="addCreditEntry" class="inline-flex w-full justify-center rounded-md bg-blue-500 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-gray-700">Añadir al haber</button>
        </div>
      </div>
    </div>
  </div>

  <div class="-mx-8 mt-8 flex overflow-x-auto">
    <div class="inline-block min-w-full py-2 px-8 align-middle">
      <div class="overflow-hidden rounded-lg shadow ring-1 ring-black ring-opacity-5">
        <table class="min-w-full divide-y divide-gray-300">
          <thead class="bg-gray-700">
            <tr>
              <th scope="col" class="whitespace-nowrap py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-white sm:pl-6">Cuenta</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Divisa</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Fecha</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Concepto</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Debe</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Haber</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200 bg-white">
            <tr v-for="(entry, i) in entriesStore.state.entriesCreation.reverse()" :key="entry.id" :class="i % 2 === 0 ? undefined : 'bg-gray-50'">
              <td class="whitespace-nowrap py-2 pl-4 pr-3 text-sm text-gray-500 sm:pl-6">{{ entry.displayName }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm font-medium text-gray-900">{{ entry.currency }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-900">{{ entry.displayDate }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-500">{{ entry.concept }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-500">{{ entry.debitAmount }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-500">{{ entry.creditAmount }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style>
.p-treeselect {
  @apply w-full rounded-md border-gray-300 pl-7 pr-12 text-sm focus:border-blue-500 focus:ring-blue-500;
}

.p-treenode-content {
  @apply bg-gray-100 text-sm text-gray-700;
}

.p-calendar-w-btn .p-datepicker-trigger {
  @apply border-transparent bg-gray-700 text-sm text-white hover:bg-gray-700 focus:border-gray-700 focus:ring-gray-700;
}
</style>
