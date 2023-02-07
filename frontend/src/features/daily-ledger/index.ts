import { DefaultDailyLedgerHttpInfra } from '@/features/daily-ledger/infra/http/http-infra';
import { DefaultDailyLedgerHttpAdapter } from '@/features/daily-ledger/adapters/http/http-adapter';
import { DefaultDailyLedgerUseCase } from '@/features/daily-ledger/application/daily-ledger-usecase';

const httpInfra = new DefaultDailyLedgerHttpInfra();
const httpAdapter = new DefaultDailyLedgerHttpAdapter(httpInfra);
export const dailyLedgerUseCase = new DefaultDailyLedgerUseCase(httpAdapter);
