import Vue from 'vue'
import helper from './helper.js'

/**
 * Functions for HTML content.
 * Functions for JavaScript can be found in helper.js
 */
const mixin = Vue.mixin({
  methods: {
    ////////////////
    // User
    ////////////////
    getUserByID(id) {
      return this.$store.state.users.find(u => {
        return u.ID == id;
      });
    },
    displayUserName(user) {
      return helper.displayUserName(user);
    },
    displayUserNameFull(user) {
      return helper.displayUserNameFull(user);
    },
    selectUser(user) {
      this.$store.commit("selectUser", user);
    },
    ////////////////
    // Item
    ////////////////
    getItemByID(items, id) {
      return helper.getItemByID(items, id);
    },
    displayItem(item) {
      return helper.displayItem(item);
    },
    selectItem(item) {
      this.$store.commit("selectMultipleItems", item);
    },
    ////////////////
    // Misc
    ////////////////
    printDateTime(dateTime) {
      var d = new Date(Date.parse(dateTime));
      return (
        ("0" + d.getDate()).slice(-2) +
        "." +
        ("0" + (d.getMonth() + 1)).slice(-2) +
        "." +
        d.getFullYear() +
        " " +
        ("0" + d.getHours()).slice(-2) +
        ":" +
        ("0" + d.getMinutes()).slice(-2) +
        ":" +
        ("0" + d.getSeconds()).slice(-2)
      );
    },
    buttonColor(selectedUser, user) {
      if (selectedUser === user) {
        return "is-link";
      } else if (user.Balance >= user.MaxDebt) {
        return "is-danger";
      } else if (user.MaxDebt - user.Balance <= user.MaxDebt * 0.1) {
        return "is-warning";
      } else {
        return "";
      }
    }
  }
})

export default mixin