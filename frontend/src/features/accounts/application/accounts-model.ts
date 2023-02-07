export interface AccountsAdapter {
  listAccountsHttpResponse(): Promise<Account[]>;
  describeAccountHttpResponse(id: string): Promise<Account>;
}

export interface Account {
  id: string;
  level: string;
  name: string;
  type: string;
  displayName: string;
}
