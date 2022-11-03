// STORAGE

export interface AccountStorageInterface {
  get(input: GetAccountStorageInput): Promise<GetAccountStorageOutput>;
  list(): Promise<ListAccountsStorageOutput>;
}

export type GetAccountStorageInput = {
  id: string;
};

export type GetAccountStorageOutput = {
  id: string;
  name: string;
  group: number;
  category: string;
};

export type ListAccountsStorageOutput = {
  items: GetAccountStorageOutput[];
};

// USE CASE

export interface AccountUseCaseInterface {
  GetAccountUseCase(input: GetAccountUseCaseInput): Promise<GetAccountUseCaseOutput>;
  ListAccountsUseCase(): Promise<ListAccountsUseCaseOutput>;
}

export type GetAccountUseCaseInput = {
  id: string;
};

export type GetAccountUseCaseOutput = {
  id: string;
  name: string;
  group: number;
  category: string;
};

export type ListAccountsUseCaseOutput = {
  items: GetAccountUseCaseOutput[];
};

// CONTROLLER

export type GetAccountControllerOutput = {
  id: string;
  name: string;
  group: number;
  category: string;
};

export type ListAccountsControllerOutput = {
  items: GetAccountControllerOutput[];
};

// UTILS

export type DataError = {
  statusCode: number;
  message: string;
};
