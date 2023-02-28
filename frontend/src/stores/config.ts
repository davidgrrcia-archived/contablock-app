import { defineStore } from 'pinia';

//TODO: apply clean architecture
interface Navigation {
  name: string;
  href: string;
  current: boolean;
}

interface ConfigState {
  navigation: Navigation[];
  project: string;
}

const navigation: Navigation[] = [
  { name: 'Libro Diario', href: '/projects/', current: false },
  { name: 'Libro Mayor', href: '/general-ledger', current: false },
  { name: 'Balance General', href: '/balances', current: false },
];

const state: ConfigState = {
  navigation: navigation,
  project: '',
};

export const useConfigStore = defineStore('configStore', {
  state: () => {
    return { state };
  },

  actions: {
    activateCurrent(project: string, url: string) {
      this.state.project = project;

      this.state.navigation = [
        { name: 'Libro Diario', href: '/projects/' + project, current: false },
        { name: 'Libro Mayor', href: '/projects/' + project + '/general-ledger', current: false },
        { name: 'Balance General', href: '/projects/' + project + '/balances', current: false },
      ];

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
