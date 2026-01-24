import { Product } from "./product"

export interface CartContent {
  product: Product
  quantity: number
  price: number
  tax: number
}
//ToDo: Save Tax like Price individually