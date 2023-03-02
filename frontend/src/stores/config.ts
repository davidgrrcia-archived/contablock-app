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
  isAuthenticated: boolean;
  userId: string;
  businessId: string;
  name: string;
  role: string;
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
  isAuthenticated: false,
  userId: '',
  businessId: '',
  name: '',
  role: '',
};

export const useConfigStore = defineStore('configStore', {
  state: () => {
    return { state };
  },

  actions: {
    async activateCurrent(project: string, url: string) {
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

      //Hydrate authentication status, user id and business id from local storage
      const isAuthenticated = localStorage.getItem(project + 'Authentication') === 'true';
      const userId = localStorage.getItem(project + 'UserId')?.replace(/['"]+/g, '') || '';
      const businessId = localStorage.getItem(project + 'BusinessId')?.replace(/['"]+/g, '') || '';
      const name = localStorage.getItem(project + 'Name')?.replace(/['"]+/g, '') || '';
      const role = localStorage.getItem(project + 'Role')?.replace(/['"]+/g, '') || '';
      this.state.project = project;
      this.state.isAuthenticated = isAuthenticated;
      this.state.userId = userId;
      this.state.businessId = businessId;
      this.state.name = name;
      this.state.role = role;
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
        localStorage.setItem(response.project + 'Name', JSON.stringify(response.name));
        localStorage.setItem(response.project + 'Role', JSON.stringify(response.role));

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
