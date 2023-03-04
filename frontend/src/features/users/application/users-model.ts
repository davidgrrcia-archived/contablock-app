export interface UsersAdapter {
  loginUserHttpResponse(input: LoginUserInput): Promise<User | Error>;
}

export interface User {
  project: string;
  id: string;
  contractId: string;
  name: string;
  email: string;
  password: string;
  role: string;
  ethAddress: string;
  defaultPasswordChanged: boolean;
  network: string;
  businessId: string;
  creationDate: number;
  lastUpdate: number;
}

export interface LoginUserInput {
  email: string;
  password: string;
}
