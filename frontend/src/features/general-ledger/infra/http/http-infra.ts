import type { GeneralLedgerListResponse } from '@/features/general-ledger/infra/http/models';

export interface GeneralLedgerHttpInfra {
  listGeneralLedger(): Promise<GeneralLedgerListResponse>;
}

export class DefaultGeneralLedgerHttpInfra implements GeneralLedgerHttpInfra {
  listGeneralLedger() {
    return fetch('/api/general-ledger')
      .then(response => response.json())
      .then(response => response as GeneralLedgerListResponse);
  }
}
