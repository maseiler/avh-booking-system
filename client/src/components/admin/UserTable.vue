<template>
  <div>
    <div class="table-container">
      <table class="table is-hoverable is-striped">
        <thead>
          <tr>
            <th>{{$t('generic.id')}}</th>
            <th>{{$t('user.firstName')}}</th>
            <th>{{$t('user.nickname')}}</th>
            <th>{{$t('user.lastName')}}</th>
            <th>{{$t('user.boatName')}}</th>
            <th>{{$t('user.status')}}</th>
            <th>{{$t('user.eMail')}}</th>
            <th>{{$t('user.phoneNumber')}}</th>
            <th style="text-align: right">{{$t('generic.balance')}}</th>
            <th style="text-align: right">{{$t('generic.max')}} {{$t('generic.debt')}}</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="user in users"
            :key="user"
            @click="selectUser(user)"
            :class="[selectedUser === user ? 'is-selected' : '']"
          >
            <th>{{ user.ID }}</th>
            <td>{{ user.FirstName }}</td>
            <td>{{ user.BierName }}</td>
            <td>{{ user.LastName }}</td>
            <td>{{ user.BoatName }}</td>
            <td>{{ user.Status }}</td>
            <td>{{ user.Email }}</td>
            <td>{{ user.Phone }}</td>
            <td style="text-align: right">{{ user.Balance }} €</td>
            <td style="text-align: right">{{ user.MaxDebt }} €</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    users() {
      let temp = this.$store.state.users;
      return temp.sort((a, b) => {
        if (a.ID < b.ID) return -1;
        return 1;
      });
    },
    selectedUser() {
      return this.$store.state.selectedUser;
    },
  },
  methods: {
    selectUser(user) {
      this.$store.commit("selectUser", user);
    },
  },
};
</script>
