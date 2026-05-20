<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAccountStore } from '../../store/AccountStore'
import { useCategoryStore } from '../../store/CategoryStore'
import { Account } from '../../composables/account'
import Buttons from '../../composables/elements/Buttons.vue'
import Button from '../../composables/elements/Button.vue'
import { useSocketStore } from '../../store/socketStore'
import ToggleSwitch from '../../composables/elements/ToggleSwitch.vue'
import ErrorModal from '../../components/ErrorModal.vue'

const route = useRoute()
const router = useRouter()
const account$ = useAccountStore()
const category$ = useCategoryStore()
const socket$ = useSocketStore()

const account = ref<Account>(new Account({} as Account))
const doneMounting = ref(false)
const showAdvanced = ref(false)
const saveStatus = ref<'idle' | 'pending' | 'success' | 'error'>('idle')

let saveStartTime = 0
let saveTimeoutId: number | null = null
let mutationHandler: ((res: any) => void) | null = null
let wsErrorHandler: ((err: any) => void) | null = null

const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

const isEdit = computed(() => (route.params.accountId?.toString().length ?? 0) > 0)
const categoryIcon = computed(() => category$.byId(account.value.category)?.icon)
const actionButton = computed(() => isEdit.value ? 'Speichern' : 'Neu Erstellen')
const maxDebt = computed({
  get: () => account.value.maxDebt / 100,
  set: (val: number) => { account.value.maxDebt = val * 100 }
})

onMounted(() => {
  if (isEdit.value) {
    account.value = account$.byId(parseInt(route.params.accountId.toString())).copy()
  } else {
    account.value = new Account({} as Account)
  }
  doneMounting.value = true
})

onBeforeUnmount(() => {
  cleanupSaveListeners()
})

function actionButtonClicked() {
  if (isEdit.value) {
    startSave()
    socket$.updateAccount(account.value)
    return
  }
  account.value.maxDebt = Math.floor(account.value.maxDebt)
  account.value.enabled = false
  account.value.balance = 0
  startSave()
  socket$.addAccount(account.value)
}

function cleanupSaveListeners() {
  if (mutationHandler) {
    socket$.wsClient.off('mutationResult', mutationHandler)
    mutationHandler = null
  }
  if (wsErrorHandler) {
    socket$.wsClient.off('wsError', wsErrorHandler)
    wsErrorHandler = null
  }
  if (saveTimeoutId !== null) {
    window.clearTimeout(saveTimeoutId)
    saveTimeoutId = null
  }
}

function startSave() {
  saveStatus.value = 'pending'
  saveStartTime = Date.now()

  mutationHandler = (res: any) => {
    if (res.table !== 'account') return
    if (isEdit.value && res.id !== account.value.id) return
    cleanupSaveListeners()

    const elapsed = Date.now() - saveStartTime
    window.setTimeout(() => {
      saveStatus.value = 'success'
      window.setTimeout(() => {
        router.push({ name: 'AccountSettings' })
      }, 500)
    }, Math.max(0, 500 - elapsed))
  }

  wsErrorHandler = (err: any) => {
    cleanupSaveListeners()
    saveStatus.value = 'idle'
    currentError.value = err
    errorModalVisible.value = true
  }

  socket$.wsClient.on('mutationResult', mutationHandler)
  socket$.wsClient.on('wsError', wsErrorHandler)

  saveTimeoutId = window.setTimeout(() => {
    cleanupSaveListeners()
    saveStatus.value = 'error'
    window.setTimeout(() => { saveStatus.value = 'idle' }, 500)
  }, 5000)
}
</script>


<template>
  <h1 class="title" v-if="doneMounting && isEdit">
    Edit Account:
    {{ account.getFullName() }}
  </h1>
  <h1 class="title" v-if="doneMounting && !isEdit">
    Neuen Account Erstellen
  </h1>
  <div class="columns" v-if="isEdit">
    <div class="column is-3">ID:</div>
    <div class="column">
      <input type="text" class="input" :value="account.id" disabled>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">First Name:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="account.firstName">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'user']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Last Name:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="account.lastName">
        <span class="icon is-small is-left">
        <icon :icon="['fas', 'id-card']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Nickname:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="account.nickname">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'user-secret']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">E-Mail:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="email" class="input" v-model="account.email">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'envelope']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Phone Number:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="tel" class="input" v-model="account.phone">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'phone']" />
        </span>
      </p>
    </div>
  </div>


  <div class="columns">
    <div class="column is-3">Category:</div>
    <div class="column">
      <div class="control has-icons-left">
        <div class="select">
          <select v-model="account.category">
            <option v-for="category in category$.accountCategories" :value="category.id" class="has-icons-left">
              {{ category.title }}
            </option>
          </select>
          </div>
           <div class="icon is-small is-left">
            <icon :icon="categoryIcon"/>
          </div>
        </div>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Max Debt Allowance:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="number" class="input" v-model="maxDebt" step="1">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'usd']" />
        </span>
      </p>
    </div>
  </div>

  <hr class="divider"></hr>
  <div class="columns">
    <div class="column is-3">Show Advanced Settings</div>
    <div class="column">
      <ToggleSwitch v-model="showAdvanced"> </ToggleSwitch>
    </div>
  </div>

  <div class="columns" v-if="showAdvanced">
    <div class="column is-3" title="Useful if you want an account to show up in one category, but it should not be restricted by the categories visibilities.">Always show all Products</div>
    <div class="column" >
      <ToggleSwitch>Not yet implemented</ToggleSwitch>
    </div>
  </div>

  <hr class="divider" v-if="showAdvanced"></hr>
  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <div class="is-flex is-align-items-center">
        <Buttons>
          <Button :fa-icon="['fas', 'times']" icon-position="left" @click="router.go(-1)">
            Cancel
          </Button>

          <Button class="is-primary" @click="actionButtonClicked" :fa-icon="['fas', 'save']" icon-position="right" :disabled="saveStatus === 'pending'">
            {{ actionButton }}
          </Button>
        </Buttons>
        <span v-if="saveStatus === 'pending'" class="ml-3 icon has-text-grey">
          <icon :icon="['fas', 'spinner']" :spin="true" />
        </span>
        <span v-else-if="saveStatus === 'success'" class="ml-3 icon-text has-text-success">
          <span class="icon"><icon :icon="['fas', 'check']" /></span>
          <span>Saved</span>
        </span>
        <span v-else-if="saveStatus === 'error'" class="ml-3 icon-text has-text-danger">
          <span class="icon"><icon :icon="['fas', 'times']" /></span>
          <span>Saving failed</span>
        </span>
      </div>
    </div>
  </div>

  <ErrorModal v-model="errorModalVisible" :error="currentError" />
</template>

<style scoped>
.dropdown-item.is-active {
  background-color: transparent;
  color: inherit;
  font-weight: 600;
  border-left: 3px solid hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l));
  padding-left: calc(1rem - 3px);
}
.columns{
  align-items: center;
}
</style>