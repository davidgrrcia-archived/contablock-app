import * as userEntities from "../entities/User";
import UserRepository from "../repositories/user.repository";
import { promisePool } from "../database";

export default class UserPostgres implements UserRepository {
  public async get(id: string, businessId: string): Promise<userEntities.UserOutput> {
    try {
      const sqlQuery = "SELECT * FROM users WHERE id = ? AND business_id = ?";
      const [rows] = await promisePool.query(sqlQuery, [id, businessId]);
      const result = Object.values(rows);

      const user: userEntities.UserOutput = {
        id: result[0]["id"],
        businessId: result[0]["business_id"],
        name: result[0]["name"],
        email: result[0]["email"],
        password: result[0]["password"],
        phone: result[0]["phone"],
        role: result[0]["role"],
        state: result[0]["state"],
        creationDate: result[0]["creation_date"],
        lastUpdate: result[0]["last_update"],
      };

      return user;
    } catch (e) {
      throw e;
    }
  }

  public async list(businessId: string): Promise<userEntities.UserOutput[]> {
    try {
      const sqlQuery = "SELECT * FROM users WHERE business_id = ?";
      const [rows] = await promisePool.query(sqlQuery, [businessId]);
      const result = Object.values(rows);

      let user: userEntities.UserOutput[] = [];

      for (let i = 0; i < result.length; i++) {
        let u: userEntities.UserOutput = {
          id: result[i]["id"],
          businessId: result[i]["business_id"],
          name: result[i]["name"],
          email: result[i]["email"],
          password: result[i]["password"],
          phone: result[i]["phone"],
          role: result[i]["role"],
          state: result[i]["state"],
          creationDate: result[i]["creation_date"],
          lastUpdate: result[i]["last_update"],
        };
        user.push(u);
      }

      return user;
    } catch (e) {
      throw e;
    }
  }

  public async create(userInput: userEntities.User): Promise<userEntities.UserOutput> {
    try {
      const sqlQuery =
        "INSERT INTO users (" +
          "id," +
          "business_id," +
          "name," +
          "email," +
          "password," +
          "phone," +
          "role," +
          "state," +
          "creation_date," +
          "last_update" +
          ") " +
          "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";
        await promisePool.query(sqlQuery, [
          userInput.id,
          userInput.businessId,
          userInput.name,
          userInput.email,
          userInput.password,
          userInput.phone,
          userInput.role,
          userInput.state,
          userInput.creationDate,
          userInput.lastUpdate,
        ]
      );

      const userOutput: userEntities.UserOutput = await this.get(userInput.id!, userInput.businessId);

      return userOutput;
    } catch (e) {
      throw e;
    }
  }

  public async update(userInput: userEntities.User): Promise<userEntities.UserOutput> {
    try {
      const sqlQuery =
        "UPDATE users SET " +
          "name = ?," +
          "email = ?," +
          "password = ?," +
          "phone = ?," +
          "role = ?," +
          "state = ?," +
          "last_update = ?" +
          " WHERE id = ? AND business_id = ?";
      await promisePool.query(sqlQuery, [
        userInput.name,
        userInput.email,
        userInput.password,
        userInput.phone,
        userInput.role,
        userInput.state,
        userInput.lastUpdate,
        userInput.id,
        userInput.businessId,
      ]);

      const userOutput: userEntities.UserOutput = await this.get(userInput.id!, userInput.businessId);

      return userOutput;
    } catch (e) {
      throw e;
    }
  }

  public async delete(id: string, businessId: string): Promise<userEntities.UserOutput> {
    try {

      const userOutput: userEntities.UserOutput = await this.get(id, businessId);

      const sqlQuery = "DELETE FROM users WHERE id = ? AND business_id = ?";
      await promisePool.query(sqlQuery, [id, businessId]);

      return userOutput;
    } catch (e) {
      throw e;
    }
  }
}
