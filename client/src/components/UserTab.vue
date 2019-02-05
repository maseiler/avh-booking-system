<template>
    <div>
        <p class="title">Users</p>
        <div class="tabs">
            <ul>
                <li :class="[ activeTab === 'tab1' ? 'is-active' : '']"><a @click="activeTab='tab1'">Altherren</a></li>
                <li :class="[ activeTab === 'tab2' ? 'is-active' : '']"><a @click="activeTab='tab2'">Aktive B</a></li>
                <li :class="[ activeTab === 'tab3' ? 'is-active' : '']"><a @click="activeTab='tab3'">Aktive KA</a></li>
                <li :class="[ activeTab === 'tab4' ? 'is-active' : '']"><a @click="activeTab='tab4'">Gäste</a></li>
            </ul>
        </div>

        <p v-if="activeTab ==='tab1'">
            <button class="button" is-primary v-for="user in users" :key="user">{{ user }}</button>
        </p>
        <p v-if="activeTab ==='tab2'">
            tab 2 content
        </p>
        <p v-if="activeTab ==='tab3'">
            tab 3 content
        </p>
        <p v-if="activeTab ==='tab4'">
            tab 4 content
        </p>
        <div class="field">
            <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Search" >
                <span class="icon is-small is-left">
                    <font-awesome-icon icon="search"/>
                </span>
            </div>
        </div>
        <button @click="showAddUserForm = true">Add User</button>
        <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false"></AddUserForm>
    </div>
</template>

<script src="./UserTab.js"></script>