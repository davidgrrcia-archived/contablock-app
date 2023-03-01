import type { LoginUserInput, User, UsersAdapter } from '@/features/users/application/users-model';

export interface UsersUseCase {
  loginUser(input: LoginUserInput): Promise<User | Error>;
}

export class DefaultUsersUseCase implements UsersUseCase {
  constructor(private readonly usersAdapter: UsersAdapter) {}

  async loginUser(input: LoginUserInput) {
    return await this.usersAdapter.loginUserHttpResponse(input);
  }
}
