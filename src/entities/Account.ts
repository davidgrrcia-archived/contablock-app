// CONTROLLER

export type GetAccountControllerOutput = {
  id: string;
  name: string;
  type: string;
  kind: string;
};

export type ListAccountsControllerOutput = {
  items: GetAccountControllerOutput[];
};

// USE CASE

export type GetAccountUseCaseInput = {
  id: string;
};

export type GetAccountUseCaseOutput = {
  [x: string]: any;
  id: string;
  name: string;
  type: string;
  kind: string;
};

export type ListAccountsUseCaseOutput = {
  items: GetAccountUseCaseOutput[];
};

// STORAGE

export type GetAccountStorageInput = {
  id: string;
};

export type GetAccountStorageOutput = {
  id: string;
  name: string;
  type: string;
  kind: string;
}

export type ListAccountsStorageOutput = {
  items: GetAccountStorageOutput[];
};

// UTILS

export type DataError = {
  statusCode: number;
  message: string;
};
