import { defineStore } from 'pinia';
import type { Entry } from '@/features/entries/application/entries-model';

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
    addEntry(entry: Entry) {
      this.state.entriesCreation.push(entry);
    },
    updateTotalBalance(amount: number) {
      this.state.totalBalance = this.state.totalBalance + amount;
    },
  },
});
