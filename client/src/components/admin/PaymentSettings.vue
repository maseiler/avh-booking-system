<template>
    <div>
      <br />
      <div class="tile is-ancestor is-narrow has-text-left">
        <div class="tile is-parent is-vertical">
          <article class="tile is-child">
            <div class="field">
                <label for="stripeAPI" class="label">{{ $t("admin.paymentSettings.stripeAPILabel") }}:</label>
                <div class="control has-icons-left">
                  <input id="stripeAPI" class="input" type="text" :placeholder="$t('admin.paymentSettings.stripeAPILabel')" v-model.lazy="stripeAPIKey" @change="changeAPI" />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="key" />
                  </span>
                </div>
            </div>
            <div class="field">
                <label for="stripeCardReader" class="label">{{ $t("admin.paymentSettings.stripeCardReaderLabel") }}:</label>
                <div class="control has-icons-left">
                  <div class="select is-normal">
                    <select id="stripeCardReader" v-model.lazy="selectedCardReader" @change="changeCardReader">
                        <!-- <option value="tmr_Fjr3RwpN5Jg1xD">Simulierter Test-Leser</option> -->
                        <option
                            v-for="reader in availableCardReaders"
                            :key="reader.id"
                            :value="reader.id"
                            >
                            {{ reader.label }}
                        </option>
                    </select>
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="cash-register" />
                    </span>
                  </div>
                </div>
            </div>
          </article>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    components: {
    },
    data() {
      return {
        stripeAPIKey: "",
        availableCardReaders: [],
        selectedCardReader: ""
      };
    },
    mounted: function () {
        this.$http
            .post("getSettings")
            .then((resp) => {
            let settings = resp.data;
            for (const setting of settings) {
                if(setting.Name == "StripeAPIKey"){
                    this.stripeAPIKey = setting.Value;
                }
            }
            });
        this.$http.post("getStripeCardReader").then(
            (resp) => {
                this.availableCardReaders = resp.data.data;
                console.log(this.availableCardReaders)
            }
        )
        this.selectedCardReader = localStorage.getItem("StripeCardReader");
    },
    methods: {
        changeAPI(){
            let set = {
                "Name": "StripeAPIKey",
                "Value": this.stripeAPIKey
            }
            this.$http.post("updateSetting", set);
        },
        changeCardReader(){
            localStorage.setItem("StripeCardReader", this.selectedCardReader)
        }
    }
  };
  </script>
  