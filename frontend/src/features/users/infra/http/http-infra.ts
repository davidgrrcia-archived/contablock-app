import type { UserDescribeResponse, UserLoginBody } from '@/features/users/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface UsersHttpInfra {
  loginUser(body: UserLoginBody): Promise<UserDescribeResponse | Error>;
}

export class DefaultUsersHttpInfra implements UsersHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async loginUser(body: UserLoginBody): Promise<UserDescribeResponse | Error> {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/${configStore.state.project}/users:login`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email: body.email, password: body.password }),
    });

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as UserDescribeResponse;
  }
}
