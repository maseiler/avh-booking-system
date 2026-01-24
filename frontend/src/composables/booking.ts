import { type CartContent } from "./cartContent";
import { Account } from "./account";

export const BookingType = {
  BOOKING_PAYED_BY_ACCOUNT: 0,
  BOOKING_PAYED_BY_CASH: 1,
  BOOKING_PAYED_BY_CARD: 2,
  CANCELLATION: 3,
} as const;
export type BookingType = (typeof BookingType)[keyof typeof BookingType];

export type BookingTotals = [total: number, tax: number];


export interface Booking {
  id: number,
  products: CartContent[],
  account: Account,
  timestamp: string,
  type: BookingType,
  reference?: Booking  // When a booking gets cancelled the original booking should get the canellation booking as a reference so that it can be marked as cancelled. The cancellation order will handle the rest.
}

export class Booking implements Booking{

  constructor(booking: Booking){
    this.products = booking.products;
    this.timestamp = booking.timestamp;
    this.account = booking.account;
    this.type = booking.type
    this.reference = booking.reference;
  }

  public getTotals(): BookingTotals{
    let total = 0;
    let tax = 0;

    this.products.forEach((product) => {
      let subTotal = product.price * product.quantity;
      let subTax = subTotal*(product.tax / 100);
      total += subTotal;
      tax += subTax;
    })
    return [total, tax] as BookingTotals;
  }
}