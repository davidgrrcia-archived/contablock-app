type Entry = {
  id: string;
  businessId: string;
  currency: string;
  description: string;
  participant: string;
  fileUpload: string;
  date: Date;
  type: string;
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
  businessId: string;
};

export type GetEntryStorageOutput = Entry;

export type ListEntryStorageInput = {
  businessId: string;
};
export type ListEntryStorageOutput = {
  items: GetEntryStorageOutput[];
};

export type CreateEntryStorageInput = Entry;
export type CreateEntryStorageOutput = GetEntryStorageOutput;

export type UpdateEntryStorageInput = {
  id: string;
  businessId: string;
  currency: string;
  description: string;
  participant: string;
  fileUpload: string;
  date: Date;
  type: string;
  debitBalance: number;
  creditBalance: number;
  totalBalance: number;
  state: string;
  lastUpdate: Date;
};
export type UpdateEntryStorageOutput = GetEntryStorageOutput;

export type DeleteEntryStorageInput = {
  id: string;
  businessId: string;
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
  businessId: string;
};
export type GetEntryUseCaseOutput = Entry;

export type ListEntryUseCaseInput = {
  businessId: string;
};
export type ListEntryUseCaseOutput = {
  items: GetEntryUseCaseOutput[];
};

export type CreateEntryUseCaseInput = {
  id: string;
  businessId: string;
  currency: string;
  description: string;
  participant: string;
  fileUpload: string;
  date: Date;
  type: string;
  debitBalance: number;
  creditBalance: number;
  totalBalance: number;
};
export type CreateEntryUseCaseOutput = GetEntryUseCaseOutput;

export type UpdateEntryUseCaseInput = {
  id: string;
  businessId: string;
  currency?: string;
  description?: string;
  participant?: string;
  fileUpload?: string;
  date?: Date;
  type?: string;
  debitBalance?: number;
  creditBalance?: number;
  totalBalance?: number;
};
export type UpdateEntryUseCaseOutput = GetEntryUseCaseOutput;

export type DeleteEntryUseCaseInput = {
  id: string;
  businessId: string;
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