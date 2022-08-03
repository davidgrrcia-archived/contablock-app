enum UserRole {
  Contable = "Contable",
  Auditor = "Auditor",
  Stakeholder = "Stakeholder",
}

export interface User {
  id?: string;
  businessId: string;
  name: string;
  email: string;
  password: string;
  phone: string;
  role: string;
  state?: string;
  creationDate?: Date;
  lastUpdate?: Date;
}

export interface UserOutput {
  id: string;
  businessId: string;
  name: string;
  email: string;
  password: string;
  phone: string;
  role: string;
  state: string;
  creationDate: Date;
  lastUpdate: Date;
}
