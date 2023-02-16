import { defineStore } from 'pinia';
import type { Balance } from '@/features/balances/application/balances-model';
import { balancesUseCase } from '@/features/balances';
import { useMessagesStore } from '@/stores/messages';

//TODO: apply clean architecture
interface BalanceState {
  storedBalance: Balance[];
  stateMessage: string | null;
}

const state: BalanceState = {
  storedBalance: [],
  stateMessage: null,
};

export const useBalanceStore = defineStore('balance', {
  state: () => {
    return { state };
  },

  actions: {
    async loadBalance() {
      try {
        this.state.storedBalance = [];
        this.state.stateMessage = 'loading balances';

        const balance = await balancesUseCase.listBalances();

        this.state.storedBalance = balance;
        this.state.stateMessage = 'balances loaded';

        setInterval(async () => {
          this.state.storedBalance = [];
          this.state.stateMessage = 'loading balances';

          const balance = await balancesUseCase.listBalances();

          this.state.storedBalance = balance;
          this.state.stateMessage = 'balances loaded';
        }, 3000);
      } catch (err) {
        this.state.stateMessage = 'error loading balances';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },
});
