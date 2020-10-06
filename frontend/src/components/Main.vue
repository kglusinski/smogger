<template>
  <div>
    <Cities />
  </div>
</template>

<script>
import Cities from "@/components/Cities";
export default {
  name: 'Main',
  components: {Cities},
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
