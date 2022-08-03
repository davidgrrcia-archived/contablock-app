import * as entryEntities from "../entities/Entry";
import EntryRepository from "../repositories/entry.repository";
import { promisePool } from "../database";

export default class EntryPostgres implements EntryRepository {
  public async get(id: string, businessId: string): Promise<entryEntities.EntryOutput> {
    try {
      const sqlQuery = "SELECT * FROM entries WHERE id = ? AND business_id = ?";
      const [rows] = await promisePool.query(sqlQuery, [id, businessId]);
      const result = Object.values(rows);

      const entry: entryEntities.EntryOutput = {
        id: result[0]["id"],
        businessId: result[0]["business_id"],
        currency: result[0]["currency"],
        description: result[0]["description"],
        participant: result[0]["participant"],
        fileUpload: result[0]["file_upload"],
        date: result[0]["date"],
        type: result[0]["type"],
        debitBalance: result[0]["debit_balance"],
        creditBalance: result[0]["credit_balance"],
        totalBalance: result[0]["total_balance"],
        state: result[0]["state"],
        creationDate: result[0]["creation_date"],
        lastUpdate: result[0]["last_update"],
      };

      return entry;
    } catch (e) {
      throw e;
    }
  }

  public async list(businessId: string): Promise<entryEntities.EntryOutput[]> {
    try {
      const sqlQuery = "SELECT * FROM entries WHERE business_id = ?";
      const [rows] = await promisePool.query(sqlQuery, [businessId]);
      const result = Object.values(rows);

      let entry: entryEntities.EntryOutput[] = [];

      for (let i = 0; i < result.length; i++) {
        let e: entryEntities.EntryOutput = {
          id: result[i]["id"],
          businessId: result[i]["business_id"],
          currency: result[i]["currency"],
          description: result[i]["description"],
          participant: result[i]["participant"],
          fileUpload: result[i]["file_upload"],
          date: result[i]["date"],
          type: result[i]["type"],
          debitBalance: result[i]["debit_balance"],
          creditBalance: result[i]["credit_balance"],
          totalBalance: result[i]["total_balance"],
          state: result[i]["state"],
          creationDate: result[i]["creation_date"],
          lastUpdate: result[i]["last_update"],
        };
        entry.push(e);
      }

      return entry;
    } catch (e) {
      throw e;
    }
  }

  public async create(entryInput: entryEntities.Entry): Promise<entryEntities.EntryOutput> {
    try {
      const sqlQuery =
        "INSERT INTO entries (" +
        "id," +
        "business_id," +
        "currency," +
        "description," +
        "participant," +
        "file_upload," +
        "date," +
        "type," +
        "debit_balance," +
        "credit_balance," +
        "total_balance," +
        "state," +
        "creation_date," +
        "last_update" +
        ") " +
        "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";
      await promisePool.query(sqlQuery, [
        entryInput.id,
        entryInput.businessId,
        entryInput.currency,
        entryInput.description,
        entryInput.participant,
        entryInput.fileUpload,
        entryInput.date,
        entryInput.type,
        entryInput.debitBalance,
        entryInput.creditBalance,
        entryInput.totalBalance,
        entryInput.state,
        entryInput.creationDate,
        entryInput.lastUpdate,
      ]);

      const entryOutput: entryEntities.EntryOutput = await this.get(entryInput.id!, entryInput.businessId);

      return entryOutput;
    } catch (e) {
      throw e;
    }
  }

  public async update(entryInput: entryEntities.Entry): Promise<entryEntities.EntryOutput> {
    try {
      const sqlQuery =
        "UPDATE entries SET " +
        "currency = ?," +
        "description = ?," +
        "participant = ?," +
        "file_upload = ?," +
        "date = ?," +
        "type = ?," +
        "debit_balance = ?," +
        "credit_balance = ?," +
        "total_balance = ?," +
        "state = ?," +
        "creation_date = ?," +
        "last_update = ?" +
        " WHERE id = ? AND business_id = ?";
      await promisePool.query(sqlQuery, [
        entryInput.currency,
        entryInput.description,
        entryInput.participant,
        entryInput.fileUpload,
        entryInput.date,
        entryInput.type,
        entryInput.debitBalance,
        entryInput.creditBalance,
        entryInput.totalBalance,
        entryInput.state,
        entryInput.creationDate,
        entryInput.lastUpdate,
        entryInput.id,
        entryInput.businessId,
      ]);

      const entryOutput: entryEntities.EntryOutput = await this.get(entryInput.id!, entryInput.businessId);

      return entryOutput;
    } catch (e) {
      throw e;
    }
  }

  public async delete(id: string, businessId: string): Promise<entryEntities.EntryOutput> {
    try {
      const entryOutput: entryEntities.EntryOutput = await this.get(id, businessId);

      const sqlQuery = "DELETE FROM entries WHERE id = ? AND business_id = ?";
      await promisePool.query(sqlQuery, [id, businessId]);

      return entryOutput;
    } catch (e) {
      throw e;
    }
  }
}
