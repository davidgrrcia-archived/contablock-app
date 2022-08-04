import BusinessRepository from "../repositories/business.repository";
import * as businessEntities from "../entities/Business";

export const GetBusinessUseCase = (businessRepository: BusinessRepository) => async (businessId: string) => {
  const business: businessEntities.BusinessOutput = await businessRepository.get(businessId);
  return business;
};

export const ListBusinessUseCase = (businessRepository: BusinessRepository) => async () => {
  const business: businessEntities.BusinessOutput[] = await businessRepository.list();
  return business;
};

export const CreateBusinessUseCase = (businessRepository: BusinessRepository) => async (createBusinessControllerInput: businessEntities.CreateBusinessControllerInput) => {
  
  // TODO add validators and check if the business already exists
  
  if (createBusinessControllerInput.id === "" || createBusinessControllerInput.id === undefined) {
    createBusinessControllerInput.id = uid();
  }
  createBusinessControllerInput.state = "CREATED";
  createBusinessControllerInput.creationDate = new Date();
  createBusinessControllerInput.lastUpdate =  createBusinessControllerInput.creationDate
  const business: businessEntities.BusinessOutput = await businessRepository.create(createBusinessControllerInput);
  
  return business;
};

export const UpdateBusinessUseCase = (businessRepository: BusinessRepository) => async (businessInput: businessEntities.Business ) => {
  
// TODO improve types and mapers

  const business: businessEntities.BusinessOutput = await businessRepository.get(businessInput.id!);
  
  businessInput.state = "UPDATED";
  businessInput.creationDate = business.creationDate;
  businessInput.lastUpdate = new Date();
  const businessOutput: businessEntities.BusinessOutput = await businessRepository.update(businessInput);
  return businessOutput;
};

export const DeleteBusinessUseCase = (businessRepository: BusinessRepository) => async (businessId: string) => {
  const business: businessEntities.BusinessOutput = await businessRepository.delete(businessId);
  return business;
};

const uid = () => {
  return String.fromCharCode(Math.floor(Math.random() * 26) + 97) + Math.random().toString(16).slice(2) + Date.now().toString(16).slice(5);
};

// TODO balance or stadistics endpoint fot getting: 
// debit_balance 
// credit_balance 
// total_balance 
// total_entries 
// total_transactions 