import type { Business, BusinessAdapter } from '@/features/business/application/business-model';

export interface BusinessUseCase {
  describeBusiness(id: string): Promise<Business>;
}

export class DefaultBusinessUseCase implements BusinessUseCase {
  constructor(private readonly businessAdapter: BusinessAdapter) {}

  async describeBusiness(id: string) {
    return await this.businessAdapter.describeBusinessHttpResponse(id);
  }
}
