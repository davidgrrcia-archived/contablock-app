export interface EntriesAdapter {
  createEntryHttpResponse(entry: Entry): Promise<Entry>;
}

export interface Entry {
  id?: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: string;
  quarter: number;
  debitAmount: number | null;
  creditAmount: number | null;
  totalAmount: number;
  displayName?: string;
  displayDate?: string;
  state: string;
  stateReason: string;
  creationDate: string;
  lastUpdate: string;
}
