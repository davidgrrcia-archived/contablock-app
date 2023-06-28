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
    case 'UPDATED':
      return 'Actualizado';
    case 'REMOVED':
      return 'Eliminado';
    case 'PENDING_REVIEW':
      return 'Pendiente de revisión';
    case 'APPROVED':
      return 'Aprobado';
    case 'REJECTED':
      return 'Rechazado';
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

export function getRoleColor(role: string): string {
  switch (role) {
    case 'accountant':
      return 'bg-gray-800';
    case 'auditor':
      return 'bg-cyan-900';
    case 'viewer':
      return 'bg-indigo-900';
    default:
      return 'bg-black';
  }
}

export function getTotalBalanceColor(amount1: number, amount2: number): string {
  const roundedAmount1 = Math.round(amount1);
  const roundedAmount2 = Math.round(amount2);

  if (roundedAmount1 < roundedAmount2) {
    return 'text-red-300';
  } else if (roundedAmount1 > roundedAmount2) {
    return 'text-green-300';
  } else {
    return 'text-white';
  }
}

export function isActionAllowed(role: string, state: string, action: string): boolean {
  if (role === 'admin') {
    return true;
  }

  if (role === 'viewer') {
    return false;
  }

  if (role === 'accountant') {
    if (action === 'create') {
      return true;
    }

    if (action === 'modify') {
      return state === 'CREATED' || state === 'UPDATED' || state === 'REJECTED';
    }

    if (action === 'publish') {
      return state === 'APPROVED';
    }

    if (action === 'remove') {
      return state === 'CREATED' || state === 'UPDATED' || state === 'REJECTED';
    }

    if (action === 'review') {
      return state === 'CREATED' || state === 'UPDATED';
    }
  }

  if (role === 'auditor') {
    if (action === 'approve') {
      return state === 'PENDING_REVIEW';
    }

    if (action === 'reject') {
      return state === 'PENDING_REVIEW';
    }
  }

  return false;
}

export function showEntryActions(role: string, state: string): boolean {
  if (role === 'admin') {
    return true;
  }

  if (role === 'viewer') {
    return false;
  }

  if (role === 'accountant') {
    return state === 'CREATED' || state === 'UPDATED' || state === 'REJECTED' || state === 'APPROVED';
  }

  if (role === 'auditor') {
    return state === 'PENDING_REVIEW';
  }

  return false;
}

export function getEntryStateColor(state: string): string {
  switch (state) {
    case 'CREATED':
      return 'bg-yellow-100 text-yellow-800';
    case 'PUBLISHED':
      return 'bg-green-100 text-green-800';
    case 'UPDATED':
      return 'bg-yellow-100 text-yellow-800';
    case 'REJECTED':
      return 'bg-red-100 text-red-800';
    case 'PENDING_REVIEW':
      return 'bg-orange-100 text-orange-800';
    case 'APPROVED':
      return 'bg-blue-100 text-blue-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
}
