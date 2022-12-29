import type { Account, AccountsAdapter } from '@/features/accounts/application/accounts-model';

export interface AccountsUseCase {
  listAccounts(): Promise<Account[]>;
  describeAccount(id: string): Promise<Account>;
}

export class DefaultAccountsUseCase implements AccountsUseCase {
  constructor(private readonly accountsAdapter: AccountsAdapter) {}

  async listAccounts() {
    return await this.accountsAdapter.listAccountsHttpResponse();
  }

  async describeAccount(id: string) {
    return await this.accountsAdapter.describeAccountHttpResponse(id);
  }
}
