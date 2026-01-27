
export interface Unit {
  id?: number
  name: string
}

export class Unit implements Unit {

  constructor (unit: Unit) {
    this.id = unit.id;
    this.name = unit.name;
  }
}
