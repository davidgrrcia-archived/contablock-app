import type { DefaultAccountsHttpInfra } from '@/features/accounts/infra/http/http-infra';
import type { Account, AccountsAdapter } from '@/features/accounts/application/accounts-model';

export class DefaultAccountsHttpAdapter implements AccountsAdapter {
  constructor(private readonly accountsHttpInfra: DefaultAccountsHttpInfra) {}

  async listAccountsHttpResponse() {
    const accountHttpResponse = await this.accountsHttpInfra.listAccounts();

    return accountHttpResponse.items.map(item => ({
      id: item.id,
      name: item.name,
      level: item.level,
      type: item.type,
    }));
  }

  async describeAccountHttpResponse(id: string) {
    const input = await this.accountsHttpInfra.describeAccount(id);

    const output: Account = {
      id: input.id,
      name: input.name,
      level: input.level,
      type: input.type,
    };

    return output;
  }
}
