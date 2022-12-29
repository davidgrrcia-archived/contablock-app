export interface AccountsAdapter {
  listAccountsHttpResponse(): Promise<Account[]>;
  describeAccountHttpResponse(id: string): Promise<Account>;
}

export interface Account {
  id: string;
  name: string;
  type: string;
  kind: string;
}
