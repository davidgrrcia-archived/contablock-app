import { DefaultAccountsHttpInfra } from '@/features/accounts/infra/http/http-infra';
import { DefaultAccountsHttpAdapter } from '@/features/accounts/adapters/http/http-adapter';
import { DefaultAccountsUseCase } from '@/features/accounts/application/accounts-usecase';
import { AccountsStore } from "@/features/accounts/application/accounts-store";

const httpInfra = new DefaultAccountsHttpInfra();
const httpAdapter = new DefaultAccountsHttpAdapter(httpInfra);
const accountsUseCase = new DefaultAccountsUseCase(httpAdapter);
export const accountsPLoC = new AccountsStore(accountsUseCase)