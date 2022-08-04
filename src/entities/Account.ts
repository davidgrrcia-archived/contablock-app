export type GetAccountControllerInput = {
  id: string;
};

export type GetAccountControllerOutput = {
  id: string;
  name: string;
  type: string;
  kind: string;
};

export type ListAccountsControllerOutput = {
  items: GetAccountControllerOutput[];
};

export type DataError = {
  statusCode: number;
  message: string;
};
