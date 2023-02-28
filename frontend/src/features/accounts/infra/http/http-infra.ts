import type { AccountDescribeResponse, AccountListResponse } from '@/features/accounts/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface AccountsHttpInfra {
  listAccounts(): Promise<AccountListResponse | Error>;
  describeAccount(id: string): Promise<AccountDescribeResponse | Error>;
}

export class DefaultAccountsHttpInfra implements AccountsHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async listAccounts() {
    const configStore = useConfigStore();

    const response = await fetch('/api/projects/' + configStore.state.project + '/accounts');

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as AccountListResponse;
  }
  async describeAccount(id: string) {
    const configStore = useConfigStore();

    const response = await fetch('/api/projects/' + configStore.state.project + '/accounts/' + id);

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as AccountDescribeResponse;
  }
}
