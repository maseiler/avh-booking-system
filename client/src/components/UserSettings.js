import AddUserFormAdmin from "./AddUserFormAdmin.vue"
import ModifyUserForm from "./ModifyUserForm.vue"
import DeleteUserForm from "./DeleteUserForm.vue"
import UserInfo from "./UserInfo.vue"

export default {
  components: {
    AddUserFormAdmin,
    ModifyUserForm,
    DeleteUserForm,
    UserInfo
  }, data: function () {
    return {
      allUsers: [],
      showAddUserFormAdmin: false,
      showModifyUserForm: false,
      showDeleteUserForm: false,
      search: '',
      searchResults: [],
      selectedUser: {}
    };
  },
  methods: {
    getUsers: function () {
      this.$http.get("/getUsers").then(response => {
        var temp = response.body
        this.allUsers = [].concat.apply([], temp)
      });
    }, searchUsers: function () {
      if (this.search != '') {
        var tmpSearch = this.search.toLowerCase()
        this.searchResults = this.allUsers.filter(user => (user['BierName'].toLowerCase().includes(tmpSearch)) | (user['FirstName'].toLowerCase().includes(tmpSearch)) | (user['LastName'].toLowerCase().includes(tmpSearch)))
      } else {
        this.searchResults = []
      }
    },
    displayName: function (user) {
      if (user.BierName != '') {
        return user.BierName
      } else {
        if (user.LastName != '') {
          return user.FirstName + ' ' + user.LastName[0] + '.'
        }
        return user.FirstName
      }
    }
  }, created() {
    this.$nextTick(this.getUsers())
  }
}