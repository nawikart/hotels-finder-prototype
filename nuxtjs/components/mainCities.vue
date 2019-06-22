<template>
  <v-container id="gridCities" fluid grid-list-sm style="max-width: 1360px;">

    <v-flex xs11 md8 style="margin: 50px auto 0;">
      <p class="display-1 text-xs-center orange--text">{{ topTitle }}</p>
      <p class="subheading text-xs-center">{{ topDes }}</p>
      <p align="center">
        <v-btn large dark :outline="!showCountries" color="green" v-on:click="countriesClicked()">Top Country</v-btn>
        <v-btn large dark :outline="!showCities" color="green" v-on:click="citiesClicked()">Top Cities</v-btn>
      </p>
    </v-flex>

    <v-layout wrap gallery v-if="isCity && !ccLoader">
      <v-flex item xs6 sm4 md3 v-for="(item, i) in cities" :key="i" v-if="i < 12">
        <div class="triangle"></div>
        <router-link :to="'/' + item.City_key + '-' + item.CountryIsoCode.toLowerCase()  + '-hotels.html'">
            <v-flex img :style="'background-image: url('+ $store.state.mediaPath +'/media/images/destination/'+ item.City_key +'.jpg); background-size: cover;'">
              <div class="overlay">
                <p class="subheading text-xs-center">{{item.City}} ({{ item.HotelsCount }})</p>
              </div>
            </v-flex>
        </router-link>
      </v-flex>
    </v-layout>

    <v-layout wrap gallery v-if="!isCity && !ccLoader">
      <v-flex item xs6 sm4 md3 v-for="(item, i) in countries" :key="i" v-on:click="countryClicked(item.Country_id, item.Country_key, item.Country)">
            <v-flex img :style="'background-image: url('+ $store.state.mediaPath +'/media/images/destination/'+ item.Country_key +'.jpg); background-size: cover; cursor: pointer;'">
              <div class="overlay">
                <p class="subheading text-xs-center">{{item.Country}} ({{ item.HotelsCount }})</p>
              </div>
            </v-flex>
      </v-flex>
    </v-layout>
    <!-- <p align="center"><v-btn large dark outline color="grey" v-on:click="" style="border: none;">show more</v-btn></p> -->


    <v-layout wrap gallery v-if="ccLoader">
      <v-flex item xs6 sm4 md3 p-2 v-for="i in 12" :key="i" style="height: 180px; background-color: #fafafa; border: 6px solid white;">
              <br />
              <br />
              <br />
              <br />
              <br />
              <br />
              <v-progress-linear :indeterminate="true" value="15" height="1" color="black lighten-1" style="width: 80%; margin: auto;"></v-progress-linear>
              <br />
      </v-flex>
    </v-layout>

    <p>&nbsp;</p><p>&nbsp;</p>
  </v-container>
  </v-container>
</template>

<script>
  import axios from 'axios'
  export default {
    data: () => ({
      hide: false,
      ccLoader: false,
      showCities: false,
      showCountries: false,
      isCity: false,
      topTitle: '',
      topDes: '',
      cities: [],

      countries: [],

      citiesInaCountry: [],                
    }),

        methods: {
          topCountries () {
              this.ccLoader = true
              axios.get(this.$store.state.apiPath +'/getTop24Countries').then(response => {
                console.log(response.data)
                this.countries = response.data
                this.ccLoader = false                
              })              
          },          
          citiesClicked () {
              this.ccLoader = true
              axios.get(this.$store.state.apiPath +'/getTop24Cities').then(response => {
                // console.log(response.data)
                this.cities = response.data
                // console.log(this.cities)

                this.isCity = true
                this.showCities = true
                this.showCountries = false
                this.topTitle = 'top destinations you will love'
                this.topDes = 'Whoever you are, whatever you want, we’ve got the perfect place for you. From cozy country homes to funky city apartments, are located in 133,839 destinations across 210+ countries and territories.'
                this.ccLoader = false                
              })              
          },
          countriesClicked () {
              this.topCountries()
              this.isCity = false
              this.showCities = false
              this.showCountries = true
              this.topTitle = '400,000+ properties, worldwide'
              this.topDes = 'Our huge selection of accommodations will let you plan the perfect trip. From adventure travel and backpacking to honeymoons and family vacations, we\'ve got you covered. Let\'s explore the 400,000+ properties in 22,000+ cities across 210+ countries, worldwide.'
          },                   
          countryClicked (country_id, country_key, country_name) {
              this.ccLoader = true
              axios.get(this.$store.state.apiPath +'/getTop24CitiesPerCountry/'+ country_key).then(response => {
                this.cities = response.data
                this.isCity = true
                this.showCities = false
                this.showCountries = false
                this.topTitle = 'recommended cities in ' + country_name
                this.topDes = 'From up-and-coming cities to charming small towns, these are the destinations in '+ country_name +' you need to visit this year.'
                this.ccLoader = false               
              })
          }
        },
         beforeMount(){
            this.countriesClicked()
         }         
  }
</script>

<style>
#gridCities .gallery{
  width: 94%!important;
  margin: auto!important;
}
#gridCities .gallery .item{
  padding: 7px!important;
  position: relative;
}
#gridCities .gallery .item a {
    display: block;
    position: relative;
    overflow: hidden;
}
#gridCities .gallery .item .img{
  height: 180px;
}
#gridCities .gallery .item .subheading{
  position: absolute;
  color: #333;
  font-weight: bold;
  bottom: -16px;
  right: 0;
  text-align: center;
  display: table;
  width: 100%;
  padding: 9px 0;
  background-color: rgba(255,255,255,.7);
}
#gridCities .gallery .item:hover .subheading{
  background-color: rgba(255,255,255,.9);
}
#gridCities .gallery .item:hover .overlay{
  background-color: rgba(255,255,255,.1);  
}
#gridCities .gallery .item img {
    zoom: 1;
    width: 100%;
    -webkit-transition: all .9s ease;
    -moz-transition: all .9s ease;
    -ms-transition: all .9s ease;
    transition: all .9s ease;
}
#gridCities .gallery .item>a:hover img{
  zoom: 2;
}
#gridCities .gallery .item .overlay {
    width: 100%;
    height: 100%;
    padding: 0;
    position: absolute;
    background-color: rgba(255,255,255,.25);
    -webkit-transition: all 600ms ease-out 0s;
    -moz-transition: all 600ms ease-out 0s;
    -o-transition: all 600ms ease-out 0s;
}

@media only screen and (max-width: 1024px) {
  #gridCities .gallery .item .img{
    height: 140px;
  }
}
@media only screen and (max-width: 1023px) {
  #gridCities .gallery{
    width: 94%!important;
    margin: auto!important;
  }
}
@media only screen and (max-width: 599px) {
  #gridCities .gallery{
    width: auto!important;
  }
  #gridCities .gallery .item{
    padding: 5px!important;
  }
  #gridCities .gallery .item .overlay {
      background-color: rgba(255,255,255,.1);  
  }
  #gridCities .gallery .item .img{
    height: 120px;
  }  
}
</style>