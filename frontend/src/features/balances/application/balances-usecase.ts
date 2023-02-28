import type { Balance, BalancesAdapter } from '@/features/balances/application/balances-model';

export interface BalancesUsecase {
  listBalances(): Promise<Balance[] | Error>;
}

export class DefaultBalancesUseCase implements BalancesUsecase {
  constructor(private readonly balanceAdapter: BalancesAdapter) {}

  async listBalances() {
    return await this.balanceAdapter.listBalancesHttpResponse();
  }
}
