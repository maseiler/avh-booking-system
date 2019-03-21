import UserList from "./UserList.vue"

export default {
    components: {
        UserList
    },
    props: {
        allUsers: [],
    },
    data: function () {
        return {
            activeTab: 'tab0',
            selectedUser: {}
        };
    },
    computed: {
        usersAH: function () {
            return this.allUsers.filter(user => (user['Status'] == 'AH'))
        },
        usersAktivB: function () {
            return this.allUsers.filter(user => (user['Status'] == 'Aktiv B'))
        },
        usersAktivKA: function () {
            return this.allUsers.filter(user => (user['Status'] == 'Aktiv KA'))
        },
        usersGaeste: function () {
            return this.allUsers.filter(user => (user['Status'] == 'Gast'))
        }
    },
    methods: {
        displayName: function (user) {
            if (user.BierName != '') {
                return user.BierName
            } else {
                if (user.LastName != '') {
                    return user.FirstName + ' ' + user.LastName[0] + '.'
                }
                return user.FirstName
            }
        },
        selectUser: function (user) {
            this.selectedUser = user
            this.$emit('selectUser', user)
            this.$userEventBus.$emit('sendToBus', user);
        },
        receiveFromEventBus(user) {
            this.selectedUser = user
        }
    },
    created: function () {
        this.$userEventBus.$on('sendToBus', this.receiveFromEventBus);
    }
};