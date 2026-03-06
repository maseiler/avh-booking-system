<template>

  <Buttons>
    <Button
    class="is-warning"
    @click="$emit('cancelOrder')"
    title="discard cart and unselect account"
    :fa-icon="['fas', 'trash']"
    icon-position="left"
    >
    Cancel Order
    </Button>

    <Button
    v-if="!cart$.isOverdrawn"
    class="is-success"
    @click="checkoutOrder"
    title="discard cart and unselect account"
    :fa-icon="['fas', 'beer']"
    icon-position="right"
    >
    Book now
    </Button>

    <Button
    v-if="cart$.isOverdrawn"
    class="is-danger"
    title="discard cart and unselect account"
    :fa-icon="['fas', 'coins']"
    icon-position="right"
    >
    Pay for this order
    </Button>
  </Buttons>

  <Message
  class="is-danger hint"
  v-if="cart$.isOverdrawn"
  >
    <icon :icon="['fas', 'warning']" />
    The Cart exceeds the allowance! <br>
    <router-link to="/payment">Top up the account</router-link> or pay for this order immediately with the button above.
  </Message>
</template>

<style scoped>
.buttons{
  margin-top:.5rem;
  justify-content: space-between;
}
</style>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCartStore } from '../../store/CartStore';
import Button from '../../composables/elements/Button.vue';
import Buttons from '../../composables/elements/Buttons.vue';
import Message from '../../composables/elements/Message.vue';
import { useBookingStore } from '../../store/BookingStore';

export default{
  data() {
    return {
      account$: useAccountStore(),
      cart$: useCartStore(),
      booking$: useBookingStore(),
    }
  },
  components: {
    Button,
    Buttons,
    Message,
  },
  methods: {
    checkoutOrder(){
      this.booking$.addBooking(this.cart$.cartContents, this.account$.selected[0], 0)
      this.$emit('cancelOrder')
      // ToDo: this
    }
  },
  emits: {
    cancelOrder: null
  }
}
</script>