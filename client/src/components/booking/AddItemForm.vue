<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">
              {{ $t("booking.addItem.title") }}
              </h1>
            <hr />
          </div>

          <div class="modal-body">
            <div class="field">
              <label class="label">{{ $t("item.name") }}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('item.name')"
                  v-model.lazy="newItem.Name"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="font" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{ $t("item.price") }}</label>
              <div class="control has-icons-left">
                <input
                  class="input no-controls"
                  type="number"
                  v-bind:placeholder="$t('item.price')"
                  v-model.number="newItem.Price"
                  step=".01"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="euro-sign" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{ $t("item.size") }}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('item.size')"
                  v-model.number="newItem.Size"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="expand-arrows-alt" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{ $t("item.unit") }}</label>
              <div class="control">
                <div class="select">
                  <select v-model="newItem.Unit">
                    <option disabled value>{{ $t("item.unit") }}</option>
                    <option value="l">{{ $t("item.unitL") }}</option>
                    <option value="piece">{{ $t("item.unitPiece") }}</option>
                  </select>
                </div>
              </div>
            </div>

            <div class="field">
              <label class="label">{{ $t("item.type") }}</label>
              <div class="control">
                <div class="select">
                  <select v-model="newItem.Type">
                    <option disabled value>{{ $t("item.type") }}</option>
                    <option value="alcoholic">{{ $t("generic.alcoholic") }}</option>
                    <option value="non-alcoholic">{{ $t("generic.nonAlcoholic") }}</option>
                    <option value="food">{{ $t("generic.food") }}</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <article v-if="validationError !== ''" class="message is-danger">
              <div class="message-header">
                <div class="field is-grouped">
                  <p class="icon is-small is-left">
                    <font-awesome-icon icon="exclamation" size="lg" />
                  </p>
                  <p>{{ validationError }}</p>
                </div>
              </div>
            </article>
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="submitItem">
                    {{ $t("generic.submit") }}
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="resetAndCloseForm">
                    {{ $t("generic.cancel") }}
                  </button>
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
  data() {
    return {
      newItem: {
        Name: "",
        Price: 0,
        Size: 0,
        Unit: "",
        Type: "",
      },
      validationError: "",
    };
  },
  methods: {
    submitItem() {
      this.$http
        .post("addItem", this.newItem)
        .then(() => {
          let message = `${this.$t("messages.success.newItemAdded")}: ${this.displayItem(this.newItem)}`;
          this.resetAndCloseForm();
          this.$store.commit("getItems");
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch((response) => {
          //ToDo: Internationalize this message
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
    },
  },
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
