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
          var sortedDict = this.sortByValue(response.data);
          var itemIDs = sortedDict.map(pair => {
            return pair.key;
          });
          var itemCount = sortedDict.map(pair => {
            return pair.val;
          });

          this.chartData = {
            labels: itemIDs.map(id => {
              return this.displayItem(this.getItemByID(id));
            }),
            datasets: [
              {
                label: "Drinks Bought",
                data: itemCount,
                backgroundColor: "rgba(50, 115, 220, 0.5)",
                borderColor: "rgba(50, 115, 220)",
                borderWidth: 1
              }
            ]
          };
          this.loaded = true;
        })
        .catch(response => {
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    sortByValue(dict) {
      var keys = Object.keys(dict);
      var sortedValues = Object.values(dict).sort((a, b) => {
        return b - a;
      });
      var sortedDict = [];
      for (var v = 0; v < sortedValues.length; v++) {
        for (var k = 0; k < keys.length; k++) {
          if (dict[keys[k]] == sortedValues[v]) {
            sortedDict.push({ key: keys[k], val: sortedValues[v] });
            keys.splice(k, 1);
            break;
          }
        }
      }
      return sortedDict;
    }
  },
  mounted() {
    this.fillData();
  }
};
</script>