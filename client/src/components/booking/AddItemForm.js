export default {
  data: function () {
    return {
      newItem: {
        Name: '',
        Price: '0',
        Size: '0',
        Unit: '',
        Type: '',
      },
      validationError: '',
    };
  },
  methods: {
    submitItem() {
      this.$http.post('/addItem', this.newItem).then(function (response) {
        console.log("Added new item:", this.newItem.Name, this.newItem.Price + '€', this.newItem.Size, this.newItem.Unit, this.newItem.Type)
        this.resetAndCloseForm()
        this.$router.go()
      }).catch(function (response) {
        this.validationError = response.data
      });
    },
    cancelSubmission() {
      console.log("canceled submission")
      this.resetAndCloseForm()
    },
    resetAndCloseForm() {
      this.Name = '',
      this.Price = '',
      this.Size = '',
      this.Unit = '',
      this.Type = '',
      this.validationError = ''
      this.$emit("close");
    },
  }
};