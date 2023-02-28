export interface AccountsAdapter {
  listAccountsHttpResponse(): Promise<Node[] | Error>;
  describeAccountHttpResponse(id: string): Promise<Account | Error>;
}

export interface Account {
  id: string;
  level: string;
  name: string;
  balanceId: string;
  displayName: string;
}

export interface Node {
  key: string;
  label: string;
  data?: string;
  icon?: string;
  selectable?: boolean;
  children?: Node[];
}
