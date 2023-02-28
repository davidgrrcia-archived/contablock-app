import type { BalanceListResponse } from '@/features/balances/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface BalancesHttpInfra {
  listBalances(): Promise<BalanceListResponse | Error>;
}

export class DefaultBalancesHttpInfra implements BalancesHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async listBalances() {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/${configStore.state.project}/balances?year=${new Date().getFullYear()}`);

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as BalanceListResponse;
  }
}
