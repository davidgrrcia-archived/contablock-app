import type { Account } from '@/features/accounts/infra/http/models';

export interface Entry {
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: string;
  quarter: string;
  debitBalance: number | null;
  creditBalance: number | null;
  totalBalance: number;
  state: string;
  stateReason: string;
  creationDate: string;
  lastUpdate: string;
}

export interface EntryCreateBody {
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: string;
  debitBalance: number | null;
  creditBalance: number | null;
}

export interface EntryListResponse {
  items: {
    entry: Entry;
    account: Account | null;
  }[];
}

export interface EntryDescribeResponse {
  entry: Entry;
}
