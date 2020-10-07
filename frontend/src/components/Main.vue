<template>
  <div>
    <Cities @updated-city="fetchMeasurements($event)"/>
    <Chart />
  </div>
</template>

<script>
import Cities from "@/components/Cities";
import Chart from "@/components/Chart";
import axios from 'axios';
export default {
  name: 'Main',
  components: {Chart, Cities},
  props: {
    msg: String
  },
  methods: {
    fetchMeasurements: async function (e) {
      const res = await axios.get(`http://localhost:8080/v1/measurements?city=${e}&date_from=2020-01-01&date_to=2020-10-01`);

      console.log(res.data, e)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
