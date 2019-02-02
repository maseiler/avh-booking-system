import AddUserForm from "./AddUserForm.vue";
import ItemList from "./ItemList.vue";
export default {
  components: {
    AddUserForm,
    ItemList,
  },
  data: function() {
    return {
      showAddUserForm: false,
      users: [],
      stringCheckout: "",
    };
  },
  methods: {
    getUsers: function() {
      this.$http.get("/users").then(response => {
        this.users = response.body;
      });
      console.log("getUsers");
    },
    checkout: function() {
      this.$http.get("/checkout").then(response => {
        this.stringCheckout = response.body;
      });
      console.log("checkout");
    },
  },
};