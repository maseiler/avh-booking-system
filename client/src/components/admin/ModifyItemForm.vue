<!-- TODO: not DRY. Same as AddItemForm -->
<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">{{$t('admin.itemSettings.modify')}}</h1>
            <hr />
          </div>

          <div class="modal-body">
            <div class="field">
              <label class="label">{{$t('item.name')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('item.name')"
                  v-model.lazy="item.Name"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="font" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{$t('item.price')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="number"
                  v-bind:placeholder="$t('item.price')"
                  v-model.number="item.Price"
                  step=".01"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="euro-sign" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{$t('item.size')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('item.size')"
                  v-model.number="item.Size"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="expand-arrows-alt" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{$t('item.unit')}}</label>
              <div class="control">
                <div class="select">
                  <select v-model="item.Unit">
                    <option disabled value>{{$t('item.unit')}}</option>
                    <option value="l">{{$t('item.unitL')}}</option>
                    <option value="piece">{{$t('item.unitPiece')}}</option>
                  </select>
                </div>
              </div>
            </div>

            <div class="field">
              <label class="label">{{$t('item.type')}}</label>
              <div class="control">
                <div class="select">
                  <select v-model="item.Type">
                    <option disabled value>{{$t('item.type')}}</option>
                    <option value="alcoholic">{{$t('generic.alcoholic')}}</option>
                    <option value="non-alcoholic">{{$t('generic.nonAlcoholic')}}</option>
                    <option value="food">{{$t('generic.food')}}</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
          <div class="field">
            <label class="label">{{$t('item.enabled')}}</label>
            <div class="control">
              <input type="checkbox" id="ItemEnabled" v-model="item.Enabled" class="cbSwitch">
              <label for="ItemEnabled"></label>
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
                  <button class="button is-link" @click="modifyItem">
                    {{$t('generic.submit')}}
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="closeForm">
                    {{$t('generic.cancel')}}
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
  computed: {
    item() {
      return this.$store.state.selectedSingleItem;
    },
  },
  data() {
    return {
      validationError: "",
    };
  },
  methods: {
    modifyItem() {
      //if(this.item.Enabled == true){this.item.Enabled = 1} else if(this.item.Enabled == false) {this.item.Enabled = 0}
      if (Object.keys(this.item).length === 0) {
        this.validationError = this.$t('admin.itemSettings.selectFirst');
      } else {
        this.$http
          .post("modifyItem", this.item)
          .then(() => {
            let message = `${this.$t('messages.success.modifiedItem')}: ${this.displayItem(this.item)}`;
            this.closeForm();
            this.$store.commit("getItems");
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch((response) => {
            if (response.data !== undefined) {
              //ToDo: Internationalize this failure Message
              this.validationError = response.data;
            } else {
              this.$store.commit("getItems");
              this.closeForm();
            }
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
