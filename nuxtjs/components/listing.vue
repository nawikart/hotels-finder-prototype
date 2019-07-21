<template>
	<div id="listing_contents">
            <v-layout wrap v-if="loader && listing=='list'">
              <v-flex xs12 list v-for="i in 12" :key="i" px-1 py-2>
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
            <v-layout wrap v-if="!loader && listing=='list'">
              <v-flex xs12 list v-for="(item, i) in items" :key="i" px-1 py-2>
                <div >
                  <v-card class="item" v-if="listing=='list' && item.Hotel_id != '' && item.Photo != ''">
                      <v-layout wrap>

                        <v-flex my-1 xs4 img :style="'background: url('+ item.Photo +') no-repeat; background-size: cover;'">
                          <div class="heart mt-2 ml-2"><v-icon medium color="white">favorite_border</v-icon></div>
                          <div class="price subheading">${{ item.Rates_from }}</div>
                        </v-flex>   

                        <v-flex xs8 sm6 px-3 pt-1 pb-2 style="position: relative;">
                            
                            <p class="title"><span class="subheading"><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode + '.html'" style="color: #444;"><b>{{ item.Hotel_name }}</b></router-link> &nbsp; </span></p>

                            <p class="stars"><span>
                                <v-icon color="orange" v-for="i in parseInt(item.Star_rating)" :key="i">star</v-icon>
                                <v-icon color="orange" v-if="parseInt(item.Star_rating) < item.Star_rating">star_half</v-icon>
                            </span> <span class="caption"> &nbsp; (<b>{{ item.No_reviews }}</b> reviews)</span></p>

                            <p class="body-1 address"><v-icon small>location_on</v-icon>Uluwatu, Bali, ID &nbsp; <code class="viewonmap blue--text">view on map</code></p>
                            
                            <v-flex column green--text fitur>
                              <div><v-icon small color="green">wifi</v-icon> wifi</div>
                              <div><v-icon small color="blue">check</v-icon> order now pay later</div>
                              <div><v-icon small color="pink">fastfood</v-icon> include breakfast</div>
                            </v-flex>

                            <v-btn v-on:click="showDetailClicked(item.Hotel_id)" outline fab absolute right bottom color="blue" style="border: none; bottom: -5px; right: -5px;">
                                <v-icon v-if="showDetail !== item.Hotel_id">more_vert</v-icon>
                                <v-icon v-if="showDetail == item.Hotel_id" close="red">close</v-icon>
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

                          <div align="center"><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode + '.html'"><v-btn outline small color="blue lighten-3"><v-icon>map</v-icon> &nbsp; map</v-btn></router-link></div>

                          <div align="center"><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode + '.html'"><v-btn outline small color="orange lighten-2">more &nbsp; <v-icon>arrow_right_alt</v-icon></v-btn></router-link></div>
                        </v-flex>

                        <v-flex xs12 detailProgress v-if="detailProgress == item.Hotel_id">
                            <v-progress-linear :indeterminate="true" value="15" height="3" color="blue"></v-progress-linear>
                        </v-flex>
                      </v-layout>

                      <v-layout v-if="showDetail == item.Hotel_id" itemDetail wrap mt-2 style="padding: 20px 30px 10px; background-color: #fafafa;">
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
                            <v-btn outline color="blue" style="float: right;" class="hidden-sm-and-up">view more</v-btn>
                          </v-flex>                           
                      </v-layout>
                  </v-card>
                </div>
              </v-flex>
            </v-layout>



            <v-layout wrap v-if="loader && listing=='grid'">
              <v-flex xs12 sm6 md4 lg3 grid v-for="i in 8" :key="i" px-1 py-2>
                <v-flex class="item">
                  <v-container fluid grid-list-md>
                    <v-layout wrap>
                      <v-flex xs4 sm12 img :style="'background-color: #eee; background-size: cover; height: 185px; margin-bottom: 9px;'"></v-flex>                  
                      <v-flex xs8 sm12 white>
                        <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 90%;"></v-progress-linear><br />
                        <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 40%;"></v-progress-linear><br />
                        <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 70%;"></v-progress-linear><br />              
                      </v-flex>
                    </v-layout>
                  </v-container>                
                </v-flex>
              </v-flex>                      
            </v-layout>
            <v-layout wrap v-if="!loader && listing=='grid'">
              <v-flex xs12 sm6 md4 lg3 grid v-for="(item, i) in items" :key="i" px-1 py-2 v-if="item.Hotel_id != '' && item.Photo != ''">
                <v-card class="item">
                  <v-container fluid grid-list-md>
                    <v-layout wrap>
                      <v-flex xs4 sm12 img :style="'background: url('+ item.Photo +') no-repeat; background-size: cover; height: 185px; margin-bottom: 9px;'">
                        <div class="heart"><v-icon medium color="white">favorite_border</v-icon></div>
                        <div class="price subheading">${{ item.Rates_from }}</div>
                      </v-flex>                  
                      <v-flex xs8 sm12>
                        <div class="itemDes">
                            <div class="subheading"><b><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode + '.html'" style="color: #444;">{{ item.Hotel_name }}</router-link></b></div>

                            <p class="stars">
                                <v-icon color="orange" v-for="i in parseInt(item.Star_rating)" :key="i">star</v-icon>
                                <v-icon color="orange" v-if="parseInt(item.Star_rating) < item.Star_rating">star_half</v-icon>
                              </span> <span class="caption"> &nbsp; (<b>{{ item.No_reviews }}</b> reviews)</span></p>

                            <p class="body-1 address"><v-icon small>location_on</v-icon>Uluwatu, Bali, ID &nbsp; <code class="viewonmap">view on map</code></p>
                            
                            <v-layout row wrap green--text>
                              <v-flex class="xs3"><v-icon small color="green">wifi</v-icon> wifi</v-flex>
                              <v-flex class="xs9"><v-icon small color="blue">check</v-icon> include breakfast</v-flex>
                            </v-layout>
                        </div>                       
                      </v-flex>
                    </v-layout>
                  </v-container>

                    <v-flex footer py-0 px-1 sx12>                      
                      <v-flex py-0 px-0 xs6 rating_average><v-icon small>grade</v-icon> <em><b class="mr-1">{{ item.Rating_average }}</b></em> 
                                <span v-if="item.Rating_average >= 9">Exceptional</span>
                                <span v-if="item.Rating_average >= 8 && item.Rating_average < 9">Exellent</span>
                                <span v-if="item.Rating_average >= 7 && item.Rating_average < 8">Very good</span>
                                <span v-if="item.Rating_average >= 6 && item.Rating_average < 7">Good</span>
                      </v-flex>
                      <v-flex xs6 py-0 px-0 mr-0 text-xs-right><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode + '.html'"><v-icon small>more_vert</v-icon></router-link></v-flex>
                    </v-flex>                   
                </v-card>
              </v-flex>                      
            </v-layout>



            <v-layout wrap hidden-xs-only grid2 v-if="loader && listing=='grid2'">
              <v-flex xs12 sm6 item md4 v-for="i in 9" :key="i" px-1 py-2>
                <v-flex class="item">
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
            <v-layout wrap hidden-xs-only grid2 v-if="!loader && listing=='grid2'">
              <v-flex xs12 sm6 item md4 v-for="(item, i) in items" :key="i" v-if="item.Hotel_id != '' && item.Photo != ''">
                <v-layout wrap white style="padding: 20px 3% 30px">
                  <v-flex xs4 sm12 img :style="'background: url('+ item.Photo +') no-repeat; background-size: cover; height: 280px; margin-bottom: 9px;'">
                    <div class="heart ml-2 mt-2"><v-icon medium color="white">favorite_border</v-icon></div>

                    <div class="badge-price"><div class="discount-ribbon middle"><span class="discount-ribbon-discount">45%</span><span class="discount-ribbon-text">Discount</span></div><div class="price-ribbon middle"><span class="price-ribbon-from">Starting at</span>
                    <div style="display: table; margin: 0 auto;"><span class="price-ribbon-currency">USD</span>&nbsp;<span class="price-ribbon-price">287</span></div></div></div>

                    <div style="width: 100%; padding: 12px 10px 10px; position: absolute; bottom: 0; background-color: rgba(0,0,0,.5); color: white;">
                        <div class="body-1"><b>{{ item.Hotel_name }}</b></div>

                        <div class="stars">
                            <v-icon small color="yellow">star</v-icon>
                            <v-icon small color="yellow">star</v-icon>
                            <v-icon small color="yellow">star</v-icon>
                            <v-icon small color="yellow">star</v-icon>
                            <v-icon small color="yellow">star</v-icon>
                          </span> <span class="caption"> &nbsp; (<b>192</b> reviews)</span></div>
                    </div>
                  </v-flex>                  
                  <v-flex xs8 sm12>
                    <div class="itemDes">
                        <p class="body-1 address"><v-icon small>location_on</v-icon>Uluwatu, Bali, ID &nbsp; <code class="viewonmap">view on map</code></p>
                        
                        <v-layout row wrap green--text>
                          <v-flex class="xs3"><v-icon small color="green">wifi</v-icon> wifi</v-flex>
                          <v-flex class="xs9"><v-icon small color="blue">check</v-icon> <nowrap>order now pay later</nowrap></v-flex>
                        </v-layout>

                        <v-btn outline fab absolute right bottom color="blue" style="border: none; bottom: -5px; right: -5px;">
                            <v-icon>more_vert</v-icon>
                        </v-btn>                                       
                    </div>                       
                  </v-flex>
                </v-layout>
              </v-flex>
            </v-layout>

	        <v-container v-if="!loader">
	          <nav class="pagination">                                            
	              <ul> 
	                  <li><router-link :to="'/' + citykey + '-' + countryisocode + '-hotels-p' + (P - 1) + '.html'" v-if="P != 1">PREV</router-link></li>
	                  <li><router-link v-bind:class="{ active: P == 1 }" :to="'/' + citykey + '-' + countryisocode + '-hotels.html'">1</router-link></li>
	                  <li v-if="P > 4 && Ps > PsMax"><a>...</a></li>
	                  <li v-for="i in Pss" :key="i"><router-link v-bind:class="{ active: P == (Pstart + i - 1) }" :to="'/' + citykey + '-' + countryisocode + '-hotels-p' + (Pstart + i - 1) + '.html'">{{ Pstart + i - 1 }}</router-link></li>
	                  <li v-if="P < Ps - 6 && Ps > PsMax"><a>...</a></li>
	                  <li><router-link v-bind:class="{ active: P == Ps }" :to="'/' + citykey + '-' + countryisocode + '-hotels-p' + Ps + '.html'"><b>{{ Ps }}</b></router-link></li>
	                  <li><router-link :to="'/' + citykey + '-' + countryisocode + '-hotels-p' + (P + 1) + '.html'" v-if="P != Ps">NEXT</router-link></li>
	              </ul>
	          </nav>
	        </v-container>
	</div>
</template>


<script>
  import axios from 'axios'
  export default {
    data: () => ({
      windowWidth: null,
      listFilter: true,
      gridFilter: false,
      listing: 'grid2',
      showDetail: false,
      detailProgress: false,
      showFilter: false,
      itemDetailTab: 'overview',      
      itemDetail: [],
      loader: true,

      citykey: false,
      countryisocode: false,

      //PAGINATION
      P: 0,
      Pstart: 0,
      PsMax: 10,
      Ps: 0,
      Pss: 0,
      ////////////
      items: [],
    }),

        methods: {
          showDetailClicked (Hotel_id) {
              if(this.showDetail == Hotel_id){
                  this.showDetail = false
              }else{                  
                  this.detailProgress = Hotel_id
                  axios.get(this.$store.state.apiPath +'/front?mod=hotel&hotel_id='+ Hotel_id).then(response => {
                    this.showDetail = Hotel_id
                    this.itemDetail = response.data
                    this.detailProgress = false
                  })
              }         
          },

          listListingClicked(){
              this.listing = 'list'
              this.listFilter = true
              this.gridFilter = false
          },

          gridListingClicked(){
              this.listing = 'grid'
              this.listFilter = false
              this.gridFilter = false
          },

          list(){
            this.loader = true
            // axios.get(this.$store.state.apiPath +'/front?mod=hotels&city_cc='+ this.$route.params.city).then(response => {
            axios.get(this.$store.state.apiPath +'/front?mod=hotels&city_cc=bali-id-hotels.html').then(response => {
              this.items = response.data.Hotels
              this.loader = false
              console.log(response.data)

              this.P = response.data.P
              this.Ps = response.data.Ps
              this.citykey = response.data.Citykey
              this.countryisocode = response.data.Countryisocode

              /////////PAGINATION///////////////////////
              if(this.PsMax > this.Ps){
                  this.PsMax = this.Ps
              }
              this.Pss = this.PsMax - 2 // dikurangi first dan last number
              this.Pstart = this.P - 2
              if(this.Pstart < 2){
                  this.Pstart = 2
              }
              if(this.Pstart > (this.Ps - this.PsMax + 2)){
                  this.Pstart = this.Ps - this.PsMax + 2             
              }
              //////////////////////////////////////////
              return false;
            })
          }
        },
         beforeMount(){
            this.windowWidth = window.innerWidth
            this.list()           
         }     
  }
</script>

<style type="text/css">
  a{
    text-decoration: none;
  }
  #listing .grid .rating_average span{
    font-size: 12px;
  }

  .tabs button{
    font-size: 10px;
    margin: 0 8px 0 0;
    padding: 2px;
  } 
  .toolbar{
    margin-bottom: 10px!important;
  }

  .address i,
  .fitur i{
    font-size: 16px!important;
    margin-bottom: 4px;
    margin-right: 5px;
  }
  .start{
    white-space: nowrap;
    display: inline-flex;
  }
  .stars i{
    font-size: 16px;
  }
  .title,
  .stars,
  .address{
    margin-bottom: 8px;
  }
  .itemGallery img{
    border: 8px solid #fff;
  }
  .address .viewonmap{
    /*text-decoration: underline;*/
    font-size: 10px!important;
    font-weight: bold;
    white-space: nowrap;
    cursor: pointer;
  }
  .img{
    position: relative;
  }
  .img .heart{
    position: absolute;
    z-index: 1;
  }
  .img .price{
    position: absolute;
    z-index: 0;
    right: -4px;
    bottom: 10px;
    background-color: #EE595D;
    font-weight: bold;
    padding: 5px 20px;
    color: white;
  }

  .detailProgress{
    position: relative;
  }
  .detailProgress .progress-linear{
    position: absolute;
    top: -16px;
  }


.grid2 .img{
  position: relative;
}
.grid2 .img .heart{
  position: absolute;
  z-index: 1;
}
  .grid2 .badge-price{
      position: absolute;
      right: -2px;
      bottom: 80px;
      width: 200px;
      text-align: center;
      z-index: 99;
  }
  .grid2 .badge-price .discount-ribbon {
      background: #EE595D;
      width: 50%;
      height: 50px;    
      float: right;
      line-height: 90%;
      font-weight: bold;
  }
    .grid2 .badge-price .discount-ribbon .discount-ribbon-discount{
        display: block;
        font-size: 20px;
        line-height: 20px;
        color: #fff;
        margin-top: 9px;
    }
    .grid2 .discount-ribbon-text{
        font-size: 10px;
        line-height: 10px;
        color: #fff;
        text-transform: uppercase;
    }
  .grid2 .badge-price .price-ribbon {
      background: #fff;
      float: left;
      text-align: center;
      width: 50%;
      height: 50px;
  }
    .grid2 .price-ribbon-from{
        display: block;
        font-size: 10px;
        line-height: 5px;
        color: #777;
        margin-top: 12px;
    }
    .grid2 .price-ribbon-currency{
        color: #EE595D;
        font-size: 12px;
    }
    .grid2 .price-ribbon-price{
        color: #EE595D;
        font-weight: 700;
        font-size: 20px;
    }

.grid2 .itemDes{
  padding: 10px;
  position: relative;
}
.grid2 .stars i.icon{
  font-size: 14px;
}
.grid2 .itemDes i.icon{
  font-size: 16px;
  margin-top: -4px;
}

.grid2 .subheading,
.grid2 .stars,
.grid2 .address{
  margin-bottom: 8px;
}
.grid2 .address .viewonmap{
  font-size: 10px!important;
  font-weight: bold;
  white-space: nowrap;
  cursor: pointer;
}

  nav.pagination {
      width: auto;
      display: table;
      float: right;
      clear: both;
      margin-top: 50px;
      margin-bottom: 0;
      line-height: 180%;
      padding: 0 20px;
  }
  nav.pagination ul{
    margin: 0;
    list-style: none;
    padding: 0;
  }
  nav.pagination ul li{
    float: left;
    margin: 3px 1px;
  }
  nav.pagination ul li a {
      background-color: #fff;
      padding: 7px 15px;
      border: 1px solid rgba(72,139,248, .5);
  }
  nav.pagination ul li a.active {
      background-color: rgba(72,139,248, .5); 
      color: #fff;
  }

  @media only screen and (min-width: 600px) {
    .grid .card.item{
        margin: 4px!important;
        box-shadow: none;
        background-color: #fcfcfc;
        border-right: 1px solid #eee;
        border-bottom: 1px solid #eee;
    }
  }
  @media only screen and (max-width: 599px) {
    .loader img{
      height: 90px;
    }
    .itemDetail .tabs button{
      margin: 1px;
    }  
    .overviewText{
      margin-left: 0!important;
    }  
  }
</style>