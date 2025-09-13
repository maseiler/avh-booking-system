export interface Account {
  id?: number
  firstName: string
  lastName: string
  nickName?: string
  email: string
  phone?: string
  balance: number
  maxDebt: number
  category?: number
  enabled: boolean
  createdAt?: string
}
export class Account implements Account{

  constructor(acc: Account){
    this.id = acc.id;
    this.firstName = acc.firstName;
    this.lastName = acc.lastName ;
    this.nickName = acc.nickName ;
    this.email = acc.email ;
    this.phone = acc.phone ;
    this.balance = acc.balance ;
    this.maxDebt = acc.maxDebt ;
    this.category = acc.category ;
    this.enabled = acc.enabled ;
    this.createdAt = acc.createdAt ;
  }

  public getShortName(): string{
    if(this.hasNickname()){
      return `${this.nickName}`;
    }
    return `${this.firstName} ${this.lastName[0]}.`;
  }

  public getFullName(): string {
    if(this.hasNickname()){
      return `${this.firstName} (${this.nickName}) ${this.lastName}`;
    }
    return `${this.firstName} ${this.lastName}`;
  }

  private hasNickname(): boolean{
    return (
      this.nickName !== null &&
      this.nickName !== undefined &&
      this.nickName !== ""
    );
  }
  
}

export function generateTestData(): Account[]{
  let a1 = {
    firstName: "Nomen",
    lastName: "Omen",
    email: "test@test.de",
    balance: Math.trunc(Math.random() * 10000),
    maxDebt: 0,
    enabled: true,
    category: 1
  } as Account
  a1 = new Account(a1);

  let a2 = {
    firstName: "Peter",
    lastName: "Super",
    email: "test@test.de",
    balance: Math.trunc(Math.random() * 10000),
    maxDebt: 0,
    enabled: true,
    category: 2
  } as Account
  a2 = new Account(a2)

  let a3= {
    firstName: "Andi",
    nickName: "Saufi",
    lastName: "Theke",
    email: "test@test.de",
    balance: Math.trunc(Math.random() * 10000),
    maxDebt: 0,
    enabled: true,
    category: 2
  } as Account
  a3 = new Account(a3)

  return [a3, a1, a2];
}