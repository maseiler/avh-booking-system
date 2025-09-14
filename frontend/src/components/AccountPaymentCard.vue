<template>
  <div class="panel" v-if="account$.selected.length > 0">
    <p class="panel-heading">{{ account$.selected[0].getFullName() }}</p>
    <div class="panel-block">
      <div class="account-id-area">
        <div class="account-id-grid">
          <icon class="" :icon="['fas', 'user-secret']" /><span>{{ account$.selected[0].nickName }}</span>
          <icon class="" :icon="['fas', 'user']" /><span>{{ account$.selected[0].firstName }}</span>
          <icon class="" :icon="['fas', 'id-card']" /><span>{{ account$.selected[0].lastName }}</span>
          <icon class="" :icon="categoryIcon" /><span>{{ category$.byId(account$.selected[0].category)?.title }}</span>
        </div>
        <div class="balance-area">
          <span>Balance</span><br>
          <span class="balance">{{ $n(account$.selected[0].balance / 100, 'currency', 'de-DE') }}</span>
        </div>
      </div>
      <div class="buttons">
        <button class="button" @click="$emit('cancelOrder')" title="discard cart and unselect account">
          <span class="icon"><icon :icon="['fas', 'list']" /></span>
          <span>List Orders</span>
        </button>
        <button class="button is-success is-inverted is-outlined">
          <span>Pay now</span>
          <span class="icon"><icon :icon="['fas', 'coins']"/></span>
        </button>
      </div>
    </div>
  </div>

  <!-- Skeleton -->
  <div class="panel" v-if="account$.selected.length == 0">
    <p class="panel-heading is-skeleton">Select an Account</p>
    <div class="panel-block">
      <div class="account-id-area">
        <div class="account-id-grid">
          <icon class="icon is-skeleton" :icon="['fas', 'user-secret']" /><span class="skeleton-lines"><div></div></span>
          <icon class="icon is-skeleton" :icon="['fas', 'user']" /><span class="skeleton-lines"><div></div></span>
          <icon class="icon is-skeleton" :icon="['fas', 'id-card']" /><span class="skeleton-lines"><div></div></span>
          <icon class="icon is-skeleton" :icon="['fas', 'user']" /><span><span class="skeleton-lines"><div></div></span></span>
        </div>
        <div class="balance-area">
          <span class="skeleton-lines"><div></div></span>
          <span class="balance is-skeleton">BALANCE</span>
        </div>
      </div>
      <div class="buttons">
        <button class="button is-skeleton" title="discard cart and unselect account">
          <span class="icon"><icon :icon="['fas', 'list']" /></span>
          <span>List Orders</span>
        </button>
        <button class="button is-skeleton">
          <span class="icon"><icon :icon="['fas', 'coins']"/></span>
        </button>
      </div>
    </div>
  </div>

</template>

<style scoped>
.account-id-area{
  display:flex;
  justify-content: space-between;
}
.account-id-grid{
  display:grid;
  grid-template-columns: 1.5rem auto;
  align-items: center;
}
.columns{
  width:100%;
}
.panel{
  max-width: 80ch;
  margin-inline: auto;
}
.panel-block{
  display:block;
}
.buttons{
  margin-top:.5rem;
  justify-content: end;
}
.balance-area{
  text-align: center;
}
.balance{
  font-size:2rem;
  font-weight:600;
}
</style>

<script lang="ts">
import { useAccountStore } from '../store/AccountStore';
import { useCategoryStore } from '../store/CategoryStore';

export default {
  data(){
    return {
      account$: useAccountStore(),
      category$: useCategoryStore()
    }
  },
  computed: {
    categoryIcon(){
      let hasSelectedAccount = this.account$.selected.length > 0;
      let undefindedCategoryIcon = ['fas', 'circle-info'];
      if (!hasSelectedAccount){
        return undefindedCategoryIcon
      }
      return this.category$.byId(this.account$.selected[0].category)?.icon ?? undefindedCategoryIcon;
    }
  }
}

</script>