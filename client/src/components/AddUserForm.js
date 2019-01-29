export default {
    data: function() {
      return {
        newUser:{
          bierName: '',
          firstName: '',
          lastName: '',
          email: '',
          phone: '',
          status: ''
        }
      };
    },
    methods: {
      submitUser(){
        this.$http.post('/newUser', this.newUser).then(function (response) {
          // Success
          console.log("Added new user:", this.newUser.bierName, this.newUser.firstName, this.newUser.lastName, this.newUser.email, this.newUser.phone, this.newUser.status)
        },function (response) {
            // Error
            console.log(response.data)
        });
        // TODO empty fields
      },
      cancelSubmission(){
        // TODO empty fields
        this.$emit("close");
        console.log("canceled submission")
      }
    }
  };