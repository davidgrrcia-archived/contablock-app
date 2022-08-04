import * as entities from "../entities/Account";

export default interface AccountRepository {
  get(input: entities.GetAccountStorageInput): Promise<entities.GetAccountStorageOutput>;
  list(): Promise<entities.ListAccountsStorageOutput>;
}
