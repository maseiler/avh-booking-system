<template>
  <article class="order message is-info">

    <div class="message-header" @click="toggleOrderDetails">
      <span>
        <icon :class="showOrderDetails ? 'showDetails' : ''" class="order-icon details" :icon="['fas', 'arrow-up-short-wide']" />
        <icon :class="showOrderDetails ? 'showDetails' : ''" class="order-icon" :icon="['fas', 'arrow-down-short-wide']" />
        {{ $t('transaction.cart') }} {{ dateString }}
      </span>
      <button v-if="!allowEdit" class="button is-small is-ghost receipt-btn" @click.stop="showReceiptModal = true" :title="$t('transaction.showReceipt', 'Kassenbon anzeigen')">
          <span class="icon"><icon :icon="['fas', 'receipt']" /></span>
        </button>
        <button v-if="accounts.length > 0 && allowEdit" class="delete" @click="cancelOrder" :title="$t('transaction.discardCartTooltip')"></button>
    </div>

    <div :class="showOrderDetails ? 'showDetails' : ''" class="message-body fixed-grid has-3-cols">
      <AccountTagList :accounts="accounts" :allowEdit="allowEdit" @unselect="(a) => unselectAccount(a)"></AccountTagList>

      <CartList  :allowEdit="allowEdit" :contents="contents" />
      <CartSums :totals="totals" v-if="accounts.length > 0"/>
      <CartControl v-if="allowEdit && accounts.length > 0" @cancelOrder="cancelOrder"/>
    </div>
    <div :class="showOrderDetails ? 'showDetails' : ''" class="order-ripped-teaser"> </div>
  </article>

  <Teleport to="body">
    <ReceiptModal
      v-if="showReceiptModal"
      :products="contents"
      :accounts="accounts"
      :timestamp="timestamp ?? ''"
      :booking-id="bookingId"
      @close="showReceiptModal = false"
    />
  </Teleport>
</template>

<style scoped>
.order{
  position:relative;
}
.order-icon{
  margin-right:.5em;
  &.details{
    display:none;
  }
}
.showDetails.order-icon{
  display:none;
  &.details{
    display:inline-block;
  }
}
.message-header{
  cursor:pointer;
  font-size:1.05rem;
  letter-spacing: 0.02em;
}
.receipt-btn{
  color: inherit;
  opacity: 0.75;
  &:hover{ opacity: 1; }
}
.message-body{
  display:none;
  /* position:absolute; */
  width:100%;
  z-index:20;
  /* background-color: hsl(var(--bulma-scheme-h), var(--bulma-scheme-s), var(--bulma-scheme-main-l)); */
  box-shadow: 0 4px 16px rgba(0,0,0,0.12);
  border-top: none;
}
.order-ripped-teaser{
  --_bg-color:hsl(var(--bulma-message-h),var(--bulma-message-s),var(--bulma-message-background-l));
  background-color:var(--_bg-color);
  width:100%;
  height:0.5rem;
  position:relative;
}
.order-ripped-teaser::after {
  content: '';
  position: absolute;
  right: 0;
  left: -0%;
  top: 100%;
  z-index: 10;
  display: block;
  height: 20px;
  background-size: 20px 100%;
  background-image: linear-gradient(135deg, var(--_bg-color) 25%, transparent 25%), linear-gradient(225deg, var(--_bg-color) 25%, transparent 25%);
  background-position: 10% 0;
}
.showDetails{
  &.order-ripped-teaser{
    display: none;
  }
  &.message-body{
    display:block;
  }
}
.tag{
  margin-bottom:.25em;
  margin-right:.25em;
}
@media screen and (min-width: 768px) {
  .order-icon,
  .showDetails.order-icon.details{
    display:none;
  }
  .message-header{
    cursor:auto;
    pointer-events:none;
  }
  .receipt-btn{
    pointer-events: auto;
    cursor: pointer;
  }
  .order-ripped-teaser{
    display:none;
  }
  .message-body{
    display:block;
  }
}
</style>

<script lang="ts" setup>
import type { BookingTotals } from '../../composables/booking';
import { Account } from '../../composables/account.ts';
import CartContent from '../../composables/cartContent.ts';
import { computed } from 'vue';

const props = defineProps<{
  totals: BookingTotals,
  accounts: Account[],
  contents: CartContent[],
  allowEdit: Boolean,
  timestamp?: string,
  bookingId?: number,
}>()

const dateString = computed(() => {
  if(!props.timestamp){
    return 
  }
  let date = new Date(props.timestamp);
  let dow = date.getDay();
  let dows = ['Sonntag', 'Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag', 'Samstag'];
  let day = date.getDate();
  let month = date.getMonth();
  let year = date.getFullYear();
  let hrs = date.getHours();
  let mnts = date.getMinutes();
  let scnds = date.getSeconds();
  let result = `${dows[dow]} ${day.toString().length == 1 ? '0' + day : day}.${(month+1).toString().length == 1 ? '0' + (month+1) : month}.${year} - ${hrs}:${mnts}:${scnds}`;
  return result;
})

</script>


<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCartStore } from '../../store/CartStore';
import CartList from './CartList.vue';
import Message from '../../composables/elements/Message.vue';
import AccountTagList from './AccountTagList.vue';
import CartSums from './CartSums.vue';
import CartControl from './CartControl.vue';
import ReceiptModal from '../Receipt/ReceiptModal.vue';


export default {
  data() {
    return {
      account$: useAccountStore(),
      showOrderDetails: false,
      showReceiptModal: false,
      cart$: useCartStore(),
    }
  },
  components: {
    CartList,
    Message,
    AccountTagList,
    CartSums,
    CartControl,
    ReceiptModal,
  },
  methods: {
    unselectAccount(account: Account){
      this.account$.selectSubstract(account);
    },
    toggleOrderDetails(){
      this.showOrderDetails = !this.showOrderDetails;
    },
    cancelOrder(){
      this.account$.unselect();
      this.cart$.cartContents = [];
    },
    checkoutOrder(){
      alert("new Order received");
      this.cancelOrder();
    }
  }
}
</script>