import { AccountUseCaseInterface } from "../entities/Account";
import * as entities from "../entities/Account";
import AccountStorage from "../dataSources/accounts.datasource";

const accountRepository = new AccountStorage();

export default class AccountUseCase implements AccountUseCaseInterface {
  public async GetAccountUseCase(input: entities.GetAccountUseCaseInput) {
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
  }

  public async ListAccountsUseCase() {
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
  }
}
