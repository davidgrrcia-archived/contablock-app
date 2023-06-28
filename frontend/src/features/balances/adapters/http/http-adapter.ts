import type { DefaultBalancesHttpInfra } from '@/features/balances/infra/http/http-infra';
import type { BalancesAdapter } from '@/features/balances/application/balances-model';
import type { BalanceListResponse } from '@/features/balances/application/balances-model';

export class DefaultBalancesHttpAdapter implements BalancesAdapter {
  constructor(private readonly balanceHttpInfra: DefaultBalancesHttpInfra) {}

  async listBalancesHttpResponse() {
    const balanceHttpResponse = await this.balanceHttpInfra.listBalances();

    if (balanceHttpResponse instanceof Error) {
      return balanceHttpResponse;
    }

    const balanceListResponse: BalanceListResponse = {
      items: [],
      currentDebitTotalAmount: 0,
      previousDebitTotalAmount: 0,
      currentCreditTotalAmount: 0,
      previousCreditTotalAmount: 0,
    };

    balanceListResponse.currentDebitTotalAmount = balanceHttpResponse.currentDebitTotalAmount;
    balanceListResponse.previousDebitTotalAmount = balanceHttpResponse.previousDebitTotalAmount;
    balanceListResponse.currentCreditTotalAmount = balanceHttpResponse.currentCreditTotalAmount;
    balanceListResponse.previousCreditTotalAmount = balanceHttpResponse.previousCreditTotalAmount;

    balanceListResponse.items = balanceHttpResponse.items.map(item => ({
      id: item.id,
      name: item.name,
      currentYear: item.currentYear,
      previousYear: item.previousYear,
      currentAmount: item.currentAmount,
      previousAmount: item.previousAmount,
      accounts: item.accounts,
      child: item.child
        ? item.child.map(item2 => ({
            id: item2.id,
            name: item2.name,
            currentYear: item2.currentYear,
            previousYear: item2.previousYear,
            currentAmount: item2.currentAmount,
            previousAmount: item2.previousAmount,
            accounts: item2.accounts,
            child: item2.child
              ? item2.child.map(item3 => ({
                  id: item3.id,
                  name: item3.name,
                  currentYear: item3.currentYear,
                  previousYear: item3.previousYear,
                  currentAmount: item3.currentAmount,
                  previousAmount: item3.previousAmount,
                  accounts: item3.accounts,
                  child: item3.child
                    ? item3.child.map(item4 => ({
                        id: item4.id,
                        name: item4.name,
                        currentYear: item4.currentYear,
                        previousYear: item4.previousYear,
                        currentAmount: item4.currentAmount,
                        previousAmount: item4.previousAmount,
                        accounts: item4.accounts,
                        child: null,
                      }))
                    : null,
                }))
              : null,
          }))
        : null,
    }));

    return balanceListResponse;
  }
}
