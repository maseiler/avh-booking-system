<script setup lang="ts">
import Message from '../../composables/elements/Message.vue';
import { Booking } from '../../composables/booking';
import { computed } from 'vue';

const props = defineProps<{
  booking: Booking
}>()

const dateString = computed(() => {
  
  let date = new Date(props.booking.timestamp);
  let dow = date.getDay();
  let dows = ['Sonntag', 'Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag', 'Samstag'];
  let day = date.getDate();
  let month = date.getMonth();
  let year = date.getFullYear();
  let hrs = date.getHours();
  let mnts = date.getMinutes();
  let scnds = date.getSeconds();
  let result = `${dows[dow]} ${hrs}:${mnts}:${scnds} - ${day}.${month+1}.${year}`;
  return result;
})

</script>

<template>
  <Message>
    <template #header>
      Booking from {{ dateString }}
    </template>

    Booking for: {{ props.booking.account.getFullName() }} <br>
    Booking Total: {{ props.booking.getTotals()[0] }} <br>
    Booking Tax: {{ props.booking.getTotals()[1] }} <br>
    <div v-for="product in props.booking.products">
      Product: {{ product.product.name }} {{ product.product.size }} {{ product.product.unit }}
      Price: {{ product.quantity }}x * {{ product.price }} = {{ product.quantity*product.price/100 }}
    </div>
  </Message>
</template>