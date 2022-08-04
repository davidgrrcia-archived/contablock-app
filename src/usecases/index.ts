import AccountPostgres from "../dataSources/accounts.datasource";
import BusinessPostgres from "../dataSources/business.datasource";
import UserPostgres from "../dataSources/user.datasource";
import EntryPostgres from "../dataSources/entry.datasource";
import TransactionPostgres from "../dataSources/transaction.datasource";
import * as account from "./account.usecase";
import * as business from "./business.usecase";
import * as user from "./user.usecase";
import * as entry from "./entry.usecase";
import * as transaction from "./transaction.usecase";

// Repositories
const accountRepository = new AccountPostgres();
const businessRepository = new BusinessPostgres();
const userRepository = new UserPostgres();
const entryRepository = new EntryPostgres();
const transactionRepository = new TransactionPostgres();

// UseCases
export const getAccount = account.GetAccountUseCase(accountRepository);
export const listAccountsUseCase =
  account.ListAccountsUseCase(accountRepository);

export const getBusiness = business.GetBusinessUseCase(businessRepository);
export const listBusiness = business.ListBusinessUseCase(businessRepository);
export const createBusiness =
  business.CreateBusinessUseCase(businessRepository);
export const updateBusiness =
  business.UpdateBusinessUseCase(businessRepository);
export const deleteBusiness =
  business.DeleteBusinessUseCase(businessRepository);

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
