import type { Entry, EntriesAdapter } from '@/features/entries/application/entries-model';

export interface EntriesUseCase {
  createEntry(entry: Entry): Promise<Entry>;
}

export class DefaultEntriesUseCase implements EntriesUseCase {
  constructor(private readonly entriesAdapter: EntriesAdapter) {}

  async createEntry(entry: Entry) {
    return await this.entriesAdapter.createEntryHttpResponse(entry);
  }
}
