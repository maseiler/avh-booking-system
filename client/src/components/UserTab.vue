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
            <button class="button" is-primary v-for="user in usersAH" :key="user">{{ user.BierName }}</button>
        </p>
        <p v-if="activeTab ==='tab2'">
            <button class="button" is-primary v-for="user in usersAktivB" :key="user">{{ user.BierName }}</button>
        </p>
        <p v-if="activeTab ==='tab3'">
            <button class="button" is-primary v-for="user in usersAktivKA" :key="user">{{ user.BierName }}</button>
        </p>
        <p v-if="activeTab ==='tab4'">
            <button class="button" is-primary v-for="user in usersGaeste" :key="user">{{ user.FirstName + " " + user.LastName }}</button>
        </p>

        <br/>
        <div class="columns">
            <div class="column">
                <div class="field">
                    <div class="control has-icons-left">
                        <input class="input" type="text" placeholder="Search user" v-model="search" v-on:keyup="searchUsers">
                        <span class="icon is-small is-left">
                            <font-awesome-icon icon="search"/>
                        </span>
                    </div>
                </div>
            </div>
            <div class="column is-2">
                <button class="button is-link" @click="showAddUserForm = true">Add User</button>
            </div>
        </div>
        <button class="button" is-primary v-for="user in searchResults" :key="user">{{ user.BierName }}</button>
        
        <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false"></AddUserForm>
    </div>
</template>

<script src="./UserTab.js"></script>