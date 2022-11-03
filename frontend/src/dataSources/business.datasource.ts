import * as entities from "../entities/Business";
import { BusinessStorageInterface } from "../entities/Business";
import { promisePool } from "../database";

export default class BusinessStorage implements BusinessStorageInterface {
  public async get(
    input: entities.GetBusinessStorageInput
  ): Promise<entities.GetBusinessStorageOutput> {
    try {
      const sqlQuery = "SELECT * FROM business WHERE id = ?";
      const [rows] = await promisePool.query(sqlQuery, [input.id]);
      const result = Object.values(rows);

      const getBusinessStorageOutput: entities.GetBusinessStorageOutput = {
        id: result[0]["id"],
        name: result[0]["name"],
        shortDescription: result[0]["short_description"],
        longDescription: result[0]["long_description"],
        logo: result[0]["logo"],
        stockPrice: result[0]["stock_price"],
        urlWebsite: result[0]["url_website"],
        ceoName: result[0]["ceo_name"],
        numberOfEmployees: result[0]["number_of_employees"],
        country: result[0]["country"],
        nif: result[0]["nif"],
        fiscalName: result[0]["fiscal_name"],
        address: result[0]["address"],
        phone: result[0]["phone"],
        email: result[0]["email"],
        state: result[0]["state"],
        creationDate: result[0]["creation_date"],
        lastUpdate: result[0]["last_update"],
      };

      return getBusinessStorageOutput;
    } catch (e) {
      throw e;
    }
  }

  public async list(): Promise<entities.ListBusinessStorageOutput> {
    try {
      const sqlQuery = "SELECT * FROM business";
      const [rows] = await promisePool.query(sqlQuery);
      const result = Object.values(rows);

      const listBusinessStorageOutput: entities.ListBusinessStorageOutput = {
        items: result.map((item) => ({
          id: item["id"],
          name: item["name"],
          shortDescription: item["short_description"],
          longDescription: item["long_description"],
          logo: item["logo"],
          stockPrice: item["stock_price"],
          urlWebsite: item["url_website"],
          ceoName: item["ceo_name"],
          numberOfEmployees: item["number_of_employees"],
          country: item["country"],
          nif: item["nif"],
          fiscalName: item["fiscal_name"],
          address: item["address"],
          phone: item["phone"],
          email: item["email"],
          state: item["state"],
          creationDate: item["creation_date"],
          lastUpdate: item["last_update"],
        })),
      };

      return listBusinessStorageOutput;
    } catch (e) {
      throw e;
    }
  }

  public async create(
    input: entities.CreateBusinessStorageInput
  ): Promise<entities.CreateBusinessStorageOutput> {
    try {
      const sqlQuery =
        "INSERT INTO business (" +
        "id," +
        "name," +
        "short_description," +
        "long_description," +
        "logo," +
        "stock_price," +
        "url_website," +
        "ceo_name," +
        "number_of_employees," +
        "country," +
        "nif," +
        "fiscal_name," +
        "address," +
        "phone," +
        "email," +
        "state," +
        "creation_date," +
        "last_update" +
        ") " +
        "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";

      await promisePool.query(sqlQuery, [
        input.id,
        input.name,
        input.shortDescription,
        input.longDescription,
        input.logo,
        input.stockPrice,
        input.urlWebsite,
        input.ceoName,
        input.numberOfEmployees,
        input.country,
        input.nif,
        input.fiscalName,
        input.address,
        input.phone,
        input.email,
        input.state,
        input.creationDate,
        input.lastUpdate,
      ]);

      const getBusinessStorageInput: entities.GetBusinessStorageInput = {
        id: input.id,
      };
      const createBusinessStorageOutput: entities.CreateBusinessStorageOutput =
        await this.get(getBusinessStorageInput);

      return createBusinessStorageOutput;
    } catch (e) {
      throw e;
    }
  }

  public async update(
    input: entities.UpdateBusinessStorageInput
  ): Promise<entities.UpdateBusinessStorageOutput> {
    try {
      const sqlQuery =
        "UPDATE business SET " +
        "name = ?," +
        "short_description = ?," +
        "long_description = ?," +
        "logo = ?," +
        "stock_price = ?," +
        "url_website = ?," +
        "ceo_name = ?," +
        "number_of_employees = ?," +
        "country = ?," +
        "nif = ?," +
        "fiscal_name = ?," +
        "address = ?," +
        "phone = ?," +
        "email = ?," +
        "state = ?," +
        "last_update = ?" +
        " WHERE id = ?";

      await promisePool.query(sqlQuery, [
        input.name,
        input.shortDescription,
        input.longDescription,
        input.logo,
        input.stockPrice,
        input.urlWebsite,
        input.ceoName,
        input.numberOfEmployees,
        input.country,
        input.nif,
        input.fiscalName,
        input.address,
        input.phone,
        input.email,
        input.state,
        input.lastUpdate,
        input.id,
      ]);

      const getBusinessStorageInput: entities.GetBusinessStorageInput = {
        id: input.id,
      };
      const updateBusinessStorageOutput: entities.UpdateBusinessStorageOutput =
        await this.get(getBusinessStorageInput);

      return updateBusinessStorageOutput;
    } catch (e) {
      throw e;
    }
  }

  public async delete(
    input: entities.DeleteBusinessStorageInput
  ): Promise<entities.DeleteBusinessStorageOutput> {
    try {
      const getBusinessStorageInput: entities.GetBusinessStorageInput = {
        id: input.id,
      };
      const deleteBusinessStorageOutput: entities.DeleteBusinessStorageOutput =
        await this.get(getBusinessStorageInput);

      const sqlQuery = "DELETE FROM business WHERE id = ?";
      await promisePool.query(sqlQuery, [input.id]);

      return deleteBusinessStorageOutput;
    } catch (e) {
      throw e;
    }
  }
}
