import type { DefaultAccountsHttpInfra } from '@/features/accounts/infra/http/http-infra';
import type { AccountsAdapter } from '@/features/accounts/application/accounts-model';

export class DefaultAccountsHttpAdapter implements AccountsAdapter {
  constructor(private readonly accountsHttpInfra: DefaultAccountsHttpInfra) {}

  async listAccountsHttpResponse() {
    const accountHttpResponse = await this.accountsHttpInfra.listAccounts();

    return accountHttpResponse.items.map(item => ({
      id: item.id,
      name: item.name,
      type: item.name,
      kind: item.kind,
    }));
  }
}
