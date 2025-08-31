import { Product } from "./product"

export interface CartContent {
  product: Product
  quantity: number
  productPrice: number
}