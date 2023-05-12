import { defineStore } from 'pinia';
import type { DailyLedger } from '@/features/daily-ledger/application/daily-ledger-model';
import { dailyLedgerUseCase } from '@/features/daily-ledger';
import { useMessagesStore } from '@/stores/messages';

//TODO: apply clean architecture
interface DailyLedgerState {
  storedDailyLedger: DailyLedger[];

  stateMessage: string | null;
}

const state: DailyLedgerState = {
  storedDailyLedger: [],
  stateMessage: null,
};

export const useDailyLedgerStore = defineStore('dailyLedger', {
  state: () => {
    return { state };
  },

  actions: {
    async loadDailyLedger() {
      try {
        this.state.stateMessage = 'loading daily ledger';

        const dailyLedger = await dailyLedgerUseCase.listDailyLedger();

        if (dailyLedger instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addMessage('Fallo al cargar el libro diario');
          return;
        }

        this.state.storedDailyLedger = dailyLedger.reverse();

        setInterval(async () => {
          this.state.stateMessage = 'loading daily ledger';

          const dailyLedger = await dailyLedgerUseCase.listDailyLedger();

          if (dailyLedger instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addMessage('Fallo al cargar el libro diario');
            return;
          }

          this.state.storedDailyLedger = dailyLedger.reverse();

          this.state.stateMessage = 'daily ledger loaded';
        }, 3000);
      } catch (err) {
        this.state.stateMessage = 'error loading daily ledger';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },

  getters: {
    getDailyLedger(): DailyLedger[] {
      return this.state.storedDailyLedger;
    },
  },
});
