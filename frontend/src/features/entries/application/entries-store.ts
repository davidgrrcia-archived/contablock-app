import { defineStore } from 'pinia';
import type { CreateEntryInput } from '@/features/entries/application/entries-model';
import { accountsUseCase } from '@/features/accounts';
import { entriesUseCase } from '@/features/entries';
import { useMessagesStore } from '@/stores/messages';
import type { Account } from '@/features/accounts/application/accounts-model';
import { findCurrencyByCode } from '@/features/utils/utils';

interface EntryWithAccount {
  entry: CreateEntryInput;
  account: Account;
}

//TODO: apply clean architecture
interface EntriesState {
  entriesCreation: EntryWithAccount[];
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
    async addDebitEntry(entry: CreateEntryInput) {
      const messagesStore = useMessagesStore();

      if (entry.accountId == undefined) {
        messagesStore.addWarnMessage('Selecciona una cuenta en el debe', 'Para poder crear un asiento debes seleccionar una cuenta en la partida del debe');
        return;
      }

      if (entry.debitAmount == undefined) {
        messagesStore.addWarnMessage('Introduce una cantidad en el debe', 'Para poder crear un asiento debes introducir una cantidad en la partida del debe');
        return;
      }

      //match account id
      const match = JSON.stringify(entry.accountId).match(/"(\d+)"/);
      entry.accountId = match ? match[1] : '';

      const account = await accountsUseCase.describeAccount(entry.accountId);

      if (account instanceof Error) {
        messagesStore.addWarnMessage('Fallo al cargar la cuenta', 'No se ha podido cargar la cuenta seleccionada');
        return;
      }

      const entryWithAccount: EntryWithAccount = {
        entry: entry,
        account: account,
      };

      this.state.entriesCreation.push(entryWithAccount);
      this.state.totalBalance = this.state.totalBalance + entry.debitAmount;
    },
    async addCreditEntry(entry: CreateEntryInput) {
      const messagesStore = useMessagesStore();

      if (entry.accountId == undefined) {
        messagesStore.addWarnMessage('Selecciona una cuenta en el haber', 'Para poder crear un asiento debes seleccionar una cuenta en la partida del haber');
        return;
      }

      if (entry.creditAmount == undefined) {
        messagesStore.addWarnMessage('Introduce una cantidad en el haber', 'Para poder crear un asiento debes introducir una cantidad en la partida del haber');
        return;
      }

      //match account id
      const match = JSON.stringify(entry.accountId).match(/"(\d+)"/);
      entry.accountId = match ? match[1] : '';

      const account = await accountsUseCase.describeAccount(entry.accountId);

      if (account instanceof Error) {
        messagesStore.addWarnMessage('Fallo al cargar la cuenta', 'No se ha podido cargar la cuenta seleccionada');
        return;
      }

      const entryWithAccount: EntryWithAccount = {
        entry: entry,
        account: account,
      };

      this.state.entriesCreation.push(entryWithAccount);
      this.state.totalBalance = this.state.totalBalance - entry.creditAmount;
    },
    async createEntries() {
      const messagesStore = useMessagesStore();

      try {
        if (this.state.totalBalance > 0) {
          messagesStore.addWarnMessage('El principio de partida doble es obligatorio', 'Para poder crear un asiento el debe y el haber deben estar balanceados, actualmente el debe es mayor que el haber en ' + this.state.totalBalance + findCurrencyByCode(this.state.entriesCreation[0].entry.currency));
        }
        if (this.state.totalBalance < 0) {
          messagesStore.addWarnMessage('El principio de partida doble es obligatorio', 'Para poder crear un asiento el debe y el haber deben estar balanceados, actualmente el debe es menor que el haber en ' + this.state.totalBalance + findCurrencyByCode(this.state.entriesCreation[0].entry.currency));
        }
        if (this.state.entriesCreation.length === 0) {
          messagesStore.addWarnMessage('No hay asientos creados', 'Para poder crear un asiento debes crear al menos una partida en el debe y otra en el haber');
        }

        this.state.stateMessage = 'creating entries';

        this.state.entriesCreation.map(async entry => {
          const response = await entriesUseCase.createEntry(entry.entry);

          if (response instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addErrorMessage('Error creando asiento', 'Ha ocurrido un error creando el asiento');
            return;
          }

          return response;
        });

        this.state.stateMessage = 'entries created';
        this.clearEntries();
      } catch (err) {
        this.state.stateMessage = 'error creating entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error creando asientos', 'Ha ocurrido un error creando los asientos');
      }
    },
    async publishEntries(entriesExpanded: any) {
      try {
        this.state.stateMessage = 'publishing entries';

        entriesExpanded.map(async (entry: any) => {
          const response = await entriesUseCase.publishEntry(entry.entry.id);

          if (response instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addErrorMessage('Error publicando asiento', 'Ha ocurrido un error publicando el asiento');
            return;
          }

          return response;
        });
        this.state.stateMessage = 'entries published';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage('Asiento publicado', 'El asiento ha sido publicado correctamente');
      } catch (err) {
        this.state.stateMessage = 'error publishing entries';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
    async removeEntries(entriesExpanded: any) {
      try {
        this.state.stateMessage = 'removing entries';

        entriesExpanded.map(async (entry: any) => {
          const response = await entriesUseCase.removeEntry(entry.entry.id);

          if (response instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addErrorMessage('Error eliminando asientos', 'Ha ocurrido un error eliminando los asientos');
            return;
          }

          return response;
        });
        this.state.stateMessage = 'entries removed';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage('Asiento eliminado', 'El asiento ha sido eliminado correctamente');
      } catch (err) {
        this.state.stateMessage = 'error removing entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error eliminando asientos', 'Ha ocurrido un error eliminando los asientos');
      }
    },
    clearEntries() {
      this.state.entriesCreation = [];
      this.state.multipleEntryId = '';
      this.state.totalBalance = 0;
    },
  },
});
