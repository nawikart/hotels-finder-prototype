<template>
	<div class="stars">
		<v-icon color="orange" v-for="i in parseInt(Star_rating)" :key="i">star</v-icon>
		<v-icon color="orange" v-if="parseInt(Star_rating) < Star_rating">star_half</v-icon>
	</div>
</template>

<script>
  export default {
    props:{ Star_rating: Number }
  }
</script>