import { defineStore } from 'pinia'
import { Account, type Account, generateTestData } from '../composables/account'
import { useSocketStore } from './socketStore'

export const useAccountStore = defineStore('account', {
  state: () => {
    return {
      accounts: [] as Account[],
      selected: [] as Account[]
    }
  },
  actions: {
    generateTestData(){
      // let testData = generateTestData();
      // this.accounts.push(...testData);
      // this.accounts.sort((a, b) => {
      //   return a.getShortName() < b.getShortName() ? -1 : 1;
      // })
    },
    select(acc: Account){
      this.selected = [acc];
    },
    selectAdd(acc: Account){
      if(this.selected.includes(acc)){ return; }
      this.selected.push(acc);
    },
    selectSubstract(acc: Account){
      this.selected = this.selected.filter((a) => a != acc);
    },
    unselect(){
      this.selected = [];
    },
    getByCategory(categoryId: number, all?: boolean): Account[]{
      let enabledUsers = this.accounts as Account[];
      if(!all){
        enabledUsers = this.accounts.filter((acc) => acc.enabled) as Account[];
      }
      if(categoryId == 0){
        return enabledUsers;
      }
      return enabledUsers.filter((acc) => acc.category == categoryId );
    },
    getBySearchAndCategory(searchString: string, categoryId: number, all?: boolean): Account[]{      
      let search = searchString.toLowerCase();
      let byCategory = this.getByCategory(categoryId, all);
      let searchResults = byCategory.filter((acc) => {
        return (
          acc.firstName.toLowerCase().includes(search) ||
          acc.nickname?.toLowerCase().includes(search) ||
          acc.lastName.toLowerCase().includes(search)
        )
      })
      return searchResults;
    },
    byId(id: number | undefined): Account | undefined {
      let foundAcc = this.accounts.find((acc) => acc.id == id);
      return foundAcc
    },
    loadAllAccounts(accJsonObj: any[]){
      let accObjArray = [] as Account[];
      accJsonObj.forEach((accObj) => {
        // let newAccount = {} as Account;
        let newAccount = new Account({} as Account);
        newAccount.jsonToClass(accObj);
        accObjArray.push(newAccount);
      })
      this.accounts = accObjArray;
    },
    addAccount(newAccount: Account){
      this.accounts.push(newAccount);
      // const socket$ = useSocketStore();
      // socket$.addAccount(newAccount);
      // socket$.queryAccounts();
    },
    patchAccounts(newAccounts: Account[]){
      this.$patch(state => {
        newAccounts.forEach(newAcc => {
          const newAccObj = new Account(newAcc);
          const existing = state.accounts.find(a => a.id === newAccObj.id);
          if(existing){
            Object.assign(existing, newAccObj);
          } else {
            state.accounts.push(newAccObj);
          }
        })
      })
    }
  }
})