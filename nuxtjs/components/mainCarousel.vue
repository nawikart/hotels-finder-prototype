<template>
  <v-carousel>
    <v-carousel-item v-for="(item,i) in items2" :key="i" :src="item.src"></v-carousel-item>
  </v-carousel>
</template>

<script>
  export default {
    data () {
      return {
        items2: [
          {
            src: '/squirrel.jpg'
          },
          {
            src: '/sky.jpg'
          },
          {
            src: '/planet.jpg'
          }
        ]
      }
    }
  }
</script>

<style>
.carousel{
	height: auto;
}
</style>