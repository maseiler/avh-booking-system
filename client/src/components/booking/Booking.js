import UserTab from "./UserTab.vue"
import ItemTab from "./ItemTab.vue";

export default {
  components: {
    UserTab,
    ItemTab
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