import type { EntryListResponse, EntryDescribeResponse, EntryCreateBody } from '@/features/entries/infra/http/models';

export interface EntriesHttpInfra {
  listEntries(): Promise<EntryListResponse>;
  createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse>;
}

export class DefaultEntriesHttpInfra implements EntriesHttpInfra {
  listEntries() {
    return fetch('/api/entries')
      .then(response => response.json())
      .then(response => response as EntryListResponse);
  }
  createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse> {
    return fetch('/api/entries', {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(entry),
    })
      .then(response => response.json())
      .then(response => response as EntryDescribeResponse);
  }
}
