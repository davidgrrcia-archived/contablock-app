import type { DailyLedgerListResponse } from '@/features/daily-ledger/infra/http/models';

export interface DailyLedgerHttpInfra {
  listDailyLedger(): Promise<DailyLedgerListResponse>;
}

export class DefaultDailyLedgerHttpInfra implements DailyLedgerHttpInfra {
  listDailyLedger() {
    return fetch('/api/daily-ledger')
      .then(response => response.json())
      .then(response => response as DailyLedgerListResponse);
  }
}
