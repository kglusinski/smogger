<template>
  <form>
    <select v-model="selected" @input="updateSelected($event.target.value)">
      <option v-for="city in cities" :key="city.name">
        {{ city.name }}
      </option>
    </select>
  </form>
</template>

<script>
import axios from 'axios'

export default {
  name: "Cities",
  data () {
    return {
      cities: [],
      selected: null
    }
  },
  methods: {
    fetchData: async function () {
      const res = await axios.get("http://localhost:8080/v1/cities?country=PL")

      this.cities = res.data
    },
    updateSelected: function(s) {
      this.$emit('updated-city', s)
    }
  },
  mounted() {
    this.fetchData()
  }
}
</script>

<style scoped>

</style>