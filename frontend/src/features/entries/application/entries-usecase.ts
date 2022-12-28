import type { Entry, EntriesAdapter } from '@/features/entries/application/entries-model';

export interface EntriesUseCase {
  listEntries(): Promise<Entry[]>;
}

export class DefaultEntriesUseCase implements EntriesUseCase {
  constructor(private readonly entriesAdapter: EntriesAdapter) {}

  async listEntries() {
    return await this.entriesAdapter.listEntriesHttpResponse();
  }
}
