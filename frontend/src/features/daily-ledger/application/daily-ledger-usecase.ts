import type { DailyLedger, DailyLedgerAdapter } from '@/features/daily-ledger/application/daily-ledger-model';

export interface DailyLedgerUsecase {
  listDailyLedger(): Promise<DailyLedger[] | Error>;
}

export class DefaultDailyLedgerUseCase implements DailyLedgerUsecase {
  constructor(private readonly dailyLedgerAdapter: DailyLedgerAdapter) {}

  async listDailyLedger() {
    return await this.dailyLedgerAdapter.listDailyLedgerHttpResponse();
  }
}
