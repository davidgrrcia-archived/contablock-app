import type { Entry } from '@/features/entries/application/entries-model';
import type { Account } from '@/features/accounts/application/accounts-model';

export interface DailyLedgerAdapter {
  listDailyLedgerHttpResponse(): Promise<DailyLedger[] | Error>;
}

export interface DailyLedger {
  id: string;
  totalDebitAmount: number;
  totalCreditAmount: number;
  entriesExpanded: {
    entry: Entry;
    account: Account;
  }[];
}
