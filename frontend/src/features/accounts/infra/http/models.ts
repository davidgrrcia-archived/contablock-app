export interface Account {
  id: string;
  level: string;
  name: string;

  type: string;
  displayName: string;
}

export interface AccountListResponse {
  items: Account[];
}

export interface AccountDescribeResponse {
  id: string;
  level: string;
  name: string;

  type: string;
  displayName: string
}
