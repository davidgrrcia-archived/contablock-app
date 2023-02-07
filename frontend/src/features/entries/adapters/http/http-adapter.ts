import type { DefaultEntriesHttpInfra } from '@/features/entries/infra/http/http-infra';
import type { EntriesAdapter } from '@/features/entries/application/entries-model';
import type { Entry as EntryUseCase } from '@/features/entries/application/entries-model';
import type { EntryCreateBody as EntryInfra } from '@/features/entries/infra/http/models';

export class DefaultEntriesHttpAdapter implements EntriesAdapter {
  constructor(private readonly entriesHttpInfra: DefaultEntriesHttpInfra) {}

  async createEntryHttpResponse(entry: EntryUseCase) {
    const input: EntryInfra = {
      accountId: entry.accountId,
      multipleEntryId: entry.multipleEntryId,
      currency: entry.currency,
      concept: entry.concept,
      date: entry.date,
      debitBalance: entry.debitAmount,
      creditBalance: entry.creditAmount,
    };

    const entryHttpResponse = await this.entriesHttpInfra.createEntry(input);

    const output: EntryUseCase = {
      accountId: entryHttpResponse.entry.accountId,
      multipleEntryId: entryHttpResponse.entry.multipleEntryId,
      currency: entryHttpResponse.entry.currency,
      concept: entryHttpResponse.entry.concept,
      date: entryHttpResponse.entry.date,
      debitAmount: entryHttpResponse.entry.debitBalance,
      creditAmount: entryHttpResponse.entry.creditBalance,
    };

    return output;
  }
}
