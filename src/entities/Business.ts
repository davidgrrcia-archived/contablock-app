type Business = {
  id: string;
  name: string;
  socialDenomination: string;
  cif: string;
  logo: string;
};

// STORAGE

export interface BusinessStorageInterface {
  get(input: GetBusinessStorageInput): Promise<GetBusinessStorageOutput>;
  list(): Promise<ListBusinessStorageOutput>;
}

export type GetBusinessStorageInput = {
  id: string;
};

export type GetBusinessStorageOutput = Business;

export type ListBusinessStorageOutput = {
  items: GetBusinessStorageOutput[];
};

// USE CASE

export interface BusinessUseCaseInterface {
  GetBusinessUseCase(
    input: GetBusinessUseCaseInput
  ): Promise<GetBusinessUseCaseOutput>;
  ListBusinessUseCase(): Promise<ListBusinessUseCaseOutput>;
}

export type GetBusinessUseCaseInput = {
  id: string;
};

export type GetBusinessUseCaseOutput = Business;

export type ListBusinessUseCaseOutput = {
  items: GetBusinessUseCaseOutput[];
};

// CONTROLLER

export type GetBusinessControllerOutput = Business;

export type ListBusinessControllerOutput = {
  items: GetBusinessControllerOutput[];
};

// UTILS

export type DataError = {
  statusCode: number;
  message: string;
};
