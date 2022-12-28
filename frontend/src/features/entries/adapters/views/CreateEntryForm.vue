<script setup lang="ts">
import AccountsDropDown from '@/features/accounts/adapters/views/AccountsDropDown.vue';
import { ref } from 'vue';
import { v4 as uuid } from 'uuid';
import type { Entry } from '@/features/entries/application/entries-model';
import { useEntriesStore } from '@/stores/entries';

const entriesStore = useEntriesStore();

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
  const newEntry: Entry = {
    id: id,
    accountId: '', //TODO: retrieve
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: currency.value,
    date: date.value.toISOString(),
    concept: concept.value,
    debitAmount: debitAmount.value,
    creditAmount: null,
  };

  entriesStore.addEntry(newEntry);
  entriesStore.updateTotalBalance(debitAmount.value);
}

function addCreditEntry() {
  const newEntry: Entry = {
    id: id,
    accountId: '', //TODO: retrieve
    multipleEntryId: entriesStore.state.multipleEntryId,
    currency: currency.value,
    date: date.value.toISOString(),
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
      <Suspense> <AccountsDropDown></AccountsDropDown></Suspense>
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
    <!-- TODO implement -->
  </div>
</template>

<style scoped>
.form-body {
  display: flex;
  margin-top: 20px;
}
</style>
