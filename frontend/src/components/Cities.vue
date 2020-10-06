<template>
  <form>

  </form>
</template>

<script>
import axios from 'axios'

export default {
  name: "Cities",
  mounted() {
    axios.get("http://localhost:8080/v1/cities?country=PL")
        .then(function (res) {
          console.log(res)
        })
  }
}
</script>

<style scoped>

</style>