export default {
  data: function () {
    return {
      newUser: {
        bierName: '',
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        status: ''
      },
      validationError : '',
    };
  },
  methods: {
    submitUser() {
      this.$http.post('/addUser', this.newUser).then(function (response) {
        console.log("Added new user:", this.newUser.bierName, this.newUser.firstName, this.newUser.lastName, this.newUser.email, this.newUser.phone, this.newUser.status)
        this.resetAndCloseForm()
        this.$router.go()
        // todo reload page
      }).catch(function (response) {
        this.validationError = response.data
        // todo visible output
      });
    },
    cancelSubmission() {
      console.log("canceled submission")
      this.resetAndCloseForm()
    },
    resetAndCloseForm() {
      this.bierName = '',
        this.firstName = '',
        this.lastName = '',
        this.email = '',
        this.phone = '',
        this.status = '',
        this.validationError = ''
      this.$emit("close");
    },
  }
};