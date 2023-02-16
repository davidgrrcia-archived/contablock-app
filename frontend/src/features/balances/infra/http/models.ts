export interface Balance {
  id: string;
  name: string;
  currentYear: number;
  previousYear: number;
  currentAmount: number;
  previousAmount: number;
  child: Balance[] | null;
}

export interface BalanceListResponse {
  items: Balance[];
}
