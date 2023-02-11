import Vue from 'vue'
import Vuex from 'vuex'
import VueResource from 'vue-resource'
import helper from './helper.js'

Vue.use(Vuex)
Vue.use(VueResource);

const store = new Vuex.Store({
  state: {
    users: [],
    usersAsDict: {},
    items: [],
    lastNBookEntries: [],
    feedback: [],
    selectedUser: {},
    selectedSingleItem: {},
    selectedMultipleItems: []
  },
  mutations: {
    getUsers(state) {
      Vue.http.get("getUsers").then(response => {
        var users = [].concat.apply([], response.body);
        users = users.filter(Boolean);
        helper.sortUsersByName(users);
        state.users = users;
      })
    },
    getItems(state) {
      Vue.http.get("getAllItems").then(response => {
        var items = [].concat.apply([], response.body);
        items = items.filter(Boolean);
        helper.sortItemsByName(items);
        state.items = items;
      });
    },
    getLastNBookEntries(state, n) {
      Vue.http.post("getLastNBookEntries", n).then(response => {
        var bookEntries = [].concat.apply([], response.body);
        bookEntries = bookEntries.filter(Boolean);
        state.lastNBookEntries = bookEntries;
      });
    },
    getFeedbackList(state) {
      Vue.http.get("getFeedback").then(response => {
        var feedback = [].concat.apply([], response.body);
        feedback = feedback.filter(Boolean);
        state.feedback = feedback;
      });
    },
    selectUser(state, user) {
      if (state.selectedUser === user) {
        state.selectedUser = {};
      } else {
        state.selectedUser = user;
      }
    },
    selectSingleItem(state, item) {
      if (state.selectedSingleItem === item) {
        state.selectedSingleItem = {};
      } else {
        state.selectedSingleItem = item;
      }
    },
    selectMultipleItems(state, item) {
      state.selectedMultipleItems = state.selectedMultipleItems.concat(item);
    }
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
      return state.items.filter(function(item){
        if(item["Type"] === "alcoholic" && item["Enabled"] == 1){
          return true;
        }
        return false;
      })
    },
    itemsNonAlc: state => {
      return state.items.filter(function(item) {
        if(item["Type"] === "non-alcoholic" && item["Enabled"] == 1){
          return true;
        }
        return false;
      })
    },
    itemsFood: state => {
      return state.items.filter(function (item) {
        if(item["Type"] === "food" && item["Enabled"] == 1){
          return true;
        }
        return false;
      })
    }
  }
})

store.commit("getUsers");
store.commit("getItems");
store.commit("getFeedbackList");
store.commit("getLastNBookEntries", 5);

export default store
