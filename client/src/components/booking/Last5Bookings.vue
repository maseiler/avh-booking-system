<template>
  <div>
    <ul v-for="entry in bookings" :key="entry">
      <li class="dropdown is-hoverable is-fullwidth">
        <div class="dropdown-trigger is-fullwidth">
          <button class="button is-multiline">
            {{ printDateTime(entry.TimeStamp) }}<br />
            <span v-if="entry.ItemID === 0 && entry.PaymentMethod.length > 0">
              {{ displayUserName(getUserByID(entry.UserID)) }} ~ Payment
            </span>
            <span
              v-else-if="entry.ItemID === 0 && entry.Comment.startsWith('Undo')"
            >
              {{ displayUserName(getUserByID(entry.UserID)) }} ~ Undid book
              entry
            </span>
            <span v-else>
              {{ displayUserName(getUserByID(entry.UserID)) }} ~
              {{ entry.Amount }}x
              {{ displayItem(getItemByID(items, entry.ItemID)) }}
            </span>
          </button>
        </div>
        <div class="dropdown-menu" role="menu">
          <div class="dropdown-content has-background-primary-light">
            <div class="dropdown-item has-background-primary-light">
              <p>{{ displayUserNameFull(getUserByID(entry.UserID)) }}</p>
              <span v-if="entry.ItemID === 0 && entry.PaymentMethod.length > 0">
                {{ entry.TotalPrice * -1 }}€ {{ entry.PaymentMethod }}
              </span>
              <span
                v-else-if="
                  entry.ItemID === 0 && entry.Comment.startsWith('Undo')
                "
              >
                {{ entry.TotalPrice * -1 }}€
              </span>
              <span v-else>
                {{ entry.Amount }}x
                {{ displayItem(getItemByID(items, entry.ItemID)) }} =
                {{ entry.TotalPrice }}€
              </span>
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  computed: {
    bookings() {
      return this.$store.state.last5Bookings;
    },
    items() {
      return this.$store.state.items;
    },
  },
};
</script>

<style lang="scss">
.dropdown {
  width: 100%;

  .dropdown-trigger {
    width: 100%;
  }

  .button {
    display: block;
    width: 100%;
    justify-content: space-between;
  }

  .dropdown-menu {
    width: 100%;
  }
}

.button.is-multiline {
  min-height: 2.25em;
  white-space: unset;
  height: auto;
  flex-direction: column;
}
</style>