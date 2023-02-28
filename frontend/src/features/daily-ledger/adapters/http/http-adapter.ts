import type { DefaultDailyLedgerHttpInfra } from '@/features/daily-ledger/infra/http/http-infra';
import type { DailyLedgerAdapter } from '@/features/daily-ledger/application/daily-ledger-model';

export class DefaultDailyLedgerHttpAdapter implements DailyLedgerAdapter {
  constructor(private readonly dailyLedgerHttpInfra: DefaultDailyLedgerHttpInfra) {}

  async listDailyLedgerHttpResponse() {
    const dailyLedgerHttpResponse = await this.dailyLedgerHttpInfra.listDailyLedger();

    if (dailyLedgerHttpResponse instanceof Error) {
      return dailyLedgerHttpResponse;
    }

    return dailyLedgerHttpResponse.items.map(item => ({
      id: item.id,
      entriesExpanded: item.entriesExpanded.map(item2 => ({
        entry: {
          id: item2.entry.id,
          accountId: item2.entry.accountId,
          multipleEntryId: item2.entry.multipleEntryId,
          currency: item2.entry.currency,
          concept: item2.entry.concept,
          date: item2.entry.date,
          quarter: item2.entry.quarter,
          debitAmount: item2.entry.debitBalance,
          creditAmount: item2.entry.creditBalance,
          totalAmount: item2.entry.totalBalance,
          state: item2.entry.state,
          stateReason: item2.entry.stateReason,
          creationDate: item2.entry.creationDate,
          lastUpdate: item2.entry.lastUpdate,
        },
        account: {
          id: item2.account.id,
          name: item2.account.name,
          level: item2.account.level,
          balanceId: item2.account.balanceId,
          displayName: item2.account.displayName,
        },
      })),
    }));
  }
}
