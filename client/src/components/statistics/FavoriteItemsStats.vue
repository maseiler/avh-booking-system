<template>
  <div>
    <br />
    <BarChart
      v-if="loaded"
      :chartdata="chartData"
      :options="options"
      style="position: relative; height:75vh"
    />
  </div>
</template>

<script>
import BarChart from "./FavoriteItemsStats.js";

export default {
  components: {
    BarChart
  },
  data() {
    return {
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
        },
        legend: {
          display: false
        }
      }
    };
  },
  methods: {
    fillData() {
      this.loaded = false;
      this.$http
        .get("getFavoriteItemsStats")
        .then(response => {
          var itemStats = this.sortByCount(response.data);
          var labels = this.getLabels(itemStats);
          var values = this.getValues(itemStats);
          var barColors = this.getBarColors(itemStats);

          this.chartData = {
            labels: labels,
            datasets: [
              {
                data: values,
                backgroundColor: barColors,
                borderWidth: 1
              }
            ]
          };

          this.loaded = true;
        })
        .catch(response => {
          //ToDo: Internationalize failure message
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    sortByCount(itemStats) {
      return itemStats.sort((a, b) => {
        return b["Count"] > a["Count"] ? 1 : b["Count"] < a["Count"] ? -1 : 0;
      });
    },
    getLabels(itemStats) {
      var labels = [];
      itemStats.forEach(stat => {
        var label = "".concat(stat["Name"]);
        if (stat["Unit"] == "l") {
          label = label.concat(" ", stat["Size"]);
        }
        labels.push(label);
      });
      return labels;
    },
    getValues(itemStats) {
      var values = [];
      itemStats.forEach(stat => {
        values.push(stat["Count"]);
      });
      return values;
    },
    getBarColors(itemStats) {
      var colors = [];
      itemStats.forEach(stat => {
        switch (stat["Type"]) {
          case "non-alcoholic":
            colors.push("rgb(153, 212, 244, 0.9)");
            break;
          case "alcoholic":
            colors.push("rgb(103, 163, 193, 0.9)");
            break;
          case "food":
            colors.push("rgb(204, 255, 255, 0.9)");
            break;
          default:
            colors.push("rgb(218, 240, 251, 0.9)");
        }
      });
      return colors;
    }
  },
  mounted() {
    this.fillData();
  }
};
</script>