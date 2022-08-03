import * as businessEntities from "../entities/Business";
import BusinessRepository from "../repositories/business.repository";
import { promisePool } from "../database";

export default class BusinessPostgres implements BusinessRepository {
  public async get(id: string): Promise<businessEntities.BusinessOutput> {
    try {
      const sqlQuery = "SELECT * FROM business WHERE id = ?";
      const [rows] = await promisePool.query(sqlQuery, [id]);
      const result = Object.values(rows);

      const business: businessEntities.BusinessOutput = {
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

      return business;
    } catch (e) {
      throw e;
    }
  }

  public async list(): Promise<businessEntities.BusinessOutput[]> {
    try {
      const sqlQuery = "SELECT * FROM business";
      const [rows] = await promisePool.query(sqlQuery);
      const result = Object.values(rows);

      let business: businessEntities.BusinessOutput[] = [];

      for (let i = 0; i < result.length; i++) {
        let b: businessEntities.BusinessOutput = {
          id: result[i]["id"],
          name: result[i]["name"],
          shortDescription: result[i]["short_description"],
          longDescription: result[i]["long_description"],
          logo: result[i]["logo"],
          stockPrice: result[i]["stock_price"],
          urlWebsite: result[i]["url_website"],
          ceoName: result[i]["ceo_name"],
          numberOfEmployees: result[i]["number_of_employees"],
          country: result[i]["country"],
          nif: result[i]["nif"],
          fiscalName: result[i]["fiscal_name"],
          address: result[i]["address"],
          phone: result[i]["phone"],
          email: result[i]["email"],
          state: result[i]["state"],
          creationDate: result[i]["creation_date"],
          lastUpdate: result[i]["last_update"],
        };
        business.push(b);
      }

      return business;
    } catch (e) {
      throw e;
    }
  }

  public async create(businessInput: businessEntities.Business): Promise<businessEntities.BusinessOutput> {
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
        businessInput.id,
        businessInput.name,
        businessInput.shortDescription,
        businessInput.longDescription,
        businessInput.logo,
        businessInput.stockPrice,
        businessInput.urlWebsite,
        businessInput.ceoName,
        businessInput.numberOfEmployees,
        businessInput.country,
        businessInput.nif,
        businessInput.fiscalName,
        businessInput.address,
        businessInput.phone,
        businessInput.email,
        businessInput.state,
        businessInput.creationDate,
        businessInput.lastUpdate,
      ]);

      const businessOutput: businessEntities.BusinessOutput = await this.get(businessInput.id!);

      return businessOutput;
    } catch (e) {
      throw e;
    }
  }

  public async update(businessInput: businessEntities.Business): Promise<businessEntities.BusinessOutput> {
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
        "creation_date = ?," +
        "last_update = ?" +
        " WHERE id = ?";

      await promisePool.query(sqlQuery, [
        businessInput.name,
        businessInput.shortDescription,
        businessInput.longDescription,
        businessInput.logo,
        businessInput.stockPrice,
        businessInput.urlWebsite,
        businessInput.ceoName,
        businessInput.numberOfEmployees,
        businessInput.country,
        businessInput.nif,
        businessInput.fiscalName,
        businessInput.address,
        businessInput.phone,
        businessInput.email,
        businessInput.state,
        businessInput.creationDate,
        businessInput.lastUpdate,
        businessInput.id,
      ]);

      const businessOutput: businessEntities.BusinessOutput = await this.get(businessInput.id!);

      return businessOutput;
    } catch (e) {
      throw e;
    }
  }

  public async delete(id: string): Promise<businessEntities.BusinessOutput> {
    try {
      const businessOutput: businessEntities.BusinessOutput = await this.get(id);

      const sqlQuery = "DELETE FROM business WHERE id = ?";
      await promisePool.query(sqlQuery, [id]);

      return businessOutput;
    } catch (e) {
      throw e;
    }
  }
}
