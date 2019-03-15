import DeleteItemForm from "./DeleteItemForm.vue"

export default {
    components: {
        DeleteItemForm
    },
    props: {
        item: {}
    },
    methods: {
        deleteItem() {
            if (Object.keys(this.item).length === 0) {
                this.error = 'Please select an item first.'
            }
            else {
                this.$http.post('/deleteItem', this.item).then(function (response) {
                    console.log('Deleted Item: ', this.item.Name, this.item.Price + '€' + this.item.Size + this.item.Unit + this.item.Type)
                    this.$router.go()
                  }).catch(function (response) {
                    console.log(response)
                    console.log('Error: Couldn\'t delete item.')
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