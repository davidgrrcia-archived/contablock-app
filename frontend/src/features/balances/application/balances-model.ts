export interface BalancesAdapter {
  listBalancesHttpResponse(): Promise<Balance[]>;
}

export interface Balance {
  id: string;
  name: string;
  currentYear: number;
  previousYear: number;
  currentAmount: number;
  previousAmount: number;
  child: Balance[] | null;
}
