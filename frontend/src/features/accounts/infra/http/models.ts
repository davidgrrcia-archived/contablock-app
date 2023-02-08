export interface Account {
  id: string;
  level: string;
  name: string;

  type: string;
  displayName: string;
}

export interface Group {
  id: string;
  level: string;
  name: string;

  type: string;
  displayName: string;
  subgroups: Subgroup[];
}

export interface Subgroup {
  id: string;
  level: string;
  name: string;

  type: string;
  displayName: string;
  accounts: Account[];
}

export interface AccountListResponse {
  items: Group[];
}

export interface AccountDescribeResponse {
  id: string;
  level: string;
  name: string;

  type: string;
  displayName: string;
}
