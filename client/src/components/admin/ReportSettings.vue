<template>
    <div>
      <br />
      <div class="tile is-ancestor is-narrow">
        <div class="tile is-parent is-vertical">
          <article class="tile is-child">
            <div class="box">
              <div class="columns">
                <div class="column is-2">
                  <div class="buttons">
                    <button
                      class="button is-link is-fullwidth"
                      @click="sendCurrentDebts"
                    >
                    {{$t('admin.reportSettings.sendCurrentDebts')}}
                    </button>
                  </div>
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
      };
    },
    methods: {
        sendCurrentDebts() {
            let date = new Date()
            let year = date.getFullYear();
            let month = '' + (date.getMonth() + 1);
            let day = '' + date.getDate();
            if (month.length < 2){
                month = "0" + month;
            }
            if (day.length < 2) {
                day = "0" + day;
            }
            this.$http
            .post("sendCurrentDebts", `${year}-${month}-${day}`)
            .then(() => {
                let message = `${this.$t('messages.success.emailSend')}`;
                this.$responseEventBus.$emit("successMessage", message);
            })
            .catch((response) => {
                if (response.data !== undefined) {
                //ToDo: Internationalize this failure Message
                this.validationError = response.data;
                }
            });
        },
    }
  };
  </script>
  