<script setup lang="ts">
import { useBookingStore } from '../store/BookingStore';
import OrderViewer from '../components/OrderViewer/OrderViewer.vue';

const booking$ = useBookingStore();

</script>

<template>
  <h1 class="title">Order Viewer</h1>
  <div class="fixed-grid has-3-cols">
    <div class="grid">
      <OrderViewer class="cell" v-for="booking in booking$.bookings" :allowEdit="false" :accounts="booking.account" :totals="booking.getTotals()" :contents="booking.products" :timestamp="booking.timestamp"/>
    </div>
  </div>
</template>