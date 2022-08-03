enum TransactionType {
  Debit = "Debit",
  Credit = "Credit",
}

export interface Transaction {
  id?: string;
  businessId: string;
  entryId: string;
  accountId: string;
  type?: string;
  concept: string;
  debitAmount?: number;
  creditAmount?: number;
  totalAmount?: number;
  state?: string;
  creationDate?: Date;
  lastUpdate?: Date;
}

export interface TransactionOutput {
  id: string;
  businessId: string;
  entryId: string;
  accountId: string;
  type: string;
  concept: string;
  debitAmount: number;
  creditAmount: number;
  totalAmount: number;
  state: string;
  creationDate: Date;
  lastUpdate: Date;
}