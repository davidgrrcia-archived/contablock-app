import type { GeneralLedgerListResponse } from '@/features/general-ledger/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface GeneralLedgerHttpInfra {
  listGeneralLedger(): Promise<GeneralLedgerListResponse | Error>;
}

export class DefaultGeneralLedgerHttpInfra implements GeneralLedgerHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async listGeneralLedger() {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/${configStore.state.project}/general-ledger?contractId=${configStore.state.contractId}&userId=${configStore.state.userId}`);

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as GeneralLedgerListResponse;
  }
}
