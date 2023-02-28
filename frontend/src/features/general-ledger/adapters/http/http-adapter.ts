import type { DefaultGeneralLedgerHttpInfra } from '@/features/general-ledger/infra/http/http-infra';
import type { GeneralLedgerAdapter } from '@/features/general-ledger/application/general-ledger-model';

export class DefaultGeneralLedgerHttpAdapter implements GeneralLedgerAdapter {
  constructor(private readonly generalLedgerHttpInfra: DefaultGeneralLedgerHttpInfra) {}

  async listGeneralLedgerHttpResponse() {
    const generalLedgerHttpResponse = await this.generalLedgerHttpInfra.listGeneralLedger();

    if (generalLedgerHttpResponse instanceof Error) {
      return generalLedgerHttpResponse;
    }

    return generalLedgerHttpResponse.items.map(item => ({
      account: {
        id: item.account.id,
        name: item.account.name,
        level: item.account.level,
        balanceId: item.account.balanceId,
        displayName: item.account.displayName,
      },
      entries: item.entries.map(item2 => ({
        id: item2.id,
        accountId: item2.accountId,
        multipleEntryId: item2.multipleEntryId,
        currency: item2.currency,
        concept: item2.concept,
        date: item2.date,
        quarter: item2.quarter,
        debitAmount: item2.debitBalance,
        creditAmount: item2.creditBalance,
        totalAmount: item2.totalBalance,
        state: item2.state,
        stateReason: item2.stateReason,
        creationDate: item2.creationDate,
        lastUpdate: item2.lastUpdate,
      })),
    }));
  }
}
