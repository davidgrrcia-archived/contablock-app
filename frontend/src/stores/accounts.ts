import { defineStore } from 'pinia';
import { accountsUseCase } from '@/features/accounts';
import type { Account } from '@/features/accounts/application/accounts-model';

//TODO: apply clean architecture
interface AccountsState {
  accounts: Account[];
  stateMessage: string | null;
}

const state: AccountsState = {
  accounts: [],
  stateMessage: null,
};

export const useAccountsStore = defineStore('accounts', {
  state: () => {
    return { state };
  },

  actions: {
    async loadAccounts() {
      try {
        this.state = {
          accounts: [],
          stateMessage: 'loading accounts',
        };

        const accounts = await accountsUseCase.listAccounts();

        this.state = {
          accounts: accounts,
          stateMessage: 'accounts loaded',
        };
      } catch (err) {
        this.state = {
          accounts: [],
          stateMessage: 'error loading accounts',
        };
      }
    },
  },
});
