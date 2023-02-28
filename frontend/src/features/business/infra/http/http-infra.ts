import type { BusinessDescribeResponse } from '@/features/business/infra/http/models';
import { useConfigStore } from '@/stores/config';

export interface BusinessHttpInfra {
  describeBusiness(id: string): Promise<BusinessDescribeResponse | Error>;
}

export class DefaultBusinessHttpInfra implements BusinessHttpInfra {
  // constructor(private readonly project: string) {} TODO use project constructor

  async describeBusiness(id: string) {
    const configStore = useConfigStore();

    const response = await fetch('/api/projects/' + configStore.state.project + '/business/' + id);

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as BusinessDescribeResponse;
  }
}
