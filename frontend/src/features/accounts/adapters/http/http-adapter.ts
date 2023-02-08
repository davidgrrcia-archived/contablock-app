import type { DefaultAccountsHttpInfra } from '@/features/accounts/infra/http/http-infra';
import type { Account, AccountsAdapter } from '@/features/accounts/application/accounts-model';

export class DefaultAccountsHttpAdapter implements AccountsAdapter {
  constructor(private readonly accountsHttpInfra: DefaultAccountsHttpInfra) {}

  async listAccountsHttpResponse() {
    const accountHttpResponse = await this.accountsHttpInfra.listAccounts();

    return accountHttpResponse.items.map(item => ({
      key: item.id,
      label: item.displayName,
      data: item.id,
      selectable: false,
      children: item.subgroups.map(subgroup => ({
        key: subgroup.id,
        label: subgroup.displayName,
        data: item.id,
        selectable: false,
        children: subgroup.accounts.map(account => ({
          key: account.id,
          label: account.displayName,
          data: item.id,
          selectable: true,
        })),
      })),
    }));
  }

  async describeAccountHttpResponse(id: string) {
    const input = await this.accountsHttpInfra.describeAccount(id);

    const output: Account = {
      id: input.id,
      name: input.name,
      level: input.level,
      type: input.type,
      displayName: input.displayName,
    };

    return output;
  }
}
