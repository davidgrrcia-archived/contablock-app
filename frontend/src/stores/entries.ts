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

const currencies = [
  {
    code: 'EUR',
    symbol: '€',
  },
  {
    code: 'GBP',
    symbol: '£',
  },
  {
    code: 'USD',
    symbol: '$',
  },
];

function findCurrencyByCode(code: string): string {
  console.log('code', code)

  const c = currencies.find(currency => currency.code === code);

  return c ? c.symbol : '';
}

export const useEntriesStore = defineStore('entries', {
  state: () => {
    return { state };
  },

  actions: {
    async addDebitEntry(entry: Entry) {
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
      const displayName = entry.accountId + ' - ' + account.name;
      entry.displayName = displayName;

      this.state.entriesCreation.push(entry);
      this.state.totalBalance = this.state.totalBalance + entry.debitAmount;
    },
    async addCreditEntry(entry: Entry) {
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
      const displayName = entry.accountId + ' - ' + account.name;
      entry.displayName = displayName;

      this.state.entriesCreation.push(entry);
      this.state.totalBalance = this.state.totalBalance - entry.creditAmount;
    },
    async createEntries() {
      const messagesStore = useMessagesStore();

      try {
        if (this.state.totalBalance > 0) {
          messagesStore.addWarnMessage('El principio de partida doble es obligatorio', 'Para poder crear un asiento el debe y el haber deben estar balanceados, actualmente el debe es mayor que el haber en ' + this.state.totalBalance + findCurrencyByCode(this.state.entriesCreation[0].currency));
        }
        if (this.state.totalBalance < 0) {
          messagesStore.addWarnMessage('El principio de partida doble es obligatorio', 'Para poder crear un asiento el debe y el haber deben estar balanceados, actualmente el debe es menor que el haber en ' +  this.state.totalBalance + findCurrencyByCode(this.state.entriesCreation[0].currency));
        }
        if (this.state.entriesCreation.length === 0) {
          messagesStore.addWarnMessage('No hay asientos creados', 'Para poder crear un asiento debes crear al menos una partida en el debe y otra en el haber');
        }

        this.state.stateMessage = 'creating entries';

        this.state.entriesCreation.map(async entry => {
          return await entriesUseCase.createEntry(entry);
        });

        this.state.stateMessage = 'entries created';
        this.clearEntries();
      } catch (err) {
        this.state.stateMessage = 'error creating entries';
        const messagesStore = useMessagesStore();
        messagesStore.addErrorMessage(err.toString());
      }
    },
    clearEntries() {
      this.state.entriesCreation = [];
      this.state.multipleEntryId = '';
      this.state.totalBalance = 0;
    },
  },
});
