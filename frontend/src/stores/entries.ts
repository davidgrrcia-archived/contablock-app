import { defineStore } from 'pinia';
import type { Entry } from '@/features/entries/application/entries-model';
import { accountsUseCase } from '@/features/accounts';
import { entriesUseCase } from '@/features/entries';
import { useMessagesStore } from '@/stores/messages';

//TODO: apply clean architecture
interface EntriesState {
  entriesCreation: Entry[];
  multipleEntryId: string;
  totalBalance: number;
  storedEntries: Entry[];
  stateMessage: string | null;
  differentEntry: boolean;
}

const state: EntriesState = {
  entriesCreation: [],
  multipleEntryId: '',
  totalBalance: 0,
  storedEntries: [],
  stateMessage: null,
  differentEntry: false,
};

export const useEntriesStore = defineStore('entries', {
  state: () => {
    return { state };
  },

  actions: {
    async addEntry(entry: Entry) {
      const account = await accountsUseCase.describeAccount(entry.accountId);
      const displayName = entry.accountId + ' - ' + account.name;
      entry.displayName = displayName;

      this.state.entriesCreation.push(entry);
    },
    updateTotalBalance(amount: number) {
      this.state.totalBalance = this.state.totalBalance + amount;
    },
    async createEntries() {
      try {
        this.state.stateMessage = 'creating entries';

        this.state.entriesCreation.map(async entry => {
          return await entriesUseCase.createEntry(entry);
        });

        this.state.stateMessage = 'entries created';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage(this.state.stateMessage);

        this.clearEntries();
      } catch (err) {
        this.state.stateMessage = 'error creating entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage(this.state.stateMessage);
      }
    },
    async loadEntries() {
      try {
        this.state.storedEntries = [];
        this.state.stateMessage = 'loading entries';

        const entries = await entriesUseCase.listEntries();

        this.state.storedEntries = entries;
        this.state.stateMessage = 'entries loaded';
      } catch (err) {
        this.state.stateMessage = 'error loading entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage(this.state.stateMessage);
      }
    },
    clearEntries() {
      this.state.entriesCreation = [];
      this.state.multipleEntryId = '';
    },
  },
});
