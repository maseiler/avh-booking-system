<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Delete Item</h1>
            <hr />
          </div>

          <div class="modal-body">
            <article v-if="Object.keys(item).length === 0" class="message is-danger">
              <div class="message-header">
                <div class="field is-grouped">
                  <p class="icon is-small is-left">
                    <font-awesome-icon icon="exclamation" size="lg" />
                  </p>
                  <p>Select item first.</p>
                </div>
              </div>
            </article>
            <div v-else>
              <p>
                Do you really want to delete
                <b>{{displayItem(item)}}</b>?
              </p>
            </div>
          </div>

          <div class="modal-footer">
            <div class="field is-grouped">
              <div class="control" v-if="Object.keys(item).length !== 0">
                <button class="button is-link" @click="deleteItem">Delete</button>
              </div>
              <div class="control">
                <button class="button" @click="closeForm">Cancel</button>
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
  props: {
    item: {}
  },
  methods: {
    deleteItem() {
      if (Object.keys(this.item).length === 0) {
        this.error = "Please select an item first.";
      } else {
        this.$http
          .post("deleteItem", this.item)
          .then(function(response) {
            var message = "".concat(
              "Deleted item: ",
              this.displayItem(this.item)
            );
            this.closeForm();
            this.$store.commit("getItems");
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch(function(response) {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't delete item."
            );
          });
      }
    },
    closeForm() {
      this.$emit("close");
    }
  }
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
