/**
 * Functions for JavaScript.
 * Functions for HTML can be found in mixin.js
 */
const helper = {
  ////////////////
  // User
  ////////////////
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
  displayUserNameFull(user) {
    var name = ""
    if (user === undefined) {
      return "???";
    } else {
      if (user.FirstName !== "") {
        name = name.concat(user.FirstName);
      }
      if (user.BierName !== "") {
        name = name.concat(" (", user.BierName, ")")
      }
      if (user.LastName !== "") {
        name = name.concat(" ", user.LastName)
      }
    }
    return name;
  },
  sortUsersByName(users) {
    users.sort((a, b) => {
      if (helper.displayUserName(a).toLowerCase() < helper.displayUserName(b).toLowerCase()) return -1;
      if (helper.displayUserName(a).toLowerCase() > helper.displayUserName(b).toLowerCase()) return 1;
      return 0;
    });
  },
  getUsersAsDict(users) {
    var dict = {};
    users.forEach(user => {
      var char;
      if (user.BierName !== "") {
        char = user.BierName[0].toUpperCase();
      } else {
        char = user.FirstName[0].toUpperCase();
      }
      var charCode = char.charCodeAt(0);
      if (charCode >= 65 && charCode <= 90) { // A-Z
      } else if (charCode >= 48 && charCode <= 57) { // 0-9
        char = "#";
      } else {
        char = "?";
      }
      if (dict[char] === undefined) {
        dict[char] = [user]
      } else {
        dict[char].push(user);
      }
    })
    return dict;
  },
  ////////////////
  // Item
  ////////////////
  getItemByID(items, id) {
    if (items !== undefined) {
      return items.find(i => {
        return i.ID == id;
      });
    }
  },
  displayItem(item) {
    if (item === undefined) {
      return "???"
    } else if (item.Type === "alcoholic" || item.Type === "non-alcoholic") {
      return item.Name + " " + item.Size + " " + item.Unit;
    } else if (item.Type === "food") {
      return item.Name;
    } else {
      return "???";
    }
  },
  sortItemsByName(items) {
    items.sort((a, b) => {
      var nameA = a["Name"].toLowerCase(),
        nameB = b["Name"].toLowerCase();
      var sizeA = a.Size,
        sizeB = b.Size;
      if (nameA < nameB) return -1;
      if (nameA > nameB) return 1;
      if (sizeA < sizeB) return -1;
      if (sizeA > sizeB) return 1;
      return 0;
    });
  },
  getItemsAsDict(items) {
    var dict = {};
    items.forEach(item => {
      var char = item.Name[0].toUpperCase();
      var charCode = char.charCodeAt(0);
      if (charCode >= 65 && charCode <= 90) { // A-Z
      } else if (charCode >= 48 && charCode <= 57) { // 0-9
        char = "#";
      } else {
        char = "?";
      }
      if (dict[char] === undefined) {
        dict[char] = [item]
      } else {
        dict[char].push(item);
      }
    })
    return dict;
  },
  getFavoriteItemsAsDict(items) {
    var dict = {};
    var i = 1;
    items.forEach(item => {
      dict[i] = [item]
      i++;
    })
    return dict;
  },
}

export default helper