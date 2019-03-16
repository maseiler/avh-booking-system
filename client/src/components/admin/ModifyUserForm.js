import ModifyUserForm from "./ModifyUserForm.vue"

export default {
    components: {
        ModifyUserForm
    },
    props: {
        user: {}
    },
    data() {
        return {
            statusOptions: [ 'Aktiv B', 'Aktiv KA', 'AH', 'Gast'],
            selectedStatus: this.user.Status,
            validationError: ''
        }
    },
    methods: {
        modifyUser() {
            if (Object.keys(this.user).length === 0) {
                this.error = 'Please select a user first.'
                console.log('error')
            }
            else {
                this.$http.post('/modifyUser', this.user).then(function (response) {
                    console.log('Modified User: ' + this.user.BierName + ' ' + this.user.FirstName + ' ' + this.user.LastName + '.')
                    this.$router.go()
                }).catch(function (response) {
                    console.log('Error: Couldn\'t modify user.')
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