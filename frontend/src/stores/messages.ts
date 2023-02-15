import { defineStore } from 'pinia';

//TODO: apply clean architecture

export enum Severities {
  'success',
  'info',
  'warning',
  'error',
  'confirm',
  'dangerousConfirm',
}

interface Message {
  id: string;
  title: string;
  description: string;
  severity: Severities;
  lifeInSeconds?: number;
  show: boolean;
  actionCallback?: (props?: any) => void;
}

const state = {
  messages: [] as Message[],
};

export const useMessagesStore = defineStore('messagesStore', {
  state: () => {
    return { state };
  },

  actions: {
    addSuccessMessage(title: string, description?: string, lifeInSeconds?: number) {
      if (description == undefined || '') {
        description = '';
      }
      if (lifeInSeconds == undefined || 0 || lifeInSeconds > 60) {
        lifeInSeconds = 60;
      }

      if (title !== undefined || '') {
        const msg: Message = {
          id: '',
          title: title,
          description: description,
          severity: Severities.success,
          lifeInSeconds: lifeInSeconds,
          show: true,
        };
        this.state.messages.push(msg);
      }
    },
    addErrorMessage(title: string, description?: string, lifeInSeconds?: number) {
      if (description == undefined || '') {
        description = '';
      }
      if (lifeInSeconds == undefined || 0 || lifeInSeconds > 60) {
        lifeInSeconds = 60;
      }

      if (title !== undefined || '') {
        const msg: Message = {
          id: '',
          title: title,
          description: description,
          severity: Severities.error,
          lifeInSeconds: lifeInSeconds,
          show: true,
        };
        this.state.messages.push(msg);
      }
    },
    addMessage(title: string, description?: string, lifeInSeconds?: number) {
      if (description == undefined || '') {
        description = '';
      }
      if (lifeInSeconds == undefined || 0 || lifeInSeconds > 60) {
        lifeInSeconds = 60;
      }

      if (title !== undefined || '') {
        const msg: Message = {
          id: '',
          title: title,
          description: description,
          severity: Severities.info,
          lifeInSeconds: lifeInSeconds,
          show: true,
        };
        this.state.messages.push(msg);
      }
    },
    addWarnMessage(title: string, description?: string, lifeInSeconds?: number) {
      if (description == undefined || '') {
        description = '';
      }
      if (lifeInSeconds == undefined || 0 || lifeInSeconds > 60) {
        lifeInSeconds = 60;
      }

      if (title !== undefined || '') {
        const msg: Message = {
          id: '',
          title: title,
          description: description,
          severity: Severities.warning,
          lifeInSeconds: lifeInSeconds,
          show: true,
        };
        this.state.messages.push(msg);
      }
    },
    addConfirmMessage(title: string, description?: string, actionCallback?: (props?: any) => void) {
      if (description == undefined || '') {
        description = '';
      }
      if (title !== undefined || '') {
        const msg: Message = {
          id: '',
          title: title,
          description: description,
          severity: Severities.confirm,
          show: true,
          actionCallback: actionCallback,
        };
        this.state.messages.push(msg);
      }
    },
    addDangerousConfirmMessage(title: string, description?: string, actionCallback?: (props?: any) => void) {
      if (description == undefined || '') {
        description = '';
      }
      if (title !== undefined || '') {
        const msg: Message = {
          id: '',
          title: title,
          description: description,
          severity: Severities.dangerousConfirm,
          show: true,
          actionCallback: actionCallback,
        };
        this.state.messages.push(msg);
      }
    },
    clearMessages() {
      this.state.messages = [];
    },
  },
});
