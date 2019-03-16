import ModifyItemForm from "./ModifyItemForm.vue"

export default {
    components: {
        ModifyItemForm
    },
    props: {
        item: {}
    },
    data: function () {
        return {
            validationError: ''
        }
    },
    methods: {
        modifyItem() {
            console.log(this.item)
            if (Object.keys(this.item).length === 0) {
                this.error = 'Please select a user first.'
            } else {
                this.$http.post('/modifyItem', this.item).then(function (response) {
                    console.log('Modified Item: ', this.item.Name, this.item.Price + '€', this.item.Size, this.item.Unit, this.item.Type)
                    this.$router.go()
                }).catch(function (response) {
                    this.validationError = response.data
                    console.log('Error: Couldn\'t modify item.')
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