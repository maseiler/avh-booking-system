import Vue from 'vue'
import Vuex from 'vuex'
import VueResource from 'vue-resource';

Vue.use(Vuex)
Vue.use(VueResource);

function sortUsersByName(array) {
  array.sort(function (a, b) {
    var bierNameA = a.BierName.toLowerCase(),
      bierNameB = b.BierName.toLowerCase();
    var firstNameA = a.FirstName.toLowerCase(),
      firstNameB = b.FirstName.toLowerCase();
    var lastNameA = a.LastName.toLowerCase(),
      lastNameB = b.LastName.toLowerCase();
    if (bierNameA < bierNameB) return -1;
    if (bierNameA > bierNameB) return 1;
    if (firstNameA < firstNameB) return -1;
    if (firstNameA > firstNameB) return 1;
    if (lastNameA < lastNameB) return -1;
    if (lastNameA > lastNameB) return 1;
    return 0;
  });
}

function sortItemsByName(array) {
  array.sort(function (a, b) {
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
}

const store = new Vuex.Store({
  state: {
    users: [],
    items: [],
    last5Bookings: [],
    feedback: []
  },
  mutations: {
    getUsers(state) {
      Vue.http.get("getUsers").then(response => {
        var users = [].concat.apply([], response.body);
        users = users.filter(Boolean);
        sortUsersByName(users);
        state.users = users;
      })
    },
    getItems(state) {
      Vue.http.get("getAllItems").then(response => {
        var items = [].concat.apply([], response.body);
        items = items.filter(Boolean);
        sortItemsByName(items);
        state.items = items;
      });
    },
    getLast5Bookings(state) {
      Vue.http.post("getLastNBookings", 5).then(response => {
        var bookings = [].concat.apply([], response.body);
        bookings = bookings.filter(Boolean);
        state.last5Bookings = bookings;
      });
    },
    getFeedbackList(state) {
      Vue.http.get("getFeedback").then(response => {
        var feedback = [].concat.apply([], response.body);
        feedback = feedback.filter(Boolean);
        state.feedback = feedback;
      });
    },
  },
  getters: {
    usersAH: state => {
      return state.users.filter(user => user["Status"] === "AH");
    },
    usersAktivB: state => {
      return state.users.filter(user => user["Status"] === "Aktiv B");
    },
    usersAktivKA: state => {
      return state.users.filter(user => user["Status"] === "Aktiv KA");
    },
    usersSteganleger: state => {
      return state.users.filter(user => user["Status"] === "Steganleger");
    },
    usersGast: state => {
      return state.users.filter(user => user["Status"] === "Gast");
    },
    itemsAlc: state => {
      return state.items.filter(item => item["Type"] === "alcoholic")
    },
    itemsNonAlc: state => {
      return state.items.filter(item => item["Type"] === "non-alcoholic")
    },
    itemsFood: state => {
      return state.items.filter(item => item["Type"] === "food")
    },
    itemsBoat: state => {
      return state.items.filter(item => item["Type"] === "boat")
    }
  }
})

store.commit("getUsers");
store.commit("getItems");
store.commit("getFeedbackList");
store.commit("getLast5Bookings");

export default store
