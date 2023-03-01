import { usersUseCase } from '@/features/users';
import type { LoginUserInput } from '@/features/users/application/users-model';
import { defineStore } from 'pinia';
import { useMessagesStore } from './messages';

//TODO: apply clean architecture
interface Navigation {
  name: string;
  href: string;
  current: boolean;
}

interface ConfigState {
  stateMessage: string | null;
  navigation: Navigation[];
  project: string;
  isAuthenticated: Map<string, boolean>;
  userId: Map<string, string>;
  businessId: Map<string, string>;
  rolePermission: Map<string, string>;
}

const navigation: Navigation[] = [
  { name: 'Libro Diario', href: '/projects/', current: false },
  { name: 'Libro Mayor', href: '/general-ledger', current: false },
  { name: 'Balance General', href: '/balances', current: false },
];

const state: ConfigState = {
  stateMessage: null,
  navigation: navigation,
  project: '',
  isAuthenticated: new Map(),
  userId: new Map(),
  businessId: new Map(),
  rolePermission: new Map(),
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
    getIsAuthenticated(project: string): boolean {
      return this.state.isAuthenticated.get(project) || false;
    },
    async loginUser(email: string, password: string) {
      try {
        this.state.stateMessage = 'publishing entries';

        const input: LoginUserInput = {
          email: email,
          password: password,
        };

        const response = await usersUseCase.loginUser(input);

        if (response instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addErrorMessage('Error al iniciar sesión', 'No se ha podido iniciar sesión');
          return;
        }

        console.log('response', response);

        this.state.isAuthenticated.set(response.project, true);
        this.state.userId.set(response.project, response.id);
        this.state.businessId.set(response.project, response.businessId);
        this.state.rolePermission.set(response.project, response.role);
      } catch (err) {
        this.state.stateMessage = 'error al iniciar sesión';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },
});
