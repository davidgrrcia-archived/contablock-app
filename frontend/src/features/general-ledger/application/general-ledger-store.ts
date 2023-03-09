import { defineStore } from 'pinia';
import type { GeneralLedger } from '@/features/general-ledger/application/general-ledger-model';
import { generalLedgerUseCase } from '@/features/general-ledger';
import { useMessagesStore } from '@/stores/messages';

//TODO: apply clean architecture
interface GeneralLedgerState {
  storedGeneralLedger: GeneralLedger[];
  stateMessage: string | null;
}

const state: GeneralLedgerState = {
  storedGeneralLedger: [],
  stateMessage: null,
};

export const useGeneralLedgerStore = defineStore('generalLedger', {
  state: () => {
    return { state };
  },

  actions: {
    async loadGeneralLedger() {
      try {
        this.state.stateMessage = 'loading general ledger';

        const generalLedger = await generalLedgerUseCase.listGeneralLedger();

        if (generalLedger instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addMessage('Fallo al cargar el libro mayor');
          return;
        }

        this.state.storedGeneralLedger = generalLedger.reverse();
        this.state.stateMessage = 'general ledger loaded';

        setInterval(async () => {
          this.state.stateMessage = 'loading general ledger';

          const generalLedger = await generalLedgerUseCase.listGeneralLedger();

          if (generalLedger instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addMessage('Fallo al cargar el libro mayor');
            return;
          }

          if (generalLedger.length !== this.state.storedGeneralLedger.length) {
            this.state.storedGeneralLedger = generalLedger.reverse();
          }

          this.state.stateMessage = 'general ledger loaded';
        }, 3000);
      } catch (err) {
        this.state.stateMessage = 'error loading general ledger';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },
});
