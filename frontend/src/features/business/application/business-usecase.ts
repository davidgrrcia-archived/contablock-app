import type { Business, BusinessAdapter } from '@/features/business/application/business-model';

export interface BusinessUseCase {
  describeBusiness(project: string, id: string): Promise<Business | Error>;
}

export class DefaultBusinessUseCase implements BusinessUseCase {
  constructor(private readonly businessAdapter: BusinessAdapter) {}

  async describeBusiness(project: string, id: string) {
    return await this.businessAdapter.describeBusinessHttpResponse(project, id);
  }
}
