<template>
  <AccountPaymentCard />
  <AccountSelector show="button"/>
</template>

<script lang="ts">
import AccountSelector from '../components/AccountSelector/AccountSelector.vue';
import AccountPaymentCard from '../components/AccountPaymentCard.vue';

export default {
  data() {
    return {

    }   
  },
  components: {
    AccountSelector,
    AccountPaymentCard
  }
}
</script>