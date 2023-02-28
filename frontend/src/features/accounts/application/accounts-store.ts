import { defineStore } from 'pinia';
import { accountsUseCase } from '@/features/accounts';
import type { Node } from '@/features/accounts/application/accounts-model';
import { useMessagesStore } from '@/stores/messages';

//TODO: apply clean architecture
interface AccountsState {
  accounts: Node[];
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

        if (accounts instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addMessage('Fallo al cargar las cuentas');
          return;
        }

        this.state.accounts = accounts;
        this.state.stateMessage = 'accounts loaded';
      } catch (err) {
        this.state.accounts = [];
        this.state.stateMessage = 'error loading accounts';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },
});
