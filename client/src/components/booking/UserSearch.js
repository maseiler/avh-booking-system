import UserSearch from "./UserSearch.vue";

export default {
    components: {
        UserSearch,
    },
    props: {
        allUsers: [],
    },
    data: function () {
        return {
            selectedUser: {},
            search: '',
            searchResults: []
        };
    },
    methods: {
        searchUsers: function () {
            if (this.search != '') {
                var tmpSearch = this.search.toLowerCase()
                this.searchResults = this.allUsers.filter(user => (user['BierName'].toLowerCase().includes(tmpSearch)) | (user['FirstName'].toLowerCase().includes(tmpSearch)) | (user['LastName'].toLowerCase().includes(tmpSearch)))
            } else {
                this.searchResults = []
            }
        },
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