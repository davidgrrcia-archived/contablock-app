import { DefaultBalancesHttpInfra } from '@/features/balances/infra/http/http-infra';
import { DefaultBalancesHttpAdapter } from '@/features/balances/adapters/http/http-adapter';
import { DefaultBalancesUseCase } from '@/features/balances/application/balances-usecase';

const httpInfra = new DefaultBalancesHttpInfra();
const httpAdapter = new DefaultBalancesHttpAdapter(httpInfra);
export const balancesUseCase = new DefaultBalancesUseCase(httpAdapter);
