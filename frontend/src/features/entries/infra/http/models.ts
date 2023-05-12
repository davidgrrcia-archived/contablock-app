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
  state: string;
  stateReason: string;
  creationDate: number;
  lastUpdate: number;
}

export interface EntryCreateBody {
  contractId: string;
  userId: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: number;
  debitBalance: number | null;
  creditBalance: number | null;
}

export interface EntryPublishBody {
  id: string;
  contractId: string;
  userId: string;
}

export interface EntryRemoveBody {
  id: string;
  contractId: string;
  userId: string;
}

export interface EntryUpdateBody {
  id: string;
  state: string;
}

export interface EntryReviewBody {
  id: string;
  contractId: string;
  userId: string;
}

export interface EntryApproveBody {
  id: string;
  contractId: string;
  userId: string;
}

export interface EntryRejectBody {
  id: string;
  contractId: string;
  userId: string;
}

export interface EntryDescribeResponse {
  id: string;
  accountId: string;
  multipleEntryId: string;
  currency: string;
  concept: string;
  date: number;
  quarter: number;
  debitBalance: number | null;
  creditBalance: number | null;
  state: string;
  stateReason: string;
  creationDate: number;
  lastUpdate: number;
}
