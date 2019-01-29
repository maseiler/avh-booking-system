// import AddUserForm from "./AddUserForm.vue";
// import ItemList from "./ItemList.vue";
// export default {
//   components: {
//     AddUserForm,
//     ItemList
//   },
//   data: function() {
//     return {
//       boolAddUser: false,
//       users: [],
//       items: [],
//       firstItem: ``,
//       stringCheckout: ""
//     };
//   },
//   computed:{
//     itemsAlc: function(){
//       return this.items.filter(function(item){
//         return item.Type == 'alcohol'
//       })
//     },
//     itemsNonAlc: function(){
//       return this.items.filter(function(item){
//         return item.Type == 'non-alcoholic'
//       })
//     },
//     itemsFood: function(){
//       return this.items.filter(function(item){
//         return item.Type == 'food'
//       })
//     },
//     itemsBoat: function(){
//       return this.items.filter(function(item){
//         return item.Type == 'boat'
//       })
//     }
//   },
//   methods: {
//     showAddUserForm() {
//       this.boolAddUser = true;
//       console.log("showAddUserForm");
//     },
//     closeAddUserForm() {
//       this.boolAddUser = false;
//       console.log("closeAddUserForm");
//     },
//     getUsers: function() {
//       this.$http.get("/users").then(response => {
//         this.users = response.body;
//       });
//       console.log("getUsers");
//     },
//     checkout: function() {
//       this.$http.get("/checkout").then(response => {
//         this.stringCheckout = response.body;
//       });
//       console.log("checkout");
//     },
//     getItems: function() {
//       this.$http.get("/items").then(response => {
//         this.items = response.body;
//       });
//       console.log("getItems");
//     },
//     // TODO remove
//     getFirstItem: function() {
//       this.$http.get("/items").then(response => {
//         this.items = response.body;
//         this.firstItem = response.body[0];
//         console.log(this.firstItem);
//       });
//     }
//   },
//   created(){
//     this.$nextTick(this.getItems())
//     // this.$nextTick(this.getUsers())
//   }
// };

import AddUserForm from "./AddUserForm.vue";
import ItemList from "./ItemList.vue";
export default {
  components: {
    AddUserForm,
    ItemList
  },
  data: function() {
    return {
      boolAddUser: false,
      users: [],
      stringCheckout: ""
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