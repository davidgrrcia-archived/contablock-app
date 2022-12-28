import type { EntryListResponse } from '@/features/entries/infra/http/models';

export interface EntriesHttpInfra {
  listEntries(): Promise<EntryListResponse>;
}

export class DefaultEntriesHttpInfra implements EntriesHttpInfra {
  listEntries() {
    return fetch('/api/entries')
      .then(response => response.json())
      .then(response => response as EntryListResponse);
  }
}
