import { defineStore } from 'pinia';

//TODO: apply clean architecture

interface Message {
  id: string;
  severity: 'success' | 'info' | 'warn' | 'error' | undefined;
  content: string;
  life: number;
  sticky: boolean;
}

const state = {
  messages: [] as Message[],
};

export const useMessagesStore = defineStore('messagesStore', {
  state: () => {
    return { state };
  },

  actions: {
    addErrorMessage(message: string) {
      if (message !== undefined || '') {
        const msg: Message = {
          id: '',
          severity: 'error',
          content: message,
          life: 3000,
          sticky: false,
        };
        this.state.messages.push(msg);
      }
    },
    addSuccessMessage(message: string) {
      if (message !== undefined || '') {
        const msg: Message = {
          id: '',
          severity: 'success',
          content: message,
          life: 3000,
          sticky: false,
        };
        this.state.messages.push(msg);
      }
    },
    clearMessages() {
      this.state.messages = [];
    },
  },
});
