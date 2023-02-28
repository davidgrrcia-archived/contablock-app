import type { DailyLedgerListResponse } from '@/features/daily-ledger/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface DailyLedgerHttpInfra {
  listDailyLedger(): Promise<DailyLedgerListResponse | Error>;
}

export class DefaultDailyLedgerHttpInfra implements DailyLedgerHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async listDailyLedger() {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/${configStore.state.project}/daily-ledger`);

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as DailyLedgerListResponse;
  }
}
