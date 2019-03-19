import UserTab from "./UserTab.vue"
import ItemTab from "./ItemTab.vue"
import Way from "./Way.vue"

export default {
  components: {
    UserTab,
    ItemTab,
    Way
  },
  data: function() {
    return {
      stringCheckout: "",
    };
  },
  methods: {
    checkout: function() {
      this.$http.get("/checkout").then(response => {
        this.stringCheckout = response.body;
      });
      console.log("checkout");
    },
  },
};