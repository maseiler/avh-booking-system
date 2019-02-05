import AddUserForm from "./AddUserForm.vue";

export default {
    components: {
      AddUserForm,
    },
    data: function() {
      return {
        showAddUserForm: false,
        users: [],
        activeTab: 'tab1'
      };
    },
    methods: {
      getUsers: function() {
        this.$http.get("/users").then(response => {
          this.users = response.body;
        });
        console.log("getUsers");
      },
    },
    created(){
        this.$nextTick(this.getUsers())
      }
  };