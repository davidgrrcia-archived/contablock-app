import AccountRepository from "../repositories/account.repository";
import Account from "../entities/Account";

export const GetAccountUseCase = (accountRepository: AccountRepository) => async (id: string) => {
  const account: Account = await accountRepository.get(id);
  return account;
};

export const ListAccountsUseCase = (accountRepository: AccountRepository) => async () => {
  const accounts: Account[] = await accountRepository.list();
  return accounts;
};

export const ListAccountsByTypeUseCase = (accountRepository: AccountRepository) => async (type: any) => {
  const accounts: Account[] = await accountRepository.listByType(type);
  return accounts;
};
