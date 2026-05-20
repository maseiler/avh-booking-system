<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { CategoryType, CategoryTypeToString, StringToCategoryType } from '../../composables/category'
import { useCategoryStore } from '../../store/CategoryStore'
import Buttons from '../../composables/elements/Buttons.vue'
import Button from '../../composables/elements/Button.vue'
import { useSocketStore } from '../../store/socketStore'
import ErrorModal from '../../components/ErrorModal.vue'

const route = useRoute()
const router = useRouter()
const socket$ = useSocketStore()
const category$ = useCategoryStore()

const categoryType = computed(() => StringToCategoryType[route.params.type as string])
const typeLabel = computed(() => categoryType.value === CategoryType.ACCOUNT ? 'Account' : 'Produkt')
const isEdit = computed(() => !!route.params.categoryId)
const editId = computed(() => isEdit.value ? parseInt(route.params.categoryId as string) : null)
const backRoute = computed(() => {
  if (isEdit.value) return { name: 'CategorySettings' }
  return categoryType.value === CategoryType.ACCOUNT
    ? { name: 'AccountSettings' }
    : { name: 'ProductSettings' }
})

const name = ref('')
const iconName = ref('')
const iconPreview = computed<[string, string]>(() => ['fas', iconName.value.trim() || 'circle-question'])

const saveStatus = ref<'idle' | 'pending' | 'success' | 'error'>('idle')

let saveStartTime = 0
let saveTimeoutId: number | null = null
let mutationHandler: ((res: any) => void) | null = null
let wsErrorHandler: ((err: any) => void) | null = null

const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

onMounted(() => {
  if (isEdit.value && editId.value !== null) {
    const cat = category$.byId(editId.value)
    if (cat) {
      name.value = cat.title
      iconName.value = Array.isArray(cat.icon) ? cat.icon[1] : (typeof cat.icon === 'string' ? cat.icon : '')
    }
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
    if (res.table !== 'category') return
    if (isEdit.value && res.id !== editId.value) return
    cleanupSaveListeners()

    const elapsed = Date.now() - saveStartTime
    window.setTimeout(() => {
      saveStatus.value = 'success'
      window.setTimeout(() => {
        router.push(backRoute.value)
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
    socket$.updateCategory({
      id: editId.value,
      name: name.value.trim(),
      icon: ['fas', iconName.value.trim()],
      type: CategoryTypeToString[categoryType.value]
    })
  } else {
    socket$.addCategory({
      name: name.value.trim(),
      icon: ['fas', iconName.value.trim()],
      type: CategoryTypeToString[categoryType.value]
    })
  }
}
</script>

<template>
  <h1 class="title" v-if="isEdit">{{ typeLabel }}-Kategorie bearbeiten</h1>
  <h1 class="title" v-else>Neue {{ typeLabel }}-Kategorie erstellen</h1>

  <div class="columns">
    <div class="column is-3">Name:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="name" placeholder="Kategoriename">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'tag']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Icon:</div>
    <div class="column">
      <div class="is-flex is-align-items-center">
        <div class="control has-icons-left">
          <input type="text" class="input" v-model="iconName" placeholder="z.B. user, star, home">
          <span class="icon is-small is-left">
            <icon :icon="['fas', 'image']" />
          </span>
        </div>
        <span class="ml-4 icon is-large" :class="iconName.trim() ? 'has-text-primary' : 'has-text-grey-light'">
          <icon :icon="iconPreview" size="2x" />
        </span>
      </div>
    </div>
  </div>

  <hr class="divider">
  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <div class="is-flex is-align-items-center">
        <Buttons>
          <Button :fa-icon="['fas', 'times']" icon-position="left" @click="router.push(backRoute)">
            Cancel
          </Button>
          <Button
            class="is-primary"
            @click="save"
            :fa-icon="['fas', 'save']"
            icon-position="right"
            :disabled="saveStatus === 'pending' || !name.trim() || !iconName.trim()"
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
