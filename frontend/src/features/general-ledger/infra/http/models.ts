import type { Account } from '@/features/accounts/infra/http/models';
import type { Entry } from '@/features/entries/infra/http/models';

export interface GeneralLedger {
  totalDebitAmount: number;
  totalCreditAmount: number;
  account: Account;
  entries: Entry[];
}

export interface GeneralLedgerListResponse {
  items: GeneralLedger[];
}
