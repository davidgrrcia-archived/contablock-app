import type { BusinessDescribeResponse } from '@/features/business/infra/http/models';

export interface BusinessHttpInfra {
  describeBusiness(project: string, id: string): Promise<BusinessDescribeResponse | Error>;
}

export class DefaultBusinessHttpInfra implements BusinessHttpInfra {
  async describeBusiness(project: string, id: string) {
    const response = await fetch(`/api/projects/${project}/business/${id}`);

    const data = await response.json();

    if (data.statusCode === 500) {
      return new Error(data.error);
    }

    return data as BusinessDescribeResponse;
  }
}
