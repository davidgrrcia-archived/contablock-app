import type { Entry } from '@/features/entries/application/entries-model';
import type { Account } from '@/features/accounts/application/accounts-model';

export interface GeneralLedgerAdapter {
  listGeneralLedgerHttpResponse(): Promise<GeneralLedger[]>;
}

export interface GeneralLedger {
  entries: Entry[];
  account: Account;
}
