<template>
	<div>
		<hr style="border: 10px solid #fafafa; box-shadow: 0px 5px 25px #ddd;" />
		<div style="max-width: 1280px; margin: 40px auto 0; display: table; width: 100%;">
			<v-container class="display-1 text-xs-center blue--text" style="text-shadow: 2px 2px 5px #aaa; margin-bottom: 10px;">Hotel Deals of the Day</v-container>

			<p align="center"><v-btn :outline="cityId != item.City_id" small v-for="(item, i) in items" :key="i" color="orange" dark v-on:click="onClickChild(item.City_id, item.Citykey)">{{ item.City }}</v-btn></p>

			<listingMdGrid v-if="listing_type == 'md-grid'" :items="hotels" :loader="loader"></listingMdGrid>
			<listingList v-if="listing_type == 'list'" :items="hotels" :loader="loader"></listingList>
		</div>
	</div>
</template>

<script>
  import axios from 'axios'

  export default {    
    data: () => ({
      loader: false,
      cityId: '17193',
      citykey: 'bali',
      listing_type: 'md-grid',
      items: [],
      hotels: []
    }),

        methods: {
          onClickChild (cityId, citykey) {

            this.cityId = cityId

            var d = new Date()
            var mo = (d.getMonth() + 1)
            if(mo < 10){
                mo = '0' + mo
            }
            var dt = d.getDate()
            if(dt < 10){
                dt = '0' + dt
            }
            var sd = d.getFullYear() + '-' + mo + '-' + dt

            d.setDate(d.getDate() + 1)
            var mo = (d.getMonth() + 1)
            if(mo < 10){
                mo = '0' + mo
            }
            var dt = d.getDate()
            if(dt < 10){
                dt = '0' + dt
            }      
            var ed = d.getFullYear() + '-' + mo + '-' + dt      


            this.cityId = cityId
            this.loader = true

            var dataFilter = JSON.stringify({
                                 "criteria": {
                                               "additional": {
                                                               "maxResult": 30,
                                                               // "discountOnly": true,
                                                               "sortBy": "Recommended"
                                                             },
                                               "checkInDate": sd,
                                               "checkOutDate": ed,
                                               "cityId": parseInt(cityId)
                                              }
                              })

            axios.get(this.$store.state.apiPath +'/front?mod=agoda&data='+ dataFilter).then(response => {
              var agoda = response.data.results              
              var agoda_hids = []

              for (var i = 0; i < agoda.length; i++) {
                  agoda_hids.push(agoda[i].hotelId)
              }

                  axios.get(this.$store.state.apiPath +'/front?mod=hotels&citykey='+ citykey +'&hids='+ agoda_hids.join()).then(response => {
                      var hotels = response.data.Hotels

                      for (var i = 0; i < hotels.length; i++) { 
                        for (var j = 0; j < agoda.length; j++) {                     
                          if(hotels[i]['Hotel_id'] == agoda[j]['hotelId']){
                              hotels[i]['starRating'] = agoda[j]['starRating']
                              hotels[i]['reviewScore'] = agoda[j]['reviewScore']
                              hotels[i]['reviewCount'] = agoda[j]['reviewCount']
                              hotels[i]['currency'] = agoda[j]['currency']
                              hotels[i]['dailyRate'] = agoda[j]['dailyRate']
                              hotels[i]['discountPercentage'] = agoda[j]['discountPercentage']
                              hotels[i]['Photo'] = agoda[j]['imageURL']
                              hotels[i]['includeBreakfast'] = agoda[j]['includeBreakfast']
                              hotels[i]['freeWifi'] = agoda[j]['freeWifi']
                          }
                        }
                      }
                      
                      this.hotels = hotels
                      this.loader = false
                  })
            })
          }
        },
         beforeMount(){

         	  axios.get(this.$store.state.apiPath +'/front?mod=cities').then(response => {
                // console.log(response.data)
                this.items = response.data.Cities              
              }) 

            if(window.innerWidth <= 599){
            	this.listing_type = 'list'
            }else{
            	this.listing_type = 'md-grid'
            }

            this.onClickChild(this.cityId, this.citykey)
         }     
  }
</script>