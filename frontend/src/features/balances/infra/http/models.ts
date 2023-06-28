export interface Balance {
  id: string;
  name: string;
  currentYear: number;
  previousYear: number;
  currentAmount: number;
  previousAmount: number;
  child: Balance[] | null;
  accounts: string[];
}

export interface BalanceListResponse {
  items: Balance[];
  currentDebitTotalAmount: number;
  previousDebitTotalAmount: number;
  currentCreditTotalAmount: number;
  previousCreditTotalAmount: number;
}
