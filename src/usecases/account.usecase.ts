import AccountRepository from "../repositories/account.repository";
import * as entities from "../entities/Account";

//TODO: define functions output type

export function GetAccountUseCase(accountRepository: AccountRepository) {
  return async (input: entities.GetAccountUseCaseInput) => {
    if (input.id === "") {
      throw new Error("id is required");
    }

    const getAccountStorageInput: entities.GetAccountStorageInput = {
      id: input.id,
    };

    const storedAccount: entities.GetAccountStorageOutput =
      await accountRepository.get(getAccountStorageInput);

    const getAccountUseCaseOutput: entities.GetAccountUseCaseOutput = {
      id: storedAccount.id,
      name: storedAccount.name,
      type: storedAccount.type,
      kind: storedAccount.kind,
    };

    return getAccountUseCaseOutput;
  };
}

export function ListAccountsUseCase(accountRepository: AccountRepository) {
  return async () => {
    const storedAccounts: entities.ListAccountsStorageOutput =
      await accountRepository.list();

    const listAccountsUseCaseOutput: entities.ListAccountsUseCaseOutput = {
      items: storedAccounts.items.map((item) => ({
        id: item.id,
        name: item.name,
        type: item.type,
        kind: item.kind,
      })),
    };

    return listAccountsUseCaseOutput;
  };
}
