import { defineStore } from 'pinia';

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
  },
});
