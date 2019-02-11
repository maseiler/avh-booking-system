import AddUserForm from "./AddUserForm.vue";

export default {
    components: {
      AddUserForm,
    },
    data: function() {
      return {
        showAddUserForm: false,
        allUsers: [],
        usersAH: [],
        usersAktivB: [],
        usersAktivKA: [],
        usersGaeste: [],
        activeTab: 'tab1',
        search: '',
        searchResults: []
      };
    },
    methods: {
        getUsers: function() {
            this.$http.get("/users").then(response => {
            var temp = response.body
            this.allUsers = [].concat.apply([], temp)
            this.usersAH = temp[0]
            this.usersAktivB = temp[1]
            this.usersAktivKA = temp[2]
            this.usersGaeste = temp[3]
            this.sortBy(this.usersAH, "BierName")
            this.sortBy(this.usersAktivB, "BierName")
            this.sortBy(this.usersAktivKA, "BierName")
            this.sortBy(this.usersGaeste, "FirstName")
            });
            console.log("getUsers");
        },
        sortBy: function(array, param){
            array.sort(function(a, b){
                var nameA=a[param].toLowerCase(), nameB=b[param].toLowerCase()
                if (nameA < nameB)
                    return -1 
                if (nameA > nameB)
                    return 1
                return 0
            })
        },
        searchUsers: function(){
            if(this.search != ''){
                var tmpSearch = this.search.toLowerCase()
                this.searchResults = this.allUsers.filter(user => (user['BierName'].toLowerCase().includes(tmpSearch)) | (user['FirstName'].toLowerCase().includes(tmpSearch)) | (user['LastName'].toLowerCase().includes(tmpSearch)))
                console.log(this.searchResults)
            } else {
                this.searchResults = []
            }
        }
    },
    created(){
        this.$nextTick(this.getUsers())
    }
};