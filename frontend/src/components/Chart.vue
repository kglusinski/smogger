<template>
  <div id="chart">
    <D3BarChart :config="config" :datum="data" />
  </div>
</template>

<script>
//import { D3BarChart } from 'vue-d3-charts';
export default {
name: "Chart",
  //components: {D3BarChart},
  data() {
    return {
      config: {
        key: "date",
        value: "value",
        color: "steelblue"
      },
      width: 800,
      height: 300,
      margin: {
        top: 50,
        right: 50,
        bottom: 50,
        left: 50
      },
      data: [
        {
          date: "2020-01-01",
          value: 0.5
        },
        {
          date: "2020-01-02",
          value: 0.7
        }
      ]
    }
  },
  methods: {
    createChart: function () {

    }
  },
  init() {

  }
}
</script>

<style scoped>

</style>