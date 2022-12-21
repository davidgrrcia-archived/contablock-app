import type { AccountListResponse } from '@/features/accounts/infra/http/models';

export interface AccountsHttpInfra {
  listAccounts(): Promise<AccountListResponse>;
}

export class DefaultAccountsHttpInfra implements AccountsHttpInfra {
  listAccounts() {
    return fetch('http://localhost:4000/api/accounts')
      .then(response => response.json())
      .then(response => response as AccountListResponse);
  }
}
