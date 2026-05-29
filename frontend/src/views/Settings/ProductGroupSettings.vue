<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { useProductGroupStore } from '../../store/ProductGroupStore'
import { useSocketStore } from '../../store/socketStore'
import ErrorModal from '../../components/ErrorModal.vue'

const router = useRouter()
const productGroup$ = useProductGroupStore()
const socket$ = useSocketStore()

const pendingDeleteId = ref<number | null>(null)
const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

let deleteTimeoutId: number | null = null
let mutationHandler: ((res: any) => void) | null = null
let wsErrorHandler: ((err: any) => void) | null = null

function cleanupDeleteListeners() {
  if (mutationHandler) {
    socket$.wsClient.off('mutationResult', mutationHandler)
    mutationHandler = null
  }
  if (wsErrorHandler) {
    socket$.wsClient.off('wsError', wsErrorHandler)
    wsErrorHandler = null
  }
  if (deleteTimeoutId !== null) {
    window.clearTimeout(deleteTimeoutId)
    deleteTimeoutId = null
  }
}

function deleteGroup(id: number, name: string) {
  if (!window.confirm(`Produktgruppe "${name}" wirklich löschen?`)) return

  pendingDeleteId.value = id

  mutationHandler = (res: any) => {
    if (res.table !== 'product_group') return
    cleanupDeleteListeners()
    productGroup$.removeById(id)
    pendingDeleteId.value = null
  }

  wsErrorHandler = (err: any) => {
    cleanupDeleteListeners()
    pendingDeleteId.value = null
    currentError.value = err
    errorModalVisible.value = true
  }

  socket$.wsClient.on('mutationResult', mutationHandler)
  socket$.wsClient.on('wsError', wsErrorHandler)

  deleteTimeoutId = window.setTimeout(() => {
    cleanupDeleteListeners()
    pendingDeleteId.value = null
  }, 5000)

  socket$.deleteProductGroup(id)
}

onBeforeUnmount(() => {
  cleanupDeleteListeners()
})
</script>

<template>
  <h1 class="title">Produktgruppen</h1>

  <div class="panel">
    <p class="panel-heading has-text-primary-dark">Produktgruppen</p>
    <div class="panel-block">
      <table class="table is-fullwidth is-striped is-hoverable">
        <thead>
          <tr>
            <th>Name</th>
            <th class="has-text-right">Aktionen</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="group in productGroup$.all.filter(g => g.id !== 0)" :key="group.id">
            <td>{{ group.name }}</td>
            <td class="has-text-right">
              <div class="buttons is-right">
                <button
                  class="button is-small"
                  :disabled="pendingDeleteId === group.id"
                  @click="router.push({ name: 'ProductGroupSettingsEdit', params: { groupId: group.id } })">
                  <span class="icon is-small"><icon :icon="['fas', 'pen']" /></span>
                </button>
                <button
                  class="button is-small is-danger is-light"
                  :disabled="pendingDeleteId !== null"
                  :class="{ 'is-loading': pendingDeleteId === group.id }"
                  @click="deleteGroup(group.id, group.name)">
                  <span class="icon is-small"><icon :icon="['fas', 'trash']" /></span>
                </button>
              </div>
            </td>
          </tr>
          <tr v-if="productGroup$.all.filter(g => g.id !== 0).length === 0">
            <td colspan="2" class="has-text-grey has-text-centered is-italic">
              Keine Produktgruppen vorhanden
            </td>
          </tr>
          <tr>
            <td colspan="2" class="has-text-centered">
              <button class="button" @click="router.push({ name: 'ProductGroupSettingsAdd' })">
                <span class="icon"><icon :icon="['fas', 'plus']" /></span>
                <span>Produktgruppe hinzufügen</span>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

  <ErrorModal v-model="errorModalVisible" :error="currentError" />
</template>

<style scoped>
td {
  vertical-align: middle;
}
</style>
