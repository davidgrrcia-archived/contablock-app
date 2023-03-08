import type * as infraModel from '@/features/entries/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface EntriesHttpInfra {
  createEntry(entry: infraModel.EntryCreateBody): Promise<infraModel.EntryDescribeResponse | Error>;
  publishEntry(input: infraModel.EntryPublishBody): Promise<infraModel.EntryDescribeResponse | Error>;
  removeEntry(input: infraModel.EntryRemoveBody): Promise<infraModel.EntryDescribeResponse | Error>;
  reviewEntry(input: infraModel.EntryPublishBody): Promise<infraModel.EntryDescribeResponse | Error>;
  approveEntry(input: infraModel.EntryPublishBody): Promise<infraModel.EntryDescribeResponse | Error>;
  rejectEntry(input: infraModel.EntryPublishBody): Promise<infraModel.EntryDescribeResponse | Error>;
}

export class DefaultEntriesHttpInfra implements EntriesHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async createEntry(entry: infraModel.EntryCreateBody): Promise<infraModel.EntryDescribeResponse | Error> {
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

    return data as infraModel.EntryDescribeResponse;
  }
  async publishEntry(input: infraModel.EntryPublishBody): Promise<infraModel.EntryDescribeResponse | Error> {
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

    return data as infraModel.EntryDescribeResponse;
  }
  async removeEntry(input: infraModel.EntryRemoveBody): Promise<infraModel.EntryDescribeResponse | Error> {
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

    return data as infraModel.EntryDescribeResponse;
  }

  async reviewEntry(input: infraModel.EntryReviewBody): Promise<infraModel.EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    input.contractId = configStore.state.contractId;
    input.userId = configStore.state.userId;

    const response = await fetch(`/api/projects/${configStore.state.project}/entries:review`, {
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

    return data as infraModel.EntryDescribeResponse;
  }

  async approveEntry(input: infraModel.EntryApproveBody): Promise<infraModel.EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    input.contractId = configStore.state.contractId;
    input.userId = configStore.state.userId;

    const response = await fetch(`/api/projects/${configStore.state.project}/entries:approve`, {
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

    return data as infraModel.EntryDescribeResponse;
  }

  async rejectEntry(input: infraModel.EntryRejectBody): Promise<infraModel.EntryDescribeResponse | Error> {
    const configStore = useConfigStore();

    input.contractId = configStore.state.contractId;
    input.userId = configStore.state.userId;

    const response = await fetch(`/api/projects/${configStore.state.project}/entries:reject`, {
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

    return data as infraModel.EntryDescribeResponse;
  }
}
