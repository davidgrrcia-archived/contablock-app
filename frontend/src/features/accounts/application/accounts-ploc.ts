import type { AccountsUseCase } from '@/features/accounts/application/accounts-usecase';
import type { Account } from '@/features/accounts/application/accounts-model';

interface AccountsState {
  accounts: Account[];
  stateMessage: string | null
}

const INITIAL_STATE: AccountsState = {
  accounts: [],
  stateMessage: null,
};

export class AccountsPLoC {
  private state: AccountsState;

  constructor(private readonly accountsUseCase: AccountsUseCase) {
    this.state = INITIAL_STATE;
  }

  async loadAccounts(){
    try {
      this.state = {
        accounts: [],
        stateMessage: "loading accounts",
      }

      const accounts = await this.accountsUseCase.listAccounts()

      this.state = {
        accounts: accounts,
        stateMessage: "accounts loaded",
      }

        }

    catch (err) {
      this.state = {
        accounts: [],
        stateMessage: "error loading accounts",
      }
    }
  }
}
