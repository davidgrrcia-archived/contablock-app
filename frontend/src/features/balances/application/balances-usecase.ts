import type { BalancesAdapter, BalanceListResponse } from '@/features/balances/application/balances-model';

export interface BalancesUsecase {
  listBalances(): Promise<BalanceListResponse | Error>;
}

export class DefaultBalancesUseCase implements BalancesUsecase {
  constructor(private readonly balanceAdapter: BalancesAdapter) {}

  async listBalances() {
    return await this.balanceAdapter.listBalancesHttpResponse();
  }
}
