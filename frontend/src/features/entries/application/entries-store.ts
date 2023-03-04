import { defineStore } from 'pinia';
import type { CreateEntryInput, ModifyEntryInput } from '@/features/entries/application/entries-model';
import { accountsUseCase } from '@/features/accounts';
import { entriesUseCase } from '@/features/entries';
import { useMessagesStore } from '@/stores/messages';
import type { Account } from '@/features/accounts/application/accounts-model';

interface EntryWithAccount {
  entry: CreateEntryInput;
  account: Account;
}

//TODO: apply clean architecture
interface EntriesState {
  creationModalOpen: boolean;
  modificationModalOpen: boolean;
  entriesCreation: EntryWithAccount[];
  entriesModification: EntryWithAccount[];
  multipleEntryId: string;
  totalBalance: number;
  stateMessage: string | null;
  differentEntry: boolean;
}

const state: EntriesState = {
  creationModalOpen: false,
  modificationModalOpen: false,
  entriesCreation: [],
  entriesModification: [],
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
    async addDebitEntryForModification(entry: CreateEntryInput) {
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

      const response = await entriesUseCase.createEntry(entry);

      if (response instanceof Error) {
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error creando asiento', 'Ha ocurrido un error creando el asiento');
        return;
      }

      const entryWithAccount: EntryWithAccount = {
        entry: entry,
        account: account,
      };

      this.state.entriesModification.push(entryWithAccount);
      this.state.totalBalance = this.state.totalBalance + entry.debitAmount;
    },
    async addCreditEntryForModification(entry: CreateEntryInput) {
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

      const response = await entriesUseCase.createEntry(entry);

      if (response instanceof Error) {
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error creando asiento', 'Ha ocurrido un error creando el asiento');
        return;
      }

      const entryWithAccount: EntryWithAccount = {
        entry: entry,
        account: account,
      };

      this.state.entriesModification.push(entryWithAccount);
      this.state.totalBalance = this.state.totalBalance - entry.creditAmount;
    },
    async modifyDebitEntry(entry: ModifyEntryInput) {
      const messagesStore = useMessagesStore();

      if (entry.accountId == undefined) {
        messagesStore.addWarnMessage('Selecciona una cuenta en el debe', 'Para poder modificar un asiento debes seleccionar una cuenta en la partida del debe');
        return;
      }

      if (entry.debitAmount == undefined) {
        messagesStore.addWarnMessage('Introduce una cantidad en el debe', 'Para poder modificar un asiento debes introducir una cantidad en la partida del debe');
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

      this.state.entriesModification.push(entryWithAccount);
      this.state.totalBalance = this.state.totalBalance + entry.debitAmount;
    },
    async modifyCreditEntry(entry: ModifyEntryInput) {
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
        this.clearEntriesAndHideModals();
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
    modifyEntries(entries: any) {
      this.state.modificationModalOpen = true;
      this.state.entriesModification = entries;
      this.state.multipleEntryId = entries[0].entry.multipleEntryId;
    },
    showCreationModal() {
      this.state.creationModalOpen = true;
    },
    showModificationModal() {
      this.state.modificationModalOpen = true;
    },
    clearEntriesAndHideModals() {
      this.state.entriesCreation = [];
      this.state.multipleEntryId = '';
      this.state.totalBalance = 0;
      this.state.creationModalOpen = false;
      this.state.modificationModalOpen = false;
    },
  },
});
