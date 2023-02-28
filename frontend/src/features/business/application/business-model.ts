export interface BusinessAdapter {
  describeBusinessHttpResponse(id: string): Promise<Business | Error>;
}

export interface Business {
  id: string;
  name: string;
  socialDenomination: string;
  registeredOffice: string;
  cif: string;
  logo: string;
  ethAddress: string;
}
