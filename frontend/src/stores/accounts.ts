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
        this.state.accounts = [];
        this.state.stateMessage = 'loading accounts';

        const accounts = await accountsUseCase.listAccounts();

        this.state.accounts = accounts;
        this.state.stateMessage = 'accounts loaded';
      } catch (err) {
        this.state.accounts = [];
        this.state.stateMessage = 'error loading accounts';
      }
    },
  },
});
