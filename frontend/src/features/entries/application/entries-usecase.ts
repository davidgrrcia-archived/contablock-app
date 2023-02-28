import type { CreateEntryInput, Entry, EntriesAdapter } from '@/features/entries/application/entries-model';

export interface EntriesUseCase {
  createEntry(entry: CreateEntryInput): Promise<Entry | Error>;
  publishEntry(id: string): Promise<Entry | Error>;
  removeEntry(id: string): Promise<Entry | Error>;
}

export class DefaultEntriesUseCase implements EntriesUseCase {
  constructor(private readonly entriesAdapter: EntriesAdapter) {}

  async createEntry(entry: CreateEntryInput) {
    return await this.entriesAdapter.createEntryHttpResponse(entry);
  }

  async publishEntry(id: string) {
    return await this.entriesAdapter.publishEntryHttpResponse(id);
  }
  async removeEntry(id: string): Promise<Entry | Error> {
    return await this.entriesAdapter.removeEntryHttpResponse(id);
  }
}
