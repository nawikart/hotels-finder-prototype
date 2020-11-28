<template>
	<div id="list">
      <v-layout wrap v-if="loader">
        <v-flex xs12 item v-for="i in 12" :key="i" px-1 py-2>
          <div >  
            <v-layout wrap loader>
              <v-flex xs4 img style="height: 175px;background-color: #fafafa;">
              </v-flex>
              <v-flex xs8 pl-4>
                  <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 70%;"></v-progress-linear><br />
                  <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 40%;"></v-progress-linear><br />
                  <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 95%;"></v-progress-linear><br />
              </v-flex>                                        
            </v-layout>
          </div>
        </v-flex>
      </v-layout>

      <v-layout wrap v-if="!loader">
        <v-flex xs12 item v-for="(item, i) in items" :key="i" px-1 py-2 v-if="item.Hotel_id != '' && item.Photo != '' && item.Hotel_id != hotelExclude">
          <div>
            <v-card>
                <v-layout wrap>
                  <v-flex my-1 xs4 img>
                    <v-img
                        :src="item.Photo1"
                        :lazy-src="$store.state.mediaPath +'/media/images/loading.gif'"
                        aspect-ratio="1.7"
                        class="grey lighten-2" />

                      <div class="heart mt-2 ml-2"><v-icon medium color="white">favorite_border</v-icon></div>
                      <div class="badge-price" v-if="item.Rates_from != 0 || item.dailyRate">

                        <div class="discount-ribbon middle" v-if="item.discountPercentage && item.discountPercentage != 0"><span class="discount-ribbon-discount">{{ item.discountPercentage }}%</span><span class="discount-ribbon-text">Discount</span></div>

                        <div class="price-ribbon middle"><span class="price-ribbon-from">Starting at</span>
                          <div style="display: table; margin: 0 auto;"><span class="price-ribbon-currency" v-if="item.currency">{{ item.currency }}</span>&nbsp;
                            <span class="price-ribbon-price" v-if="item.dailyRate">{{ item.dailyRate }}</span>
                            <span class="price-ribbon-price" v-if="!item.dailyRate">${{ item.Rates_from }}</span>
                          </div>
                        </div>
                      </div>
                  </v-flex>   

                  <v-flex xs8 sm6 px-3 pt-1 pb-2 style="position: relative;">
                      
                      <p><span class="subheading">
                        <router-link :to="item.Hotel_name_key + '/hotel/' + item.City_key + '-' + item.Countryisocode.toLowerCase() + '.html'" style="color: #444;"><b>{{ item.Hotel_name }}</b></router-link></span></p>

                      <p><stars :Star_rating="item.Star_rating"></stars> 
                      <span class="caption"> &nbsp; (<b>{{ item.No_reviews }}</b> reviews)</span></p>
                      <p class="body-1 address">
                        <v-icon small>location_on</v-icon>                        
                        <span>{{ item.Addressline1 }}, {{ item.City }}, {{ item.Countryisocode.toLowerCase() }} &nbsp; <small class="viewonmap blue--text">view on map</small></span>                        
                      </p>
                      
                      <v-layout class="pl-1" row wrap green--text fitur v-if="!item.freeWifi || !item.includeBreakfast">
                        <div class="xs5 sm12 mr-3"><v-icon small color="green">wifi</v-icon> wifi</div>
                        <div class="xs7 sm12"><v-icon small color="green">check</v-icon> include breakfast</div>
                      </v-layout>

                      <v-layout class="pl-1" row wrap green--text fitur v-if="item.freeWifi && item.includeBreakfast">
                        <div class="xs5 sm12 mr-3"><v-icon small color="green">wifi</v-icon> 
                            <span v-if="item.freeWifi">Wifi</span>
                            <strike v-if="!item.freeWifi">Wifi</strike>
                        </div>
                        <div class="xs7 sm12">
                            <v-icon small color="green" v-if="item.includeBreakfast">check</v-icon> 
                            <span v-if="item.includeBreakfast" class="green--text">Include Breakfast</span>
                            <strike v-if="!item.includeBreakfast" class="grey--text">Include Breakfast</strike>
                        </div>
                      </v-layout>                      

                      <v-btn v-if="showDetail !== item.Hotel_id" v-on:click="showDetailClicked(item.Hotel_id)" outline fab absolute right bottom color="blue" style="border: none; bottom: 17px; right: -5px;">
                          <v-icon>more_vert</v-icon>
                      </v-btn>                                                 
                  </v-flex>

                  <v-flex xs2 pt-4 my-1 hidden-xs-only style="background-color: #f6f6f6;">
                    <p align="center" class="rating_average"><v-icon small>grade</v-icon><em><b class="mr-1">{{ item.Rating_average }}</b></em> 

                          <span v-if="item.Rating_average >= 9">Exceptional</span>
                          <span v-if="item.Rating_average >= 8 && item.Rating_average < 9">Exellent</span>
                          <span v-if="item.Rating_average >= 7 && item.Rating_average < 8">Very good</span>
                          <span v-if="item.Rating_average >= 6 && item.Rating_average < 7">Good</span>

                    </p>
                    <!-- <p align="center"><b class="caption red--text">Our last room at this price</b></p> -->

                    <div align="center"><router-link :to="item.Hotel_name_key + '/hotel/' + item.City_key + '-' + item.Countryisocode.toLowerCase() + '.html'"><v-btn outline small color="orange lighten-1"><v-icon>map</v-icon> &nbsp; map</v-btn></router-link></div>

                    <div align="center"><router-link :to="item.Hotel_name_key + '/hotel/' + item.City_key + '-' + item.Countryisocode.toLowerCase() + '.html'"><v-btn outline small color="blue lighten-2">more &nbsp; <v-icon>arrow_right_alt</v-icon></v-btn></router-link></div>
                  </v-flex>

                  <v-flex xs12 detailProgress v-if="detailProgress == item.Hotel_id">
                      <v-progress-linear :indeterminate="true" value="15" height="3" color="blue"></v-progress-linear>
                  </v-flex>
                </v-layout>

                <v-layout v-if="showDetail == item.Hotel_id" itemDetail wrap mt-2 style="padding: 35px 30px 10px; background-color: #fafafa;">

                    <v-btn v-on:click="showDetailClicked(item.Hotel_id)" outline fab absolute right top color="blue" style="border: none; top: -10px; right: -5px;">
                        <v-icon close="red">close</v-icon>
                    </v-btn> 

                    <v-flex xs12 mb-4 tabs>
                      <v-btn dark color="pink" :outline="itemDetailTab !== 'overview'" v-on:click="itemDetailTab = 'overview'" small>overview</v-btn>
                      <v-btn dark color="pink" :outline="itemDetailTab !== 'compare'" v-on:click="itemDetailTab = 'compare'" small>price compare</v-btn>
                      <v-btn dark color="pink" v-if="item.No_reviews > 0" :outline="itemDetailTab !== 'reviews'" v-on:click="itemDetailTab = 'reviews'" small>reviews</v-btn>

                      <v-btn dark color="blue" outline small style="float: right;" class="m-0 hidden-xs-only">view more</v-btn>
                    </v-flex>

                    <v-layout wrap v-if="itemDetailTab == 'overview'">
                      <v-flex xs12>
                        <v-flex overviewText>
                          <p class="title pink--text lighten-1">Hotel Overview</p>
                          <p class="body-1" style="color: #555;">{{ itemDetail.Overview }}</p>      
                        </v-flex>                                                                               
                      </v-flex>
                      <v-flex xs12>
                        <v-layout row wrap mb-4 itemGallery>
                          <v-flex item xs6 sm2>
                            <a href="#"><img :src="itemDetail.Photo1" width="100%" class="d-flex" /></a>
                          </v-flex>
                          <v-flex item xs6 sm2>
                            <a href="#"><img :src="itemDetail.Photo2" width="100%" class="d-flex" /></a>
                          </v-flex>
                          <v-flex item xs6 sm2>
                            <a href="#"><img :src="itemDetail.Photo3" width="100%" class="d-flex" /></a>
                          </v-flex>
                          <v-flex item xs6 sm2>
                            <a href="#"><img :src="itemDetail.Photo4" width="100%" class="d-flex" /></a>
                          </v-flex>                                                            
                        </v-layout>
                      </v-flex>                            
                    </v-layout>

                    <v-layout wrap v-if="itemDetailTab == 'compare'">
                      <p>compare....</p>
                    </v-layout>
                    <v-layout wrap v-if="itemDetailTab == 'reviews'">
                      <p>reviews....</p>
                    </v-layout>


                    <v-flex xs12 mt-4>
                      <router-link :to="item.Hotel_name_key + '/hotel/' + item.City_key + '-' + item.Countryisocode.toLowerCase() + '.html'"><v-btn outline color="blue" style="float: right;" class="hidden-sm-and-up">view more</v-btn></router-link>
                    </v-flex>                           
                </v-layout>
            </v-card>
          </div>
        </v-flex>
      </v-layout>
	</div>
</template>


<script>
  import axios from 'axios'
  export default {

    props:{ items: Array, hotelExclude: String, loader: Boolean },

    data: () => ({
      showDetail: false,
      detailProgress: false,
      itemDetailTab: 'overview',      
      itemDetail: [],
    }),

        methods: {
          showDetailClicked (Hotel_id) {
              if(this.showDetail == Hotel_id){
                  this.showDetail = false
              }else{                  
                  this.detailProgress = Hotel_id
                  axios.get(this.$store.state.apiPath +'/hotel/'+ Hotel_id).then(response => {
                    this.showDetail = Hotel_id
                    this.itemDetail = response.data
                    this.detailProgress = false
                  })
              }         
          }
        }  
  }
</script>

<style type="text/css">
  #list .tabs button{
    font-size: 10px;
    margin: 0 8px 0 0;
    padding: 2px;
  }

  #list .address i,
  #list .fitur i{
    font-size: 16px!important;
    margin-top: 1px;
    margin-right: 5px;
  }
  #list .address{
    display: table;
    width: 100%;
  }  
  #list .address i,
  #list .address span{
    float: left;
  }
  #list .address span{
    display: table;
    max-width: calc(100% - 30px);
    width: auto;
    margin-right: 8px;
  }

  #list .stars{
    display: table;
    float: left;
  }
  #list .stars i{
    font-size: 16px;
  }
  #list .title,
  #list .stars,
  #list .address{
    margin-bottom: 8px;
  }
  #list .itemGallery img{
    border: 8px solid #fff;
  }
  #list .address .viewonmap{
    /*text-decoration: underline;*/
    font-size: 11px!important;
    /* font-weight: bold; */
    border-bottom: 1px dotted;
    white-space: nowrap;
    cursor: pointer;
  }
  #list .img{
    position: relative;
  }
  #list .img .heart{
    position: absolute;
    /* z-index: 1; */
  }
  /*#list .img .price{
    position: absolute;
    z-index: 0;
    right: -4px;
    bottom: 10px;
    background-color: #EE595D;
    font-weight: bold;
    padding: 5px 20px;
    color: white;
  }*/

  #list .badge-price{
      position: absolute;
      right: 0;
      bottom: 10px;
      width: 80px;
      text-align: center;
      /* z-index: 1; */
  }
  #list .badge-price .discount-ribbon {
      background: #EE595D;
      /*width: 50%;*/
      padding: 4px 8px;  
      float: right;
      line-height: 1px;
      font-weight: bold;
  }
    #list .badge-price .discount-ribbon .discount-ribbon-discount{
        display: block;
        font-size: 16px;
        line-height: 20px;
        color: #fff;
    }
    #list .discount-ribbon-text{
        font-size: 8px;
        line-height: 10px;
        color: #fff;
        text-transform: uppercase;
    }
  #list .badge-price .price-ribbon {
      background: #fff;
      float: right;
      text-align: center;
      /*width: 50%;*/
      padding: 2px 8px;
  }
    #list .price-ribbon-from{
        display: block;
        font-size: 10px;
        line-height: 5px;
        color: #777;
        display: none;
    }
    #list .price-ribbon-currency{
        color: #EE595D;
        font-size: 10px;
    }
    #list .price-ribbon-price{
        color: #EE595D;
        font-weight: 700;
        font-size: 14px;
    }


  #list .detailProgress{
    position: relative;
  }
  #list .detailProgress .progress-linear{
    position: absolute;
    top: -16px;
  }

  #list .itemDetail{
    position: relative;
  }


@media only screen and (max-width: 599px) {
  #list .loader img{
    height: 90px;
  }
  #list .itemDetail .tabs button{
    margin: 1px;
  }  
  #list .overviewText{
    margin-left: 0!important;
  }  
}
</style>