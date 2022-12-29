import { defineStore } from 'pinia';
import { entriesUseCase } from '@/features/entries';

//TODO: apply clean architecture
interface Navigation {
  name: string;
  href: string;
  current: boolean;
}

const state = {
  navigation: [
    { name: 'Libro Diario', href: '/', current: false },
    { name: 'Libro Mayor', href: '/general-ledger', current: false },
  ] as Navigation[],
};

export const useConfigStore = defineStore('configStore', {
  state: () => {
    return { state };
  },

  actions: {
    activateCurrent(url: string) {
      this.state.navigation.map(async item => {
        if (item.href === url) {
          item.current = true;
        } else {
          item.current = false;
        }
      });
    },
  },
});
