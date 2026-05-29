<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProductGroupStore } from '../../store/ProductGroupStore'
import { useSocketStore } from '../../store/socketStore'
import Buttons from '../../composables/elements/Buttons.vue'
import Button from '../../composables/elements/Button.vue'
import ErrorModal from '../../components/ErrorModal.vue'

const route = useRoute()
const router = useRouter()
const productGroup$ = useProductGroupStore()
const socket$ = useSocketStore()

const isEdit = computed(() => !!route.params.groupId)
const editId = computed(() => isEdit.value ? parseInt(route.params.groupId as string) : null)

const name = ref('')

const saveStatus = ref<'idle' | 'pending' | 'success' | 'error'>('idle')
const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

let saveStartTime = 0
let saveTimeoutId: number | null = null
let mutationHandler: ((res: any) => void) | null = null
let wsErrorHandler: ((err: any) => void) | null = null

onMounted(() => {
  if (isEdit.value && editId.value !== null) {
    const group = productGroup$.byId(editId.value)
    if (group) name.value = group.name
  }
})

onBeforeUnmount(() => {
  cleanupSaveListeners()
})

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

function save() {
  saveStatus.value = 'pending'
  saveStartTime = Date.now()

  mutationHandler = (res: any) => {
    if (res.table !== 'product_group') return
    cleanupSaveListeners()

    const elapsed = Date.now() - saveStartTime
    window.setTimeout(() => {
      saveStatus.value = 'success'
      window.setTimeout(() => {
        router.push({ name: 'ProductGroupSettings' })
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

  if (isEdit.value && editId.value !== null) {
    console.log(editId.value);
    socket$.updateProductGroup(editId.value, name.value.trim())
  } else {
    socket$.addProductGroup(name.value.trim())
  }
}
</script>

<template>
  <h1 class="title" v-if="isEdit">Produktgruppe bearbeiten</h1>
  <h1 class="title" v-else>Neue Produktgruppe erstellen</h1>

  <div class="columns">
    <div class="column is-3">Name:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="name" placeholder="Gruppenname">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'list']" />
        </span>
      </p>
    </div>
  </div>

  <hr class="divider">
  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <div class="is-flex is-align-items-center">
        <Buttons>
          <Button :fa-icon="['fas', 'times']" icon-position="left" @click="router.push({ name: 'ProductGroupSettings' })">
            Cancel
          </Button>
          <Button
            class="is-primary"
            @click="save"
            :fa-icon="['fas', 'save']"
            icon-position="right"
            :disabled="saveStatus === 'pending' || !name.trim()"
          >
            {{ isEdit ? 'Speichern' : 'Erstellen' }}
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
.columns {
  align-items: center;
}
</style>
