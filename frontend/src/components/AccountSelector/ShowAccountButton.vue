<template>
  <div class="accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">
    <div class="dictionary" v-for="(dict, key) of accountsInOrder" :key="key">
      <span class="title is-1">{{ key }}</span>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <Button
          v-for="account in dict"
          @mousemove="changeSelectMode"
          @click="selectAccount($event, account)"
          :class="account.getShortName() + (account$.selected.includes(account) ? ' is-primary' : '')"
          title="select account">
          {{ account.getShortName() }}
        </Button>
        <!-- <div v-for="account in dict" @mousemove="changeSelectMode" @click="selectAccount($event, account)" >
          <button class="button is-fullwidth" :class="account$.selected.includes(account) ? 'is-primary' : ''" title="select account">{{ account.getShortName() }}</button>
        </div> -->
      </div>
    </div>
  </div>
</template>

<style scoped>
.accountList{
  height: var(--_height);
  overflow-y: scroll;
}
.dictionary{
  display:grid;
  grid-template-columns: 10% 90%;
  margin-bottom:1em;
  .title{
    justify-self: center;
    margin-bottom:0;
  }
}

button.Brenner {
  cursor:url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="40px" height="30px" style="font-size: 20px;"><text y="15">🔥</text></svg>'), auto;
}
button.Käptn {
  cursor:url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="40px" height="30px" style="font-size: 20px;"><text y="15">⚓</text></svg>'), auto;
}


</style>

<script lang="ts">
import { Account } from '../../composables/account';
import { useAccountStore } from '../../store/AccountStore';
import type { PropType } from 'vue';
import { useResizeObserver } from '@vueuse/core';
import { useCartStore } from '../../store/CartStore';
import Button from '../../composables/elements/Button.vue';

export interface dictionary {
    [key: string]: Account[];
}

export default {
  data(){
    return {
      account$: useAccountStore(),
      selectMode: "single",
      height: 0,
      resizeElement: {} as HTMLElement,
      cart$: useCartStore()
    }
  },
  props: {
    accounts: {
      type: Array as PropType<Account[]>
    }
  },
  computed: {
    accountsInOrder(): dictionary {
      var dict: dictionary = {};
      this.accounts?.forEach(acc => {
        var char = acc.getShortName()[0].toUpperCase();
        var charCode = char.charCodeAt(0);
        if (charCode >= 65 && charCode <= 90) { // A-Z
        } else if (charCode >= 48 && charCode <= 57) { // 0-9
          char = "#";
        } else {
          char = "?";
        }
        if (dict[char] === undefined) {
          dict[char] = [acc]
        } else {
          dict[char].push(acc);
        }
      })
      return dict;
    },
  },
  components: {
    Button,
  },
  methods: {
    selectAccount(e: MouseEvent, account: Account) {
      this.changeSelectMode(e);
      this.cart$.cartContents = [];
      if(this.selectMode == "single") {
        this.account$.select(account);
        return;
      }
      if(this.selectMode == "multiple") {
        this.account$.selectAdd(account);
      }
    },
    changeSelectMode(e: MouseEvent){
      if(e.ctrlKey){
        this.selectMode = "multiple";
        (e.target as HTMLElement).style.cursor = "copy";
        return;
      }
      this.selectMode = "single";
      (e.target as HTMLElement).style.cursor = "";
    },
    onResize(){
      let y = window.innerHeight;
      let _y = this.resizeElement.getBoundingClientRect().top;
      let dy = y - _y;
      this.height = dy -15 ;
    }
  },
  mounted() {
    this.resizeElement = this.$refs.resizeBox as HTMLElement
    // Watch for resizing
    useResizeObserver(this.resizeElement, () => {
      this.onResize();
    })
  },
}
</script>