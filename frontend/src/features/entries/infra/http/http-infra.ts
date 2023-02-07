import type { EntryDescribeResponse, EntryCreateBody } from '@/features/entries/infra/http/models';

export interface EntriesHttpInfra {
  createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse>;
}

export class DefaultEntriesHttpInfra implements EntriesHttpInfra {
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
