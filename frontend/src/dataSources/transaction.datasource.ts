import * as transactionEntities from "../entities/Transaction";
import TransactionRepository from "../repositories/transaction.repository";
import { promisePool } from "../database";

export default class TransactionPostgres implements TransactionRepository {
  public async get(id: string, businessId: string): Promise<transactionEntities.TransactionOutput> {
    try {
      const sqlQuery = "SELECT * FROM transactions WHERE id = ? AND business_id = ?";
      const [rows] = await promisePool.query(sqlQuery, [id, businessId]);
      const result = Object.values(rows);

      const transaction: transactionEntities.TransactionOutput = {
        id: result[0]["id"],
        businessId: result[0]["business_id"],
        entryId: result[0]["entry_id"],
        accountId: result[0]["account_id"],
        type: result[0]["type"],
        concept: result[0]["concept"],
        debitAmount: result[0]["debit_amount"],
        creditAmount: result[0]["credit_amount"],
        totalAmount: result[0]["total_amount"],
        state: result[0]["state"],
        creationDate: result[0]["creation_date"],
        lastUpdate: result[0]["last_update"],
      };

      return transaction;
    } catch (e) {
      throw e;
    }
  }

  public async list(businessId: string): Promise<transactionEntities.TransactionOutput[]> {
    try {
      const sqlQuery = "SELECT * FROM transactions WHERE business_id = ?";
      const [rows] = await promisePool.query(sqlQuery, [businessId]);
      const result = Object.values(rows);

      let transaction: transactionEntities.TransactionOutput[] = [];

      for (let i = 0; i < result.length; i++) {
        let t: transactionEntities.TransactionOutput = {
          id: result[i]["id"],
          businessId: result[i]["business_id"],
          entryId: result[i]["entry_id"],
          accountId: result[i]["account_id"],
          type: result[i]["type"],
          concept: result[i]["concept"],
          debitAmount: result[i]["debit_amount"],
          creditAmount: result[i]["credit_amount"],
          totalAmount: result[i]["total_amount"],
          state: result[i]["state"],
          creationDate: result[i]["creation_date"],
          lastUpdate: result[i]["last_update"],
        };
        transaction.push(t);
      }

      return transaction;
    } catch (e) {
      throw e;
    }
  }

  public async create(transactionInput: transactionEntities.Transaction): Promise<transactionEntities.TransactionOutput> {
    try {
      const sqlQuery = 
        "INSERT INTO transactions (" +
        "id," +
        "business_id," +
        "entry_id," +
        "account_id," +
        "type," +
        "concept," +
        "debit_amount," +
        "credit_amount," +
        "total_amount," +
        "state," +
        "creation_date," +
        "last_update" +
        ") " +
        "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";
        await promisePool.query(sqlQuery, [
        transactionInput.id,
        transactionInput.businessId,
        transactionInput.entryId,
        transactionInput.accountId,
        transactionInput.type,
        transactionInput.concept,
        transactionInput.debitAmount,
        transactionInput.creditAmount,
        transactionInput.totalAmount,
        transactionInput.state,
        transactionInput.creationDate,
        transactionInput.lastUpdate,
      ]
      );

      const transactionOutput: transactionEntities.TransactionOutput = await this.get(transactionInput.id!, transactionInput.businessId);

      return transactionOutput;
    } catch (e) {
      throw e;
    }
  }

  public async update(transactionInput: transactionEntities.Transaction): Promise<transactionEntities.TransactionOutput> {
    try {
      const sqlQuery =
        "UPDATE transactions SET " +
        "entry_id = ?," +
        "account_id = ?," +
        "type = ?," +
        "concept = ?," +
        "debit_amount = ?," +
        "credit_amount = ?," +
        "total_amount = ?," +
        "state = ?," +
        "creation_date = ?," +
        "last_update = ? " +
        "WHERE id = ? AND business_id = ?";
      await promisePool.query(sqlQuery, [
        transactionInput.entryId,
        transactionInput.accountId,
        transactionInput.type,
        transactionInput.concept,
        transactionInput.debitAmount,
        transactionInput.creditAmount,
        transactionInput.totalAmount,
        transactionInput.state,
        transactionInput.creationDate,
        transactionInput.lastUpdate,
        transactionInput.id,
        transactionInput.businessId,
        ]
      );

      const transactionOutput: transactionEntities.TransactionOutput = await this.get(transactionInput.id!, transactionInput.businessId);

      return transactionOutput;
    } catch (e) {
      throw e;
    }
  }

  public async delete(id: string, businessId: string): Promise<transactionEntities.TransactionOutput> {
    try {

      const transactionOutput: transactionEntities.TransactionOutput = await this.get(id, businessId);

      const sqlQuery = "DELETE FROM transactions WHERE id = ? AND business_id = ?";
      await promisePool.query(sqlQuery, [id, businessId]);

      return transactionOutput;
    } catch (e) {
      throw e;
    }
  }
}
