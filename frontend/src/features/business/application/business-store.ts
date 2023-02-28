import { defineStore } from 'pinia';
import { businessUseCase } from '@/features/business';
import type { Business } from '@/features/business/application/business-model';
import { useMessagesStore } from '@/stores/messages';

//TODO: apply clean architecture
interface BusinessState {
  business: Business;
  stateMessage: string | null;
}

const emptyBusiness: Business = {
  id: '',
  name: '',
  socialDenomination: '',
  registeredOffice: '',
  cif: '',
  logo: '',
  ethAddress: '',
};

const state: BusinessState = {
  business: emptyBusiness,
  stateMessage: null,
};

export const useBusinessStore = defineStore('business', {
  state: () => {
    return { state };
  },

  actions: {
    async loadBusiness(id: string): Promise<BusinessState> {
      try {
        this.state.stateMessage = 'loading business';

        const response = await businessUseCase.describeBusiness(id);

        if (response instanceof Error) {
          const messagesStore = useMessagesStore();
          messagesStore.addMessage('Fallo al cargar los datos de la empresa');
          return this.state;
        }

        this.state.business = response;

        this.state.stateMessage = 'business loaded';

        return this.state;
      } catch (err) {
        this.state.stateMessage = 'error loading business';
        const messagesStore = useMessagesStore();
        messagesStore.addMessage(this.state.stateMessage);
        return this.state;
      }
    },
  },
});
