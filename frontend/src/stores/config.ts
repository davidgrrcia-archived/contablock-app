import { usersUseCase } from '@/features/users';
import type { LoginUserInput, User } from '@/features/users/application/users-model';
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
  user: User | null;
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
  isAuthenticated: new Map<string, boolean>(),
  userId: new Map<string, string>(),
  user: null,
};

export const useConfigStore = defineStore('configStore', {
  state: () => {
    return { state };
  },

  actions: {
    async activateCurrent(project: string, url: string) {
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

      //Save project in local storage
      localStorage.setItem('currentProject', project);
    },
    getIsAuthenticated(): boolean {
      return this.state.isAuthenticated.get(this.state.project) || false;
    },
    async loginUser(email: string, password: string, redirect: string) {
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

        //Save authentication status, user id and business id in local storage
        localStorage.setItem(response.project + 'Authentication', JSON.stringify(response.id !== undefined));
        localStorage.setItem(response.project + 'UserId', JSON.stringify(response.id));
        localStorage.setItem(response.project + 'BusinessId', JSON.stringify(response.businessId));

        //Redirect to the page the user was trying to access
        window.location.href = redirect;
      } catch (err) {
        this.state.stateMessage = 'error al iniciar sesión';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
      }
    },
  },
});
