<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useBookingStore } from '../store/BookingStore'
import OrderViewer from '../components/OrderViewer/OrderViewer.vue'
import OrderFilterBar from '../components/OrderFilter/OrderFilterBar.vue'

const booking$ = useBookingStore()
const route = useRoute()
const router = useRouter()

// Filter state
const filterAccountId = ref<number | null>(null)
const filterDateFrom = ref('')
const filterDateTo = ref('')

// Pre-fill account filter from ?account=<id> URL param
onMounted(() => {
  const param = route.query.account
  const raw = Array.isArray(param) ? param[0] : param
  if (raw) {
    const id = Number(raw)
    if (!isNaN(id)) filterAccountId.value = id
  }
})

// Keep URL in sync whenever the account filter changes
watch(filterAccountId, (id) => {
  router.replace({
    query: {
      ...route.query,
      account: id !== null ? String(id) : undefined,
    },
  })
})

// Filtered bookings
const filteredBookings = computed(() => {
  let foundBookings = booking$.bookings.filter((booking) => {
    // Account filter
    if (filterAccountId.value !== null) {
      const accounts = Array.isArray(booking.account) ? booking.account : [booking.account]
      const matches = accounts.some((a) => a.id === filterAccountId.value)
      if (!matches) return false
    }

    // Date from — include from start of that day (local time)
    if (filterDateFrom.value) {
      const from = new Date(filterDateFrom.value + 'T00:00:00')
      if (new Date(booking.timestamp) < from) return false
    }

    // Date to — include until end of that day (local time)
    if (filterDateTo.value) {
      const to = new Date(filterDateTo.value + 'T23:59:59')
      if (new Date(booking.timestamp) > to) return false
    }

    return true
  })
  let sortedBookings = foundBookings.sort((a, b) => {
    return Date.parse(a.timestamp) - Date.parse(b.timestamp) <= 0 ? 1 : -1;
  })
  return sortedBookings;
})
</script>

<template>
  <h1 class="title">Order Viewer</h1>

  <OrderFilterBar
    :account-id="filterAccountId"
    :date-from="filterDateFrom"
    :date-to="filterDateTo"
    @update:account-id="filterAccountId = $event"
    @update:date-from="filterDateFrom = $event"
    @update:date-to="filterDateTo = $event"
  />

  <div v-if="filteredBookings.length === 0" class="notification is-light">
    <span class="icon"><icon :icon="['fas', 'inbox']" /></span>
    Keine Bestellungen gefunden.
  </div>

  <div class="fixed-grid has-3-cols">
    <div class="grid">
      <OrderViewer
        class="cell"
        v-for="booking in filteredBookings"
        :key="booking.timestamp"
        :allowEdit="false"
        :accounts="booking.account"
        :totals="booking.getTotals()"
        :contents="booking.products"
        :timestamp="booking.timestamp"
        :booking-id="booking.id"
      />
    </div>
  </div>
</template>
