<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box max-height-scroll-y">
          <UserTab />
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box max-height-scroll-y" style="width: 500px">
          <MidTab />
        </div>
      </div>
      <div class="column">
        <div class="box max-height-scroll-y">
          <ItemTab />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserTab from "../components/booking/UserTab.vue";
import MidTab from "../components/booking/MidTab.vue";
import ItemTab from "../components/booking/ItemTab.vue";

export default {
  components: {
    UserTab,
    MidTab,
    ItemTab,
  },
};
</script>
