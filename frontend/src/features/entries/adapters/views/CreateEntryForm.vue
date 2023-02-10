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
  const match = JSON.stringify(debitAccount.value).match(/"(\d+)"/);
  const accountId = match ? match[1] : '';

  const newEntry: Entry = {
    id: id,
    accountId: accountId,
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
  const match = JSON.stringify(creditAccount.value).match(/"(\d+)"/);
  const accountId = match ? match[1] : '';

  const newEntry: Entry = {
    id: id,
    accountId: accountId,
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
  },
  {
    code: 'GBP',
  },
  {
    code: 'USD',
  },
]);
</script>

<template>
  <div class="mt-0 flex space-x-10 overflow-hidden rounded-md p-6 shadow">
    <div class="w-1/4">
      <label for="currency" class="block text-sm font-medium text-gray-700">Currency</label>
      <select v-model="selectedCurrency" id="currency" name="currency" autocomplete="currency-name" class="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
        <option v-for="(currency, i) in currencies" :key="i">{{ currency.code }}</option>
      </select>
    </div>
    <div class="w-1/4">
      <label for="concept" class="block text-sm font-medium text-gray-700">Concepto</label>
      <input type="text" v-model="concept" maxlength="64" name="concept" id="concept" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" />
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

  <div class="mt-2 mb-10 flex space-x-4">
    <div class="w-1/2">
      <div class="mt-0 overflow-hidden rounded-md shadow">
        <div class="bg-gray-50 py-3 px-6">
          <h4 class="text-2xl">Debe</h4>
        </div>
        <div class="space-y-4 bg-white p-6">
          <div>
            <label for="debit-account" class="block text-sm font-medium text-gray-700">Cuenta</label>
            <TreeSelectPrimeVue v-model="debitAccount" :options="accounts" id="account" name="account"></TreeSelectPrimeVue>
          </div>
          <div>
            <label for="debit-amount" class="block text-sm font-medium text-gray-700">Importe</label>
            <InputNumberPrimeVue v-model="debitAmount" mode="currency" :currency="selectedCurrency" locale="es-ES" />
          </div>
        </div>
        <div class="bg-gray-50 py-3 px-6">
          <button type="submit" id="debit-amount" name="debit-amount" @click="addDebitEntry" class="inline-flex justify-center rounded-md border border-transparent bg-[var(--blue)] hover:bg-white hover:text-[var(--blue)] py-2 px-4 text-sm font-medium text-white shadow-sm">Añadir al debe</button>
        </div>
      </div>
    </div>

    <div class="w-1/2">
      <div class="mt-0 overflow-hidden rounded-md text-right shadow">
        <div class="bg-gray-50 py-3 px-6">
          <h4 class="text-2xl">Haber</h4>
        </div>
        <div class="space-y-4 bg-white p-6">
          <div>
            <label for="debit-account" class="block text-sm font-medium text-gray-700">Cuenta</label>
            <TreeSelectPrimeVue v-model="creditAccount" :options="accounts" id="account" name="account"></TreeSelectPrimeVue>
          </div>
          <div>
            <label for="credit-amount" class="block text-sm font-medium text-gray-700">Importe</label>
            <InputNumberPrimeVue v-model="creditAmount" mode="currency" :currency="selectedCurrency" locale="es-ES" />
          </div>
        </div>
        <div class="bg-gray-50 py-3 px-6">
          <button type="submit" id="credit-amount" name="credit-amount" @click="addCreditEntry" class="inline-flex justify-center rounded-md border border-transparent bg-[var(--blue)] hover:bg-white hover:text-[var(--blue)] py-2 px-4 text-sm font-medium text-white shadow-sm">Añadir al haber</button>
        </div>
      </div>
    </div>
  </div>

  <div class="-mx-8 mt-8 flex overflow-x-auto">
    <div class="inline-block min-w-full py-2 px-8 align-middle">
      <div class="overflow-hidden rounded-lg shadow ring-1 ring-black ring-opacity-5">
        <table class="min-w-full divide-y divide-gray-300">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="whitespace-nowrap py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">Cuenta</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-gray-900">Divisa</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-gray-900">Fecha</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-gray-900">Concepto</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-gray-900">Debe</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-gray-900">Haber</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200 bg-white">
            <tr v-for="(entry, i) in entriesStore.state.entriesCreation" :key="entry.id" :class="i % 2 === 0 ? undefined : 'bg-gray-50'">
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
