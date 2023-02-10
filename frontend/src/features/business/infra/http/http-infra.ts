import type { BusinessDescribeResponse } from '@/features/business/infra/http/models';

export interface BusinessHttpInfra {
  describeBusiness(id: string): Promise<BusinessDescribeResponse>;
}

export class DefaultBusinessHttpInfra implements BusinessHttpInfra {
  describeBusiness(id: string) {
    return fetch('/api/business/' + id)
      .then(response => response.json())
      .then(response => response as BusinessDescribeResponse);
  }
}
