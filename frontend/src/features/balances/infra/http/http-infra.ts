import type { BalanceListResponse } from '@/features/balances/infra/http/models';

export interface BalancesHttpInfra {
  listBalances(): Promise<BalanceListResponse>;
}

export class DefaultBalancesHttpInfra implements BalancesHttpInfra {
  listBalances() {
    return fetch('/api/balances?year='+ new Date().getFullYear())
      .then(response => response.json())
      .then(response => response as BalanceListResponse);
  }
}
