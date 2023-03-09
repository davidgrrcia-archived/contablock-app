import type { DefaultEntriesHttpInfra } from '@/features/entries/infra/http/http-infra';
import type { EntriesAdapter } from '@/features/entries/application/entries-model';
import type { Entry as EntryUseCase } from '@/features/entries/application/entries-model';
import type * as infraModel from '@/features/entries/infra/http/models';

export class DefaultEntriesHttpAdapter implements EntriesAdapter {
  constructor(private readonly entriesHttpInfra: DefaultEntriesHttpInfra) {}

  async createEntryHttpResponse(entry: EntryUseCase) {
    const input: infraModel.EntryCreateBody = {
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
      id: entryHttpResponse.id,
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }

  async publishEntryHttpResponse(id: string) {
    const input: infraModel.EntryPublishBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.publishEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.id, //todo: this returning undefined
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }

  async removeEntryHttpResponse(id: string) {
    const input: infraModel.EntryRemoveBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.removeEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.id,
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }

  async reviewEntryHttpResponse(id: string) {
    const input: infraModel.EntryReviewBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.reviewEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.id, //todo: this returning undefined
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }

  async approveEntryHttpResponse(id: string) {
    const input: infraModel.EntryApproveBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.approveEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.id, //todo: this returning undefined
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }

  async rejectEntryHttpResponse(id: string) {
    const input: infraModel.EntryRejectBody = {
      contractId: '',
      userId: '',
      id: id,
    };

    const entryHttpResponse = await this.entriesHttpInfra.rejectEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.id, //todo: this returning undefined
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }

  async updateEntryHttpResponse(id: string, state: string): Promise<EntryUseCase | Error> {
    const input: infraModel.EntryUpdateBody = {
      id: id,
      state: state,
    };

    const entryHttpResponse = await this.entriesHttpInfra.updateEntry(input);

    if (entryHttpResponse instanceof Error) {
      return entryHttpResponse;
    }

    const output: EntryUseCase = {
      id: entryHttpResponse.id,
      accountId: entryHttpResponse.accountId,
      multipleEntryId: entryHttpResponse.multipleEntryId,
      currency: entryHttpResponse.currency,
      concept: entryHttpResponse.concept,
      date: entryHttpResponse.date,
      quarter: entryHttpResponse.quarter,
      debitAmount: entryHttpResponse.debitBalance,
      creditAmount: entryHttpResponse.creditBalance,
      totalAmount: entryHttpResponse.totalBalance,
      state: entryHttpResponse.state,
      stateReason: entryHttpResponse.stateReason,
      creationDate: entryHttpResponse.creationDate,
      lastUpdate: entryHttpResponse.lastUpdate,
    };

    return output;
  }
}
