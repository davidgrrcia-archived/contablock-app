import Account from "../entities/Account";

export default interface AccountRepository {
  get(id: string): Promise<Account>;
  list(): Promise<Account[]>;
  listByType(type: string): Promise<Account[]>;
}
