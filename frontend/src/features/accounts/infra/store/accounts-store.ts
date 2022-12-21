import { defineStore } from 'pinia';

export const useAccountsStore = defineStore('accountsStore', {
  state: () => {
    return { accounts: [] };
  },

  actions: {
    loadAccounts() {
      this.accounts = listAccountsHttpResponse;
    },
    clear() {
      this.accounts = [];
    },
  },
});
