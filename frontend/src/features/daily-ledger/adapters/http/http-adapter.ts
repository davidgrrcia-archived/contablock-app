import type { DefaultDailyLedgerHttpInfra } from '@/features/daily-ledger/infra/http/http-infra';
import type { DailyLedgerAdapter } from '@/features/daily-ledger/application/daily-ledger-model';

export class DefaultDailyLedgerHttpAdapter implements DailyLedgerAdapter {
  constructor(private readonly dailyLedgerHttpInfra: DefaultDailyLedgerHttpInfra) {}

  async listDailyLedgerHttpResponse() {
    const dailyLedgerHttpResponse = await this.dailyLedgerHttpInfra.listDailyLedger();

    return dailyLedgerHttpResponse.items.map(item => ({
      id: item.id,
      entriesExpanded: item.entriesExpanded.map(item2 => ({
        entry: {
          accountId: item2.entry.accountId,
          multipleEntryId: item2.entry.multipleEntryId,
          currency: item2.entry.currency,
          concept: item2.entry.concept,
          date: item2.entry.date,
          debitAmount: item2.entry.debitBalance,
          creditAmount: item2.entry.creditBalance,
        },
        account: {
          id: item2.account.id,
          name: item2.account.name,
          level: item2.account.level,
          type: item2.account.type,
          displayName: item2.account.displayName,
        },
      })),
    }));
  }
}
