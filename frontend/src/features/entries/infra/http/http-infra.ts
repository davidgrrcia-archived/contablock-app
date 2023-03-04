import type { EntryDescribeResponse, EntryCreateBody, EntryPublishBody, EntryRemoveBody } from '@/features/entries/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface EntriesHttpInfra {
  createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse | Error>;
  publishEntry(input: EntryPublishBody): Promise<EntryDescribeResponse | Error>;
  removeEntry(input: EntryRemoveBody): Promise<EntryDescribeResponse | Error>;
}

export class DefaultEntriesHttpInfra implements EntriesHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async createEntry(entry: EntryCreateBody): Promise<EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    entry.contractId = configStore.state.contractId;
    entry.userId = configStore.state.userId;

    const response = await fetch(`/api/projects/${configStore.state.project}/entries`, {
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
  async publishEntry(input: EntryPublishBody): Promise<EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    input.contractId = configStore.state.contractId;
    input.userId = configStore.state.userId;

    const response = await fetch(`/api/projects/${configStore.state.project}/entries:publish`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(input),
    });

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as EntryDescribeResponse;
  }
  async removeEntry(input: EntryRemoveBody): Promise<EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    input.contractId = configStore.state.contractId;
    input.userId = configStore.state.userId;

    const response = await fetch(`/api/projects/${configStore.state.project}/entries/${input.id}`, {
      method: 'DELETE',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(input),
    });

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as EntryDescribeResponse;
  }
}
