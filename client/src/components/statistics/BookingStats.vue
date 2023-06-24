<template>
  <div>
    <br />
    <div class="level">
      <div class="level-left">
        <div clas="level-item">
          <input
            class="input"
            type="number"
            v-bind:placeholder="$t('statistics.numDays')"
            style="width: 6em; text-align: right"
            v-model="days"
          />
          <Button class="button is-link" @click="fillData">{{$t('statistics.createChart')}}</Button>
          <p class="p has-text-grey-light">
            &nbsp;&nbsp;{{$t('statistics.selectDays')}}
          </p>
        </div>
      </div>
    </div>
    <div clas="level-item has-text-centered">
      <p class="p has-text-grey-light">
        &nbsp;&nbsp;{{$t('statistics.clickToDisEn')}}
      </p>
    </div>
    <div clas="level-item"></div>
    <LineChart
      v-if="loaded"
      :chartdata="chartData"
      :options="options"
      style="position: relative; height: 75vh"
    />
  </div>
</template>

<script>
import LineChart from "./BookingStats.js";

export default {
  components: {
    LineChart,
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
          xAxes: [
            {
              ticks: {
                reverse: true,
              },
            },
          ],
          yAxes: [
            {
              ticks: {
                beginAtZero: true,
              },
            },
          ],
        },
      },
    };
  },
  computed: {
    items() {
      return this.$store.state.items;
    },
  },
  methods: {
    fillData() {
      this.loaded = false;
      this.$http
        .post("getBookingStats", this.days)
        .then((response) => {
          var datasets = this.createDataset(response.data);
          this.chartData = {
            labels: this.cropTime(response.data.timeStamp),
            datasets: datasets,
          };
          this.loaded = true;
        })
        .catch((response) => {
          //ToDo: internationalize this failure message
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    createDataset(resp) {
      var keys = Object.keys(resp);
      var values = Object.values(resp);
      var datasets = [];
      datasets.push({
        label: "Total",
        data: values[values.length - 1],
        backgroundColor: "#000000",
        borderColor: "#000000",
        fill: false,
        borderJoinStyle: "miter",
        lineTension: 0,
        radius: 2,
      });
      keys.splice(keys.length - 2, 2);
      values.splice(values.length - 2, 2);
      for (var i = 0; i < keys.length; i++) {
        var color = this.getRandomColor(keys[i]);
        datasets.push({
          label: `${this.displayItem(
            this.getItemByID(this.$store.state.items, keys[i])
          )}`,
          data: values[i],
          backgroundColor: color,
          borderColor: color,
          fill: false,
          borderJoinStyle: "miter",
          lineTension: 0,
          radius: 2,
        });
      }
      return datasets;
    },
    getRandomColor(itemID) {
      var keys = Object.keys(Colors);
      return Colors[keys[itemID % keys.length]];
    },
    cropTime(arr) {
      var result = [];
      arr.forEach((element) => {
        result = [].concat(result, element.substring(0, 10));
      });
      return result;
    },
  },
  mounted() {
    this.fillData();
  },
};

const Colors = {
  aqua: "#00ffff",
  // azure: "#f0ffff",
  // beige: "#f5f5dc",
  // black: "#000000",
  blue: "#0000ff",
  brown: "#a52a2a",
  cyan: "#00ffff",
  darkblue: "#00008b",
  darkcyan: "#008b8b",
  darkgrey: "#a9a9a9",
  darkgreen: "#006400",
  darkkhaki: "#bdb76b",
  darkmagenta: "#8b008b",
  darkolivegreen: "#556b2f",
  darkorange: "#ff8c00",
  darkorchid: "#9932cc",
  darkred: "#8b0000",
  darksalmon: "#e9967a",
  darkviolet: "#9400d3",
  fuchsia: "#ff00ff",
  gold: "#ffd700",
  green: "#008000",
  indigo: "#4b0082",
  khaki: "#f0e68c",
  lightblue: "#add8e6",
  lightcyan: "#e0ffff",
  lightgreen: "#90ee90",
  lightgrey: "#d3d3d3",
  lightpink: "#ffb6c1",
  // lightyellow: "#ffffe0",
  lime: "#00ff00",
  magenta: "#ff00ff",
  maroon: "#800000",
  navy: "#000080",
  olive: "#808000",
  orange: "#ffa500",
  pink: "#ffc0cb",
  purple: "#800080",
  violet: "#800080",
  red: "#ff0000",
  silver: "#c0c0c0",
  //     white: "#ffffff",
  yellow: "#ffff00",
};
</script>