import type { CreateEntryInput, Entry, EntriesAdapter } from '@/features/entries/application/entries-model';

export interface EntriesUseCase {
  createEntry(entry: CreateEntryInput): Promise<Entry>;
}

export class DefaultEntriesUseCase implements EntriesUseCase {
  constructor(private readonly entriesAdapter: EntriesAdapter) {}

  async createEntry(entry: CreateEntryInput) {
    return await this.entriesAdapter.createEntryHttpResponse(entry);
  }
}
