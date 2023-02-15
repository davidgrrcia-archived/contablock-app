import { DefaultGeneralLedgerHttpInfra } from '@/features/general-ledger/infra/http/http-infra';
import { DefaultGeneralLedgerHttpAdapter } from '@/features/general-ledger/adapters/http/http-adapter';
import { DefaultGeneralLedgerUseCase } from '@/features/general-ledger/application/general-ledger-usecase';

const httpInfra = new DefaultGeneralLedgerHttpInfra();
const httpAdapter = new DefaultGeneralLedgerHttpAdapter(httpInfra);
export const generalLedgerUseCase = new DefaultGeneralLedgerUseCase(httpAdapter);
