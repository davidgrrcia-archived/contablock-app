export interface AccountsAdapter {
  listAccountsHttpResponse(): Promise<Node[]>;
  describeAccountHttpResponse(id: string): Promise<Account>;
}

export interface Account {
  id: string;
  level: string;
  name: string;
  type: string;
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
