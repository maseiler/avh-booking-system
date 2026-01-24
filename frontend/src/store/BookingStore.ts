import { defineStore } from 'pinia'
import { Booking } from '../composables/booking';
import type { CartContent } from '../composables/cartContent';
import type { BookingType } from '../composables/booking';
import type { Account } from '../composables/account';

export const useBookingStore = defineStore('booking', {
  state: () => {
    return {
      bookings: [] as Booking[],
      
    }
  },
  actions:{
    addBooking(cartContent: CartContent[], account: Account, type: BookingType, reference?: Booking){
      const timestamp = new Date(Date.now()).toISOString();
      let newBooking = new Booking({
        products: cartContent,
        account: account,
        timestamp: timestamp,
        type: type,
        reference: reference,
      } as Booking);

      this.bookings.push(newBooking);
      // ToDo: add listener to store mutations to communicate with websocket
      // Let Websocket send mutations to the accounts, so that these will be patched automatically
    }
  }
});
