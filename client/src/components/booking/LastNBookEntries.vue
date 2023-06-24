<template>
  <div>
    <ul v-for="entry in bookEntries" :key="entry">
      <li class="dropdown is-hoverable is-fullwidth">
        <div class="dropdown-trigger is-fullwidth">
          <button class="button is-multiline">
            {{ printDateTime(entry.TimeStamp) }}<br />
            <span v-if="entry.ItemID === 0 && entry.PaymentMethod.length > 0">
              {{ displayUserName(getUserByID(entry.UserID)) }} ~ {{$t('booking.entries.payment')}}
            </span>
            <span
              v-else-if="entry.ItemID === 0 && entry.Comment.startsWith('Undo')"
            >
              {{ displayUserName(getUserByID(entry.UserID)) }} ~ {{$t('booking.entries.undid')}}
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
                {{ $n(entry.TotalPrice * -1, "currency", "de-DE") }} {{ entry.PaymentMethod }}
              </span>
              <span
                v-else-if="
                  entry.ItemID === 0 && entry.Comment.startsWith('Undo')
                "
              >
                {{ $n(entry.TotalPrice * -1, "currency", "de-DE") }}
              </span>
              <span v-else>
                {{ entry.Amount }}x
                {{ displayItem(getItemByID(items, entry.ItemID)) }} =
                {{ $n(entry.TotalPrice, "currency", "de-DE") }}
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
    bookEntries() {
      return this.$store.state.lastNBookEntries;
    },
    items() {
      return this.$store.state.items;
    },
  },
};
</script>