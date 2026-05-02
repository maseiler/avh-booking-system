<script lang="ts" setup>
import { ref } from 'vue';
import DetailsArea from '../../components/DetailsArea.vue';
import { useSettingStore } from '../../store/SettingStore';
import { computed } from 'vue';

const setting$ = useSettingStore();

const companyName =  computed( {
  get() {
    return setting$.get("compTitle").value;
  },
  set(newValue: string) {
    setting$.set("compTitle", newValue);
  }
});

const companySlogan =  computed( {
  get() {
    return setting$.get("compSlogan").value;
  },
  set(newValue: string) {
    setting$.set("compSlogan", newValue);
  }
});

const companyLogo = computed({
  get(){
    if (setting$.get("compLogo") == -1){
      return null;
    }
    return setting$.get("compLogo").value;
  },
  set(newValue: string | ArrayBuffer | null) {
    if(newValue.length / 1024 > 200) {
      console.error("New Image uploaded with size (KB)", newValue.length / 1024, "This is too much. Reduce image Size so that is is below 200 KB.");
      return;
    }
    setting$.set("compLogo", newValue);
  }
});

function changeLogo(e){
  const reader = new FileReader();
  reader.addEventListener("load", () => {
    companyLogo.value = reader.result;
  })

  reader.readAsDataURL(e.target.files[0]);
}

const companyIcon = ref("");
function changeIcon(e){
  const reader = new FileReader();
  reader.addEventListener("load", () => {
    companyIcon.value = reader.result;
    console.log(companyIcon.value.length);
  })

  reader.readAsDataURL(e.target.files[0]);
}
</script>

<template>

<h1 class="title">General Settings</h1>

<DetailsArea>
  <template #summary>
    <h2 class="title is-4">Company Settings</h2>
  </template>

  <div class="columns">
    <div class="column is-3">
      Firmenname
    </div>
    <div class="column">
      <p class="control has-icons-left">
          <input type="text" v-model="companyName" class="input" placeholder="SOS Children's Villages">
          <span class="icon is-small is-left">
            <icon :icon="['fas', 'id-card']" />
          </span>
        </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Slogan
    </div>
    <div class="column">
      <p class="control has-icons-left">
          <input type="text" v-model="companySlogan" class="input" placeholder="Every child a home!">
          <span class="icon is-small is-left">
            <icon :icon="['fas', 'microphone']" />
          </span>
        </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3 has-start-align">
      Logo (.png/.svg below 200kb)
    </div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="file" class="input" accept="image/png, .svg" @change="changeLogo">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'image']" />
        </span>
      </p>
      <!-- Logo Preview -->
      <img width="150px" v-if="!(companyLogo == -1 || companyLogo == null)" :src="companyLogo">

    </div>
  </div>

  <div class="columns">
    <div class="column is-3 has-start-align">
      Icon
    </div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="file" class="input" accept="image/png, .svg" @change="changeIcon">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'image']" />
        </span>
      </p>
      <!-- Icon Preview -->
      <img width="50px ":src="companyIcon.valueOf()">
    </div>
  </div>

</DetailsArea>

<DetailsArea>
  <template #summary>
    <h2 class="title is-4">Application Look</h2>
  </template>

  <div class="columns">
    <div class="column is-3">
      Application Color Scheme (Dark)
    </div>
    <div class="column">
      <div class="select">
        <select>
          <option value="">Not yet implemented</option>
        </select>
      </div>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Application Color Scheme (Light)
    </div>
    <div class="column">
      <div class="select">
        <select>
          <option value="">Not yet implemented</option>
        </select>
      </div>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Primary Color
    </div>
    <div class="column is-2">
      <input type="color" class="input">
    </div>
  </div>

</DetailsArea>

<DetailsArea>
  <template #summary>
    <h2 class="title is-4">Localisation</h2>
  </template>

  <div class="columns">
    <div class="column is-3">
      {{ $t('settings.language')}}
    </div>
    <div class="column is-2">
      <div class="select">
        <select v-model="$i18n.locale">
          <option v-for="l of $i18n.availableLocales" :key="`locale-${l}`" :value="l">{{ $t(`settings.${l}`) }}</option>
        </select>
      </div>
    </div>
  </div>

  <!-- ToDo Allow seperate changing of currency -->


</DetailsArea>



</template>

<style scoped>
.columns{
  align-items: center;
}
.has-start-align {
  align-self: start;
  margin-top:.5em;
}
</style>