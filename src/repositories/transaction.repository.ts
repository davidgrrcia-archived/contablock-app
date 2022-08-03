import * as transactionEntities from "../entities/Transaction";

export default interface TransactionRepository {
  get(id: string, businessId: string): Promise<transactionEntities.TransactionOutput>;
  list(businessId: string): Promise<transactionEntities.TransactionOutput[]>;
  create(input: transactionEntities.Transaction): Promise<transactionEntities.TransactionOutput>;
  update(input: transactionEntities.Transaction): Promise<transactionEntities.TransactionOutput>;
  delete(id: string, businessId: string): Promise<transactionEntities.TransactionOutput>;
}
