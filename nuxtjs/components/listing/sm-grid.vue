<template>
	<div id="grid">

    <v-layout wrap hidden-xs-only v-if="loader">
      <v-flex xs12 sm6 md4 lg3 item v-for="i in 9" :key="i" px-1 py-2>
        <v-flex>
          <v-container fluid grid-list-md>
            <v-layout wrap>
              <v-flex xs4 sm12 img :style="'background-color: #eee; background-size: cover; height: 155px; margin-bottom: 9px;'"></v-flex>                  
              <v-flex xs8 sm12 white>
                <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 90%;"></v-progress-linear><br />
                <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="width: 40%;"></v-progress-linear><br />              
              </v-flex>
            </v-layout>
          </v-container>                
        </v-flex>
      </v-flex>                      
    </v-layout>

    <v-layout wrap v-if="!loader">
      <v-flex xs12 sm6 md4 lg3 item v-for="(item, i) in items" :key="i" px-1 py-2 v-if="item.Hotel_id != '' && item.Photo != ''">
        <v-card class="item">
          <v-container fluid grid-list-md>
            <v-layout wrap>
              <v-flex xs4 sm12 img :style="'background: url('+ item.Photo +') no-repeat; background-size: cover; height: 185px; margin-bottom: 2px;'">
                <div class="heart"><v-icon medium color="white">favorite_border</v-icon></div>
                <div class="price subheading">${{ item.Rates_from }}</div>
              </v-flex>                  
              <v-flex xs8 sm12>
                    <div class="subheading"><b><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode_lower + '.html'" style="color: #444;">{{ item.Hotel_name }}</router-link></b></div>

                    <p class="stars">
                        <v-icon color="orange" v-for="i in parseInt(item.Star_rating)" :key="i">star</v-icon>
                        <v-icon color="orange" v-if="parseInt(item.Star_rating) < item.Star_rating">star_half</v-icon>
                        <span class="caption"> &nbsp; (<b>{{ item.No_reviews }}</b> reviews)</span>
                    </p>

                    <p class="body-1 address">
                      <v-icon small>location_on</v-icon>                        
                      <span>{{ item.Addressline1 }}, {{ item.City }}, {{ item.Countryisocode }} &nbsp; <small class="viewonmap blue--text">view on map</small></span>                        
                    </p>
                    
                    <v-layout row wrap green--text>
                      <v-flex class="xs3"><v-icon small color="green">wifi</v-icon> wifi</v-flex>
                      <v-flex class="xs9"><v-icon small color="blue">check</v-icon> include breakfast</v-flex>
                    </v-layout>                    
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
              <v-flex xs6 py-0 px-0 mr-0 text-xs-right><router-link :to="item.Hotel_namekey + '/hotel/' + item.Citykey + '-' + item.Countryisocode_lower + '.html'"><v-icon small>more_vert</v-icon></router-link></v-flex>
            </v-flex>                   
        </v-card>
      </v-flex>                      
    </v-layout>
	</div>
</template>


<script>
  export default {
    props:{ items: Array, loader: Boolean },    
  }
</script>

<style type="text/css">
  #grid .rating_average span{
    font-size: 12px;
  }

  #grid .address i,
  #grid .fitur i{
    font-size: 16px!important;
    margin-bottom: 4px;
    margin-right: 5px;
  }
  #grid .address{
    display: table;
    width: 100%;
  }
  #grid .address i,
  #grid .address span{
    float: left;
  }
  #grid .address span{
    display: table;
    max-width: calc(100% - 30px);
    width: auto;
    margin-right: 8px;
  }  
  #grid .stars{
    white-space: nowrap;
    display: inline-flex;
  }
  #grid .stars i{
    font-size: 16px;
  }
  #grid .title,
  #grid .stars,
  #grid .address{
    margin-bottom: 8px;
  }
  #grid .address{
    display: table;
    width: 100%;
  }  
  #grid .address .viewonmap{
    /*text-decoration: underline;*/
    font-size: 11px!important;
    white-space: nowrap;
    cursor: pointer;
  }
  #grid .img{
    position: relative;
  }
  #grid .img .heart{
    position: absolute;
    z-index: 1;
  }
  #grid .img .price{
    position: absolute;
    z-index: 0;
    right: -4px;
    bottom: 10px;
    background-color: #EE595D;
    font-weight: bold;
    padding: 5px 20px;
    color: white;
  }

  @media only screen and (min-width: 600px) {
    #grid .card.item{
        margin: 4px!important;
        box-shadow: none;
        background-color: #fcfcfc;
        border-right: 1px solid #eee;
        border-bottom: 1px solid #eee;
    }
  }
  @media only screen and (max-width: 599px) {
    #grid .loader img{
      height: 90px;
    }
  }
</style>