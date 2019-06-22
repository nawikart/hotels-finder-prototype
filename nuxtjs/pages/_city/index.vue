<template>
  <div id="listing">
    <v-container fluid grid-list-xs class="mainForm_style2 blue lighten-5">
        <mainForm></mainForm>
        <breadcrumb :data="breadcrumb"></breadcrumb>
    </v-container>       

        <v-layout row style="max-width: 1136px; margin: 10px auto;">
          <v-flex md3 hidden-sm-and-down v-if="listing_type=='list'">
            <v-container grid-list-lg>
              <v-card class="listFilter">
                  <div class="blue accent-2 dark py-2 px-2">
                    <v-icon color="white">filter</v-icon> &nbsp; 
                    <span class="subheading white--text">Refine your search</span>
                  </div>                
                  <filterby :cityId="cityId" @filterSubmit="onFilterSubmit" @loading="onLoading"></filterby>    
              </v-card>
            </v-container>
          </v-flex>

          <v-flex xs12 :md9="listing_type=='list'">

            <v-flex xs12 style="height: 50px; padding: 5px; background: #f9f9f9;">
              <div style="float: left;" class="mt-2 ml-3">
                  <a @click="openMap" class="blue--text" style="text-decoration: underline;">view on Map</a>
              </div>

              <div style="float: left;" class="mt-2 ml-3 hidden-sm-and-up">
                <a class="blue--text" style="text-decoration: underline;" v-if="topFilter == false" v-on:click="topFilter = true">click to filter</a>
                <a style="text-decoration: underline;" v-if="topFilter" v-on:click="topFilter = false">close filter</a>
              </div>

              <div style="float: right;">
                <span class="hidden-xs-only" v-if="listing_type == 'sm-grid' || listing_type == 'md-grid' || windowWidth <= 768">
                  <v-btn outline small dark color="blue" v-on:click="topFilter = !topFilter">                  
                      <v-icon v-if="topFilter == true" color="red">close</v-icon>
                      <v-icon v-if="topFilter == false">filter</v-icon>
                      &nbsp; Refine your search
                  </v-btn>
                </span>

                <span class="hidden-xs-only">
                  <v-menu offset-y>
                    <v-btn slot="activator" outline small color="green">
                        <span v-if="listing_type == 'sm-grid'">Small Grid Display</span>
                        <span v-if="listing_type == 'md-grid'">Medium Grid Display</span>
                        <span v-if="listing_type == 'list'">List Display</span>
                        <v-icon>keyboard_arrow_down</v-icon>
                    </v-btn>
                    <v-list>
                      <v-list-tile @click="listing_type='list'"><v-list-tile-title>List Display</v-list-tile-title></v-list-tile>                      
                      <v-list-tile @click="listing_type='sm-grid'"><v-list-tile-title>Small Grid Display</v-list-tile-title></v-list-tile>
                      <v-list-tile @click="listing_type='md-grid'"><v-list-tile-title>Medium Grid Display</v-list-tile-title></v-list-tile>                      
                    </v-list>
                  </v-menu>
                </span>
              </div>
            </v-flex>

            <div color="grey" class="topFilter" v-if="topFilter">
                <filterby :cityId="cityId" @filterSubmit="onFilterSubmit" @loading="onLoading"></filterby>
            </div>

            <hotelinListing v-if="isHotel" :hotelStar="hotelStar" :listing_type="listing_type" :city="city" :country="country" :hotel="hotel" :hotelLoading="hotelLoading"></hotelinListing>

            <listingList v-if="listing_type == 'list'" :items="hotels" :hotelExclude="hotelExclude" :loader="loader"></listingList>
            <listingSmGrid v-if="listing_type == 'sm-grid'" :items="hotels" :hotelExclude="hotelExclude" :loader="loader"></listingSmGrid>
            <listingMdGrid v-if="listing_type == 'md-grid'" :items="hotels" :hotelExclude="hotelExclude" :loader="loader"></listingMdGrid>
            <listingPagination :cc="cc" :P="P" :Ps="Ps"></listingPagination>
          </v-flex>
        </v-layout>

        <gmap :map="map"></gmap>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {

    data: () => ({
      isHotel: false,
      hotelLoading: true,
      hotel: [],
      hotelStar: 0,
      city: '',
      cityId: '',
      country: '',

      P: 1,
      Ps: 2,
      cc: '',
      dataFilter: [],
      breadcrumb: [],
      map: false,
      topFilter: false,
      windowWidth: 0,
      listing_type: 'list',
      hotels: [],
      hotelExclude: '0',
      loader: false
    }),

        methods: {
          openMap (){
              this.map = true
          },
          onFilterSubmit (value) {
            this.loader = false
            this.hotels = value      
          },
          onLoading (value) {
            this.loader = true          
          },

          listing(){
            this.loader = true            
            axios.get(this.$store.state.apiPath +'/hotels-per-city/'+ this.$route.params.city).then(response => {
              console.log(response.data)
              // alert(response.data.Ps)
              this.P = response.data.P
              this.Ps = response.data.Ps
              this.hotels = response.data.Hotels
              this.breadcrumb = {'city': response.data.City, 'citykey': response.data.Citykey, 'country': response.data.Country, 'countryisocode': response.data.Countryisocode}
              this.loader = false
              this.city = response.data.City
              this.cityId = response.data.CityId
              this.country = response.data.Country
            })
          }
        },
         beforeMount(){
            this.cc = (this.$route.params.city).replace(/-hotels(.*)/, '')
            // alert(this.$route.query.ty)
            if(this.$route.query.ty == 'h'){
                this.isHotel = true
                this.hotelExclude = this.$route.query.id
                  axios.get(this.$store.state.apiPath +'/front?mod=hotel&hotel_id='+ this.$route.query.id).then(response => {
                      // console.log('---HOTE')
                      // console.log(response.data)
                      this.hotel = response.data
                      this.hotelStar = this.hotel.Star_rating
                      this.hotel.Excerpt = this.hotel.Overview.substr(0, 200);

                      if(this.hotel.Photo2 == ""){ this.hotel.Photo2 = this.hotel.Photo1 }
                      if(this.hotel.Photo3 == ""){ this.hotel.Photo3 = this.hotel.Photo1 }
                      if(this.hotel.Photo4 == ""){ this.hotel.Photo4 = this.hotel.Photo1 }

                      this.hotelLoading = false
                      // console.log(this.hotel)
                  })                
            }
            this.windowWidth = window.innerWidth
            this.dataFilter = {'sd': this.$route.query.sd, 'ed': this.$route.query.ed, 'ty': this.$route.query.ty, 'id': this.$route.query.id }
            this.listing()  
         }     
  }
</script>

<style type="text/css">
  i.icon{
    cursor: pointer;
  }
</style>