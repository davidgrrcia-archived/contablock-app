export interface BusinessAdapter {
  describeBusinessHttpResponse(project: string, id: string): Promise<Business | Error>;
}

export interface Business {
  project: string;
  id: string;
  name: string;
  socialDenomination: string;
  registeredOffice: string;
  cif: string;
  logo: string;
  ethAddress: string;
}
