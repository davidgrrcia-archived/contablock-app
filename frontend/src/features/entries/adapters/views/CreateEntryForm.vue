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
  <div class="mt-10 sm:mt-0">
    <div class="mt-5 md:mt-0">
      <div class="overflow-hidden shadow sm:rounded-md">
        <div class="bg-white px-4 py-5 sm:p-6">
          <div class="grid grid-cols-6 gap-6">
            <div class="col-span-6 flex sm:col-span-3">
              <label for="currency" class="block text-sm font-medium text-gray-700">Currency</label>
              <select v-model="selectedCurrency" id="currency" name="currency" autocomplete="currency-name" class="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
                <option v-for="(currency, i) in currencies" :key="i">{{ currency.code }}</option>
              </select>

              <p class="text-2xl">Balance: {{ entriesStore.state.totalBalance }}</p>

              <div class="col-span-6">
                <label for="concept" class="block text-sm font-medium text-gray-700">Concepto</label>
                <input type="text" v-model="concept" maxlength="64" name="concept" id="concept" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" />
              </div>

              <div class="col-span-6">
                <label for="calendar" class="block text-sm font-medium text-gray-700">Fecha del asiento</label>
                <CalendarPrimeVue v-model="date" :showIcon="true" dateFormat="dd/mm/yy" name="calendar" id="calendar" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="mt-2 flex space-x-4">
    <div>
      <div>
        <h4 class="mt-10 text-2xl">Debe</h4>
      </div>
      <div class="mt-10 w-1/2 sm:mt-0">
        <div class="mt-5 md:mt-0">
          <div class="overflow-hidden shadow sm:rounded-md">
            <div class="bg-white px-4 py-5 sm:p-6">
              <div class="grid grid-cols-6 gap-6">
                <div class="col-span-6 sm:col-span-3">
                  <label for="debit-account" class="block text-sm font-medium text-gray-700">Cuenta</label>
                  <TreeSelectPrimeVue v-model="debitAccount" :options="accounts" id="account" name="account"></TreeSelectPrimeVue>
                </div>

                <div class="col-span-6">
                  <InputNumberPrimeVue v-model="debitAmount" mode="currency" :currency="selectedCurrency" locale="es-ES" />
                </div>
              </div>
            </div>
            <div class="bg-gray-50 px-4 py-3 text-right sm:px-6">
              <button type="submit" @click="addDebitEntry" class="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">Añadir al debe</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div>
      <div>
        <h4 class="mt-10 text-2xl">Haber</h4>
      </div>
      <div class="mt-10 w-1/2 sm:mt-0">
        <div class="mt-5 md:mt-0">
          <div class="overflow-hidden shadow sm:rounded-md">
            <div class="bg-white px-4 py-5 sm:p-6">
              <div class="grid grid-cols-6 gap-6">
                <div class="col-span-6 sm:col-span-3">
                  <label for="debit-account" class="block text-sm font-medium text-gray-700">Cuenta</label>
                  <TreeSelectPrimeVue v-model="creditAccount" :options="accounts" id="account" name="account"></TreeSelectPrimeVue>
                </div>

                <div class="col-span-6">
                  <InputNumberPrimeVue v-model="creditAmount" mode="currency" :currency="selectedCurrency" locale="es-ES" />
                </div>
              </div>
            </div>
            <div class="bg-gray-50 px-4 py-3 text-right sm:px-6">
              <button type="submit" @click="addCreditEntry" class="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">Añadir al haber</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="form-table">
    <DataTablePrimeVue :value="entriesStore.state.entriesCreation" responsiveLayout="scroll">
      <ColumnPrimeVue field="displayName" header="Cuenta"></ColumnPrimeVue>
      <ColumnPrimeVue field="currency" header="Divisa"></ColumnPrimeVue>
      <ColumnPrimeVue field="displayDate" header="Fecha"></ColumnPrimeVue>
      <ColumnPrimeVue field="concept" header="Concepto"></ColumnPrimeVue>
      <ColumnPrimeVue field="debitAmount" header="Debe"></ColumnPrimeVue>
      <ColumnPrimeVue field="creditAmount" header="Haber"></ColumnPrimeVue>
    </DataTablePrimeVue>
  </div>
</template>
