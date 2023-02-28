export interface EntriesAdapter {
  createEntryHttpResponse(entry: CreateEntryInput): Promise<Entry | Error>;
  publishEntryHttpResponse(id: string): Promise<Entry | Error>;
  removeEntryHttpResponse(id: string): Promise<Entry | Error>;
}

export interface Entry {
  id: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: number;
  quarter: number;
  debitAmount: number | null;
  creditAmount: number | null;
  totalAmount: number;
  state: string;
  stateReason: string;
  creationDate: number;
  lastUpdate: number;
}

export interface CreateEntryInput {
  id?: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: number;
  debitAmount: number | null;
  creditAmount: number | null;
}
