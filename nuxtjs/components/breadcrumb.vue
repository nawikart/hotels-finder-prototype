<template>
	<div id="breadcrumb" class="pl-3">
			<a href="/">Home</a> » <a>{{ data.country }}</a>

			<a v-if="data.city != ''" :href="'/' + data.citykey + '-' + data.countryisocode + '-hotels.html'"> » {{ data.city }}</a>

			<a v-if="data.hotel_name" :href="'/' + data.hotel_namekey + '/hotel/' + data.citykey + '-' + data.countryisocode + '.html'"> » {{ data.hotel_name }}</a>
	</div>
</template>

<script>
  export default {
    props:{ data: Object },    
  }
</script>

<style type="text/css">
	#breadcrumb{
		padding: 8px 0 7px;
		color: #666;
		width: 100%; 
		max-width: 1136px;
		margin: auto;
		text-shadow: 0 0 2px #fff;
	}
	#breadcrumb a{
		color: #666;
		font-size: 13px;
		text-decoration: none;
	}
	#breadcrumb a:hover{
		color: #ff9800;
	}
</style>