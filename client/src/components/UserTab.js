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
        }
    },
    created(){
        this.$nextTick(this.getUsers())
    }
};