export interface Account {
  id?: number
  firstName: string
  lastName: string
  nickname?: string
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
    this.nickname = acc.nickname ;
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
      return `${this.nickname}`;
    }
    return `${this.firstName} ${this.lastName[0]}.`;
  }

  public getFullName(): string {
    if(this.hasNickname()){
      return `${this.firstName} (${this.nickname}) ${this.lastName}`;
    }
    return `${this.firstName} ${this.lastName}`;
  }

  private hasNickname(): boolean{
    return (
      this.nickname !== null &&
      this.nickname !== undefined &&
      this.nickname !== ""
    );
  }

  public jsonToClass(obj: any){
    this.id = 'id' in obj ? obj.id : null;
    this.firstName = 'firstName' in obj ? obj.firstName : "";
    this.lastName = 'lastName' in obj ? obj.lastName : "";
    this.nickname = 'nickname' in obj ? obj.nickname : "";
    this.email = 'email' in obj ? obj.email : "";
    this.phone = 'phone' in obj ? obj.phone : "";
    this.balance = 'balance' in obj ? obj.balance : 0;
    this.maxDebt = 'maxDebt' in obj ? obj.maxDebt : 0;
    this.category = 'category' in obj ? obj.category : 0;
    this.enabled = 'enabled' in obj ? obj.enabled : false;
    this.createdAt = 'createdAt' in obj ? obj.createdAt : null;
  }
  
}

export function generateTestData(): Account[]{
  let a1 = {
    id: 1,
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
    id: 2,
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
    id: 3,
    firstName: "Andi",
    nickname: "Saufi",
    lastName: "Theke",
    email: "test@test.de",
    balance: 500,
    maxDebt: 0,
    enabled: true,
    category: 2
  } as Account
  a3 = new Account(a3)

  return [a3, a1, a2];
}