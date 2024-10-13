<template>
    <div>
      <br />
      <div class="tile is-ancestor is-narrow">
        <div class="tile is-parent is-vertical">
          <article class="tile is-child">
            <div class="box">
              <div class="columns">
                <div class="column">
                  <h3>{{$t('admin.reportSettings.eMailLabel')}}</h3>
                </div>
              </div>
              <div class="columns field">
                <div class="column is-2">
                  <label for="emailServerHost">{{$t('admin.reportSettings.eMailServer')}}</label>
                </div>
                <div class="column is-2">
                  <div class="control has-icons-left">
                    <input data-setting="EMailHost" id="emailServerHost" type="text" :placeholder="$t('admin.reportSettings.eMailServer')" class="input" v-model.lazy="eMailHost" @change="changeSetting">
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="server" />
                    </span>
                  </div>
                </div>
              </div>
              <div class="columns field">
                <div class="column is-2">
                  <label for="emailServerPort">{{$t('admin.reportSettings.eMailPort')}}</label>
                </div>
                <div class="column is-1">
                  <div class="control has-icons-left">
                    <input data-setting="EMailPort" id="emailServerPort" type="number" min="80" max="6000" step="1" class="input" v-model.lazy="eMailPort" @change="changeSetting">
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="sign" />
                    </span>
                  </div>
                </div>
              </div>
              <div class="columns field">
                <div class="column is-2">
                  <label for="emailUsername">{{$t('admin.reportSettings.eMailUser')}}</label>
                </div>
                <div class="column is-2">
                  <div class="control has-icons-left">
                    <input data-setting="EMailUser" id="emailUsername" type="email" :placeholder="$t('admin.reportSettings.eMailUser')" class="input" v-model.lazy="eMailUser" @change="changeSetting">
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="envelope" />
                    </span>
                  </div>
                </div>
              </div>
              <div class="columns field">
                <div class="column is-2">
                  <label for="emailPassword">{{$t('admin.reportSettings.eMailPass')}}</label>
                </div>
                <div class="column is-2">
                  <div class="control has-icons-left">
                    <input data-setting="EMailPass" id="emailPassword" type="password" class="input" v-model.lazy="eMailPass" @change="changeSetting">
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="key" />
                    </span>
                  </div>
                </div>
              </div>
              <div class="columns field">
                <div class="column is-2">
                  <label for="emailSenderName">{{$t('admin.reportSettings.eMailSender')}}</label>
                </div>
                <div class="column is-2">
                  <div class="control has-icons-left">
                    <input data-setting="EMailSender" id="emailSenderName" type="text" :placeholder="$t('admin.reportSettings.eMailSender')" class="input" v-model.lazy="eMailSender" @change="changeSetting">
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="user" />
                    </span>
                  </div>
                </div>
              </div>
              <div class="columns">
                <div class="column">
                  <h3>{{$t('admin.reportSettings.reportLabel')}}</h3>
                </div>
              </div>
              <div class="columns">
                <div class="column is-2">
                  <p>{{$t('admin.reportSettings.sendCurrentDebts')}}</p>
                </div>
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
        eMailHost: "",
        eMailPort: "",
        eMailUser: "",
        eMailPass: "",
        eMailSender: "",
      };
    },
    mounted: function () {
      this.$http
          .post("getSettings")
          .then((resp) => {
          let settings = resp.data;
          for (const setting of settings) {
              if(setting.Name == "EMailHost"){
                  this.eMailHost = setting.Value;
              }
              if(setting.Name == "EMailPort"){
                  this.eMailPort = setting.Value;
              }
              if(setting.Name == "EMailUser"){
                  this.eMailUser = setting.Value;
              }
              if(setting.Name == "EMailPass"){
                  this.eMailPass = setting.Value;
              }
              if(setting.Name == "EMailSender"){
                  this.eMailSender = setting.Value;
              }
          }
          });
    },
    methods: {
      changeSetting(e){
            let set = {
                "Name": e.target.dataset.setting,
                "Value": e.target.value
            }
            this.$http.post("updateSetting", set);
        },
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
  