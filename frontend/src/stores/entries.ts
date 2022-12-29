import { defineStore } from 'pinia';
import type { Entry } from '@/features/entries/application/entries-model';
import { accountsUseCase } from '@/features/accounts';

//TODO: apply clean architecture
interface EntriesState {
  entriesCreation: Entry[];
  multipleEntryId: string;
  totalBalance: number;
  stateMessage: string | null;
}

const state: EntriesState = {
  entriesCreation: [],
  multipleEntryId: '',
  totalBalance: 0,
  stateMessage: null,
};

export const useEntriesStore = defineStore('entries', {
  state: () => {
    return { state };
  },

  actions: {
    async addEntry(entry: Entry) {
      const account = await accountsUseCase.describeAccount(entry.accountId);
      const displayName = entry.accountId + ' - ' + account.name;
      entry.displayName = displayName

      this.state.entriesCreation.push(entry);
    },
    updateTotalBalance(amount: number) {
      this.state.totalBalance = this.state.totalBalance + amount;
    },
    async getEntries() {},
  },
});
