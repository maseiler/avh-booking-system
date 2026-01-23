<template>
  <section class="fixed-grid has-1-cols-mobile has-3-cols-tablet">
    <div class="grid">
    <OrderViewer />

    <!-- ToDo? Do this as a own component? -->
    <!-- Only visible on mobile -->
    <div class="tabs is-boxed is-centered">
      <ul>
        <li :class="visiblePart == 0 ? 'is-active' : ''">
          <a @click="setVisiblePart(0)">
            <span class="icon is-small"><icon :icon="['fas', 'user']" /></span>
            Accounts
          </a>
        </li>
        <li :class="visiblePart == 1 ? 'is-active' : ''">
          <a @click="setVisiblePart(1)">
            <span class="icon is-small"><icon :icon="['fas', 'th-large']" /></span>
            Products
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

  export default {
    components: {
      AccountSelector,
      OrderViewer,
      ProductSelector
    },
    data() {
      return {
        visiblePart: 0 as number,
        account$: useAccountStore()
      }  
    },
    methods: {
      setVisiblePart(partNr: number){
        this.visiblePart = partNr;
      }
    }
  }
</script>