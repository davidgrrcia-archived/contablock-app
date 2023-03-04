import type { DefaultEntriesHttpInfra } from '@/features/entries/infra/http/http-infra';
import type { EntriesAdapter } from '@/features/entries/application/entries-model';
import type { Entry as EntryUseCase } from '@/features/entries/application/entries-model';
import type { EntryCreateBody, EntryPublishBody, EntryRemoveBody } from '@/features/entries/infra/http/models';

export class DefaultEntriesHttpAdapter implements EntriesAdapter {
  constructor(private readonly entriesHttpInfra: DefaultEntriesHttpInfra) {}

  async createEntryHttpResponse(entry: EntryUseCase) {
    const input: EntryCreateBody = {
      contractId: '',
      userId: '',
      accountId: entry.accountId,
      multipleEntryId: entry.multipleEntryId,
      currency: entry.currency,
      concept: entry.concept,
      date: entry.date,
      debitBalance: entry.debitAmount,
      creditBalance: entry.creditAmount,
    };

    const entryHttpResponse = await this.entriesHttpInfra.createEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.entry.id,
      accountId: entryHttpResponse.entry.accountId,
      multipleEntryId: entryHttpResponse.entry.multipleEntryId,
      currency: entryHttpResponse.entry.currency,
      concept: entryHttpResponse.entry.concept,
      date: entryHttpResponse.entry.date,
      quarter: entryHttpResponse.entry.quarter,
      debitAmount: entryHttpResponse.entry.debitBalance,
      creditAmount: entryHttpResponse.entry.creditBalance,
      totalAmount: entryHttpResponse.entry.totalBalance,
      state: entryHttpResponse.entry.state,
      stateReason: entryHttpResponse.entry.stateReason,
      creationDate: entryHttpResponse.entry.creationDate,
      lastUpdate: entryHttpResponse.entry.lastUpdate,
    };

    return output;
  }

  async publishEntryHttpResponse(id: string) {
    const input: EntryPublishBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.publishEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.entry.id, //todo: this returning undefined
      accountId: entryHttpResponse.entry.accountId,
      multipleEntryId: entryHttpResponse.entry.multipleEntryId,
      currency: entryHttpResponse.entry.currency,
      concept: entryHttpResponse.entry.concept,
      date: entryHttpResponse.entry.date,
      quarter: entryHttpResponse.entry.quarter,
      debitAmount: entryHttpResponse.entry.debitBalance,
      creditAmount: entryHttpResponse.entry.creditBalance,
      totalAmount: entryHttpResponse.entry.totalBalance,
      state: entryHttpResponse.entry.state,
      stateReason: entryHttpResponse.entry.stateReason,
      creationDate: entryHttpResponse.entry.creationDate,
      lastUpdate: entryHttpResponse.entry.lastUpdate,
    };

    return output;
  }
  async removeEntryHttpResponse(id: string) {
    const input: EntryRemoveBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.removeEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.entry.id,
      accountId: entryHttpResponse.entry.accountId,
      multipleEntryId: entryHttpResponse.entry.multipleEntryId,
      currency: entryHttpResponse.entry.currency,
      concept: entryHttpResponse.entry.concept,
      date: entryHttpResponse.entry.date,
      quarter: entryHttpResponse.entry.quarter,
      debitAmount: entryHttpResponse.entry.debitBalance,
      creditAmount: entryHttpResponse.entry.creditBalance,
      totalAmount: entryHttpResponse.entry.totalBalance,
      state: entryHttpResponse.entry.state,
      stateReason: entryHttpResponse.entry.stateReason,
      creationDate: entryHttpResponse.entry.creationDate,
      lastUpdate: entryHttpResponse.entry.lastUpdate,
    };

    return output;
  }
}
