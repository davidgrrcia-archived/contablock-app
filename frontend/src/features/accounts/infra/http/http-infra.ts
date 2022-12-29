import type { AccountDescribeResponse, AccountListResponse } from '@/features/accounts/infra/http/models';

export interface AccountsHttpInfra {
  listAccounts(): Promise<AccountListResponse>;
  describeAccount(id: string): Promise<AccountDescribeResponse>;
}

export class DefaultAccountsHttpInfra implements AccountsHttpInfra {
  listAccounts() {
    return fetch('/api/accounts')
      .then(response => response.json())
      .then(response => response as AccountListResponse);
  }
  describeAccount(id: string) {
    return fetch('/api/accounts/' + id)
      .then(response => response.json())
      .then(response => response as AccountDescribeResponse);
  }
}
