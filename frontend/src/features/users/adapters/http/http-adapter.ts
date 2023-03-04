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
      project: userHttpResponse.project,
      id: userHttpResponse.id,
      contractId: userHttpResponse.contractId,
      name: userHttpResponse.name,
      email: userHttpResponse.email,
      password: userHttpResponse.password,
      role: userHttpResponse.role,
      ethAddress: userHttpResponse.ethAddress,
      defaultPasswordChanged: userHttpResponse.defaultPasswordChanged,
      network: userHttpResponse.network,
      businessId: userHttpResponse.businessId,
      creationDate: userHttpResponse.creationDate,
      lastUpdate: userHttpResponse.lastUpdate,
    };

    return output;
  }
}
