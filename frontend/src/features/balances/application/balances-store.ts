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
        this.state.stateMessage = 'loading balances';

        const balance = await balancesUseCase.listBalances();

        if (balance instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addMessage('Fallo al cargar los balances');
          return;
        }

        this.state.storedBalance = orderAllChildsByName(balance);

        this.state.stateMessage = 'balances loaded';

        setInterval(async () => {
          this.state.stateMessage = 'loading balances';

          const balance = await balancesUseCase.listBalances();

          if (balance instanceof Error) {
            const messagesStore = useMessagesStore();
            messagesStore.addMessage('Fallo al cargar los balances');
            return;
          }

          this.state.storedBalance = orderAllChildsByName(balance);

          this.state.stateMessage = 'balances loaded';
        }, 3000);
      } catch (err) {
        this.state.stateMessage = 'error loading balances';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },
  getters: {
    getCurrentYear(): number {
      return this.state.storedBalance[0].currentYear;
    },
    getPreviousYear(): number {
      return this.state.storedBalance[0].previousYear;
    },
  },
});

function orderByName(balance: Balance[] | null): Balance[] {
  if (!balance) {
    return [];
  }

  return balance.sort((a, b) => {
    if (a.name.toLowerCase() < b.name.toLowerCase()) {
      return -1;
    }
    if (a.name.toLowerCase() > b.name.toLowerCase()) {
      return 1;
    }
    return 0;
  });
}

function orderAllChildsByName(balance: Balance[]): Balance[] {
  balance = orderByName(balance);
  balance.forEach(balance => {
    balance.child = orderByName(balance.child);
    balance.child.forEach(child => {
      child.child = orderByName(child.child);
      child.child.forEach(child => {
        child.child = orderByName(child.child);
      });
    });
  });
  return balance;
}
