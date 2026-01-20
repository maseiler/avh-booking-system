<script setup lang="ts">
import { computed, useId } from 'vue';

const inputId = useId();
const props = defineProps<{
  modelValue: boolean,
  disabled?: boolean
}>()

const checkboxValue = computed({
  get() {
    return props.modelValue;
  },
  set(newValue: boolean) {
    emit('update:modelValue', newValue);
  }
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>();

</script>

<template>
  <input type="checkbox" :id="inputId" class="switch" v-model="checkboxValue" :disabled="props.disabled"/>
  <label :for="inputId">
    <slot></slot>
  </label>
</template>

<style scoped>
  .switch,
.switch + label{
  --switch-width:2em;
  --off-color-bg:red;
  --off-color-dot:pink;
  --on-color-bg:green;
  --on-color-dot:lime;
  --disabled-color-bg:grey;
  --disabled-color-dot:silver;
  user-select:none;
  cursor:pointer;
}

.switch{
  position: relative; 
  width:0px;
  opacity:0;
  transform:translateX(calc(var(--switch-width)/2));
  z-index:-1;
}
.switch + label{
  position:relative;
  padding-left:calc(var(--switch-width)*1.1);
  z-index:1;
  --color-bg:var(--off-color-bg);
  --color-dot:var(--off-color-dot);
}
.switch + label::before,
.switch + label::after{
  content:"";
  position:absolute;
  left:0;
  z-index:1;
}
.switch + label::before{
  background-color:var(--color-bg);
  width:var(--switch-width);
  border-radius:var(--switch-width);
  height:calc(var(--switch-width)/2);
  top:0.25em;
  transition:background-color 100ms ease-in-out;
}
.switch + label::after{
  background-color:var(--color-dot);
  width:calc(var(--switch-width)/2);
  height:calc(var(--switch-width)/2);
  top:0.25em;
  border-radius:100%;
  transform:translateX(0%) scale(.8) ;
  transition:transform 200ms ease-in-out,
    background-color 100ms ease-in-out;
  z-index:2;
}
.switch:checked +label{
  --color-bg:var(--on-color-bg);
  --color-dot:var(--on-color-dot);
}
.switch:checked +label::after{
  transform:translateX(100%) scale(.8);
}

.switch:disabled + label{
  /* --color-bg:var(--disabled-color-bg);
  --color-dot:var(--disabled-color-dot); */
  filter:saturate(0.3);
  opacity:0.8;
  cursor:not-allowed;
}
.switch:focus-visible + label{
  outline:2px solid black;
  box-shadow:0px 0px 0px .2em white;
  border-radius:.2em; 
}
</style>