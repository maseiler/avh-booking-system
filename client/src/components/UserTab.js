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
        activeTab: 'tab1'
      };
    },
    methods: {
        getUsers: function() {
            this.$http.get("/users").then(response => {
            this.allUsers = response.body
            this.usersAH = this.allUsers[0]
            this.usersAktivB = this.allUsers[1]
            this.usersAktivKA = this.allUsers[2]
            this.usersGaeste = this.allUsers[3]
            this.sortByBierName(this.usersAH)
            this.sortByBierName(this.usersAktivB)
            this.sortByBierName(this.usersAktivKA)
            this.sortByFirstName(this.usersGaeste)
            });
            console.log("getUsers");
            },
        sortByBierName: function(array){
            array.sort(function(a, b){
                var nameA=a.BierName.toLowerCase(), nameB=b.BierName.toLowerCase()
                if (nameA < nameB)
                    return -1 
                if (nameA > nameB)
                    return 1
                return 0
            })
        },
        sortByFirstName: function(array){
            array.sort(function(a, b){
                var nameA=a.FirstName.toLowerCase(), nameB=b.FirstName.toLowerCase()
                if (nameA < nameB) //sort string ascending
                    return -1 
                if (nameA > nameB)
                    return 1
                return 0
            })
        }
    },
    created(){
        this.$nextTick(this.getUsers())
    }
};