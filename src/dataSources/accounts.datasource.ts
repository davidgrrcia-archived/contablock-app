import * as entities from "../entities/Account";
import AccountRepository from "../repositories/account.repository";
import { promisePool } from "../database";

export default class AccountPostgres implements AccountRepository {
  public async get(
    input: entities.GetAccountStorageInput
  ): Promise<entities.GetAccountStorageOutput> {
    try {
      const sqlQuery = "SELECT * FROM accounts WHERE id = ?";
      const [rows] = await promisePool.query(sqlQuery, [input.id]);
      const result = Object.values(rows);

      const getAccountStorageOutput: entities.GetAccountStorageOutput = {
        id: result[0]["id"],
        name: result[0]["name"],
        type: result[0]["type"],
        kind: result[0]["kind"],
      };

      return getAccountStorageOutput;
    } catch (e) {
      throw e;
    }
  }

  public async list(): Promise<entities.ListAccountsStorageOutput> {
    try {
      const sqlQuery = "SELECT * FROM accounts";
      const [rows] = await promisePool.query(sqlQuery);
      const result = Object.values(rows);

      const listAccountsStorageOutput: entities.ListAccountsStorageOutput = {
        items: result.map((item) => ({
          id: item["id"],
          name: item["name"],
          type: item["type"],
          kind: item["kind"],
        })),

      return listAccountsStorageOutput;
    } catch (e) {
      throw e;
    }
  }
}
