type Entry = {
  id: string;
  accountId: string;
  doubleEntryCode: string;
  currency: string;
  concept: string;
  date: Date;
  quarter: number;
  debitBalance: number;
  creditBalance: number;
  totalBalance: number;
  state: string;
  creationDate: Date;
  lastUpdate: Date;
}


// STORAGE

export interface EntryStorageInterface {
  get(input: GetEntryStorageInput): Promise<GetEntryStorageOutput>;
  list(): Promise<ListEntryStorageOutput>;
  create(
    input: CreateEntryStorageInput
  ): Promise<CreateEntryStorageOutput>;
  update(
    input: UpdateEntryStorageInput
  ): Promise<UpdateEntryStorageOutput>;
  delete(id: DeleteEntryStorageInput): Promise<DeleteEntryStorageOutput>;
}

export type GetEntryStorageInput = {
  id: string;
};

export type GetEntryStorageOutput = Entry;

export type ListEntryStorageOutput = {
  items: GetEntryStorageOutput[];
};

export type CreateEntryStorageInput = Entry;
export type CreateEntryStorageOutput = GetEntryStorageOutput;

export type UpdateEntryStorageInput = {
  id: string;
  accountId: string;
  doubleEntryCode: string;
  currency: string;
  concept: string;
  date: Date;
  quarter: number;
  debitBalance: number;
  creditBalance: number;
  totalBalance: number;
  state: string;
  lastUpdate: Date;
};
export type UpdateEntryStorageOutput = GetEntryStorageOutput;

export type DeleteEntryStorageInput = {
  id: string;
};
export type DeleteEntryStorageOutput = GetEntryStorageOutput;

// USE CASE

export interface EntryUseCaseInterface {
  GetEntryUseCase(
    input: GetEntryUseCaseInput
  ): Promise<GetEntryUseCaseOutput>;
  ListEntryUseCase(): Promise<ListEntryUseCaseOutput>;
  CreateEntryUseCase(
    input: CreateEntryUseCaseInput
  ): Promise<CreateEntryUseCaseOutput>;
  UpdateEntryUseCase(
    input: UpdateEntryUseCaseInput
  ): Promise<UpdateEntryUseCaseOutput>;
  DeleteEntryUseCase(
    input: DeleteEntryUseCaseInput
  ): Promise<DeleteEntryUseCaseOutput>;
}

export type GetEntryUseCaseInput = {
  id: string;
};
export type GetEntryUseCaseOutput = Entry;

export type ListEntryUseCaseOutput = {
  items: GetEntryUseCaseOutput[];
};

export type CreateEntryUseCaseInput = {
  id: string;
  accountId: string;
  doubleEntryCode: string;
  currency: string;
  concept: string;
  date: Date;
  debitBalance: number;
  creditBalance: number;
  totalBalance: number;
};
export type CreateEntryUseCaseOutput = GetEntryUseCaseOutput;

export type UpdateEntryUseCaseInput = {
  id: string;
  accountId: string;
  doubleEntryCode?: string;
  currency?: string;
  concept?: string;
  date?: Date;
  debitBalance?: number;
  creditBalance?: number;
  totalBalance?: number;
};
export type UpdateEntryUseCaseOutput = GetEntryUseCaseOutput;

export type DeleteEntryUseCaseInput = {
  id: string;
};
export type DeleteEntryUseCaseOutput = GetEntryUseCaseOutput;

// CONTROLLER

export type GetEntryControllerOutput = Entry;

export type ListEntryControllerOutput = {
  items: GetEntryControllerOutput[];
};

export type CreateEntryControllerOutput = GetEntryControllerOutput;

export type UpdateEntryControllerOutput = GetEntryControllerOutput;

export type DeleteEntryControllerOutput = GetEntryControllerOutput;

// UTILS

export type DataError = {
  statusCode: number;
  message: string;
};
