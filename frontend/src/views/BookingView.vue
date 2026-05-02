<template>
  <section class="fixed-grid has-1-cols-mobile has-3-cols-tablet">
    <div class="grid">
      <div style="position:absolute; z-index:-50; left:50%; transform: translateX(-50%); opacity:.6;">
        <img v-if="companyLogo != null" :src="companyLogo"></img>
      </div>
    <OrderViewer />

    <!-- ToDo? Do this as a own component? -->
    <!-- Only visible on mobile -->
    <div class="tabs is-boxed is-centered">
      <ul>
        <li :class="visiblePart == 0 ? 'is-active' : ''">
          <a @click="setVisiblePart(0)">
            <span class="icon is-small"><icon :icon="['fas', 'user']" /></span>
            {{ $t('booking.accounts') }}
          </a>
        </li>
        <li :class="visiblePart == 1 ? 'is-active' : ''">
          <a @click="setVisiblePart(1)">
            <span class="icon is-small"><icon :icon="['fas', 'th-large']" /></span>
            {{ $t('booking.products') }}
          </a>
        </li>
      </ul>
    </div>
    
    <div class="accounts" :class="visiblePart==0 ? '' : 'unselected'">
      <AccountSelector show="button" :all="false"/>
    </div>

    <div class="items" :class="visiblePart==1 ? '' : 'unselected'">
      <!-- ToDo: Develop Item Selector -->
      <ProductSelector show="button" />
    </div>
    
    </div>
  </section>
</template>

<style scoped>
.section{
  padding-block: 0;
}
.tag {
  margin-right:.5em;
}
.unselected{
  display:none;
}
.tabs{
  margin-bottom: 0;
}

@media screen and (min-width: 768px) {
  .order{
    grid-column-start:2;
    grid-row-start: 1;
  }
  .tabs{
    display:none;
  }
  .unselected{
    display:block;
  }
  .accounts{
    grid-column-start:1;
    grid-row-start: 1;
  }
  .items{
    grid-column-start: 3;
    grid-row-start: 1;
  }
}
</style>

<script lang="ts">
import AccountSelector from '../components/AccountSelector/AccountSelector.vue';
import OrderViewer from '../components/OrderViewer/OrderViewer.vue';
import { useAccountStore } from '../store/AccountStore';
import ProductSelector from '../components/ProductSelector/ProductSelector.vue';
import { useSettingStore } from '../store/SettingStore';

  export default {
    components: {
      AccountSelector,
      OrderViewer,
      ProductSelector
    },
    data() {
      return {
        visiblePart: 0 as number,
        account$: useAccountStore(),
        setting$: useSettingStore(),
      }  
    },
    methods: {
      setVisiblePart(partNr: number){
        this.visiblePart = partNr;
      }
    },
    computed: {
      companyLogo(){
        if (this.setting$.get("compLogo") == -1){
          return null;
        }
        return this.setting$.get("compLogo").value;
      }
    }
  }
</script>