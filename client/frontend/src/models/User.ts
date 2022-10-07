export default interface User {
    Id: number
    Creation: Date
    Client: string
    BierName: string
    FirstName: string
    LastName: string
    BoatName: string
    Status: string
    Email: string
    Phone: string
    Balance: number
    MaxDebt: number
}

export const nullUser: User = {
    Id: 0,
    Creation: new Date(0),
    Client: '',
    BierName: '',
    FirstName: '',
    LastName: '',
    BoatName: '',
    Status: '',
    Email: '',
    Phone: '',
    Balance: 0,
    MaxDebt: 0
  }