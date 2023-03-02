import type { DefaultBusinessHttpInfra } from '@/features/business/infra/http/http-infra';
import type { Business, BusinessAdapter } from '@/features/business/application/business-model';

export class DefaultBusinessHttpAdapter implements BusinessAdapter {
  constructor(private readonly businessHttpInfra: DefaultBusinessHttpInfra) {}

  async describeBusinessHttpResponse(project: string, id: string) {
    const input = await this.businessHttpInfra.describeBusiness(project, id);

    if (input instanceof Error) {
      return input;
    }

    const output: Business = {
      project: input.project,
      id: input.id,
      name: input.name,
      socialDenomination: input.socialDenomination,
      registeredOffice: input.registeredOffice,
      cif: input.cif,
      logo: input.logo,
      ethAddress: input.ethAddress,
    };

    return output;
  }
}
