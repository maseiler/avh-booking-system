import Vue from 'vue'

const mixin = Vue.mixin({
  methods: {
    getUserByID(id) {
      return this.$store.state.users.find(u => {
        return u.UserID == id;
      });
    },
    getItemByID(id) {
      return this.$store.state.items.find(i => {
        return i.ItemID == id;
      });
    },
    displayUserName(user) {
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
    displayItem(item) {
      if (item === undefined) {
        return "???"
      } else if (item.Type === "alcoholic" || item.Type === "non-alcoholic") {
        return item.Name + " " + item.Size + " " + item.Unit;
      } else if (item.Type === "boat" || item.Type === "food" || item.Type === "Admin") {
        return item.Name;
      } else {
        return "???";
      }
    },
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
    }
  }
})

export default mixin