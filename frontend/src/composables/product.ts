import type { ProductGroup } from "./productGroup"
import type { Unit } from "./unit"

export interface Product {
  id?: number
  name: string
  price: number
  group?: ProductGroup
  size: number
  unit?: Unit
  tax: number
  category: number
  visibility: number[]
}

export class Product implements Product{

  public constructor(prod: Product){
    this.id = prod.id;
    this.name = prod.name;
    this.price = prod.price;
    this.group = prod.group;
    this.size = prod.size;
    this.unit = prod.unit;
    this.tax = prod.tax;
    this.category = prod.category;
    this.visibility = prod.visibility;
  }
}

export function generateTestData(){
  let p1 = {
    name: "Beer",
    price: 120,
    size: 0.5,
    tax: 0,
    category: 3,
    visibility: [2]
  } as Product
  p1 = new Product(p1);

  let p2 = {
    name: "Bread",
    price: 250,
    size: 1,
    tax: 0,
    category: 4,
    visibility: [1, 2]
  } as Product
  p2 = new Product(p2);

  return [p1, p2];
}
