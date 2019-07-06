<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Add new item</h1>
            <hr>
          </div>

          <div class="modal-body">
            <div class="field">
              <label class="label">Name</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Name" v-model.lazy="newItem.Name">
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="font"/>
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">Price</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Price" v-model.number="newItem.Price">
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="euro-sign"/>
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">Size</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Size" v-model.number="newItem.Size">
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="expand-arrows-alt"/>
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">Unit</label>
              <div class="control">
                <div class="select">
                  <select v-model="newItem.Unit">
                    <option disabled value>Unit</option>
                    <option>l</option>
                    <option>piece</option>
                  </select>
                </div>
              </div>
            </div>

            <div class="field">
              <label class="label">Type</label>
              <div class="control">
                <div class="select">
                  <select v-model="newItem.Type">
                    <option disabled value>Type</option>
                    <option>alcoholic</option>
                    <option>non-alcoholic</option>
                    <option>food</option>
                    <option>boat</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <article v-if="validationError !==''" class="message is-danger">
              <div class="message-header">
                <div class="field is-grouped">
                  <p class="icon is-small is-left">
                    <font-awesome-icon icon="exclamation" size="lg"/>
                  </p>
                  <p>{{validationError}}</p>
                </div>
              </div>
            </article>
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="submitItem">Submit</button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="resetAndCloseForm">Cancel</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data: function() {
    return {
      newItem: {
        Name: "",
        Price: 0,
        Size: 0,
        Unit: "",
        Type: ""
      },
      validationError: ""
    };
  },
  methods: {
    submitItem() {
      this.$http
        .post("addItem", this.newItem)
        .then(function(response) {
          var message = "".concat(
            "Added new item: ",
            this.displayItem(this.newItem)
          );
          this.resetAndCloseForm();
          this.$store.commit("getItems");
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch(function(response) {
          this.validationError = response.data;
        });
    },
    resetAndCloseForm() {
      (this.Name = ""),
        (this.Price = 0),
        (this.Size = 0),
        (this.Unit = ""),
        (this.Type = ""),
        (this.validationError = "");
      this.$emit("close");
    }
  }
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
