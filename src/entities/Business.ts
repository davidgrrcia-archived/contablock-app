export interface Business {
  id?: string;
  name: string;
  shortDescription: string;
  longDescription: string;
  logo: string;
  stockPrice: number;
  urlWebsite: string;
  ceoName: string;
  numberOfEmployees: number;
  country: string;
  nif: string;
  fiscalName: string;
  address: string;
  phone: string;
  email: string;
  state?: string;
  creationDate?: Date;
  lastUpdate?: Date;
}

export interface BusinessOutput {
  id: string;
  name: string;
  shortDescription: string;
  longDescription: string;
  logo: string;
  stockPrice: number;
  urlWebsite: string;
  ceoName: string;
  numberOfEmployees: number;
  country: string;
  nif: string;
  fiscalName: string;
  address: string;
  phone: string;
  email: string;
  state: string;
  creationDate: Date;
  lastUpdate: Date;
}
