<script setup lang="ts">
import AccountsDropDown from '@/features/accounts/adapters/views/AccountsDropDown.vue';
import { ref } from 'vue';
import { v4 as uuid } from 'uuid';
import type { Entry } from '@/features/entries/application/entries-model';
import { useEntriesStore } from '@/stores/entries';
import {useErrorMessagesStore} from '@/stores/errorMessages';

const entriesStore = useEntriesStore();
const errorStore = useErrorMessagesStore()

const currency = ref('EUR');
const currentDate = new Date();
const date = ref(currentDate);
const concept = ref();
const debitAmount = ref();
const creditAmount = ref();
const id = uuid();
const multipleEntryId = uuid();

if (entriesStore.state.multipleEntryId === '') {
  entriesStore.state.multipleEntryId = multipleEntryId;
}

function addDebitEntry() {
  if (debitAmount.value === undefined){
    errorStore.addErrorMessage("Introduce una cantidad en el debe")
    return
  }

  const newEntry: Entry = {
    id: id,
    accountId: '', //TODO: retrieve
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: currency.value,
    date: date.value.toLocaleDateString(),
    concept: concept.value,
    debitAmount: debitAmount.value,
    creditAmount: null,
  };

  entriesStore.addEntry(newEntry);
  entriesStore.updateTotalBalance(debitAmount.value);
}

function addCreditEntry() {
  if (creditAmount.value === undefined){
    errorStore.addErrorMessage("Introduce una cantidad en el haber")
    return
  }

  const newEntry: Entry = {
    id: id,
    accountId: '', //TODO: retrieve
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: currency.value,
    date: date.value.toLocaleDateString(),
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
</script>

<template>
  <div class="form-header">
    <div>
      <DropdownPrimeVue
        v-model="currency"
        :options="currencies"
        optionLabel="code"
        optionValue="code"
        :placeholder="currencies[0].code"
      />
      <CalendarPrimeVue v-model="date" :showIcon="true" dateFormat="dd/mm/yy" />
    </div>
    <div>
      <TextareaPrimeVue
        v-model="concept"
        :autoResize="true"
        rows="3"
        cols="50"
        placeholder="Descripción del asiento (opcional)"
      />
    </div>
  </div>
  <div class="form-body">
    <div>
      <h4>Debe</h4>
      <Suspense> <AccountsDropDown></AccountsDropDown></Suspense> <!-- TODO improve display UI -->
      <InputNumberPrimeVue v-model="debitAmount" mode="currency" :currency="currency" locale="es-ES" />
      <ButtonPrimeVue icon="fa-solid fa-plus" @click="addDebitEntry" />
    </div>
    <div>
      <DividerPrimeVue layout="vertical"> </DividerPrimeVue>
    </div>
    <div>
      <h4>Haber</h4>
      <Suspense> <AccountsDropDown></AccountsDropDown></Suspense>
      <InputNumberPrimeVue v-model="creditAmount" mode="currency" :currency="currency" locale="es-ES" />
      <ButtonPrimeVue icon="fa-solid fa-plus" @click="addCreditEntry" />
    </div>
  </div>
  <div class="form-table">
    <DataTablePrimeVue :value="entriesStore.state.entriesCreation" responsiveLayout="scroll">
      <ColumnPrimeVue field="accountId" header="Cuenta"></ColumnPrimeVue>
      <ColumnPrimeVue field="currency" header="Divisa"></ColumnPrimeVue>
      <ColumnPrimeVue field="date" header="Fecha"></ColumnPrimeVue>
      <ColumnPrimeVue field="concept" header="Concepto"></ColumnPrimeVue>
      <ColumnPrimeVue field="debitAmount" header="Debe"></ColumnPrimeVue>
      <ColumnPrimeVue field="creditAmount" header="Haber"></ColumnPrimeVue>
    </DataTablePrimeVue>
    <h5>Balance:</h5>
    <p> {{entriesStore.state.totalBalance}}</p>
  </div>
</template>

<style scoped>
.form-body {
  display: flex;
  margin-top: 20px;
}
</style>
