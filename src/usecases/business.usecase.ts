import { BusinessUseCaseInterface } from "../entities/Business";
import * as entities from "../entities/Business";
import BusinessStorage from "../dataSources/business.datasource";
import { v4 as uuidv4 } from "uuid";

const businessRepository = new BusinessStorage();

export default class BusinessUseCase implements BusinessUseCaseInterface {
  public async GetBusinessUseCase(input: entities.GetBusinessUseCaseInput) {
    if (input.id === "") {
      throw new Error("id is required");
    }

    const getBusinessStorageInput: entities.GetBusinessStorageInput = {
      id: input.id,
    };

    const storedBusiness = await businessRepository.get(
      getBusinessStorageInput
    );

    const getBusinessUseCaseOutput: entities.GetBusinessUseCaseOutput = {
      id: storedBusiness.id,
      name: storedBusiness.name,
      shortDescription: storedBusiness.shortDescription,
      longDescription: storedBusiness.longDescription,
      logo: storedBusiness.logo,
      stockPrice: storedBusiness.stockPrice,
      urlWebsite: storedBusiness.urlWebsite,
      ceoName: storedBusiness.ceoName,
      numberOfEmployees: storedBusiness.numberOfEmployees,
      country: storedBusiness.country,
      nif: storedBusiness.nif,
      fiscalName: storedBusiness.fiscalName,
      address: storedBusiness.address,
      phone: storedBusiness.phone,
      email: storedBusiness.email,
      state: storedBusiness.state,
      creationDate: storedBusiness.creationDate,
      lastUpdate: storedBusiness.lastUpdate,
    };

    return getBusinessUseCaseOutput;
  }

  public async ListBusinessUseCase() {
    const storedBusiness = await businessRepository.list();

    const listBusinessUseCaseOutput: entities.ListBusinessUseCaseOutput = {
      items: storedBusiness.items.map((item) => ({
        id: item.id,
        name: item.name,
        shortDescription: item.shortDescription,
        longDescription: item.longDescription,
        logo: item.logo,
        stockPrice: item.stockPrice,
        urlWebsite: item.urlWebsite,
        ceoName: item.ceoName,
        numberOfEmployees: item.numberOfEmployees,
        country: item.country,
        nif: item.nif,
        fiscalName: item.fiscalName,
        address: item.address,
        phone: item.phone,
        email: item.email,
        state: item.state,
        creationDate: item.creationDate,
        lastUpdate: item.lastUpdate,
      })),
    };

    return listBusinessUseCaseOutput;
  }

  public async CreateBusinessUseCase(
    input: entities.CreateBusinessUseCaseInput
  ) {
    if (input.id === undefined || input.id === "") {
      input.id = uuidv4();
    }
    if (input.name === undefined || input.name === "") {
      throw new Error("Name is required");
    }
    if (input.shortDescription === undefined || input.shortDescription === "") {
      throw new Error("Short description is required");
    }
    if (input.longDescription === undefined || input.longDescription === "") {
      throw new Error("Long description is required");
    }
    if (input.logo === undefined || input.logo === "") {
      throw new Error("Logo is required");
    }
    if (input.nif === undefined || input.nif === "") {
      throw new Error("NIF is required");
    }
    if (input.fiscalName === undefined || input.fiscalName === "") {
      throw new Error("Fiscal name is required");
    }
    if (input.address === undefined || input.address === "") {
      throw new Error("Address is required");
    }
    if (input.phone === undefined || input.phone === "") {
      throw new Error("Phone is required");
    }
    if (input.email === undefined || input.email === "") {
      throw new Error("Email is required");
    }

    const state = "CREATED";
    const creationDate = new Date();
    const lastUpdate = creationDate;

    const createBusinessStorageInput: entities.CreateBusinessStorageInput = {
      id: input.id,
      name: input.name,
      shortDescription: input.shortDescription,
      longDescription: input.longDescription,
      logo: input.logo,
      stockPrice: input.stockPrice,
      urlWebsite: input.urlWebsite,
      ceoName: input.ceoName,
      numberOfEmployees: input.numberOfEmployees,
      country: input.country,
      nif: input.nif,
      fiscalName: input.fiscalName,
      address: input.address,
      phone: input.phone,
      email: input.email,
      state: state,
      creationDate: creationDate,
      lastUpdate: lastUpdate,
    };

    const createdBusiness = await businessRepository.create(
      createBusinessStorageInput
    );

    const createBusinessUseCaseOutput: entities.CreateBusinessUseCaseOutput = {
      id: createdBusiness.id,
      name: createdBusiness.name,
      shortDescription: createdBusiness.shortDescription,
      longDescription: createdBusiness.longDescription,
      logo: createdBusiness.logo,
      stockPrice: createdBusiness.stockPrice,
      urlWebsite: createdBusiness.urlWebsite,
      ceoName: createdBusiness.ceoName,
      numberOfEmployees: createdBusiness.numberOfEmployees,
      country: createdBusiness.country,
      nif: createdBusiness.nif,
      fiscalName: createdBusiness.fiscalName,
      address: createdBusiness.address,
      phone: createdBusiness.phone,
      email: createdBusiness.email,
      state: createdBusiness.state,
      creationDate: createdBusiness.creationDate,
      lastUpdate: createdBusiness.lastUpdate,
    };

    return createBusinessUseCaseOutput;
  }

  public async UpdateBusinessUseCase(
    input: entities.UpdateBusinessUseCaseInput
  ) {
    if (input.id === "") {
      throw new Error("id is required");
    }

    const getBusinessUseCaseInput: entities.GetBusinessUseCaseInput = {
      id: input.id,
    };
    const storedBusiness = await this.GetBusinessUseCase(
      getBusinessUseCaseInput
    );

    const name = checkNotEmptyUpdate(input.name, storedBusiness.name);
    const shortDescription = checkNotEmptyUpdate(input.shortDescription, storedBusiness.shortDescription);
    const longDescription = checkNotEmptyUpdate(input.longDescription, storedBusiness.longDescription);
    const logo = checkNotEmptyUpdate(input.logo, storedBusiness.logo);
    const stockPrice = checkNotEmptyUpdate(input.stockPrice, storedBusiness.stockPrice);
    const urlWebsite = checkNotEmptyUpdate(input.urlWebsite, storedBusiness.urlWebsite);
    const ceoName = checkNotEmptyUpdate(input.ceoName, storedBusiness.ceoName);
    const numberOfEmployees = checkNotEmptyUpdate(input.numberOfEmployees, storedBusiness.numberOfEmployees);
    const country = checkNotEmptyUpdate(input.country, storedBusiness.country);
    const nif = checkNotEmptyUpdate(input.nif, storedBusiness.nif);
    const fiscalName = checkNotEmptyUpdate(input.fiscalName, storedBusiness.fiscalName);
    const address = checkNotEmptyUpdate(input.address, storedBusiness.address);
    const phone = checkNotEmptyUpdate(input.phone, storedBusiness.phone);
    const email = checkNotEmptyUpdate(input.email, storedBusiness.email);
    const state = "UPDATED";
    const lastUpdate = new Date();

    const updateBusinessStorageInput: entities.UpdateBusinessStorageInput = {
      id: input.id,
      name: name,
      shortDescription: shortDescription,
      longDescription: longDescription,
      logo: logo,
      stockPrice: stockPrice,
      urlWebsite: urlWebsite,
      ceoName: ceoName,
      numberOfEmployees: numberOfEmployees,
      country: country,
      nif: nif,
      fiscalName: fiscalName,
      address: address,
      phone: phone,
      email: email,
      state: state,
      lastUpdate: lastUpdate,
    };

    const updatedBusiness = await businessRepository.update(
      updateBusinessStorageInput
    );

    const updateBusinessUseCaseOutput: entities.UpdateBusinessUseCaseOutput = {
      id: updatedBusiness.id,
      name: updatedBusiness.name,
      shortDescription: updatedBusiness.shortDescription,
      longDescription: updatedBusiness.longDescription,
      logo: updatedBusiness.logo,
      stockPrice: updatedBusiness.stockPrice,
      urlWebsite: updatedBusiness.urlWebsite,
      ceoName: updatedBusiness.ceoName,
      numberOfEmployees: updatedBusiness.numberOfEmployees,
      country: updatedBusiness.country,
      nif: updatedBusiness.nif,
      fiscalName: updatedBusiness.fiscalName,
      address: updatedBusiness.address,
      phone: updatedBusiness.phone,
      email: updatedBusiness.email,
      state: updatedBusiness.state,
      creationDate: updatedBusiness.creationDate,
      lastUpdate: updatedBusiness.lastUpdate,
    };

    return updateBusinessUseCaseOutput;
  }

  public async DeleteBusinessUseCase(
    input: entities.DeleteBusinessUseCaseInput
  ) {
    if (input.id === "") {
      throw new Error("id is required");
    }

    const deleteBusinessStorageInput: entities.DeleteBusinessStorageInput = {
      id: input.id,
    };

    const storedBusiness = await businessRepository.delete(
      deleteBusinessStorageInput
    );

    const deleteBusinessUseCaseOutput: entities.DeleteBusinessUseCaseOutput = {
      id: storedBusiness.id,
      name: storedBusiness.name,
      shortDescription: storedBusiness.shortDescription,
      longDescription: storedBusiness.longDescription,
      logo: storedBusiness.logo,
      stockPrice: storedBusiness.stockPrice,
      urlWebsite: storedBusiness.urlWebsite,
      ceoName: storedBusiness.ceoName,
      numberOfEmployees: storedBusiness.numberOfEmployees,
      country: storedBusiness.country,
      nif: storedBusiness.nif,
      fiscalName: storedBusiness.fiscalName,
      address: storedBusiness.address,
      phone: storedBusiness.phone,
      email: storedBusiness.email,
      state: storedBusiness.state,
      creationDate: storedBusiness.creationDate,
      lastUpdate: storedBusiness.lastUpdate,
    };

    return deleteBusinessUseCaseOutput;
  }
}

function checkNotEmptyUpdate(input: any, stored: any) {
  if (input === "" || input === undefined) {
    return stored;
  }
  return input;
}
