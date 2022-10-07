<template>
    <div>
        <div class="columns">
            <div class="column">
                <UserSearch />
            </div>
            <div class="column is-2">
                <button class="button is-huette" @click="showAddUserForm = true">
                    New User
                </button>
            </div>
            <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false" />
        </div>
        <hr />
        <UserList />
    </div>
</template>

<script lang="ts">
import UserSearch from "./BaseUserSearch.vue"
import UserList from "./BaseUserList.vue"
import AddUserForm from "./BaseAddUserForm.vue"

export default {
    components: {
        UserSearch,
        UserList,
        AddUserForm,
    },
    data() {
        return {
            showAddUserForm: false
        }
    }
}
</script>