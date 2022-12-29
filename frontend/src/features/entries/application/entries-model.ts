export interface EntriesAdapter {
  listEntriesHttpResponse(): Promise<Entry[]>;
}

export interface Entry {
  id: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: string;
  debitAmount: number | null;
  creditAmount: number | null;
  displayName?: string;
}
