<script setup lang="ts">
import { ref } from 'vue';
import { v4 as uuid } from 'uuid';
import type { CreateEntryInput } from '@/features/entries/application/entries-model';
import { useEntriesStore } from '@/features/entries/application/entries-store.js';
import { useAccountsStore } from '@/features/accounts/application/accounts-store.js';
import { findCurrencyByCode, currencies, dateToUnixTimestamp, formatDate, formatAmount } from '@/features/utils/utils';
import { FilterMatchMode, FilterService } from 'primevue/api';

//store
const entriesStore = useEntriesStore();
const accountsStore = useAccountsStore();
await accountsStore.loadAccounts();

let selectedCurrency = ref('EUR');
let selectedDate = ref(new Date());
const currenciesArray = ref(currencies);
let selectedConcept = ref();
let selectedDebitAmount = ref();
let selectedCreditAmount = ref();
let selectedDebitAccount = ref();
let filteredDebitAccounts = ref();
let selectedCreditAccount = ref();
let filteredCreditAccounts = ref();
const accounts = accountsStore.state.accounts;

function addDebitEntry() {
  const newEntry: CreateEntryInput = {
    id: uuid(),
    accountId: selectedDebitAccount.value,
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: selectedCurrency.value,
    date: dateToUnixTimestamp(selectedDate.value),
    concept: selectedConcept.value,
    debitAmount: selectedDebitAmount.value,
    creditAmount: null,
  };

  entriesStore.addDebitEntry(newEntry);
}

function addCreditEntry() {
  const newEntry: CreateEntryInput = {
    id: uuid(),
    accountId: selectedCreditAccount.value,
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: selectedCurrency.value,
    date: dateToUnixTimestamp(selectedDate.value),
    concept: selectedConcept.value,
    debitAmount: null,
    creditAmount: selectedCreditAmount.value,
  };

  entriesStore.addCreditEntry(newEntry);
}

function removeEntry(id: string) {
  entriesStore.removeModificationEntry(id);
}

const searchDebit = (event: any) => {
  let query = event.query;
  let newFilteredCities = [];

  for (let country of accounts) {
    let filteredItems = FilterService.filter(country.children, ['label'], query, FilterMatchMode.CONTAINS);
    if (filteredItems && filteredItems.length) {
      newFilteredCities.push({ ...country, ...{ children: filteredItems } });
    }
  }

  filteredDebitAccounts.value = newFilteredCities;
};

const searchCredit = (event: any) => {
  let query = event.query;
  let newFilteredCities = [];

  for (let country of accounts) {
    let filteredItems = FilterService.filter(country.children, ['label'], query, FilterMatchMode.CONTAINS);
    if (filteredItems && filteredItems.length) {
      newFilteredCities.push({ ...country, ...{ children: filteredItems } });
    }
  }

  filteredCreditAccounts.value = newFilteredCities;
};
</script>

<template>
  <div class="mt-0 flex w-full space-x-10 overflow-hidden rounded-md bg-gray-50 p-6 text-left shadow">
    <div class="w-32 flex-none">
      <label for="currency" class="block text-sm font-medium text-gray-800">Divisa</label>
      <select v-model="selectedCurrency" id="currency" name="currency" autocomplete="currency-name" class="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm sm:text-sm">
        <option v-for="(currency, i) in currenciesArray" :key="i">{{ currency.code }}</option>
      </select>
    </div>
    <div class="w-44 flex-auto">
      <label for="concept" class="block text-sm font-medium text-gray-800">Concepto</label>
      <input type="text" v-model="selectedConcept" maxlength="64" name="concept" id="concept" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm sm:text-sm" />
    </div>
    <div class="flex-none">
      <label for="calendar" class="block text-sm font-medium text-gray-800">Fecha del asiento</label>
      <CalendarPrimeVue v-model="selectedDate" :showTime="true" :showSeconds="true" :showIcon="true" dateFormat="dd/mm/yy" name="calendar" id="calendar" />
    </div>
    <div class="flex-1">
      <label for="balance" class="block text-sm font-medium text-gray-800">Balance del asiento</label>
      <p v-if="entriesStore.state.entriesCreation.length == 0" class="inline-block rounded-lg bg-red-200 p-2 text-xl text-red-800">Asiento vacío ❌</p>
      <p v-else-if="entriesStore.state.totalBalance == 0" class="inline-block rounded-lg bg-green-200 p-2 text-xl text-green-800">Asiento correcto ✅</p>
      <p v-else class="inline-block rounded-lg bg-red-200 p-2 text-xl text-red-800">Asiento descuadrado ({{ formatAmount(entriesStore.state.totalBalance, selectedCurrency) }}) ❌</p>
    </div>
  </div>

  <div class="mt-2 mb-10 flex space-x-4 text-left">
    <div class="w-1/2">
      <div class="mt-0 overflow-hidden rounded-md shadow">
        <div class="bg-gray-800 py-3 px-6 text-white">
          <h4 class="text-2xl">Debe</h4>
        </div>
        <div class="space-y-4 bg-white p-6">
          <div>
            <label for="debit-account" class="block text-sm font-medium text-gray-800">Cuenta</label>
            <AutoCompletePrimeVue v-model="selectedDebitAccount" dropdown scrollHeight="500px" :suggestions="filteredDebitAccounts" @complete="searchDebit" forceSelection optionLabel="label" optionGroupLabel="label" optionGroupChildren="children" placeholder="Cuenta del debe" class="">
              <template #optiongroup="slotProps">
                <div class="align-items-center country-item flex">
                  <div>{{ slotProps.item.label }}</div>
                </div>
              </template>
            </AutoCompletePrimeVue>
          </div>
          <div>
            <label for="debit-amount" class="block text-sm font-medium text-gray-800">Importe</label>
            <div class="p-inputgroup flex-1">
              <span class="p-inputgroup-addon">{{ findCurrencyByCode(selectedCurrency) }}</span>
              <InputNumberPrimeVue v-model="selectedDebitAmount" inputId="locale-user" :minFractionDigits="2" :maxFractionDigits="2" placeholder="Importe del debe" />
              <select v-model="selectedCurrency" id="currency" name="currency" class="p-inputgroup-addon">
                <option v-for="(currency, i) in currenciesArray" :key="i">{{ currency.code }}</option>
              </select>
              <button type="submit" id="debit-amount" name="debit-amount" @click="addDebitEntry" class="bg-blue-500 py-2 px-4 text-white shadow-sm hover:bg-gray-800">Añadir al debe</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="w-1/2">
      <div class="mt-0 overflow-hidden rounded-md shadow">
        <div class="bg-gray-800 py-3 px-6 text-white">
          <h4 class="text-2xl">Haber</h4>
        </div>
        <div class="space-y-4 bg-white p-6">
          <div>
            <label for="debit-account" class="block text-sm font-medium text-gray-800">Cuenta</label>
            <AutoCompletePrimeVue v-model="selectedCreditAccount" dropdown :suggestions="filteredCreditAccounts" @complete="searchCredit" forceSelection optionLabel="label" optionGroupLabel="label" optionGroupChildren="children" placeholder="Cuenta del haber" class="">
              <template #optiongroup="slotProps">
                <div class="align-items-center country-item flex">
                  <div>{{ slotProps.item.label }}</div>
                </div>
              </template>
            </AutoCompletePrimeVue>
          </div>
          <div>
            <label for="credit-amount" class="block text-sm font-medium text-gray-800">Importe</label>
            <div class="p-inputgroup flex-1">
              <span class="p-inputgroup-addon">{{ findCurrencyByCode(selectedCurrency) }}</span>
              <InputNumberPrimeVue v-model="selectedCreditAmount" inputId="locale-user" :minFractionDigits="2" :maxFractionDigits="2" placeholder="Importe del haber" />
              <select v-model="selectedCurrency" id="currency" name="currency" class="p-inputgroup-addon">
                <option v-for="(currency, i) in currenciesArray" :key="i">{{ currency.code }}</option>
              </select>
              <button type="submit" id="credit-amount" name="credit-amount" @click="addCreditEntry" class="bg-blue-500 py-2 px-4 text-white shadow-sm hover:bg-gray-800">Añadir al haber</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="-mx-8 mt-8 flex overflow-x-auto">
    <div class="inline-block min-w-full py-2 px-8 align-middle">
      <div class="overflow-hidden rounded-lg shadow ring-1 ring-black ring-opacity-5">
        <table class="min-w-full divide-y divide-gray-300">
          <thead class="bg-gray-800">
            <tr>
              <th scope="col" class="whitespace-nowrap py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-white sm:pl-6">Cuenta</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Divisa</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Fecha</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Concepto</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Debe</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Haber</th>
              <th scope="col" class="whitespace-nowrap px-2 py-3.5 text-left text-sm font-semibold text-white">Acciones</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200 bg-white">
            <tr v-for="(entry, i) in entriesStore.state.storedEntries.reverse()" :key="entry.entry.id" :class="i % 2 === 0 ? undefined : 'bg-gray-50'">
              <td class="whitespace-nowrap py-2 pl-4 pr-3 text-sm text-gray-500 sm:pl-6">{{ entry.account.displayName }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm font-medium text-gray-900">{{ entry.entry.currency }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-900">{{ formatDate(entry.entry.date) }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-500">{{ entry.entry.concept }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-500">{{ formatAmount(entry.entry.debitAmount, entry.entry.currency) }}</td>
              <td class="whitespace-nowrap px-2 py-2 text-sm text-gray-500">{{ formatAmount(entry.entry.creditAmount, entry.entry.currency) }}</td>
              <td class="space-x-3 whitespace-nowrap px-2 py-2 text-sm text-gray-500">
                <button @click="removeEntry(entry.entry.id || '')" class="text-red-500 hover:text-red-900">Eliminar</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style>
.p-autocomplete {
  @apply w-full rounded-md border-gray-300 text-sm;
}

.p-autocomplete-panel {
  @apply bg-gray-100 text-sm text-gray-800;
}

.p-calendar-w-btn .p-datepicker-trigger {
  @apply border-transparent bg-gray-800 text-sm text-white hover:bg-gray-800;
}
</style>
