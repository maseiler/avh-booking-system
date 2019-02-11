<template>
    <div>
        <p class="title">Users</p>
        <!-- user list -->
        <div class="tabs">
            <ul>
                <li :class="[ activeTab === 'tab1' ? 'is-active' : '']"><a @click="activeTab='tab1'">Altherren</a></li>
                <li :class="[ activeTab === 'tab2' ? 'is-active' : '']"><a @click="activeTab='tab2'">Aktive B</a></li>
                <li :class="[ activeTab === 'tab3' ? 'is-active' : '']"><a @click="activeTab='tab3'">Aktive KA</a></li>
                <li :class="[ activeTab === 'tab4' ? 'is-active' : '']"><a @click="activeTab='tab4'">Gäste</a></li>
            </ul>
        </div>
        <div class="buttons" v-if="activeTab ==='tab1'">
            <button class="button" v-for="user in usersAH" :key="user">{{ displayName(user) }}</button>
        </div>
        <div class="buttons" v-if="activeTab ==='tab2'">
            <button class="button" v-for="user in usersAktivB" :key="user">{{ displayName(user) }}</button>
        </div>
        <div class="buttons" v-if="activeTab ==='tab3'">
            <button class="button" v-for="user in usersAktivKA" :key="user">{{ displayName(user) }}</button>
        </div>
        <div class="buttons" v-if="activeTab ==='tab4'">
            <button class="button" v-for="user in usersGaeste" :key="user">{{ user.FirstName + ' ' + user.LastName }}</button>
        </div>

        <br/>
        <!-- search bar -->
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
        <!-- search results -->
        <div class="buttons" v-if="searchResults.filter(user => user['BierName'] != '')">
            <button class="button is-link" v-for="user in searchResults.filter(user => user['BierName'] != '')" :key="user">{{ user.BierName }}</button>
        </div>
        <div class="buttons" v-if="searchResults.filter(user => user['BierName'] == '')">
            <button class="button is-link" v-for="user in searchResults.filter(user => user['BierName'] == '')" :key="user">{{ user.FirstName + ' ' + user.LastName }}</button>
        </div>
        <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false"></AddUserForm>
    </div>
</template>

<script src="./UserTab.js"></script>