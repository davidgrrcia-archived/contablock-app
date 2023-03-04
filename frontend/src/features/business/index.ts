import { DefaultBusinessHttpInfra } from '@/features/business/infra/http/http-infra';
import { DefaultBusinessHttpAdapter } from '@/features/business/adapters/http/http-adapter';
import { DefaultBusinessUseCase } from '@/features/business/application/business-usecase';

const httpInfra = new DefaultBusinessHttpInfra();
const httpAdapter = new DefaultBusinessHttpAdapter(httpInfra);
export const businessUseCase = new DefaultBusinessUseCase(httpAdapter);
