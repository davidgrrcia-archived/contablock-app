import type { Account } from '@/features/accounts/infra/http/models';
import type {Entry} from "@/features/entries/infra/http/models";

export interface DailyLedger {
  id: string;
  entriesExpanded: {
    entry: Entry
    account: Account
  }[];
}

export interface DailyLedgerListResponse {
  items: DailyLedger[];
}

