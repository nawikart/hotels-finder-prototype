<template>
  <div>
    <v-card id="hotel-in-listing" class="mt-2 mb-4">
        <v-layout wrap>
          <v-flex my-1 xs12 sm4 img :style="'background: url('+ hotel.Photo1 +') no-repeat; background-size: cover!important; min-height: 275px;'">
            <div class="heart mt-2 ml-2"><v-icon medium color="white">favorite_border</v-icon></div>

              <div class="badge-price" v-if="hotel.Rates_from != 0 || hotel.dailyRate">

                <div class="discount-ribbon middle" v-if="hotel.discountPercentage && hotel.discountPercentage != 0"><span class="discount-ribbon-discount">{{ hotel.discountPercentage }}%</span><span class="discount-ribbon-text">Discount</span></div>

                <div class="price-ribbon middle"><span class="price-ribbon-from">Starting at</span>
                  <div style="display: table; margin: 0 auto;"><span class="price-ribbon-currency" v-if="hotel.currency">{{ hotel.currency }}</span>&nbsp;
                    <span class="price-ribbon-price" v-if="hotel.dailyRate">{{ hotel.dailyRate }}</span>
                    <span class="price-ribbon-price" v-if="!hotel.dailyRate">usd {{ hotel.Rates_from }}</span>
                  </div>
                </div>
              </div>

              <div class="photos">
                  <router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'"><span :style="'background-image: url('+ hotel.Photo1 +');'"></span></router-link>
                  <router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'"><span :style="'background-image: url('+ hotel.Photo2 +');'"></span></router-link>
                  <router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'"><span :style="'background-image: url('+ hotel.Photo3 +');'"></span></router-link>
                  <router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'"><span :style="'background-image: url('+ hotel.Photo4 +');'"></span></router-link>
              </div>              
          </v-flex>   

          <v-flex xs12 sm6 px-3 pt-3 pb-3 style="position: relative;">
              
              <p class="title"><router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'" style="color: #444; line-height:140%!important;"><b>{{ hotel.Hotel_name }}</b></router-link></p>

              <p>
                <stars :Star_rating="hotelStar"></stars>
                <span class="caption"> &nbsp; (<b>{{ hotel.No_reviews }}</b> reviews)</span></p>
                <p class="body-1 address">
                  <v-icon small>location_on</v-icon>                        
                  <span>{{ hotel.Addressline1 }}, {{ hotel.City }}, {{ hotel.Countryisocode }} &nbsp; <small class="viewonmap blue--text">view on map</small></span>                        
                </p>
              
              <p class="pl-1"><em>"{{ hotel.Excerpt }}.." <router-link class="blue--text" :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'">more</router-link></em></p>
              <v-layout class="pl-2" row wrap green--text fitur v-if="!hotel.freeWifi || !hotel.includeBreakfast">
                <div class="xs5 sm12 mr-3"><v-icon small color="green">wifi</v-icon> wifi</div>
                <div class="xs7 sm12"><v-icon small color="green">check</v-icon> include breakfast</div>
              </v-layout>

              <v-layout class="pl-2" row wrap green--text fitur v-if="hotel.freeWifi && hotel.includeBreakfast">
                <div class="xs5 sm12 mr-3"><v-icon small color="green">wifi</v-icon> 
                    <span v-if="hotel.freeWifi">Wifi</span>
                    <strike v-if="!hotel.freeWifi">Wifi</strike>
                </div>
                <div class="xs7 sm12">
                    <v-icon small color="green" v-if="hotel.includeBreakfast">check</v-icon> 
                    <span v-if="hotel.includeBreakfast" class="green--text">Include Breakfast</span>
                    <strike v-if="!hotel.includeBreakfast" class="grey--text">Include Breakfast</strike>
                </div>
              </v-layout>                                               
          </v-flex>

          <v-flex xs2 pt-4 my-1 hidden-xs-only class="blue white--text">
            <p align="center" class="rating_average"><v-icon small>grade</v-icon><em><b class="mr-1">{{ hotel.Rating_average }}</b></em> 
                  <span v-if="hotel.Rating_average >= 9">Exceptional</span>
                  <span v-if="hotel.Rating_average >= 8 && hotel.Rating_average < 9">Exellent</span>
                  <span v-if="hotel.Rating_average >= 7 && hotel.Rating_average < 8">Very good</span>
                  <span v-if="hotel.Rating_average >= 6 && hotel.Rating_average < 7">Good</span>
            </p>

            <div align="center"><router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'"><v-btn small color="white blue--text"><v-icon>map</v-icon> &nbsp; map</v-btn></router-link></div>
            <div align="center"><router-link :to="hotel.Hotel_namekey + '/hotel/' + hotel.Citykey + '-' + hotel.Countryisocode_lower + '.html'"><v-btn small color="white blue--text">more &nbsp; <v-icon>arrow_right_alt</v-icon></v-btn></router-link></div>
          </v-flex>
        </v-layout>
    </v-card>

    <v-card class="grey darken-5 white--text py-2 px-3 subheading">
        Recommended in {{ city }}, {{ country }}
    </v-card>
  </div>
</template>

<script>
  // import axios from 'axios'
  export default {
    props:{ hotelStar: Number, listing_type: String, city: String, country: String, hotel: Array, hotelLoading: Boolean },
  }
</script>

<style>
#hotel-in-listing .img{
  background-size: cover!important;
}
  #hotel-in-listing .address i,
  #hotel-in-listing .fitur i{
    font-size: 16px!important;
    margin-top: 1px;
    margin-right: 5px;
  }
  #hotel-in-listing .address{
    display: table;
    width: 100%;
  }  
  #hotel-in-listing .address i,
  #hotel-in-listing .address span{
    float: left;
  }
  #hotel-in-listing .address span{
    display: table;
    max-width: calc(100% - 30px);
    width: auto;
    margin-right: 8px;
  }

  #hotel-in-listing .stars{
    display: table;
    float: left;
  }
  #hotel-in-listing .stars i{
    font-size: 16px;
  }
  #hotel-in-listing .title,
  #hotel-in-listing .stars,
  #hotel-in-listing .address{
    margin-bottom: 8px;
  }
  #hotel-in-listing .itemGallery img{
    border: 8px solid #fff;
  }
  #hotel-in-listing .address .viewonmap{
    /*text-decoration: underline;*/
    font-size: 11px!important;
    /* font-weight: bold; */
    border-bottom: 1px dotted;
    white-space: nowrap;
    cursor: pointer;
  }
  #hotel-in-listing .img{
    position: relative;
  }
  #hotel-in-listing .img .heart{
    position: absolute;
    z-index: 1;
  }
  /*#hotel-in-listing .img .price{
    position: absolute;
    z-index: 0;
    right: -4px;
    bottom: 10px;
    background-color: #EE595D;
    font-weight: bold;
    padding: 5px 20px;
    color: white;
  }*/

  #hotel-in-listing .photos{
      position: absolute;
      right: 0;
      bottom: 5px;
      width: 100%;
      text-align: center;
      z-index: 1;
  }
  #hotel-in-listing .photos a{
    float: left;
    margin: 1%;
    width: 23%;
  }
  #hotel-in-listing .photos span{
    display: table;
    width: 100%;
    height: 55px;
    border: 1px solid white;
    background-size: cover;
  }
  #hotel-in-listing .badge-price{
      position: absolute;
      right: 0;
      bottom: 90px;
      width: 80px;
      text-align: center;
      z-index: 1;
  }
  #hotel-in-listing .badge-price .discount-ribbon {
      background: #EE595D;
      /*width: 50%;*/
      padding: 4px 8px;  
      float: right;
      line-height: 1px;
      font-weight: bold;
  }
    #hotel-in-listing .badge-price .discount-ribbon .discount-ribbon-discount{
        display: block;
        font-size: 16px;
        line-height: 20px;
        color: #fff;
    }
    #hotel-in-listing .discount-ribbon-text{
        font-size: 8px;
        line-height: 10px;
        color: #fff;
        text-transform: uppercase;
    }
  #hotel-in-listing .badge-price .price-ribbon {
      background: #fff;
      float: right;
      text-align: center;
      /*width: 50%;*/
      padding: 2px 8px;
  }
    #hotel-in-listing .price-ribbon-from{
        display: block;
        font-size: 10px;
        line-height: 5px;
        color: #777;
        display: none;
    }
    #hotel-in-listing .price-ribbon-currency{
        color: #EE595D;
        font-size: 10px;
    }
    #hotel-in-listing .price-ribbon-price{
        color: #EE595D;
        font-weight: 700;
        font-size: 14px;
    }


</style>
