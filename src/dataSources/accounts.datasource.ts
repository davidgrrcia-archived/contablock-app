import Account from "../entities/Account";
import AccountRepository from "../repositories/account.repository";
import { promisePool } from "../database";

export default class AccountPostgres implements AccountRepository {
  public async get(id: string): Promise<Account> {
    try {
      const sqlQuery = "SELECT * FROM accounts WHERE id = ?";
      const [rows] = await promisePool.query(sqlQuery, [id]);
      const result=Object.values(rows);

      const account: Account = {
        id: result[0]["id"],
        name: result[0]["name"],
        type: result[0]["type"],
        kind: result[0]["kind"],
      };

      return account;
    } catch (e) {
      throw e;
    }
  }

  public async list(): Promise<Account[]> {
    try {
      const sqlQuery = "SELECT * FROM accounts";
      const [rows] = await promisePool.query(sqlQuery);
      const result=Object.values(rows);

      let accounts: Account[] = [];

      for (let i = 0; i < result.length; i++) {
        let a: Account = {
          id: result[i]["id"],
          name: result[i]["name"],
          type: result[i]["type"],
          kind: result[i]["kind"],
        };
        accounts.push(a);
      }

      return accounts;
    } catch (e) {
      throw e;
    }
  }

  public async listByType(type: string): Promise<Account[]> {
    try {
      const sqlQuery = "SELECT * FROM accounts WHERE type = ?";
      const [rows] = await promisePool.query(sqlQuery, [type]);
      const result=Object.values(rows);

      let accounts: Account[] = [];

      for (let i = 0; i < result.length; i++) {
        let a: Account = {
          id: result[i]["id"],
          name: result[i]["name"],
          type: result[i]["type"],
          kind: result[i]["kind"],
        };
        accounts.push(a);
      }

      return accounts;
    } catch (e) {
      throw e;
    }
  }
}
