<template>
    <div>
      <br />
      <div class="tile is-ancestor is-narrow">
        <div class="tile is-parent is-vertical">
          <article class="tile is-child">
            <div class="row">
                <label for="stripeAPI">Stripe Payment API:</label>
                <input id="stripeAPI" class="input" type="text" placeholder="Stripe API-Key" v-model.lazy="stripeAPIKey" @change="changeAPI" />
            </div>
            <div class="row">
                <label for="stripeCardReader">Card Reader:</label>
                <select id="stripeCardReader" v-model.lazy="selectedCardReader" @change="changeCardReader">
                    <option value="tmr_Fjr3RwpN5Jg1xD">Simulierter Test-Leser</option>
                    <option
                        v-for="reader in availableCardReaders"
                        :key="reader.data.id"
                        :value="reader.data.id"
                        >
                        {{ reader.data.label }}
                    </option>
                </select>
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
                this.availableCardReaders = resp.data;
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
  