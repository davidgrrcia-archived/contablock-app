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
  stateMessage: string | null;
  differentEntry: boolean;
}

const state: EntriesState = {
  entriesCreation: [],
  multipleEntryId: '',
  totalBalance: 0,
  stateMessage: null,
  differentEntry: false,
};

export const useEntriesStore = defineStore('entries', {
  state: () => {
    return { state };
  },

  actions: {
    async addDebitEntry(entry: Entry) {
      const messagesStore = useMessagesStore();

      if (entry.debitAmount == undefined) {
        messagesStore.addMessage('Introduce una cantidad en el debe');
        return;
      }

      const account = await accountsUseCase.describeAccount(entry.accountId);
      const displayName = entry.accountId + ' - ' + account.name;
      entry.displayName = displayName;

      this.state.entriesCreation.push(entry);
      this.state.totalBalance = this.state.totalBalance + entry.debitAmount;
    },
    async addCreditEntry(entry: Entry) {
      const messagesStore = useMessagesStore();

      if (entry.creditAmount == undefined ) {
        messagesStore.addMessage('Introduce una cantidad en el haber');
        return;
      }

      const account = await accountsUseCase.describeAccount(entry.accountId);
      const displayName = entry.accountId + ' - ' + account.name;
      entry.displayName = displayName;

      this.state.entriesCreation.push(entry);
      this.state.totalBalance = this.state.totalBalance - entry.creditAmount;
    },
    async createEntries() {
      const messagesStore = useMessagesStore();

      try {
        if (this.state.totalBalance !== 0) {
          throw 'El principio de partida doble es obligatorio';
        }
        if (this.state.entriesCreation.length === 0) {
          throw 'No hay asientos creados';
        }

        this.state.stateMessage = 'creating entries';

        this.state.entriesCreation.map(async entry => {
          return await entriesUseCase.createEntry(entry);
        });

        this.state.stateMessage = 'entries created';
        messagesStore.addMessage(this.state.stateMessage);

        this.clearEntries();
      } catch (err) {
        this.state.stateMessage = 'error creating entries';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(err.toString());
      }
    },
    clearEntries() {
      this.state.entriesCreation = [];
      this.state.multipleEntryId = '';
      this.state.totalBalance = 0;
    },
  },
});
