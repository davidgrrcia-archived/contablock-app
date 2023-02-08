<script setup lang="ts">
import { ref } from 'vue';
import { v4 as uuid } from 'uuid';
import type { Entry } from '@/features/entries/application/entries-model';
import { useEntriesStore } from '@/stores/entries';
import { useMessagesStore } from '@/stores/messages';
import { useAccountsStore } from '@/stores/accounts';

//store
const entriesStore = useEntriesStore();
const errorStore = useMessagesStore();
const accountsStore = useAccountsStore();
await accountsStore.loadAccounts();

//constants
const accounts = accountsStore.state.accounts.map(account => ({
  ...account,
  displayName: account.id + ' - ' + account.name,
}));

const currency = ref('EUR');
const currentDate = new Date();
const date = ref(currentDate);
const concept = ref();
const debitAmount = ref();
const creditAmount = ref();
const id = uuid();
const multipleEntryId = uuid();
const debitAccount = ref();
const creditAccount = ref();

if (entriesStore.state.multipleEntryId === '') {
  entriesStore.state.multipleEntryId = multipleEntryId;
}

function addDebitEntry() {
  if (debitAmount.value === undefined) {
    errorStore.addErrorMessage('Introduce una cantidad en el debe');
    return;
  }

  const newEntry: Entry = {
    id: id,
    accountId: debitAccount.value,
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: currency.value,
    date: date.value.toISOString(),
    displayDate: date.value.toLocaleDateString(),
    concept: concept.value,
    debitAmount: debitAmount.value,
    creditAmount: null,
  };

  entriesStore.addEntry(newEntry);
  entriesStore.updateTotalBalance(debitAmount.value);
}

function addCreditEntry() {
  if (creditAmount.value === undefined) {
    errorStore.addErrorMessage('Introduce una cantidad en el haber');
    return;
  }

  const newEntry: Entry = {
    id: id,
    accountId: creditAccount.value,
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: currency.value,
    date: date.value.toISOString(),
    displayDate: date.value.toLocaleDateString(),
    concept: concept.value,
    debitAmount: null,
    creditAmount: creditAmount.value,
  };

  entriesStore.addEntry(newEntry);
  entriesStore.updateTotalBalance(-creditAmount.value);
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

const selectedGroupedCity = ref();

const groupedCities = ref([
  {
    label: 'Germany', code: 'DE',
    items: [
      {label: 'Berlin', value: 'Berlin'},
      {label: 'Frankfurt', value: 'Frankfurt'},
      {label: 'Hamburg', value: 'Hamburg'},
      {label: 'Munich', value: 'Munich'}
    ]
  },
  {
    label: 'USA', code: 'US',
    items: [
      {label: 'Chicago', value: 'Chicago'},
      {label: 'Los Angeles', value: 'Los Angeles'},
      {label: 'New York', value: 'New York'},
      {label: 'San Francisco', value: 'San Francisco'}
    ]
  },
  {
    label: 'Japan', code: 'JP',
    items: [
      {label: 'Kyoto', value: 'Kyoto'},
      {label: 'Osaka', value: 'Osaka'},
      {label: 'Tokyo', value: 'Tokyo'},
      {label: 'Yokohama', value: 'Yokohama'}
    ]
  }
]);
</script>

<template>

  <h5>Grouped</h5>
  <DropdownPrimeVue v-model="selectedGroupedCity" :options="groupedCities" optionLabel="label" optionGroupLabel="label" optionGroupChildren="items">
    <template #optiongroup="slotProps">
      <div class="flex align-items-center country-item">
        <img src="https://www.primefaces.org/wp-content/uploads/2020/05/placeholder.png" width="18" />
        <div>{{slotProps.option.label}}</div>
      </div>
    </template>
  </DropdownPrimeVue>

  <div class="form-header">
    <DropdownPrimeVue
      v-model="currency"
      :options="currencies"
      optionLabel="code"
      optionValue="code"
      :placeholder="currencies[0].code"
    />
    <CalendarPrimeVue v-model="date" :showIcon="true" dateFormat="dd/mm/yy" />
    <TextareaPrimeVue
      v-model="concept"
      :autoResize="true"
      rows="3"
      cols="50"
      placeholder="Descripción del asiento (opcional)"
    />
  </div>
  <div class="form-body">
    <div>
      <h4>Debe</h4>
      <DropdownPrimeVue
        v-model="debitAccount"
        :options="accounts"
        optionLabel="displayName"
        optionValue="id"
        :virtualScrollerOptions="{ itemSize: 38 }"
        placeholder="Selecciona una cuenta"
      />
      <InputNumberPrimeVue v-model="debitAmount" mode="currency" :currency="currency" locale="es-ES" />
      <ButtonPrimeVue icon="fa-solid fa-plus" @click="addDebitEntry" />
    </div>
    <div>
      <DividerPrimeVue layout="vertical"> </DividerPrimeVue>
    </div>
    <div>
      <h4>Haber</h4>
      <DropdownPrimeVue
        v-model="creditAccount"
        :options="accounts"
        optionLabel="displayName"
        optionValue="id"
        :virtualScrollerOptions="{ itemSize: 38 }"
        placeholder="Selecciona una cuenta"
      />
      <InputNumberPrimeVue v-model="creditAmount" mode="currency" :currency="currency" locale="es-ES" />
      <ButtonPrimeVue icon="fa-solid fa-plus" @click="addCreditEntry" />
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
    <h5>Balance:</h5>
    <p>{{ entriesStore.state.totalBalance }}</p>
  </div>
</template>

<style scoped>
.form-body {
  display: flex;
  margin-top: 20px;
}
</style>
