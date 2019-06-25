import Vue from 'vue'

const mixin = Vue.mixin({
  methods: {
    displayUserName: function (user) {
      if (user === undefined) {
        return "???";
      } else
        if (user.BierName !== "") {
          return user.BierName;
        } else if (user.LastName !== "" && user.FirstName !== "") {
          return user.FirstName + " " + user.LastName[0] + ".";
        } else if (user.FirstName !== "") {
          return user.FirstName
        } else {
          return "???";
        }
    },
    displayItem: function (item) {
      if (item === undefined) {
        return "???"
      } else if (item.Type === "alcoholic" || item.Type === "non-alcoholic") {
        return item.Name + " " + item.Size + " " + item.Unit;
      } else if (item.Type === "boat" || item.Type === "food" || item.Type === "Admin") {
        return item.Name;
      } else {
        return "???";
      }
    }
  }
})

export default mixin