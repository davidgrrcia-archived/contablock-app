export interface Entry {
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: string;
  debitAmount: number | null;
  creditAmount: number | null;
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
  items: Entry[];
}

export interface EntryDescribeResponse {
  entry: Entry;
}
