import type { Entry } from '@/features/entries/application/entries-model';
import type { Account } from '@/features/accounts/application/accounts-model';

export interface GeneralLedgerAdapter {
  listGeneralLedgerHttpResponse(): Promise<GeneralLedger[] | Error>;
}

export interface GeneralLedger {
  totalDebitAmount: number;
  totalCreditAmount: number;
  entries: Entry[];
  account: Account;
}
