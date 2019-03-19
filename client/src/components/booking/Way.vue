<template>
  <div>
    <button @click="someValue = 'Some New Value'">Change Value</button>
    <div>Parent: {{ someValue }}</div>
    <test v-model="someValue" endpoint="foo"></Test>
    <test v-model="someValue" endpoint="bar"></Test>
  </div>
</template>

<script src="./Way.js"></script>