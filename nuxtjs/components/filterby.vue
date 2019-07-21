<template>
  <v-layout wrap mainFilter>

      <v-flex xs12 sm6 md4 :order-xs2="filterbyreorder">
          <v-subheader>Discount Only</v-subheader>
          <v-card-text style="display: table; width: 100%;">
            <v-switch class="discountOnly" color="pink" v-model="discountOnly" style="float: left; width: 50px;"></v-switch>  
            <span style="float: left; width: calc(100% - 55px);" v-if="discountOnly">Hotels with discount only</span>
            <span style="float: left; width: calc(100% - 55px);" v-if="!discountOnly">Hotels with and without discount</span>
          </v-card-text>

          <v-subheader>Minimum Star Rating</v-subheader>
          <v-card-text>
            <v-radio-group v-model="minimumStarRating">
              <v-radio color="orange" value="0" label="All"></v-radio>  

              <div class="radioRow">
                  <v-radio color="orange" value="1" class="starsRadio"></v-radio> 
                  <stars :Star_rating="1"></stars> <span class="caption">( >= 1 star )</span>
              </div>    

              <div v-for="i in 3" class="radioRow">
                  <v-radio color="orange" :value="(i+1)" class="starsRadio"></v-radio> 
                  <stars :Star_rating="(i+1)"></stars> <span class="caption">( >= {{ i+1 }} stars )</span>
              </div>

              <div class="radioRow">
                  <v-radio color="orange" value="5" class="starsRadio"></v-radio> 
                  <stars :Star_rating="5"></stars> <span class="caption">( 5 stars only )</span>
              </div>
            </v-radio-group>
          </v-card-text>                    
      </v-flex>

      <v-flex xs12 md4 :order-xs3="filterbyreorder">
          <v-subheader>Min. Review Score <span class="caption hint">{{ minimumReviewScore }} – 10</span></v-subheader>
          <v-card-text pl-2>
            <v-slider v-model="minimumReviewScore" min="0" max="10" step="1" color="red" thumb-color="red" thumb-label></v-slider>
          </v-card-text>

          <v-subheader>Daily Rate (usd) <span class="caption hint">${{ dailyRate_min }} – ${{ dailyRate_max }}</span></v-subheader>
          <v-card-text>
            <v-slider v-model="dailyRate_min" label="Min." color="green" min="0" max="900" value="10" step="10" thumb-color="green" thumb-label></v-slider>
            <v-slider v-model="dailyRate_max" label="Max." color="green" min="10" max="1000" value="1000" step="10" thumb-color="green" thumb-label></v-slider>
          </v-card-text>
      </v-flex>

      <v-flex xs12 sm6 md4 :order-xs1="filterbyreorder">
              <v-subheader>Sort by</v-subheader>
              <v-card-text class="sortby" v-if="!showFilterAgoda">
                  <v-radio-group v-model="sortby">
                  <v-radio color="indigo accent-2"
                    value="Recommended"
                    label="Recommended"
                  ></v-radio>
                  <v-radio color="indigo accent-2"
                    value="PriceDesc"
                    label="Daily Rate Desc"
                  ></v-radio>
                  <v-radio color="indigo accent-2"
                    value="PriceAsc"
                    label="Daily Rate Asc"
                  ></v-radio>
                  <v-radio color="indigo accent-2"
                    value="StarRatingDesc"
                    label="Star Rating Desc"
                  ></v-radio> 
                  <v-radio color="indigo accent-2"
                    value="StarRatingAsc"
                    label="Star Rating Asc"
                  ></v-radio>
                  <v-radio color="indigo accent-2"
                    value="ReviewScoreDesc"
                    label="Review Score Desc"
                  ></v-radio> 
                  <v-radio color="indigo accent-2"
                    value="ReviewScoreAsc"
                    label="Review Score Asc"
                  ></v-radio>
                  <v-radio color="indigo accent-2"
                    value="ReviewsCountDesc"
                    label="Number of Reviews Desc"
                  ></v-radio>
                  <v-radio color="indigo accent-2"
                    value="ReviewsCountAsc"
                    label="Number of Reviews Asc"
                  ></v-radio>                                                                                                                                                                                                                                                                                             
                </v-radio-group>
              </v-card-text>                                 
              <v-card-text class="sortby" v-if="showFilterAgoda">
                  <v-radio-group v-model="agoda_sortby">
                  <v-radio color="blue"
                    value="Recommended"
                    label="Recommended"
                  ></v-radio>
                  <v-radio color="blue"
                    value="PriceDesc"
                    label="Price Desc"
                  ></v-radio>
                  <v-radio color="blue"
                    value="PriceAsc"
                    label="Price Asc"
                  ></v-radio>
                  <v-radio color="blue"
                    value="StarRatingDesc"
                    label="Star Rating Desc"
                  ></v-radio> 
                  <v-radio color="blue"
                    value="StarRatingAsc"
                    label="Star Rating Asc"
                  ></v-radio>
                  <v-radio color="blue"
                    value="AllGuestsReviewScore"
                    label="All Guests Review Score"
                  ></v-radio> 
                  <v-radio color="blue"
                    value="BusinessTravellerReviewScore"
                    label="Business Traveller Review Score"
                  ></v-radio> 
                  <v-radio color="blue"
                    value="CouplesReviewScore"
                    label="Couples Review Score"
                  ></v-radio> 
                  <v-radio color="blue"
                    value="SoloTravllersReviewScore"
                    label="Solo Travllers Review Score"
                  ></v-radio>
                  <!-- <v-radio color="blue"
                    value="FamiliesWithYoungReviewScore"
                    label="Families With Young Review Score"
                  ></v-radio> 
                  <v-radio color="blue"
                    value="FamiliesWithTeenReviewScore"
                    label="Families With Teen Review Score"
                  ></v-radio>
                  <v-radio color="blue"
                    value="GroupsReviewScore"
                    label="Groups Review Score"
                  ></v-radio>     -->                                                                                                                                                                                                                                             
                </v-radio-group>
              </v-card-text>                    
      </v-flex>
      
      <v-flex xs12 order-xs4>
        <p v-if="!showFilterAgoda" align="center" style="margin: auto;"><v-btn large color="blue" dark @click="filter">Filter now</v-btn></p>
        <p v-if="showFilterAgoda" align="center" style="margin: auto;"><v-btn large color="blue" dark @click="filter_agoda">Filter now</v-btn></p>
      </v-flex>

      <input v-model="cityId" type="hidden" value="cityId" />
  </v-layout>
</template>

<script>

  import axios from 'axios'

  export default {
    props:{ cityId: String },
    data: () => ({
      dailyRate_max: 1000,
      dailyRate_min: 0,
      discountOnly: false,
      minimumReviewScore: 0,
      minimumStarRating: '0',          
      sortby: 'Recommended',
      agoda_sortby: 'Recommended',
      listFilter: false,
      filterbyreorder: true,
      showFilterAgoda: false
    }),

        methods:{
          filter(){
              
              var dataFilter = {}

              // dataFilter['cityId'] = this.cityId.toString()
              dataFilter['sortBy'] = this.sortby 

              if(this.dailyRate_min > 0){
                  dataFilter['dailyRate_min'] = this.dailyRate_min.toString()
              }
              if(this.dailyRate_max < 1000){
                  dataFilter['dailyRate_max'] = this.dailyRate_max.toString()
              }
              if(this.minimumReviewScore > 0){ 
                  dataFilter['minReviewScore'] = this.minimumReviewScore.toString()
              }
              if(this.minimumStarRating > 0){
                  dataFilter['minStarRating'] = this.minimumStarRating.toString()
              }

              this.$emit('loading', true)
              axios.get(this.$store.state.apiPath +'/hotels/'+ this.$route.params.city +'/filter/'+ JSON.stringify(dataFilter)).then(response => {
                console.log(response.data)
                this.$emit('filterSubmit', response.data.Hotels)
              })              
          },

          filter_agoda(){

            alert("agoda API")

            var citykey = (this.$route.params.city).replace(/-[a-z][a-z]-hotels(.*)/, '')

            var dataFilter = {
                                 "criteria": {
                                               "additional": {
                                                               "maxResult": 30,
                                                             },
                                               "checkInDate": this.$route.query.sd,
                                               "checkOutDate": this.$route.query.ed,
                                               "cityId": parseInt(this.$route.query.cityId)
                                              }
                              }

            if(1 == 2){ dataFilter['criteria']['additional']['currency'] = 'USD' }

            if(this.dailyRate_max < 1000){
                if(dataFilter['criteria']['additional']['dailyRate'] == undefined){
                    dataFilter['criteria']['additional']['dailyRate'] = {}
                }
                dataFilter['criteria']['additional']['dailyRate']['maximum'] = this.dailyRate_max
            }
            if(this.dailyRate_min > 0){
                if(dataFilter['criteria']['additional']['dailyRate'] == undefined){
                    dataFilter['criteria']['additional']['dailyRate'] = {}
                }
                dataFilter['criteria']['additional']['dailyRate']['minimum'] = this.dailyRate_min
            }
            if(this.discountOnly == true){
                dataFilter['criteria']['additional']['discountOnly'] = true
            }
            
            if(1 == 2){ dataFilter['criteria']['additional']['language'] = 'en-us' }

            if(this.minimumReviewScore > 0){ 
                dataFilter['criteria']['additional']['minimumReviewScore'] = this.minimumReviewScore 
            }
            if(this.minimumStarRating > 0){
                dataFilter['criteria']['additional']['minimumStarRating'] = this.minimumStarRating
            }
            
            dataFilter['criteria']['additional']['sortBy'] = this.agoda_sortby                                                                         

            this.$emit('loading', true)

            axios.get(this.$store.state.apiPath +'/front?mod=agoda&data='+ JSON.stringify(dataFilter)).then(response => {

              var agoda = response.data.results           
              var agoda_hids = []

              for (var i = 0; i < agoda.length; i++) {
                  agoda_hids.push(agoda[i].hotelId)
              }
                  // alert(agoda_hids.join())
                  axios.get(this.$store.state.apiPath +'/front?mod=hotels&citykey='+ citykey +'&hids='+ agoda_hids.join()).then(response => {
                      var hotels = response.data.Hotels

                      for (var i = 0; i < hotels.length; i++) { 
                        for (var j = 0; j < agoda.length; j++) {                     
                          if(hotels[i]['Hotel_id'] == agoda[j]['hotelId']){
                              hotels[i]['reviewScore'] = agoda[j]['reviewScore']
                              hotels[i]['reviewCount'] = agoda[j]['reviewCount']
                              hotels[i]['currency'] = agoda[j]['currency']
                              hotels[i]['dailyRate'] = agoda[j]['dailyRate']
                              hotels[i]['discountPercentage'] = agoda[j]['discountPercentage']
                              hotels[i]['imageURL'] = agoda[j]['imageURL']
                              hotels[i]['includeBreakfast'] = agoda[j]['includeBreakfast']
                              hotels[i]['freeWifi'] = agoda[j]['freeWifi']
                          }
                        }
                      }

                      this.$emit('filterSubmit', hotels)
                  })
            })
          }
        }    
  }
</script>

<style type="text/css">
.mainFilter .subheader{
   position: relative;
}
.mainFilter .card__text{
   padding: 5px 16px 20px;
}
.mainFilter .input-group{
   padding: 0!important;
}
.mainFilter>.flex{
  padding: 0 8px!important;
}
.mainFilter .hint{
  float: right;
  top: 17px;
  right: 10px!important;
  position: absolute;
  font-weight: bold;
}
.mainFilter .radioRow{
    display: table;
    width: 100%;
    position: relative;
}
.mainFilter .radioRow .caption{
  float: left; margin-left: 10px; color: #666; margin-top: 8px;
}
.mainFilter .starsRadio{
    width: 30px;
    float: left;
}
.mainFilter .stars{
  margin-top: 4px;
  float: left;
}
.mainFilter .stars i{
  font-size: 16px;
}
.gridFilter{
    background: #fafafa;
    border-top: 1px solid #ddd;
    border-bottom: 1px solid #ddd;
    padding: 20px 10px 0;
    margin-bottom: 20px;
}

.listFilter .xs12,
.listFilter .sm6,
.listFilter .md4{
    max-width: 100%!important;
    flex-basis: 100%!important;
}
.sortby label{
    font-size: 14px;
}
.discountOnly .input-group__details{
  display: none;
}
</style>