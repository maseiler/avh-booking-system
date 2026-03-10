<template>
  <div class="panel is-primary" v-if="account$.selected.length > 0">
    <div class="panel-heading">
      <span>{{ account$.selected[0].getFullName() }}</span>
      <button class="delete" aria-label="delete" @click="account$.unselect();"></button>
    </div>
    <div class="panel-block">
      <div class="account-id-area">
        <div class="account-id-grid">
          <icon class="" :icon="['fas', 'user-secret']" /><span>{{ account$.selected[0].nickname }}</span>
          <icon class="" :icon="['fas', 'user']" /><span>{{ account$.selected[0].firstName }}</span>
          <icon class="" :icon="['fas', 'id-card']" /><span>{{ account$.selected[0].lastName }}</span>
          <icon class="" :icon="categoryIcon" /><span>{{ category$.byId(account$.selected[0].category)?.title }}</span>
        </div>
        <div class="balance-area">
          <span>Balance</span><br>
          <span class="balance" :class="account$.selected[0].balance <= (-1* account$.selected[0].maxDebt) ? 'has-text-danger' : 'has-text-primary'">{{ $n(account$.selected[0].balance / 100, 'currency') }}</span>
        </div>
      </div>

      <Buttons>
        <Button :fa-icon="['fas', 'list']" icon-position="left" title="List all Orders of this Account">
          List Orders
        </Button>

        <Button class="is-success" :fa-icon="['fas', 'coins']" icon-position="right">
          Pay now
        </Button>
      </Buttons>
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
          <span>Pay Now</span>
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
  gap: 0.1rem 0.25rem;
}
.columns{
  width:100%;
}
.panel{
  max-width: 80ch;
  margin-inline: auto;
  position:relative;
  box-shadow: 0 4px 16px rgba(0,0,0,0.12);
}
.panel-heading{
  font-size:1.05rem;
  letter-spacing: 0.02em;
}
.panel-block{
  display:block;
  background-color: hsl(var(--bulma-scheme-h), var(--bulma-scheme-s), var(--bulma-scheme-main-l));
}
.buttons{
  margin-top:.75rem;
  justify-content: end;
}
.balance-area{
  text-align: center;
}
.balance{
  font-size:2rem;
  font-weight:700;
}
.delete{
  position:absolute;
  right:25px;
}
</style>

<script lang="ts">
import { useAccountStore } from '../store/AccountStore';
import { useCategoryStore } from '../store/CategoryStore';
import Button from '../composables/elements/Button.vue';
import Buttons from '../composables/elements/Buttons.vue';

export default {
  data(){
    return {
      account$: useAccountStore(),
      category$: useCategoryStore()
    }
  },
  components: {
    Button,
    Buttons
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