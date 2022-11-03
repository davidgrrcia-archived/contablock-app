
import UserPostgres from "../dataSources/user.datasource";
import EntryPostgres from "../dataSources/entry.datasource";
import TransactionPostgres from "../dataSources/transaction.datasource";
import * as user from "./user.usecase";
import * as entry from "./entry.usecase";
import * as transaction from "./transaction.usecase";

// Repositories

const userRepository = new UserPostgres();
const entryRepository = new EntryPostgres();
const transactionRepository = new TransactionPostgres();

export const getUser = user.GetUserUseCase(userRepository);
export const listUser = user.ListUserUseCase(userRepository);
export const createUser = user.CreateUserUseCase(userRepository);
export const updateUser = user.UpdateUserUseCase(userRepository);
export const deleteUser = user.DeleteUserUseCase(userRepository);

export const getEntry = entry.GetEntryUseCase(entryRepository);
export const listEntry = entry.ListEntryUseCase(entryRepository);
export const createEntry = entry.CreateEntryUseCase(entryRepository);
export const updateEntry = entry.UpdateEntryUseCase(entryRepository);
export const deleteEntry = entry.DeleteEntryUseCase(entryRepository);

export const getTransaction = transaction.GetTransactionUseCase(
  transactionRepository
);
export const listTransaction = transaction.ListTransactionUseCase(
  transactionRepository
);
export const createTransaction = transaction.CreateTransactionUseCase(
  transactionRepository
);
export const updateTransaction = transaction.UpdateTransactionUseCase(
  transactionRepository
);
export const deleteTransaction = transaction.DeleteTransactionUseCase(
  transactionRepository
);
