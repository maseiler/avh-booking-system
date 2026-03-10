<script lang="ts" setup>
import { ref } from 'vue';
import DetailsArea from '../../components/DetailsArea.vue';
import { useI18n } from 'vue-i18n';

const companyLogo = ref("");
function changeLogo(e){
  const reader = new FileReader();
  reader.addEventListener("load", () => {
    companyLogo.value = reader.result;
    console.log(companyLogo.value.length);
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

// const i18n = useI18n();
// console.log(i18n.availableLocales);

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
          <input type="text" class="input" placeholder="SOS Children's Villages">
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
          <input type="text" class="input" placeholder="Every child a home!">
          <span class="icon is-small is-left">
            <icon :icon="['fas', 'microphone']" />
          </span>
        </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Logo (.png below 100kb)
    </div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="file" class="input" accept="image/png, .svg" @change="changeLogo">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'image']" />
        </span>
      </p>
      <!-- Logo Preview -->
      <img width="150px" :src="companyLogo.valueOf()">

    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
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
          <option value="">Put all</option>
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
          <option value="">Put all</option>
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

  <!-- <div class="columns">
    <div class="column is-3">
      {{ $t('settings.currrency')}}
    </div>
    <div class="column is-2">
      <div class="select">
        <select v-model="$i18n.numberFormats">
          <option v-for="l of $i18n.availableLocales" :key="`locale-${l}`" :value="l">{{ l }}</option>
        </select>
      </div>
    </div>
  </div> -->
</DetailsArea>



</template>

<style scoped>
.columns{
  align-items: center;
}
</style>