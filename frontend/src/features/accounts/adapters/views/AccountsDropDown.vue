<script setup lang="ts">
import { ref } from 'vue';
import { useAccountsStore } from '@/stores/accounts';

const accountsStore = useAccountsStore();

await accountsStore.loadAccounts();

const selectedAccount = ref();

const accounts = accountsStore.state.accounts;
</script>

<template>
  <DropdownPrimeVue
    v-model="selectedAccount"
    :options="accounts"
    optionLabel="name"
    optionValue="id"
    :virtualScrollerOptions="{ itemSize: 38 }"
    placeholder="Select an Account"
  />
</template>
