import type { CartContent } from './cartContent'

// ── TSE Placeholder-Konstanten ────────────────────────────────────────────────
// Diese Werte müssen ersetzt werden, sobald die TSE-Tabelle im Backend vorliegt.
// Der Block ist bewusst kompakt gehalten damit Matthi ihn leicht findet.
const TSE_SERIAL_NUMBER       = 'BSI-TEST-TSE-00000000000000000000'
const TSE_SIGNATURE_ALGORITHM = 'ecdsa-plain-SHA256'
const TSE_LOG_TIME_FORMAT     = 'unixTime'
const TSE_PUBLIC_KEY          = 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA='
const TSE_CERTIFICATION_ID    = 'BSI-K-TR-XXXX-XXXX'
// ─────────────────────────────────────────────────────────────────────────────

type VatGroup = 'A' | 'B' | 'C' | 'D' | 'E'
const VAT_GROUP_ORDER: VatGroup[] = ['A', 'B', 'C', 'D']

export interface VatGroupEntry {
  group: VatGroup
  /** Steuersatz in Prozent, z.B. 19 oder 7 */
  rate: number
  /** Bruttobetrag in Cent (Nettobetrag + Steuerbetrag) */
  grossCents: number
  /** Steuerbetrag in Cent (aus dem Brutto herausgerechnet: gross × rate/(100+rate)) */
  taxCents: number
  /** Nettobetrag in Cent — das ist der DSFinV-K "Nennwert" (NW_x) im QR-Code */
  netCents: number
}

/**
 * Mappt beliebige MwSt-Sätze auf DSFinV-K-Gruppen A–E.
 *  - 0 % (oder undefined) → immer E
 *  - restliche Sätze absteigend sortiert → A, B, C, D
 *
 * Beispiele:
 *  [19, 7]       → A=19 %, B=7 %
 *  [21, 7]       → A=21 %, B=7 %   (nach Steuererhöhung)
 *  [19, 15, 7]   → A=19 %, B=15 %, C=7 %
 */
function mapRatesToGroups(rates: number[]): Map<number, VatGroup> {
  const map = new Map<number, VatGroup>()

  const nonZero = [...new Set(rates.filter(r => r > 0))].sort((a, b) => b - a)
  nonZero.forEach((rate, i) => {
    if (i < VAT_GROUP_ORDER.length) {
      map.set(rate, VAT_GROUP_ORDER[i])
    }
  })

  // 0 % (und undefined-Sätze, die zu 0 normalisiert werden) → E
  if (rates.some(r => r === 0)) {
    map.set(0, 'E')
  }

  return map
}

function formatEur(cents: number): string {
  return (cents / 100).toFixed(2)
}

function toUnixTime(iso: string): number {
  return Math.floor(new Date(iso).getTime() / 1000)
}

/**
 * SHA-256 der Kerntransaktionsdaten als Interim-Signatur.
 * Jeder Kassenbon bekommt einen eindeutigen, deterministischen Wert.
 * Wird durch die echte ECDSA-Signatur der TSE ersetzt.
 */
async function computeInterimSignature(input: string): Promise<string> {
  const buf = await crypto.subtle.digest('SHA-256', new TextEncoder().encode(input))
  return btoa(String.fromCharCode(...new Uint8Array(buf)))
}

export interface ReceiptInput {
  products: CartContent[]
  timestamp: string
  id?: number
}

/**
 * Berechnet die MwSt-Aufschlüsselung für die Anzeige im Kassenbon.
 * Gibt nur Gruppen zurück, die tatsächlich Umsatz haben.
 *
 * Hinweis zur Preisbasis: product.price ist der Bruttopreis (inkl. MwSt.).
 * CartContent.tax kann undefined sein (getVat()?.rate), wird dann als 0 % behandelt.
 */
export function computeVatBreakdown(products: CartContent[]): VatGroupEntry[] {
  // tax kann undefined sein → auf 0 normalisieren (Gruppe E)
  const rates = products.map(p => p.tax ?? 0)
  const rateToGroup = mapRatesToGroups(rates)

  const groups = new Map<VatGroup, VatGroupEntry>()

  for (const item of products) {
    const taxRate   = item.tax ?? 0                           // Steuersatz in %, z.B. 19
    const grossCents = item.price * item.quantity             // Brutto (inkl. MwSt.)
    // Steuer aus dem Bruttopreis herausrechnen: T = Brutto × rate / (100 + rate)
    const taxCents  = Math.round(grossCents * (taxRate / (100 + taxRate)))
    const netCents  = grossCents - taxCents                   // Netto (exkl. MwSt.)
    const group     = rateToGroup.get(taxRate) ?? 'E'

    const existing = groups.get(group)
    if (existing) {
      existing.grossCents += grossCents
      existing.taxCents   += taxCents
      existing.netCents   += netCents
    } else {
      groups.set(group, { group, rate: taxRate, grossCents, taxCents, netCents })
    }
  }

  // Sortierung A → E
  return [...groups.values()].sort((a, b) =>
    VAT_GROUP_ORDER.indexOf(a.group as any) - VAT_GROUP_ORDER.indexOf(b.group as any)
  )
}

/**
 * Erstellt den Kassenbeleg-V1-String für den QR-Code (BSI TR-03151).
 *
 * processData-Format (Feld innerhalb des QR-Strings):
 *   Beleg^[NW_A]_[NW_B]_[NW_C]_[NW_D]_[NW_E]^[Brutto]:Zahlungsart
 *
 *   NW_x       = Nettoumsatz der MwSt-Gruppe x (Brutto abzgl. MwSt.):
 *     Gruppe A: regulärer Steuersatz   (z.B. 19 %)
 *     Gruppe B: reduzierter Steuersatz (z.B.  7 %)
 *     Gruppe C: Durchschnittssatz nach §24 UStG Nr. 3
 *     Gruppe D: Durchschnittssatz nach §24 UStG Nr. 1
 *     Gruppe E: steuerfrei / 0 %       (z.B. Trinkgeld, Gutscheinaufladung)
 *   Zahlungsart: Gutsch. (Kontoguthaben) | Unbar (Kartenzahlung via Stripe)
 */
export async function buildKassenbelegV1(receipt: ReceiptInput): Promise<string> {
  const breakdown = computeVatBreakdown(receipt.products)

  // Gesamtbruttopreis
  const totalGrossCents = receipt.products.reduce(
    (sum, item) => sum + item.price * item.quantity, 0
  )

  // ── Nettoumsätze je MwSt-Gruppe ──────────────────────────────────────────
  // Alle 5 Gruppen A–E; Gruppen ohne Umsatz → 0.00
  const allGroups: VatGroup[] = ['A', 'B', 'C', 'D', 'E']
  const groupMap = new Map(breakdown.map(e => [e.group, e]))
  const groupNetAmounts = allGroups
    .map(g => {
      const entry = groupMap.get(g)
      return entry ? formatEur(entry.netCents) : '0.00'
    })
    .join('_')

  // ── Zahlungsart ──────────────────────────────────────────────────────────
  // Kontoguthaben-Buchungen → Gutsch. (Gutschein / Guthaben)
  // Kartenzahlung via Stripe          → Unbar
  const paymentLine = `${formatEur(totalGrossCents)}:Gutsch.`

  const processData = `Beleg^${groupNetAmounts}^${paymentLine}`

  // ── TSE-Daten ────────────────────────────────────────────────────────────
  const unixTime          = toUnixTime(receipt.timestamp)
  const transactionNumber = receipt.id ?? 0
  const signatureCounter  = receipt.id ?? 0

  // Interim-Signatur: SHA-256 über Kernfelder — einmalig und deterministisch
  const interimSignature = await computeInterimSignature(
    `${receipt.timestamp}|${transactionNumber}|${processData}`
  )

  const tseLine = [
    TSE_SERIAL_NUMBER,       // [2]  Seriennummer TSE      🔸 → TSE serialNumber
    'Kassenbeleg-V1',        // [3]
    processData,             // [4]
    transactionNumber,       // [5]  TransaktionsNummer    🔸 → TSE transactionNumber
    signatureCounter,        // [6]  SignaturZähler        🔸 → TSE signatureCounter
    unixTime,                // [7]  Start-Zeit            🔸 → TSE startTime
    unixTime,                // [8]  End-Zeit              🔸 → TSE endTime  (von FiskalCheck bestätigt)
    TSE_SIGNATURE_ALGORITHM, // [9]  Signaturalgorithmus  (von FiskalCheck bestätigt)
    TSE_LOG_TIME_FORMAT,     // [10] Log-Zeit-Format       (von FiskalCheck bestätigt)
    interimSignature,        // [11] Signatur              🔸 → TSE signatureBase64
    TSE_PUBLIC_KEY,          // [12] PublicKey             🔸 → TSE publicKey
  ].join(';')

  return `V0;${tseLine}`
}
