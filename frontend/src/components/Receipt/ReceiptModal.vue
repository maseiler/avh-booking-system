<script setup lang="ts">
import { ref, computed, watchEffect } from 'vue'
import QRCode from 'qrcode'
import { buildKassenbelegV1, buildDigitalReceiptText, computeVatBreakdown } from '../../composables/useReceiptQrCode'
import type { CartContent } from '../../composables/cartContent'
import type { Account } from '../../composables/account'
import { useSettingStore } from '../../store/SettingStore'

const props = defineProps<{
  products: CartContent[]
  accounts: Account[]
  timestamp: string
  bookingId?: number
}>()

const emit = defineEmits<{
  close: []
}>()

const setting$ = useSettingStore()

const activeTab    = ref<'tse' | 'receipt'>('tse')
const qrDataUrl    = ref('')
const digitalQrUrl = ref('')
const isLoading    = ref(true)
const hasError     = ref(false)

const shopName = computed(() => {
  const s = setting$.get('compTitle')
  return s === -1 ? '[Shopname]' : String(s.value)
})

watchEffect(async () => {
  isLoading.value    = true
  hasError.value     = false
  qrDataUrl.value    = ''
  digitalQrUrl.value = ''

  try {
    // ── TSE-Code ────────────────────────────────────────────────────────────
    const tseString = await buildKassenbelegV1({
      products:  props.products,
      timestamp: props.timestamp,
      id:        props.bookingId,
    })
    qrDataUrl.value = await QRCode.toDataURL(tseString, {
      width:  256,
      margin: 2,
      color:  { dark: '#000000', light: '#ffffff' },
    })

    // ── Digitaler Kassenzettel ───────────────────────────────────────────────
    const receiptText = buildDigitalReceiptText({
      products:     props.products,
      timestamp:    props.timestamp,
      id:           props.bookingId,
      accountNames: props.accounts.map(a => a.getFullName()),
      shopName:     shopName.value,
    })
    digitalQrUrl.value = await QRCode.toDataURL(receiptText, {
      width:                256,
      margin:               2,
      errorCorrectionLevel: 'L',
      color:                { dark: '#000000', light: '#ffffff' },
    })
  } catch {
    hasError.value = true
  } finally {
    isLoading.value = false
  }
})

const vatBreakdown = computed(() => computeVatBreakdown(props.products))

const totalGross = computed(() =>
  props.products.reduce((sum, item) => sum + item.price * item.quantity, 0)
)

const totalTax = computed(() =>
  vatBreakdown.value.reduce((sum, e) => sum + e.taxCents, 0)
)

const dateString = computed(() => {
  if (!props.timestamp) return '—'
  return new Date(props.timestamp).toLocaleString('de-DE', {
    dateStyle: 'full',
    timeStyle: 'medium',
  })
})

const accountNames = computed(() =>
  props.accounts.map(a => a.getFullName()).join(', ')
)

function formatEur(cents: number): string {
  return (cents / 100).toLocaleString('de-DE', { style: 'currency', currency: 'EUR' })
}
</script>

<template>
  <div class="modal is-active">
    <div class="modal-background" @click="emit('close')" />
    <div class="modal-card receipt-card">

      <header class="modal-card-head">
        <p class="modal-card-title">
          <span class="icon"><icon :icon="['fas', 'receipt']" /></span>
          Kassenbon
        </p>
        <button class="delete" aria-label="close" @click="emit('close')" />
      </header>

      <section class="modal-card-body">

        <!-- Vorläufige-Signatur-Warnung -->
        <div class="notification is-warning is-light mb-4">
          <span class="icon"><icon :icon="['fas', 'triangle-exclamation']" /></span>
          <strong>Vorläufige Signatur</strong> QR-Code enthält noch keine zertifizierte TSE-Signatur.
          Nicht für gesetzeskonforme Belegausgabe verwenden.
        </div>

        <div class="receipt-layout">

          <!-- QR-Code -->
          <div class="qr-area">
            <div class="tabs is-boxed is-small mb-2">
              <ul>
                <li :class="{ 'is-active': activeTab === 'tse' }">
                  <a @click="activeTab = 'tse'">TSE-Code</a>
                </li>
                <li :class="{ 'is-active': activeTab === 'receipt' }">
                  <a @click="activeTab = 'receipt'">Kassenzettel</a>
                </li>
              </ul>
            </div>
            <div v-if="isLoading" class="qr-placeholder is-skeleton" />
            <div v-else-if="hasError" class="qr-placeholder has-text-danger has-text-centered">
              <icon :icon="['fas', 'circle-xmark']" size="3x" /><br>
              QR-Code konnte nicht generiert werden.
            </div>
            <template v-else>
              <img v-if="activeTab === 'tse'"
                   :src="qrDataUrl"
                   alt="TSE QR-Code"
                   class="qr-image" />
              <img v-else
                   :src="digitalQrUrl"
                   alt="Kassenzettel QR-Code"
                   class="qr-image" />
            </template>
            <p class="has-text-centered is-size-7 has-text-grey mt-1">
              {{ activeTab === 'tse' ? 'Kassenbeleg-V1 (TSE)' : 'Digitaler Kassenzettel' }}
            </p>
          </div>

          <!-- Belegdaten -->
          <div class="receipt-details">

            <table class="table is-fullwidth is-size-7 mb-2">
              <tbody>
                <tr v-if="accountNames">
                  <td class="has-text-grey">Konto</td>
                  <td>{{ accountNames }}</td>
                </tr>
                <tr>
                  <td class="has-text-grey">Datum</td>
                  <td>{{ dateString }}</td>
                </tr>
                <tr v-if="bookingId !== undefined">
                  <td class="has-text-grey">Beleg-Nr.</td>
                  <td>{{ bookingId }}</td>
                </tr>
              </tbody>
            </table>

            <!-- MwSt-Aufschlüsselung -->
            <table class="table is-fullwidth is-narrow is-size-7 mb-2">
              <thead>
                <tr>
                  <th>Gr.</th>
                  <th class="has-text-right">Satz</th>
                  <th class="has-text-right">Netto</th>
                  <th class="has-text-right">MwSt.</th>
                  <th class="has-text-right">Brutto</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="entry in vatBreakdown" :key="entry.group">
                  <td><strong>{{ entry.group }}</strong></td>
                  <td class="has-text-right">{{ entry.rate }} %</td>
                  <td class="has-text-right">{{ formatEur(entry.netCents) }}</td>
                  <td class="has-text-right">{{ formatEur(entry.taxCents) }}</td>
                  <td class="has-text-right">{{ formatEur(entry.grossCents) }}</td>
                </tr>
              </tbody>
              <tfoot>
                <tr>
                  <td colspan="3" class="has-text-grey is-size-7">inkl. MwSt. gesamt</td>
                  <td class="has-text-right"><strong>{{ formatEur(totalTax) }}</strong></td>
                  <td class="has-text-right"><strong>{{ formatEur(totalGross) }}</strong></td>
                </tr>
              </tfoot>
            </table>

          </div>
        </div>

      </section>

      <footer class="modal-card-foot is-justify-content-flex-end">
        <button class="button" @click="emit('close')">Schließen</button>
      </footer>

    </div>
  </div>
</template>

<style scoped>
.receipt-card {
  max-width: 680px;
  width: 95vw;
}
.receipt-layout {
  display: flex;
  gap: 1.5rem;
  align-items: flex-start;
  flex-wrap: wrap;
}
.qr-area {
  flex: 0 0 auto;
}
.qr-image {
  display: block;
  width: 192px;
  height: 192px;
  image-rendering: pixelated;
}
.qr-placeholder {
  width: 192px;
  height: 192px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.receipt-details {
  flex: 1 1 240px;
  min-width: 0;
}
.table td:first-child {
  white-space: nowrap;
  padding-right: 0.75rem;
}
.notification {
  padding: 0.6rem 1rem;
  font-size: 0.85rem;
}
</style>
