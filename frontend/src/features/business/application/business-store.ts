import { defineStore } from 'pinia';
import { businessUseCase } from '@/features/business';
import type { Business } from '@/features/business/application/business-model';
import { useMessagesStore } from '@/stores/messages';
import { useConfigStore } from '@/stores/config';

//TODO: apply clean architecture
interface BusinessState {
  business: Business;
  stateMessage: string | null;
}

const emptyBusiness: Business = {
  project: '',
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
    async loadBusiness(): Promise<BusinessState> {
      try {
        const configStore = useConfigStore();

        this.state.stateMessage = 'loading business';

        const response = await businessUseCase.describeBusiness(configStore.state.project, configStore.state.businessId);

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
