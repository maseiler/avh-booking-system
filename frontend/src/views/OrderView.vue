<script setup lang="ts">
import { useBookingStore } from '../store/BookingStore';
import BookingViewer from '../components/BookingViewer/BookingViewer.vue';

const booking$ = useBookingStore();

</script>

<template>
  <h1 class="h1">Order Viewer</h1>
  <BookingViewer v-for="booking in booking$.bookings" :booking="booking"></BookingViewer>

</template>