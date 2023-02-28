import type { EntryDescribeResponse, EntryCreateBody } from '@/features/entries/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface EntriesHttpInfra {
  createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse | Error>;
  publishEntry(id: string): Promise<EntryDescribeResponse | Error>;
  removeEntry(id: string): Promise<EntryDescribeResponse | Error>;
}

export class DefaultEntriesHttpInfra implements EntriesHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/'${configStore.state.project}/entries`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(entry),
    });

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as EntryDescribeResponse;
  }
  async publishEntry(id: string): Promise<EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/'${configStore.state.project}/entries:publish`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ id }),
    });

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as EntryDescribeResponse;
  }
  async removeEntry(id: string): Promise<EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    const response = await fetch(`/api/projects/'${configStore.state.project}/entries/${id}`, {
      method: 'DELETE',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
    });

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as EntryDescribeResponse;
  }
}
