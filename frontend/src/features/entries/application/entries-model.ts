export interface EntriesAdapter {
  createEntryHttpResponse(entry: CreateEntryInput): Promise<Entry | Error>;
  publishEntryHttpResponse(id: string): Promise<Entry | Error>;
  removeEntryHttpResponse(id: string): Promise<Entry | Error>;
  reviewEntryHttpResponse(id: string): Promise<Entry | Error>;
  approveEntryHttpResponse(id: string): Promise<Entry | Error>;
  rejectEntryHttpResponse(id: string): Promise<Entry | Error>;
  updateEntryHttpResponse(id: string, state: string): Promise<Entry | Error>;
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
