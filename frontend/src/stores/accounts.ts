import { defineStore } from 'pinia';
import { accountsUseCase } from "@/features/accounts";
import type { Account } from "@/features/accounts/application/accounts-model";

//TODO: apply clean architecture

interface AccountsState {
  accounts: Account[];
  stateMessage: string | null;
}

const ACCOUNTS_STATE: AccountsState = {
  accounts: [],
  stateMessage: null,
};

export const useAccountsStore = defineStore('accounts', {
  state: () => {
    return { ACCOUNTS_STATE };
  },

  actions: {
    async loadAccounts() {
      try {
        this.ACCOUNTS_STATE = {
          accounts: [],
          stateMessage: 'loading accounts',
        };

        const accounts = await accountsUseCase.listAccounts();

        this.ACCOUNTS_STATE = {
          accounts: accounts,
          stateMessage: 'accounts loaded',
        };
      } catch (err) {
        this.ACCOUNTS_STATE = {
          accounts: [],
          stateMessage: 'error loading accounts',
        };
      }
    },
  }
});


