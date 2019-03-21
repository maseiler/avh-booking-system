import UserTab from "./UserTab.vue"
import ItemTab from "./ItemTab.vue"

export default {
  components: {
    UserTab,
    ItemTab
  },
  data: function() {
    return {
      selectedUser: {},
      selectedItem: {},
      stringCheckout: "",
    };
  },
  methods: {
    getUser: function(user){
      this.selectedUser = user
    },
    getItem: function(item){
      this.selectedItem = item
    },
    checkout: function() {
      this.$http.get("/checkout").then(response => {
        this.stringCheckout = response.body;
      });
      console.log("checkout");
    },
  },
};