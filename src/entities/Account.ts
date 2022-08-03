export enum AccountType {
  Grupo = "Grupo",
  Subgrupo = "Subgrupo",
  Cuenta = "Cuenta",
  Subcuenta = "Subcuenta",
}

export enum AccountKind {
  Activo = "Activo",
  Pasivo = "Pasivo",
  PatrimonioNeto = "Patrimonio Neto",
  Gasto = "Gasto",
  Ingreso = "Ingreso",
}

export default interface Account {
  id: string;
  name: string;
  type: AccountType;
  kind: AccountKind;
}
