import UserTab from "./UserTab.vue"
import ItemList from "./ItemList.vue";
export default {
  components: {
    UserTab,
    ItemList,
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