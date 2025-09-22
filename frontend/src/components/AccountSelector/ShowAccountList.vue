<template>
  <div class="table-container">
    <table class="table is-striped is-hoverable">
      <tbody>
        <tr>
          <th>ID</th>
          <th>First Name</th>
          <th>Nick Name</th>
          <th>Last Name</th>
          <th>E-Mail</th>
          <th>Phone</th>
          <th class="has-text-right">Balance</th>
          <th class="has-text-right">MaxDebt</th>
          <th>Category</th>
          <th>Enabled</th>
          <th>Created At</th>
          <th v-show="hasEditAccountRights">Actions</th>
        </tr>
        <tr :class="account$.selected.includes(account) ? 'is-primary' : ''" v-for="account in accounts" @click="account$.select(account)">
          <td>{{ account.id }}</td>
          <td>{{ account.firstName }}</td>
          <td>{{ account.nickName }}</td>
          <td>{{ account.lastName }}</td>
          <td class="has-copy-btn">{{ account.email }} <span class="icon is-small" @click="copyText(account.email)"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-copy-btn">{{ account.phone }} <span class="icon is-small"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-text-right">{{ $n(account.balance / 100, 'currency', 'de-DE') }}</td>
          <td class="has-text-right">{{ $n(account.maxDebt / 100, 'currency', 'de-DE') }}</td>
          <td>
            <button class="tag">
              <span class="icon"><icon :icon="category$.byId(account.category)?.icon" /></span>
              <span>{{ category$.byId(account.category)?.title }}</span>
            </button>  
          </td>
          <td>{{ account.enabled }}</td>
          <td>{{ account.createdAt }}</td>
          <td v-show="hasEditAccountRights">
            <button class="tag">
              <router-link :to="{ name: 'AccountSettingsSingle', params: { accountId: account.id } }">
                <span class="icon"><icon :icon="['fas', 'pen']"/></span>
              </router-link>
              </button>  
            <button class="tag">
              <span class="icon"><icon :icon="['fas', 'trash']"/></span>
            </button>
          </td>
        </tr>

      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import type { Account } from '../../composables/account';
import { useAccountStore } from '../../store/AccountStore';
import type { PropType } from 'vue';
import { useCategoryStore } from '../../store/CategoryStore';

export default {
  data(){
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      dev: false
    }
  },
  props: {
    accounts: {
      type: Array as PropType<Account[]>
    }
  },
  methods: {
    copyText(txt: string){
      navigator.clipboard.writeText(txt);
    }
  },
  computed: {
    hasEditAccountRights(){
      if(this.dev){
        console.warn("User has elevated privileges to edit products because you are running this in development environment")
        return true;
      }
      //ToDo Check if currently loged in user is allowed to edit products
      return false
    }
  },
  mounted() {
    this.dev = import.meta.env.DEV;
  }
}
</script>

<style scoped>
  .has-copy-btn{
    position:relative;

    .icon {
      position:absolute;
      left:100%;
      z-index:10;
      padding:.8em;
      background-color:rgba(0,0,0,.3);
      border-radius:var(--bulma-radius);
      visibility:hidden;
      pointer-events: none;
      cursor:pointer;
    }
    &:hover .icon {
      visibility: visible;
      pointer-events: all;
    }
  }
</style>