<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box" style="height:90vh; overflow-x:hidden; overflow-y:auto;">
          <UserTab />
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box" style="width: 500px; height:90vh; overflow-y:auto;">
          <MidTab />
        </div>
      </div>
      <div class="column">
        <div class="box" style="height:90vh; overflow-x:hidden; overflow-y:auto;">
          <ItemTab />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserTab from "./UserTab.vue";
import MidTab from "./MidTab.vue";
import ItemTab from "./ItemTab.vue";

export default {
  components: {
    UserTab,
    MidTab,
    ItemTab
  }
};
</script>