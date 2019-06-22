<template>
    <div id="hotel">

        <!-- <mainParallax></mainParallax> -->

        <v-container fluid grid-list-xs class="mainForm_style2 blue lighten-5">
            <mainForm></mainForm>
            <breadcrumb :data="breadcrumb"></breadcrumb>
        </v-container>
        
        <!-- <div style="background-color: #f6f6f6;"><breadcrumb :data="breadcrumb"></breadcrumb></div> -->

        <v-container grid-list-sm v-if="hotelLoader">
            <p>&nbsp;</p>
            <p>&nbsp;</p>
            <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="margin: auto; width: 60%;"></v-progress-linear>
            <br />
            <br />
            <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="margin: auto; width: 30%;"></v-progress-linear>
            <br />
            <br />
            <v-progress-linear :indeterminate="true" value="15" height="1" color="grey" style="margin: auto; width: 80%;"></v-progress-linear>
            <br />
            <br />

            <v-layout wrap gallery>
                <v-flex xs12 sm6>
                    <div class="flex" style="height: 400px;background-color: #fafafa;">
                        <v-progress-linear :indeterminate="true" value="15" height="1" color="black lighten-1" style="width: 80%; margin: 280px auto 0;"></v-progress-linear>
                    </div>
                </v-flex>
                <v-flex xs12 sm6>

                    <v-layout wrap>
                        <v-flex item xs6>
                            <div class="flex" style="background-color: #fafafa; height: 200px;">
                                <v-progress-linear :indeterminate="true" value="15" height="1" color="black lighten-1" style="width: 80%; margin: 180px auto 0;"></v-progress-linear>
                            </div>
                        </v-flex>
                        <v-flex item xs6>
                            <div class="flex" style="background-color: #fafafa; height: 200px;">
                                <v-progress-linear :indeterminate="true" value="15" height="1" color="black lighten-1" style="width: 80%; margin: 180px auto 0;"></v-progress-linear>
                            </div>
                        </v-flex>
                        <v-flex item xs6>
                            <div class="flex" style="background-color: #fafafa; height: 200px;">
                                <v-progress-linear :indeterminate="true" value="15" height="1" color="black lighten-1" style="width: 80%; margin: 180px auto 0;"></v-progress-linear>
                            </div>
                        </v-flex>
                        <v-flex item xs6>
                            <div class="flex" style="background-color: #fafafa; height: 200px;">
                                <v-progress-linear :indeterminate="true" value="15" height="1" color="black lighten-1" style="width: 80%; margin: 180px auto 0;"></v-progress-linear>
                            </div>
                        </v-flex>
                    </v-layout>
                </v-flex>

            </v-layout>
        </v-container>

        <v-container fluid grid-list-xs v-if="!hotelLoader">
            <v-container grid-list-sm>
                <v-flex xs11 md8 style="margin: 50px auto 0;">
                    <div class="display-1 text-xs-center orange--text">{{ detail.Hotel_name }}</div>

                    <div class="mt-2" style="display: table; margin: auto;">
                        <stars :Star_rating="detail.Star_rating"></stars> &nbsp; &nbsp; (<b>{{ detail.Number_of_reviews }}</b> reviews)
                    </div>
                    <p class="body-1 address mt-2" align="center">
                        <v-icon small>location_on</v-icon><span class="pt-2" style="margin-top: 9px; color: #777;">{{ detail.Addressline1 }}, {{ detail.City }}, {{ detail.Countryisocode }}</span> &nbsp; <a class="viewonmap blue--text" style="text-decoration: underline; font-size: 12px!important; white-space: nowrap;" v-on:click="openMap">view on map</a></p>
                    <p class="subheading overview text-xs-center" v-if="overviewFull && detail.Overview != ''">{{ detail.Overview }} <a v-on:click="overviewFull = false">less</a></p>
                    <p class="subheading overview text-xs-center" v-if="!overviewFull && detail.Overview != ''">{{ detail.Excerpt }}... <a v-on:click="overviewFull = true">more</a></p>

                    <p align="center" class="green--text my-4">
                        <v-icon color="green">wifi</v-icon> &nbsp; wifi &nbsp; &nbsp;
                        <v-icon color="blue">check</v-icon> &nbsp;
                        <nowrap>order now pay later</nowrap> &nbsp; &nbsp;
                        <nowrap><v-icon color="pink" style="font-size: 16px; margin-top: -6px;">fastfood</v-icon> &nbsp;
                        include breakfast</nowrap>
                    </p>
                </v-flex>
            </v-container>

            <v-container fluid grid-list-sm style="max-width: 1360px;">
                <v-layout wrap gallery>
                    <v-flex xs12 :sm6="photosCount > 1" style="min-height: 250px;">
                        <div class="mainphoto" :style="'height: 100%; background-image: url('+ detail.Photo1 +'); background-size: cover!important; background-color: #ddd; position: relative'">
                            <div style="position: absolute; top: 0; right: 0; background-color: rgba(0,0,0,.5); color: white; text-align: center; padding: 5px 20px; margin-top: 20px;"><b class="display-1">{{ detail.Rating_average }}</b>
                                <br />
                            
                                <span v-if="detail.Rating_average >= 9">Exceptional</span>
                                <span v-if="detail.Rating_average >= 8 && detail.Rating_average < 9">Exellent</span>
                                <span v-if="detail.Rating_average >= 7 && detail.Rating_average < 8">Very good</span>
                                <span v-if="detail.Rating_average >= 6 && detail.Rating_average < 7">Good</span>

                            </div>
                            <div style="position: absolute; top: 100px; right: 0; font-weight: bold; background-color: rgba(238,89,93, .9); color: white; text-align: center; padding: 5px 20px; margin-top: 20px;"><span>usd </span><b class="headline">{{ detail.Rates_from }}</b></div>
                        </div>
                    </v-flex>
                    <v-flex xs12 sm6>
                        <v-layout wrap>
                            <v-flex item xs6>
                                <div :style="'height: 200px; background-image: url('+ detail.Photo2 +'); background-size: cover!important; background-color: #ddd; position: relative'"></div>
                            </v-flex>
                            <v-flex item xs6>
                                <div :style="'height: 200px; background-image: url('+ detail.Photo3 +'); background-size: cover!important; background-color: #ddd; position: relative'"></div>
                            </v-flex>
                            <v-flex item xs6>
                                <div :style="'height: 200px; background-image: url('+ detail.Photo4 +'); background-size: cover!important; background-color: #ddd; position: relative'"></div>
                            </v-flex>
                            <v-flex item xs6>
                                <div :style="'height: 200px; background-image: url('+ detail.Photo1 +'); background-size: cover!important; background-color: #ddd; position: relative'"></div>
                            </v-flex>
                        </v-layout>
                    </v-flex>

                </v-layout>
            </v-container>

            <v-container grid-list-sm mt-5>
                <v-layout row wrap>
                    <v-flex xs12 sm6 md7 pl-3>
                        <v-layout row wrap>

                            <v-flex xs12 headline style="border-bottom: 1px dashed #999; padding-bottom: 10px;">Nearest essentials</v-flex>

                            <v-layout row wrap ml-1 mt-4>
                                <v-flex xs12 sm6 body-1><b>Airports</b>
                                    <br />Ngurah Rai International Airport (DPS) ( 30.14 km )
                                    <br />Lombok International Airport (LOP) ( 114.59 km )</v-flex>
                                <v-flex xs12 sm6 body-1><b>Drug stores</b>
                                    <br />Puri Mandela Pharmacy ( 8.24 km )</v-flex>
                                <v-flex xs12 sm6 body-1><b>Cash withdrawal</b>
                                    <br />ATM ( 1.06 km )</v-flex>
                                <v-flex xs12 sm6 body-1><b>Bus and rail stations</b>
                                    <br />Bangli Bus Terminal ( 11.85 km )
                                    <br />Ubung Bus Terminal ( 16.74 km )</v-flex>
                                <v-flex xs12 sm6 body-1><b>Hospitals and clinics</b>
                                    <br />Praktek Dokter Dipayana ( 5.29 km )</v-flex>
                                <v-flex xs12 sm6 body-1><b>Convenience stores</b>
                                    <br />Indomaret ( 7.11 km )</v-flex>
                            </v-layout>

                            <p>&nbsp;</p>
                            <p>&nbsp;</p>

                            <v-flex xs12 mb-3 headline style="border-bottom: 1px dashed #999; padding-bottom: 10px;">Hotel Reviews</v-flex>
                            <v-layout my-3 row wrap v-for="(item, i) in reviews" :key="i">
                                <v-flex xs3 text-xs-right pr-2 mt-3>
                                    <p class="title"><b>8.4</b></p>
                                    <p class="body-1 grey--text">
                                        mellissa from Singapore
                                        <br /> Couple
                                        <br /> Suite Pool
                                    </p>
                                </v-flex>
                                <v-flex xs9>
                                    <div class="body-1 grey--text">Stayed 3 nights in December 2016</div>
                                    <div class="title my-2"><em>“Wonderful place with excellent service”</em></div>
                                    <div class="subheading">Enjoyed our stay. We had a pool villa which was lovely. A few problems with things not working were resolved quickly and efficiently. Staff were very helpful.</div>
                                    <div class="body-1 grey--text">Reviewed December 20, 2016</div>
                                </v-flex>
                            </v-layout>
                        </v-layout>
                    </v-flex>

                    <v-flex xs12 sm5 md4 offset-sm1 banner>
                        <div class="flex" style="min-height: 300px;">
                            <p align="center" class="blue--text"><a class="viewonmap blue--text" style="text-decoration: underline; font-size: 12px!important;" v-on:click="openMap">click view full map</a></p>
                            <iframe src="http://localhost:8080/gmap" 
                                    allowTransparency="true" frameborder="0"
                                    style="background-color:lightgreen;" 
                                    width="100%" height="300">
                            </iframe>                        
                        </div>
                        <div class="flex mt-3" style="background: url('/gallery/img-3.jpg') no-repeat; background-size: 100% 90vh;min-height: 300px;"></div>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-container>

        <gmap :map="map" :mapData="mapData"></gmap>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        data: () => ({
            breadcrumb: [],
            map: false,
            mapData: [],
            overviewFull: false,
            photosCount: 4,
            detail: [],
            hotelLoader: true,
            reviews: [{}, {}, {}, {}]
        }),

        methods: {
          openMap (){
              this.map = true
              this.mapData = {'1': 'satu', '2': 'dua'}
          },            
        },

        beforeMount() {
            alert(2)
            this.hotelLoader = true
            axios.get(this.$store.state.apiPath +'/hotel/'+ this.$route.params.hotelnamekey +'/'+ this.$route.params.city).then(response => {
                this.detail = response.data
                console.log(response.data)
                this.detail.Excerpt = this.detail.Overview.substr(0, 200);

                if(this.detail.Photo2 == ""){ this.detail.Photo2 = this.detail.Photo1 }
                if(this.detail.Photo3 == ""){this.detail.Photo3 = this.detail.Photo1 }
                if(this.detail.Photo4 == ""){this.detail.Photo4 = this.detail.Photo1 }

                this.hotelLoader = false

                this.breadcrumb = {'city': response.data.City, 'citykey': response.data.Citykey, 'country': response.data.Country, 'countryisocode_lower': response.data.Countryisocode_lower, 'hotel_name': response.data.Hotel_name, 'hotel_namekey': response.data.Hotel_namekey}                
            })
        }
    }
</script>

<style>
#hotel .stars{
    float: left;
    margin-top: -2px;
}
#hotel .stars i{
  font-size: 20px;
}
#hotel .overview a{
    font-size: 90%;
    text-decoration: underline;
    color: rgb(39, 167, 247);
}
@media only screen and (min-width: 600px) {
    #hotel .banner{
        padding-right: 10px!important;
    }    
}
</style>