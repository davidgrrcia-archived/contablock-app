import { defineStore } from 'pinia';
import type { CreateEntryInput, ModifyEntryInput } from '@/features/entries/application/entries-model';
import { accountsUseCase } from '@/features/accounts';
import { entriesUseCase } from '@/features/entries';
import { useMessagesStore } from '@/stores/messages';
import type { Account } from '@/features/accounts/application/accounts-model';
import { v4 as uuid } from 'uuid';

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
  storedEntries: EntryWithAccount[];
  id: string;
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
  storedEntries: [],
  id: '',
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

      if (this.state.id == '') {
        entry.id = uuid();
      } else {
        entry.id = this.state.id;
        this.removeCreationEntry(this.state.id);
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
      this.state.storedEntries.push(entryWithAccount);
      this.state.totalBalance += entry.debitAmount;
      this.state.id = '';
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

      if (this.state.id == '') {
        entry.id = uuid();
      } else {
        entry.id = this.state.id;
        this.removeCreationEntry(this.state.id);
      }

      const entryWithAccount: EntryWithAccount = {
        entry: entry,
        account: account,
      };

      this.state.entriesCreation.push(entryWithAccount);
      this.state.storedEntries.push(entryWithAccount);
      this.state.totalBalance -= entry.creditAmount;
      this.state.id = '';
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
      try {
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
    async reviewEntries(entriesExpanded: any) {
      try {
        this.state.stateMessage = 'reviewing entries';

        entriesExpanded.map(async (entry: any) => {
          const response = await entriesUseCase.reviewEntry(entry.entry.id);

          if (response instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addErrorMessage('Error al solicitar revisión del asiento', 'Ha ocurrido un error al solicitar la revisión del asiento');
            return;
          }

          return response;
        });

        this.state.stateMessage = 'entries reviewed';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage('Asiento pendiente de revisión', 'El asiento ha sido enviado para su revisión');
      } catch (err) {
        this.state.stateMessage = 'error reviewing entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error revisando asientos', 'Ha ocurrido un error revisando los asientos');
      }
    },
    async approveEntries(entriesExpanded: any) {
      try {
        this.state.stateMessage = 'approving entries';

        entriesExpanded.map(async (entry: any) => {
          const response = await entriesUseCase.approveEntry(entry.entry.id);

          if (response instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addErrorMessage('Error aprobando asiento', 'Ha ocurrido un error aprobando el asiento');
            return;
          }

          return response;
        });

        this.state.stateMessage = 'entries approved';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage('Asiento aprobado', 'El asiento ha sido aprobado correctamente');
      } catch (err) {
        this.state.stateMessage = 'error approving entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error aprobando asientos', 'Ha ocurrido un error aprobando los asientos');
      }
    },
    async rejectEntries(entriesExpanded: any) {
      try {
        this.state.stateMessage = 'rejecting entries';

        entriesExpanded.map(async (entry: any) => {
          const response = await entriesUseCase.rejectEntry(entry.entry.id);

          if (response instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addErrorMessage('Error rechazando asiento', 'Ha ocurrido un error rechazando el asiento');
            return;
          }

          return response;
        });

        this.state.stateMessage = 'entries rejected';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage('Asiento rechazado', 'El asiento ha sido rechazado correctamente');
      } catch (err) {
        this.state.stateMessage = 'error rejecting entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error rechazando asientos', 'Ha ocurrido un error rechazando los asientos');
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
        messagesStore.addSuccessMessage('Asientos eliminados', 'Los asientos han sido eliminados correctamente');
      } catch (err) {
        this.state.stateMessage = 'error removing entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error eliminando asientos', 'Ha ocurrido un error eliminando los asientos');
      }
    },
    async removeModificationEntry(id: string) {
      try {
        this.state.stateMessage = 'removing entry';

        const response = await entriesUseCase.removeEntry(id);

        if (response instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addErrorMessage('Error eliminando el asiento', 'Ha ocurrido un error eliminando el asiento');
          return;
        }

        this.state.stateMessage = 'entries removed';
        const messagesStore = useMessagesStore();
        messagesStore.addSuccessMessage('Asiento eliminado', 'El asiento ha sido eliminado correctamente');
      } catch (err) {
        this.state.stateMessage = 'error removing entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error eliminando asientos', 'Ha ocurrido un error eliminando los asientos');
      }
    },
    removeCreationEntry(id: string) {
      this.state.entriesCreation = this.state.entriesCreation.filter(entry => entry.entry.id !== id);
    },
    requestModification(entries: any) {
      this.state.modificationModalOpen = true;
      this.state.storedEntries = entries;
      this.state.multipleEntryId = entries[0].entry.multipleEntryId;

      entries.map((entry: any) => {
        this.state.totalBalance = this.state.totalBalance + entry.entry.debitAmount - entry.entry.creditAmount;
      });
    },
    async modifyEntries() {
      try {
        this.createEntries();

        this.state.stateMessage = 'modifying entries';

        // this.state.entriesModification.map(async entry => {
        //   const response = await entriesUseCase.modifyEntry(entry.entry);

        //   if (response instanceof Error) {
        //     const messagesStore = useMessagesStore();
        //     messagesStore.addErrorMessage('Error modificando asiento', 'Ha ocurrido un error modificando el asiento');
        //     return;
        //   }

        //   return response;
        // });

        this.state.stateMessage = 'entries modified';
        this.clearEntriesAndHideModals();
      } catch (err) {
        this.state.stateMessage = 'error modifying entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage('Error modificando asientos', 'Ha ocurrido un error modificando los asientos');
      }
    },
    showCreationModal() {
      this.state.creationModalOpen = true;
    },
    showModificationModal() {
      this.state.modificationModalOpen = true;
    },
    clearEntriesAndHideModals() {
      this.state.entriesCreation = [];
      this.state.entriesModification = [];
      this.state.storedEntries = [];
      this.state.multipleEntryId = '';
      this.state.totalBalance = 0;
      this.state.creationModalOpen = false;
      this.state.modificationModalOpen = false;
    },
  },
});
