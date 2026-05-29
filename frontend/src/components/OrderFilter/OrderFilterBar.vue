<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useAccountStore } from '../../store/AccountStore'
import type { Account } from '../../composables/account'

const props = defineProps<{
  accountId: number | null
  dateFrom: string
  dateTo: string
}>()

const emit = defineEmits<{
  'update:accountId': [value: number | null]
  'update:dateFrom': [value: string]
  'update:dateTo': [value: string]
}>()

const account$ = useAccountStore()
const accountSearch = ref('')
const dropdownOpen = ref(false)
const dropdownEl = ref<HTMLElement | null>(null)

const filteredAccounts = computed(() => {
  const s = accountSearch.value.toLowerCase()
  const enabled = account$.accounts.filter(acc => acc.enabled)
  if (!s) return enabled
  return enabled.filter(acc =>
    acc.firstName.toLowerCase().includes(s) ||
    acc.lastName.toLowerCase().includes(s) ||
    (acc.nickname?.toLowerCase().includes(s) ?? false)
  )
})

const selectedAccount = computed<Account | null>(() =>
  props.accountId !== null ? (account$.byId(props.accountId) ?? null) : null
)

const hasActiveFilters = computed(() =>
  props.accountId !== null || props.dateFrom !== '' || props.dateTo !== ''
)

function selectAccount(acc: Account | null) {
  emit('update:accountId', acc?.id ?? null)
  dropdownOpen.value = false
  accountSearch.value = ''
}

function clearAll() {
  emit('update:accountId', null)
  emit('update:dateFrom', '')
  emit('update:dateTo', '')
}

function handleClickOutside(e: MouseEvent) {
  if (dropdownEl.value && !dropdownEl.value.contains(e.target as Node)) {
    dropdownOpen.value = false
  }
}

onMounted(() => document.addEventListener('click', handleClickOutside))
onBeforeUnmount(() => document.removeEventListener('click', handleClickOutside))
</script>

<template>
  <div class="box filter-bar">
    <div class="field is-grouped is-flex-wrap-wrap is-align-items-center">

      <!-- Account Filter Dropdown -->
      <div class="control">
        <div class="dropdown" :class="{ 'is-active': dropdownOpen }" ref="dropdownEl">
          <div class="dropdown-trigger">
            <button
              class="button"
              type="button"
              :class="{ 'is-primary': accountId !== null }"
              @click.stop="dropdownOpen = !dropdownOpen"
            >
              <span class="icon is-small"><icon :icon="['fas', 'user']" /></span>
              <span>{{ selectedAccount ? selectedAccount.getShortName() : 'Alle Accounts' }}</span>
              <span class="icon is-small"><icon :icon="['fas', 'angle-down']" /></span>
            </button>
          </div>
          <div class="dropdown-menu" role="menu">
            <div class="dropdown-content account-list">
              <div class="dropdown-item">
                <input
                  class="input is-small"
                  type="text"
                  v-model="accountSearch"
                  placeholder="Suchen…"
                  @click.stop
                />
              </div>
              <hr class="dropdown-divider" />
              <a
                class="dropdown-item"
                :class="{ 'is-active': accountId === null }"
                @click="selectAccount(null)"
              >
                Alle Accounts
              </a>
              <a
                v-for="acc in filteredAccounts"
                :key="acc.id"
                class="dropdown-item"
                :class="{ 'is-active': acc.id === accountId }"
                @click="selectAccount(acc)"
              >
                {{ acc.getFullName() }}
              </a>
            </div>
          </div>
        </div>
      </div>

      <!-- Date From -->
      <div class="control">
        <div class="field has-addons">
          <p class="control">
            <span class="button is-static">Von</span>
          </p>
          <p class="control">
            <input
              class="input"
              type="date"
              :value="dateFrom"
              :class="{ 'is-primary': dateFrom !== '' }"
              @input="emit('update:dateFrom', ($event.target as HTMLInputElement).value)"
            />
          </p>
        </div>
      </div>

      <!-- Date To -->
      <div class="control">
        <div class="field has-addons">
          <p class="control">
            <span class="button is-static">Bis</span>
          </p>
          <p class="control">
            <input
              class="input"
              type="date"
              :value="dateTo"
              :class="{ 'is-primary': dateTo !== '' }"
              @input="emit('update:dateTo', ($event.target as HTMLInputElement).value)"
            />
          </p>
        </div>
      </div>

      <!-- Reset -->
      <div class="control" v-if="hasActiveFilters">
        <button class="button is-light" type="button" @click="clearAll">
          <span class="icon"><icon :icon="['fas', 'times']" /></span>
          <span>Zurücksetzen</span>
        </button>
      </div>

    </div>
  </div>
</template>

<style scoped>
.filter-bar {
  padding: 0.75rem 1rem;
  margin-bottom: 1rem;
}
.account-list {
  max-height: 320px;
  overflow-y: auto;
}
</style>
