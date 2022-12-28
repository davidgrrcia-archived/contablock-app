import { DefaultAccountsHttpInfra } from '@/features/accounts/infra/http/http-infra';
import { DefaultAccountsHttpAdapter } from '@/features/accounts/adapters/http/http-adapter';
import { DefaultAccountsUseCase } from '@/features/accounts/application/accounts-usecase';

const httpInfra = new DefaultAccountsHttpInfra();
const httpAdapter = new DefaultAccountsHttpAdapter(httpInfra);
export const accountsUseCase = new DefaultAccountsUseCase(httpAdapter);
