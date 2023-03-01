import type { DefaultUsersHttpInfra } from '@/features/users/infra/http/http-infra';
import type { LoginUserInput, UsersAdapter } from '@/features/users/application/users-model';
import type { User as UserUseCase } from '@/features/users/application/users-model';

export class DefaultUsersHttpAdapter implements UsersAdapter {
  constructor(private readonly usersHttpInfra: DefaultUsersHttpInfra) {}

  async loginUserHttpResponse(input: LoginUserInput) {
    const userHttpResponse = await this.usersHttpInfra.loginUser(input);

    if (userHttpResponse instanceof Error) {
      return userHttpResponse;
    }

    const output: UserUseCase = {
      project: userHttpResponse.user.project,
      id: userHttpResponse.user.id,
      name: userHttpResponse.user.name,
      email: userHttpResponse.user.email,
      password: userHttpResponse.user.password,
      role: userHttpResponse.user.role,
      ethAddress: userHttpResponse.user.ethAddress,
      defaultPasswordChanged: userHttpResponse.user.defaultPasswordChanged,
      network: userHttpResponse.user.network,
      businessId: userHttpResponse.user.businessId,
      creationDate: userHttpResponse.user.creationDate,
      lastUpdate: userHttpResponse.user.lastUpdate,
    };

    return output;
  }
}
