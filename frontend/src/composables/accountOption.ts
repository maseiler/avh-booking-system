import {type Account} from './account'
import { useAccountStore } from '../store/AccountStore'

export interface AccountOption{
  account: number
  key: string
  value: string
  getAccount(): Account | undefined
}

export class AccountOption implements AccountOption{

  constructor(acO: AccountOption){
    this.account = acO.account;
    this.key = acO.key;
    this.value = acO.value;
  }

  public copy(): AccountOption{
    let newAccountOption = new AccountOption(this);
    return newAccountOption;
  }
  
  public getSccount(): Account | undefined{
    return useAccountStore().byId(this.account)
  }
}
