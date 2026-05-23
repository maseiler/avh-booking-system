<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount, watch } from 'vue'
import type { Account } from '../../composables/account'
import { useAccountStore } from '../../store/AccountStore'
import { useCategoryStore } from '../../store/CategoryStore'
import { useSocketStore } from '../../store/socketStore'
import ToggleSwitch from '../../composables/elements/ToggleSwitch.vue'
import ErrorModal from '../ErrorModal.vue'

const props = defineProps<{
  accounts?: Account[]
}>()

const account$ = useAccountStore()
const category$ = useCategoryStore()
const socket$ = useSocketStore()

const dev = ref(false)
const accountsSorted = ref<Account[]>([])
const sortedTo = ref('')
const pendingIds = ref<number[]>([])
const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

const hasEditAccountRights = computed(() => {
  if (dev.value) {
    console.warn('User has elevated privileges to edit products because you are running this in development environment')
    return true
  }
  // ToDo Check if currently loged in user is allowed to edit products
  return false
})

function copyText(txt: string) {
  navigator.clipboard.writeText(txt)
}

function toggleEnabled(account: Account) {
  if (pendingIds.value.includes(account.id)) return
  pendingIds.value.push(account.id)
  socket$.toggleAccountEnabled(account.id, !account.enabled)
}

function onMutationResult(res: any) {
  if (res.table !== 'account') return
  const idx = pendingIds.value.indexOf(res.id)
  if (idx !== -1) pendingIds.value.splice(idx, 1)
}

function onWsError(err: any) {
  if (pendingIds.value.length === 0) return
  pendingIds.value = []
  currentError.value = err
  errorModalVisible.value = true
}

function sortFor(sortParam: string) {
  sortedTo.value = sortParam
  const list = props.accounts ?? []
  switch (sortParam) {
    case 'id':
      accountsSorted.value = [...list].sort((a, b) => a.id - b.id); break
    case 'enabled':
      accountsSorted.value = [...list].sort((a, b) => (b.enabled ? 1 : 0) - (a.enabled ? 1 : 0)); break
    case 'fn':
      accountsSorted.value = [...list].sort((a, b) => a.firstName.localeCompare(b.firstName)); break
    case 'nn':
      accountsSorted.value = [...list].sort((a, b) => a.nickname.localeCompare(b.nickname)); break
    case 'ln':
      accountsSorted.value = [...list].sort((a, b) => a.lastName.localeCompare(b.lastName)); break
    case 'mail':
      accountsSorted.value = [...list].sort((a, b) => a.email.localeCompare(b.email)); break
    case 'phone':
      accountsSorted.value = [...list].sort((a, b) => a.phone?.localeCompare(b.phone)); break
    case 'bal':
      accountsSorted.value = [...list].sort((a, b) => a.balance - b.balance); break
    case 'md':
      accountsSorted.value = [...list].sort((a, b) => a.maxDebt - b.maxDebt); break
    case 'cat':
      accountsSorted.value = [...list].sort((a, b) => a.category - b.category); break
    case 'create':
      accountsSorted.value = [...list].sort((a, b) => {
        const aDate = new Date(a.createdAt as string)
        const bDate = new Date(b.createdAt as string)
        return aDate > bDate ? -1 : 1
      }); break
  }
}

onMounted(() => {
  dev.value = import.meta.env.DEV
  accountsSorted.value = props.accounts ?? []
  socket$.wsClient.on('mutationResult', onMutationResult)
  socket$.wsClient.on('wsError', onWsError)
})

onBeforeUnmount(() => {
  socket$.wsClient.off('mutationResult', onMutationResult)
  socket$.wsClient.off('wsError', onWsError)
})

watch(() => props.accounts, (newList) => {
  accountsSorted.value = newList ?? []
  sortFor(sortedTo.value)
})
</script>

<template>
  <div class="table-container">
    <table class="table is-striped is-hoverable">
      <tbody>
        <tr>
          <th @click="sortFor('id')">
            ID
            {{ sortedTo == "id" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('enabled')">
            Enabled
            {{ sortedTo == "enabled" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('fn')">
            First Name
            {{ sortedTo == "fn" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('nn')">
            Nickname
          {{ sortedTo == "nn" ? "⯆" : "" }}</th>
          <th @click="sortFor('ln')">
            Last Name
            {{ sortedTo == "ln" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('mail')">
            E-Mail
          {{ sortedTo == "mail" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('phone')">
            Phone
          {{ sortedTo == "phone" ? "⯆" : "" }}</th>
          <th @click="sortFor('bal')" class="has-text-right">
            Balance
          {{ sortedTo == "bal" ? "⯆" : "" }}</th>
          <th @click="sortFor('md')" class="has-text-right">
            MaxDebt
          {{ sortedTo == "md" ? "⯆" : "" }}</th>
          <th @click="sortFor('cat')">
            Category
          {{ sortedTo == "cat" ? "⯆" : "" }}</th>
          <th @click="sortFor('create')">
            Created At
          {{ sortedTo == "create" ? "⯆" : "" }}</th>
          <th v-show="hasEditAccountRights">Edit</th>
        </tr>
        <tr :class="account$.selected.includes(account) ? 'is-primary' : ''" v-for="account in accountsSorted" :key="account.id" @click="account$.select(account)">
          <td>{{ account.id }}</td>
          <td>
              <div style="display:flex;align-items:center;gap:.4rem;">
                <ToggleSwitch
                  :model-value="account.enabled"
                  :disabled="pendingIds.includes(account.id)"
                  @update:model-value="toggleEnabled(account)"
                />
                <span v-if="pendingIds.includes(account.id)" class="icon has-text-grey">
                  <icon :icon="['fas', 'spinner']" :spin="true" />
                </span>
              </div>
            </td>
          <td>{{ account.firstName }}</td>
          <td>{{ account.nickname }}</td>
          <td>{{ account.lastName }}</td>
          <td class="has-copy-btn">{{ account.email }} <span class="icon is-small" @click="copyText(account.email)"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-copy-btn">{{ account.phone }} <span class="icon is-small"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-text-right">{{ $n(account.balance / 100, 'currency') }}</td>
          <td class="has-text-right">{{ $n(account.maxDebt / 100, 'currency') }}</td>
          <td>
            <button class="tag" :class="account.getCategory() == undefined? 'is-skeleton' : ''">
              <span class="icon"><icon :icon="account.getCategory()?.icon" /></span>
              <span>{{ account.getCategory()?.title }}</span>
            </button>  
          </td>
          <td>{{ new Date(account.createdAt).toLocaleString() }}</td>
          <td v-show="hasEditAccountRights">
            <button class="button">
              <router-link :to="{ name: 'AccountSettingsSingle', params: { accountId: account.id } }">
                <span class="icon"><icon :icon="['fas', 'pen']"/></span>
              </router-link>
              </button>
          </td>
        </tr>

      </tbody>
    </table>
  </div>

  <ErrorModal v-model="errorModalVisible" :error="currentError" />
</template>



<style scoped>
  .has-copy-btn{
    position:relative;

    .icon {
      position:absolute;
      left:100%;
      z-index:10;
      padding:.8em;
      background-color:rgba(0,0,0,.3);
      border-radius:var(--bulma-radius);
      visibility:hidden;
      pointer-events: none;
      cursor:pointer;
    }
    &:hover .icon {
      visibility: visible;
      pointer-events: all;
    }
  }
  th{
    cursor: pointer;
    &:hover{
      background-color: rgba(255, 255, 255, .2);
    }
  }
  td{
    vertical-align: middle;
  }
</style>