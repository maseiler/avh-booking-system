<template>
  <div>
    <h1>User Info</h1>
    <p>User ID: {{this.user.UserID}}</p>
    <p>BierName: {{this.user.BierName}}</p>
    <p>FirstName: {{this.user.FirstName}}</p>
    <p>LastName: {{this.user.LastName}}</p>
    <p>Status: {{this.user.Status}}</p>
    <p>Email: {{this.user.Email}}</p>
    <p>Phone: {{this.user.Phone}}</p>
    <p>Balance: {{this.user.Balance}}</p>
    <p>MaxDebt: {{this.user.MaxDebt}}</p>
  </div>
</template>

<script>
export default {
  props: {
    user: Object
  }
};
</script>