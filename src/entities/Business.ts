type Business = {
  id: string;
  name: string;
  shortDescription: string;
  longDescription: string;
  logo: string;
  stockPrice: number;
  urlWebsite: string;
  ceoName: string;
  numberOfEmployees: number;
  country: string;
  nif: string;
  fiscalName: string;
  address: string;
  phone: string;
  email: string;
  state: string;
  creationDate: Date;
  lastUpdate: Date;
};

// STORAGE

export interface BusinessStorageInterface {
  get(input: GetBusinessStorageInput): Promise<GetBusinessStorageOutput>;
  list(): Promise<ListBusinessStorageOutput>;
  create(
    input: CreateBusinessStorageInput
  ): Promise<CreateBusinessStorageOutput>;
  update(
    input: UpdateBusinessStorageInput
  ): Promise<UpdateBusinessStorageOutput>;
  delete(id: DeleteBusinessStorageInput): Promise<DeleteBusinessStorageOutput>;
}

export type GetBusinessStorageInput = {
  id: string;
};

export type GetBusinessStorageOutput = Business;

export type ListBusinessStorageOutput = {
  items: GetBusinessStorageOutput[];
};

export type CreateBusinessStorageInput = Business;
export type CreateBusinessStorageOutput = GetBusinessStorageOutput;

export type UpdateBusinessStorageInput = {
  id: string;
  name: string;
  shortDescription: string;
  longDescription: string;
  logo: string;
  stockPrice: number;
  urlWebsite: string;
  ceoName: string;
  numberOfEmployees: number;
  country: string;
  nif: string;
  fiscalName: string;
  address: string;
  phone: string;
  email: string;
  state: string;
  lastUpdate: Date;
};
export type UpdateBusinessStorageOutput = GetBusinessStorageOutput;

export type DeleteBusinessStorageInput = {
  id: string;
};
export type DeleteBusinessStorageOutput = GetBusinessStorageOutput;

// USE CASE

export interface BusinessUseCaseInterface {
  GetBusinessUseCase(
    input: GetBusinessUseCaseInput
  ): Promise<GetBusinessUseCaseOutput>;
  ListBusinessUseCase(): Promise<ListBusinessUseCaseOutput>;
  CreateBusinessUseCase(
    input: CreateBusinessUseCaseInput
  ): Promise<CreateBusinessUseCaseOutput>;
  UpdateBusinessUseCase(
    input: UpdateBusinessUseCaseInput
  ): Promise<UpdateBusinessUseCaseOutput>;
  DeleteBusinessUseCase(
    input: DeleteBusinessUseCaseInput
  ): Promise<DeleteBusinessUseCaseOutput>;
}

export type GetBusinessUseCaseInput = {
  id: string;
};

export type GetBusinessUseCaseOutput = Business;

export type ListBusinessUseCaseOutput = {
  items: GetBusinessUseCaseOutput[];
};

export type CreateBusinessUseCaseInput = {
  id: string;
  name: string;
  shortDescription: string;
  longDescription: string;
  logo: string;
  stockPrice: number;
  urlWebsite: string;
  ceoName: string;
  numberOfEmployees: number;
  country: string;
  nif: string;
  fiscalName: string;
  address: string;
  phone: string;
  email: string;
};
export type CreateBusinessUseCaseOutput = GetBusinessUseCaseOutput;

export type UpdateBusinessUseCaseInput = {
  id: string;
  name?: string;
  shortDescription?: string;
  longDescription?: string;
  logo?: string;
  stockPrice?: number;
  urlWebsite?: string;
  ceoName?: string;
  numberOfEmployees?: number;
  country?: string;
  nif?: string;
  fiscalName?: string;
  address?: string;
  phone?: string;
  email?: string;
  state?: string;
  lastUpdate?: Date;
};
export type UpdateBusinessUseCaseOutput = GetBusinessUseCaseOutput;

export type DeleteBusinessUseCaseInput = {
  id: string;
};
export type DeleteBusinessUseCaseOutput = GetBusinessUseCaseOutput;

// CONTROLLER

export type GetBusinessControllerOutput = Business;

export type ListBusinessControllerOutput = {
  items: GetBusinessControllerOutput[];
};

export type CreateBusinessControllerOutput = GetBusinessControllerOutput;

export type UpdateBusinessControllerOutput = GetBusinessControllerOutput;

export type DeleteBusinessControllerOutput = GetBusinessControllerOutput;

// UTILS

export type DataError = {
  statusCode: number;
  message: string;
};
