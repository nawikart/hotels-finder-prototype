<template>
	<div id="grid2">

    <v-layout wrap hidden-xs-only v-if="loader">
      <v-flex xs12 sm6 item md4 v-for="i in 9" :key="i" px-1 py-2>
        <v-flex>
          <v-container fluid grid-list-md>
            <v-layout wrap>
              <v-flex xs4 sm12 img :style="'background-color: #eee; background-size: cover; height: 225px; margin-bottom: 9px;'"></v-flex>                  
              <v-flex xs8 sm12 white>
                <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 90%;"></v-progress-linear><br />
                <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 40%;"></v-progress-linear><br />              
              </v-flex>
            </v-layout>
          </v-container>                
        </v-flex>
      </v-flex>                      
    </v-layout>

    <v-layout wrap hidden-xs-only v-if="!loader">
      <v-flex xs12 sm6 item md4 v-for="(item, i) in items" :key="i"><!-- v-if="item.hotelId != '' && item.imageURL != undefined" -->
        <v-layout wrap white style="padding: 20px 3% 30px">
          <v-flex xs4 sm12 img :style="'background: url('+ item.Photo +') no-repeat; background-position: center; background-size: cover!important; height: 240px; margin-bottom: 9px;'">
            <div class="heart ml-2 mt-2"><v-icon medium color="white">favorite_border</v-icon></div>

            <div class="badge-price" v-if="item.Rates_from != 0 || item.dailyRate">

              <div class="discount-ribbon middle" v-if="item.discountPercentage && item.discountPercentage != 0"><span class="discount-ribbon-discount">{{ item.discountPercentage }}%</span><span class="discount-ribbon-text">Discount</span></div>

              <div class="price-ribbon middle"><span class="price-ribbon-from">Starting at</span>
                <div style="display: table; margin: 0 auto;"><span class="price-ribbon-currency" v-if="item.currency">{{ item.currency }}</span>&nbsp;
                  <span class="price-ribbon-price" v-if="item.dailyRate">{{ item.dailyRate }}</span>
                  <span class="price-ribbon-price" v-if="!item.dailyRate">${{ item.Rates_from }}</span>
                </div>
              </div>
            </div>

            <router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode_lower + '.html'" style="width: 100%; padding: 12px 10px 10px; position: absolute; bottom: 0; background-color: rgba(0,0,0,.5); color: white;">
                <div class="body-1"><b style="margin-right: 50px; display: table;">{{ item.Hotel_name }} {{ item.Star_rating }}</b></div>

                <stars v-if="!item.starRating || item.starRating == 0 || item.starRating == ''" :Star_rating="item.Star_rating" style="float: left; margin-right: 10px;"></stars>
                <stars v-if="item.starRating && !(item.starRating == 0 || item.starRating == '')" :Star_rating="item.starRating" style="float: left; margin-right: 10px;"></stars>
                 <span class="caption" style="float: left; margin-top: 4px;">{{ item.reviewCount }} reviews</span>

                <span style="position: absolute; bottom: 15px; right: 10px; text-align: right;"><b class="title">{{ item.Rating_average }}</b><br />
                          <span v-if="item.Rating_average >= 9">Exceptional</span>
                          <span v-if="item.Rating_average >= 8 && item.Rating_average < 9">Exellent</span>
                          <span v-if="item.Rating_average >= 7 && item.Rating_average < 8">Very good</span>
                          <span v-if="item.Rating_average >= 6 && item.Rating_average < 7">Good</span>
                </span>
            </router-link>
          </v-flex>                  
          <v-flex xs8 sm12>
            <div class="itemDes">
                <p class="body-1 address">
                  <v-icon small>location_on</v-icon>                        
                  <span>{{ item.Addressline1 }}, {{ item.City }}, {{ item.Countryisocode }} &nbsp; <small class="viewonmap blue--text">view on map</small></span>                        
                </p>              
                <v-layout row wrap green--text v-if="!item.freeWifi || !item.includeBreakfast">
                  <v-flex class="xs3"><v-icon small color="green">wifi</v-icon> Wifi</v-flex>
                  <v-flex class="xs9"><v-icon small color="blue">check</v-icon> Include Breakfast</v-flex>
                </v-layout>

                <v-layout row wrap green--text v-if="item.freeWifi && item.includeBreakfast">
                  <v-flex class="xs3"><v-icon small color="green">wifi</v-icon> 
                    <span v-if="item.freeWifi">Wifi</span>
                    <strike v-if="!item.freeWifi">Wifi</strike>
                  </v-flex>
                  <v-flex class="xs9">
                            <v-icon small color="green" v-if="item.includeBreakfast">check</v-icon> 
                            <span v-if="item.includeBreakfast" class="green--text">Include Breakfast</span>
                            <strike v-if="!item.includeBreakfast" class="grey--text">Include Breakfast</strike>
                  </v-flex>
                </v-layout>

                <v-btn outline fab absolute right bottom color="blue" style="border: none; bottom: -5px; right: -5px;">
                    <v-icon>more_vert</v-icon>
                </v-btn>                                       
            </div>                       
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
	</div>
</template>


<script>
  import axios from 'axios'
  export default {
    props:{ items: Array, loader: Boolean },   
  }
</script>

<style type="text/css">
#grid2 .img{
  position: relative;
}
#grid2 .img .heart{
  position: absolute;
  z-index: 1;
}
  #grid2 .badge-price{
      position: absolute;
      right: 0;
      bottom: 100px;
      width: 200px;
      text-align: center;
      z-index: 1;
  }
  #grid2 .badge-price .discount-ribbon {
      background: #EE595D;
      width: 50%;
      height: 50px;    
      float: right;
      line-height: 90%;
      font-weight: bold;
  }
    #grid2 .badge-price .discount-ribbon .discount-ribbon-discount{
        display: block;
        font-size: 20px;
        line-height: 20px;
        color: #fff;
        margin-top: 9px;
    }
    #grid2 .discount-ribbon-text{
        font-size: 10px;
        line-height: 10px;
        color: #fff;
        text-transform: uppercase;
    }
  #grid2 .badge-price .price-ribbon {
      background: #fff;
      float: right;
      text-align: center;
      width: 50%;
      height: 50px;
  }
    #grid2 .price-ribbon-from{
        display: block;
        font-size: 10px;
        line-height: 5px;
        color: #777;
        margin-top: 12px;
    }
    #grid2 .price-ribbon-currency{
        color: #EE595D;
        font-size: 12px;
    }
    #grid2 .price-ribbon-price{
        color: #EE595D;
        font-weight: 700;
        font-size: 20px;
    }

#grid2 .itemDes{
  padding: 10px;
  position: relative;
}
#grid2 .stars i.icon{
  font-size: 16px;
}
#grid2 .itemDes i.icon{
  font-size: 20px;
}

#grid2 .subheading,
#grid2 .stars,
#grid2 .address{
  margin-bottom: 8px;
}
#grid2 .address{
  display: table;
  width: 100%;
}
#grid2 .viewonmap{
  white-space: nowrap;
}
#grid2 .address i,
#grid2 .address span{
  float: left;
}
#grid2 .address span{
  display: table;
  max-width: calc(100% - 30px);
  width: auto;
  margin-right: 8px;
}
@media only screen and (max-width: 599px) {
  #grid2 .loader img{
    height: 90px;
  }
}
</style>