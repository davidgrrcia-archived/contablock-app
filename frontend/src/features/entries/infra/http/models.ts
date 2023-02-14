export interface Entry {
  id: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: number;
  quarter: number;
  debitBalance: number | null;
  creditBalance: number | null;
  totalBalance: number;
  state: string;
  stateReason: string;
  creationDate: number;
  lastUpdate: number;
}

export interface EntryCreateBody {
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: number;
  debitBalance: number | null;
  creditBalance: number | null;
}

export interface EntryDescribeResponse {
  entry: Entry;
}
