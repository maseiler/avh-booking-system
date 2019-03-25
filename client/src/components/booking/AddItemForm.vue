<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h1>Add new item</h1>
            </slot>
          </div>

          <div class="modal-body">
            <slot name="body">
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
                  <input class="input" type="text" placeholder="Price" v-model.lazy="newItem.Price">
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="euro-sign"/>
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">Size</label>
                <div class="control has-icons-left">
                  <input class="input" type="text" placeholder="Size" v-model.lazy="newItem.Size">
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
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
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
              <div class="field is-grouped">
                <div class="control">
                  <button class="button is-link" @click="submitItem">Submit</button>
                </div>
                <div class="control">
                  <button class="button is-text" @click="cancelSubmission">Cancel</button>
                </div>
              </div>
            </slot>
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
        Price: "0",
        Size: "0",
        Unit: "",
        Type: ""
      },
      validationError: ""
    };
  },
  methods: {
    submitItem() {
      this.$http
        .post("/addItem", this.newItem)
        .then(function(response) {
          console.log(
            "Added new item:",
            this.newItem.Name,
            this.newItem.Price + "€",
            this.newItem.Size,
            this.newItem.Unit,
            this.newItem.Type
          );
          this.resetAndCloseForm();
          this.$router.go();
        })
        .catch(function(response) {
          this.validationError = response.data;
        });
    },
    cancelSubmission() {
      console.log("canceled submission");
      this.resetAndCloseForm();
    },
    resetAndCloseForm() {
      (this.Name = ""),
        (this.Price = ""),
        (this.Size = ""),
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