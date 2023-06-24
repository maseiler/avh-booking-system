<template>
  <div>
    <table class="table is-hoverable is-striped">
      <thead>
        <tr>
          <th>{{$t('generic.id')}}</th>
          <th>{{$t('item.name')}}</th>
          <th style="text-align: right">{{$t('item.size')}}</th>
          <th>{{$t('item.unit')}}</th>
          <th style="text-align: right">{{$t('item.price')}}</th>
          <th>{{$t('item.enabled')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item"
          @click="selectItem(item)"
          :class="[selectedItem === item ? 'is-selected' : '']"
        >
          <th>{{ item.ID }}</th>
          <td>{{ item.Name }}</td>
          <td style="text-align: right">{{ item.Size }}</td>
          <td>{{ item.Unit }}</td>
          <td style="text-align: right">{{ $n(item.Price, "currency", "de-DE") }}</td>
          <td>
            <font-awesome-icon v-if="item.Enabled" icon="check-circle"/>
            <font-awesome-icon v-if="!item.Enabled" icon="times"/>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  computed: {
    items() {
      let temp = this.$store.state.items;
      return temp.sort((a, b) => {
        if (a.ID < b.ID) return -1;
        return 1;
      });
    },
    selectedItem() {
      return this.$store.state.selectedSingleItem;
    },
  },
  methods: {
    selectItem(item) {
      this.$store.commit("selectSingleItem", item);
    },
  },
};
</script>
