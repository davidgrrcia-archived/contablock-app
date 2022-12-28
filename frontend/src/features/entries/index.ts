import { DefaultEntriesHttpInfra } from '@/features/entries/infra/http/http-infra';
import { DefaultEntriesHttpAdapter } from '@/features/entries/adapters/http/http-adapter';
import { DefaultEntriesUseCase } from '@/features/entries/application/entries-usecase';

const httpInfra = new DefaultEntriesHttpInfra();
const httpAdapter = new DefaultEntriesHttpAdapter(httpInfra);
export const entriesUseCase = new DefaultEntriesUseCase(httpAdapter);
