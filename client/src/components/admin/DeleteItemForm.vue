<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">{{$t('admin.itemSettings.deleteItemTitle')}}</h1>
            <hr />
          </div>

          <div class="modal-body">
            <article
              v-if="Object.keys(item).length === 0"
              class="message is-danger"
            >
              <div class="message-header">
                <div class="field is-grouped">
                  <p class="icon is-small is-left">
                    <font-awesome-icon icon="exclamation" size="lg" />
                  </p>
                  <p>{{$t('admin.itemSettings.selectFirst')}}</p>
                </div>
              </div>
            </article>
            <div v-else>
              <p>
                {{$t('admin.itemSettings.deleteShure')}}
                <b>{{ displayItem(item) }}</b
                >?
              </p>
            </div>
          </div>

          <div class="modal-footer">
            <div class="field is-grouped">
              <div class="control" v-if="Object.keys(item).length !== 0">
                <button class="button is-link" @click="deleteItem">
                  {{$t('generic.delete')}}
                </button>
              </div>
              <div class="control">
                <button class="button" @click="closeForm">{{$t('generic.cancel')}}</button>
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
  computed: {
    item() {
      return this.$store.state.selectedSingleItem;
    },
  },
  methods: {
    deleteItem() {
      if (Object.keys(this.item).length === 0) {
        this.error = this.$t('admin.itemSettings.selectFirst');
      } else {
        this.$http
          .post("deleteItem", this.item)
          .then(() => {
            let message = `${this.$t('deletedItem')}: ${this.displayItem(this.item)}`;
            this.closeForm();
            this.$store.commit("getItems");
            this.$store.commit("selectSingleItem", {});
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              this.$t('messages.failure.itemDelete')
            );
          });
      }
    },
    closeForm() {
      this.$emit("close");
    },
  },
};
</script>

<style lang="scss">
@import "../../assets/modal.css";
</style>
