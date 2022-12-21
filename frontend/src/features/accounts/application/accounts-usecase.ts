import type { Account, AccountsAdapter } from '@/features/accounts/application/accounts-model';

export interface AccountsUseCase {
  listAccounts(): Promise<Account[]>;
}

export class DefaultAccountsUseCase implements AccountsUseCase {
  constructor(private readonly accountsAdapter: AccountsAdapter) {}

  async listAccounts() {
    return await this.accountsAdapter.listAccountsHttpResponse();
  }
}
