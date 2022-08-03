enum EntryType {
  Nominas = "Nóminas",
  Amortizaciones = "Amortizaciones",
  PagoImpuestos = "Pago de impuestos",
  GestionVentas = "Gestión de ventas",
  GestionCompras = "Gestión de compras",
}

export interface Entry {
  id?: string;
  businessId: string;
  currency: string;
  description: string;
  participant: string;
  fileUpload: string;
  date: Date;
  type: string;
  debitBalance?: number;
  creditBalance?: number;
  totalBalance?: number;
  state?: string;
  creationDate?: Date;
  lastUpdate?: Date;
}

export interface EntryOutput {
  id: string;
  businessId: string;
  currency: string;
  description: string;
  participant: string;
  fileUpload: string;
  date: Date;
  type: string;
  debitBalance: number;
  creditBalance: number;
  totalBalance: number;
  state: string;
  creationDate: Date;
  lastUpdate: Date;
}
