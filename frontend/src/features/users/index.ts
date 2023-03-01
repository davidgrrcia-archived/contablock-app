import { DefaultUsersHttpInfra } from '@/features/users/infra/http/http-infra';
import { DefaultUsersHttpAdapter } from '@/features/users/adapters/http/http-adapter';
import { DefaultUsersUseCase } from '@/features/users/application/users-usecase';

const httpInfra = new DefaultUsersHttpInfra();
const httpAdapter = new DefaultUsersHttpAdapter(httpInfra);
export const usersUseCase = new DefaultUsersUseCase(httpAdapter);
