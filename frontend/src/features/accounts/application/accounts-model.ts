export interface AccountsAdapter {
  listAccountsHttpResponse(): Promise<Account[]>;
}

export interface Account {
  id: string;
  name: string;
  type: string;
  kind: string;
}
