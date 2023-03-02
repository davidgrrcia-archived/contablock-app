export const currencies = [
  {
    code: 'EUR',
    symbol: '€',
  },
  {
    code: 'GBP',
    symbol: '£',
  },
  {
    code: 'USD',
    symbol: '$',
  },
];

export function findCurrencyByCode(code: string): string {
  const c = currencies.find(currency => currency.code === code);

  return c ? c.symbol : '';
}

export function getState(state: string): string {
  switch (state) {
    case 'CREATED':
      return 'Creado';
    case 'PUBLISHED':
      return 'Publicado';
    default:
      return 'N/A';
  }
}

export function getQuarter(quarter: number): string {
  switch (quarter) {
    case 1:
      return 'Primer Trimestre';
    case 2:
      return 'Segundo Trimestre';
    case 3:
      return 'Tercer Trimestre';
    case 4:
      return 'Cuarto Trimestre';
    default:
      return 'N/A';
  }
}

export function formatAmount(amount: number | null, currency: string): string {
  if (amount === null || amount === 0) {
    return '';
  }

  switch (currency) {
    case 'USD':
      return amount + '€';
    case 'EUR':
      return new Intl.NumberFormat('de-DE', {
        style: 'currency',
        currency: 'EUR',
      }).format(amount);
    default:
      return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
      }).format(amount);
  }
}

//format unix timestamp to date with time string
export function formatDate(date: number): string {
  const d = new Date(date * 1000);
  return d.toLocaleDateString() + ' - ' + d.toLocaleTimeString();
}

// date to unix timestamp
export function dateToUnixTimestamp(date: Date): number {
  return Math.floor(date.getTime() / 1000);
}

export function formatId(id: string): string {
  return id.substring(0, 8) + '...';
}

export function getRoleColor(role: string | undefined): string {
  switch (role) {
    case 'accountant':
      return 'bg-gray-800';
    case 'auditor':
      return 'bg-blue-800';
    case 'viewer':
      return 'bg-green-800';
    default:
      return 'bg-black';
  }
}
