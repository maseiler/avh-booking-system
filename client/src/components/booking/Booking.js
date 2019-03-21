import UserTab from "./UserTab.vue"
import ItemTab from "./ItemTab.vue"

export default {
  components: {
    UserTab,
    ItemTab
  },
  data: function() {
    return {
      selectedItem: {},
      stringCheckout: "",
    };
  },
  methods: {
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