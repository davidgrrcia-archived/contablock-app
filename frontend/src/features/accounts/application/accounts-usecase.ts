import type { Node, Account, AccountsAdapter } from '@/features/accounts/application/accounts-model';

export interface AccountsUseCase {
  listAccounts(): Promise<Node[] | Error>;
  describeAccount(id: string): Promise<Account | Error>;
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
