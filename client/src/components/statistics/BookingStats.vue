<template>
  <div>
    <br />
    <div class="level">
      <div class="level-left">
        <div clas="level-item">
          <input
            class="input"
            type="number"
            placeholder="# Days"
            style="width:6em; text-align:right;"
            v-model="days"
          />
        </div>
        <div clas="level-item">
          <Button class="button is-link" @click="fillData">Create Chart</Button>
        </div>
      </div>
    </div>
    <LineChart v-if="loaded" :chartdata="chartData" :options="options" />
  </div>
</template>

<script>
import LineChart from "./BookingStats.js";

export default {
  components: {
    LineChart
  },
  data() {
    return {
      days: 7,
      loaded: false,
      chartData: null,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [
            {
              ticks: {
                beginAtZero: true
              }
            }
          ]
        }
      }
    };
  },
  methods: {
    fillData() {
      this.loaded = false;
      this.$http
        .post("getBookingStats", this.days)
        .then(response => {
          this.chartData = {
            labels: this.cropTime(Object.keys(response.data)),
            datasets: [
              {
                label: "Number of Drinks Booked",
                data: Object.values(response.data),
                backgroundColor: "#3273DC",
                borderColor: "#3273DC",
                fill: false,
                borderJoinStyle: "miter",
                lineTension: 0,
                radius: 2
              }
            ]
          };
          this.loaded = true;
        })
        .catch(response => {
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    cropTime(arr) {
      var result = [];
      arr.forEach(element => {
        result = [].concat(result, element.substring(0, 10));
      });
      return result;
    }
  },
  mounted() {
    this.fillData();
  }
};
</script>