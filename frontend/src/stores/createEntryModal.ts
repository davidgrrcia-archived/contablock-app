import { defineStore } from 'pinia';

//TODO: apply clean architecture

export const useCreateEntryModalStore = defineStore('createEntryModal', {
  state: () => {
    return { isOpen: false };
  },

  actions: {
    showModal() {
      this.isOpen = true;
    },
    hideModal() {
      this.isOpen = false;
    },
  }
});


