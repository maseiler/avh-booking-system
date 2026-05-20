<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { useCategoryStore } from '../../store/CategoryStore'
import { CategoryTypeToString, type Category } from '../../composables/category'
import { useSocketStore } from '../../store/socketStore'
import ToggleSwitch from '../../composables/elements/ToggleSwitch.vue'
import ErrorModal from '../../components/ErrorModal.vue'

const router = useRouter()
const category$ = useCategoryStore()
const socket$ = useSocketStore()

const pendingIds = ref<number[]>([])
const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

function editCategory(categoryId: number, type: string) {
  router.push({ name: 'CategorySettingsSingleEdit', params: { type, categoryId } })
}

function toggleEnabled(cat: Category) {
  if (pendingIds.value.includes(cat.id)) return
  pendingIds.value.push(cat.id)
  socket$.toggleCategoryEnabled(cat.id, !cat.enabled)
}

function onMutationResult(res: any) {
  if (res.table !== 'category') return
  const idx = pendingIds.value.indexOf(res.id)
  if (idx !== -1) pendingIds.value.splice(idx, 1)
}

function onWsError(err: any) {
  if (pendingIds.value.length === 0) return
  pendingIds.value = []
  currentError.value = err
  errorModalVisible.value = true
}

onMounted(() => {
  socket$.wsClient.on('mutationResult', onMutationResult)
  socket$.wsClient.on('wsError', onWsError)
})

onBeforeUnmount(() => {
  socket$.wsClient.off('mutationResult', onMutationResult)
  socket$.wsClient.off('wsError', onWsError)
})
</script>

<template>
  <h1 class="title">Kategorieeinstellungen</h1>

  <div class="columns">
    <div class="column">
      <h2 class="subtitle">Account-Kategorien</h2>
      <table class="table is-fullwidth is-striped is-hoverable">
        <thead>
          <tr>
            <th>Aktiv</th>
            <th>Kategorie</th>
            <th>Edit</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="cat in category$.accountCategories" :key="cat.id">
            <td>
              <div class="is-flex is-align-items-center">
                <ToggleSwitch
                  :model-value="cat.enabled"
                  :disabled="pendingIds.includes(cat.id)"
                  @update:model-value="toggleEnabled(cat)"
                />
                <span v-if="pendingIds.includes(cat.id)" class="icon has-text-grey ml-2">
                  <icon :icon="['fas', 'spinner']" :spin="true" />
                </span>
              </div>
            </td>
            <td>
              <span class="tag is-medium">
                <span class="icon is-small"><icon :icon="cat.icon" /></span>
                <span>{{ cat.title }}</span>
              </span>
            </td>
            <td>
              <button class="button is-small" @click="editCategory(cat.id, CategoryTypeToString[cat.type])">
                <span class="icon is-small"><icon :icon="['fas', 'pen']" /></span>
              </button>
            </td>
          </tr>
          <tr v-if="category$.accountCategories.length === 0">
            <td colspan="3" class="has-text-grey has-text-centered is-italic">
              Keine Account-Kategorien vorhanden
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="column">
      <h2 class="subtitle">Produkt-Kategorien</h2>
      <table class="table is-fullwidth is-striped is-hoverable">
        <thead>
          <tr>
            <th>Aktiv</th>
            <th>Kategorie</th>
            <th>Edit</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="cat in category$.productCategories" :key="cat.id">
            <td>
              <div class="is-flex is-align-items-center">
                <ToggleSwitch
                  :model-value="cat.enabled"
                  :disabled="pendingIds.includes(cat.id)"
                  @update:model-value="toggleEnabled(cat)"
                />
                <span v-if="pendingIds.includes(cat.id)" class="icon has-text-grey ml-2">
                  <icon :icon="['fas', 'spinner']" :spin="true" />
                </span>
              </div>
            </td>
            <td>
              <span class="tag is-medium">
                <span class="icon is-small"><icon :icon="cat.icon" /></span>
                <span>{{ cat.title }}</span>
              </span>
            </td>
            <td>
              <button class="button is-small" @click="editCategory(cat.id, CategoryTypeToString[cat.type])">
                <span class="icon is-small"><icon :icon="['fas', 'pen']" /></span>
              </button>
            </td>
          </tr>
          <tr v-if="category$.productCategories.length === 0">
            <td colspan="3" class="has-text-grey has-text-centered is-italic">
              Keine Produkt-Kategorien vorhanden
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
