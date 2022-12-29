export interface Account {
  id: string;
  name: string;
  type: string;
  kind: string;
}

export interface AccountListResponse {
  items: Account[];
}

export interface AccountDescribeResponse {
  account: Account;
}
