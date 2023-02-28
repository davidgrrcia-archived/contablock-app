import type { GeneralLedger, GeneralLedgerAdapter } from '@/features/general-ledger/application/general-ledger-model';

export interface GeneralLedgerUsecase {
  listGeneralLedger(): Promise<GeneralLedger[] | Error>;
}

export class DefaultGeneralLedgerUseCase implements GeneralLedgerUsecase {
  constructor(private readonly generalLedgerAdapter: GeneralLedgerAdapter) {}

  async listGeneralLedger() {
    return await this.generalLedgerAdapter.listGeneralLedgerHttpResponse();
  }
}
