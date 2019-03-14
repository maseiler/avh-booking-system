import DeleteUserForm from "./DeleteUserForm.vue"

export default {
    components: {
        DeleteUserForm
    },
    props: {
        user: {}
    },
    methods: {
        deleteUser() {
            if (Object.keys(this.user).length === 0) {
                this.error = 'Please select a user first.'
            }
            else {
                this.$http.post('/deleteUser', this.user).then(function (response) {
                    console.log('Deleted User: ' + this.user.BierName + ' ' + this.user.FirstName + ' ' + this.user.LastName + '.')
                    this.$router.go()
                  }).catch(function (response) {
                    console.log('Error: Couldn\'t delete user.')
                    //TODO
                  });
                
            }
        },
        cancelSubmission() {
            console.log("Canceled submission.")
            this.$emit("close");
        }
    }
}