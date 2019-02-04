<template>
    <div>
        <p class="title">Users</p>
        <button @click="showAddUserForm = true">Add User</button>
        <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false"></AddUserForm>
    </div>
</template>

<script src="./UserTab.js"></script>