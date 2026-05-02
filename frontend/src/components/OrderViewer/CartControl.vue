<template>

  <Buttons>
    <Button
    class="is-warning"
    @click="$emit('cancelOrder')"
    :title="$t('transaction.discardCartTooltip')"
    :fa-icon="['fas', 'trash']"
    icon-position="left"
    >
    {{ $t('transaction.discardCart') }}
    </Button>

    <Button
    v-if="!cart$.isOverdrawn"
    class="is-primary"
    @click="checkoutOrder"
    :title="$t('transaction.bookNowTooltip')"
    :fa-icon="['fas', 'beer']"
    icon-position="right"
    >
    {{ $t('transaction.bookNow') }}
    </Button>

    <Button
    v-if="cart$.isOverdrawn"
    class="is-danger"
    :title="$t('transaction.payNowTooltip')"
    :fa-icon="['fas', 'coins']"
    icon-position="right"
    >
    {{ $t('transaction.payNow') }}
    </Button>
  </Buttons>

  <Message
  class="is-danger hint"
  v-if="cart$.isOverdrawn" 
  >
    <icon :icon="['fas', 'warning']" />
    {{ $t('transaction.allowanceExceeded') }} <br>
    <router-link to="/payment">{{ $t('transaction.topUp') }}</router-link> {{ $t('transaction.orPayNow') }}
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
      this.booking$.addBooking(this.cart$.cartContents, this.account$.selected, 0)
      this.$emit('cancelOrder')
      // ToDo: this
    }
  },
  emits: {
    cancelOrder: null
  }
}
</script>