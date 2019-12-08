<template><div id="mainForm">

    <v-dialog v-model="dateStartPopup" max-width="290"><v-date-picker :min="dateStartMin" :max="dateStartMax" v-model="dateStartPicker"  color="green lighten-1" @input="dateFormat(dateStartPicker, 's')"></v-date-picker></v-dialog>
    <v-dialog v-model="dateEndPopup" max-width="290"><v-date-picker :min="dateEndMin" :max="dateEndMax" v-model="dateEndPicker" color="green lighten-1" @input="dateFormat(dateEndPicker, 'e')"></v-date-picker></v-dialog>

    <div class="overlay_mainForm" v-if="overlay_mainForm" v-on:click="searchCloseClicked()"></div>
              
              <div class="mainSearch" style="position: relative; display: table; width: 100%; max-width: 1136px; margin: auto; z-index: 2;">                
                <div class="row">

                  <div class="input_wrapper full">
                    <label for="searchQuery">Enter a destination or property</label>
                    <v-icon color="grey lighten-1" v-on:click="searchQueryFocus()">location_on</v-icon>
                    <debounce :delay="delay" v-on:input="autocomplete" v-model="searchQuery" id="searchQuery" ref="searchQuery" class="input searchQuery" placeholder="eg. Bali" />
                        <span class="searchInfo">
                          <span v-if="searchMin3char">min. 3 character</span>
                          <v-btn outline fab small color="red" v-if="searchClose" v-on:click="searchCloseClicked()"><v-icon>close</v-icon></v-btn>
                        </span>

                        <div id="autocompleteResult" v-if="showAutocompleteResult">
                          <div v-if="autocLoader" class="progress">
                            <v-progress-linear :indeterminate="true" value="15" height="2" color="pink lighten-1"></v-progress-linear>
                          </div>           
                          <v-flex v-on:click="searchQueryItemClicked(item.Type, item.Id, item.Value, item.CityUrl)" item v-for="(item, i) in items" :key="i" :val="item.Id" :class="item.Type"><v-icon small v-if="item.Type == 'c'">location_city</v-icon>
                            <v-icon color="orange" v-if="item.Type == 'h'">star_border</v-icon><span class="v"> &nbsp; {{item.Value}}</span>

                            <span class="i" v-if="item.Type == 'c'">{{item.HotelsCount}}</span>
                            <span class="i" v-if="item.Type == 'h'"><v-icon small>grade</v-icon> {{item.Rating_average}}</span>
                          </v-flex>
                        </div>
                  </div>
                </div>

                <div class="row">
                  <div class="half input_wrapper ">
                    <label for="startDate">Start date</label>
                    <v-icon color="grey lighten-1" v-on:click="showDatePicker('s')">event</v-icon>
                    <input id="startDate" readonly="readonly" class="input" type="text" :value="dateStartLabel" v-on:click="showDatePicker('s')" />  
                  </div>
                  <div class="half input_wrapper "> 
                    <label for="endDate">End date</label>
                    <v-icon color="grey lighten-1" v-on:click="showDatePicker('e')">event</v-icon>
                    <input id="endDate" readonly="readonly" class="input" type="text" :value="dateEndLabel" v-on:click="showDatePicker('e')" />
                    <a :href="searchTo"><v-btn small class="btnSearch hidden-xs-only" style="position: absolute; min-width: 50px; right: -8px; top: -6px; height: 100%; box-shadow: none; border: 2px solid #27a7f7!important; /*border-radius: 0 6px 6px 0;*/ font-weight: bold;" dark color="blue"><v-icon large>search</v-icon></v-btn></a>
                  </div>                
                </div>

                <p align="center" class="row hidden-sm-and-up">
                    <a :href="searchTo"><v-btn class="btnSearch" large dark color="blue" style="width: 90%;">Search</v-btn></a>
                </p>
              </div>

</div></template>

<script>
  import axios from 'axios'
  import Debounce from "./Debounce.vue";
  export default {
    data: () => ({
      delay: '10',
      ty: '',
      id: '',
      cityId: '',
      searchTo: '/',
      searchTo_url: '',

      dateStartPopup: '',
      dateEndPopup: '',

      dateStartPicker: '',
      dateEndPicker: '',

      dateStartLabel: '',
      dateEndLabel: '',

      dateStartMin: '2018-05-16',
      dateStartMax: '2018-05-16',

      dateEndMin: '2018-05-02',
      dateEndMax: '2018-05-16',

      overlay_mainForm: false,
      searchMin3char: false,
      searchClose: false,
      searchQuery: "",
      autocLoader: false,
      showAutocompleteResult: false,

      selectedType: '',
      selectedId: '',

      items: []
    }),

      components: { Debounce },

        methods: {

          searchTo_ok(){
              this.searchTo = this.searchTo_url + '?sd=' + this.dateStartPicker + '&ed=' + this.dateEndPicker + '&ty=' + this.ty + '&id=' + this.id + '&cityId=' + this.cityId + '&q=' + this.searchQuery
          },

          autocomplete () {

            // console.log(this.searchQuery)

            this.searchMin3char = false
            this.searchClose = false
            if(this.searchQuery.length > 1){

              if(this.delay != '300'){
                this.delay = '300'
              }

              if(this.selectedId == ''){ this.overlay_mainForm = true }
              this.searchClose = true
              this.autocLoader = true           
              axios.get(this.$store.state.apiPath +'/autocomplete/' + this.searchQuery)
              .then(response => {
                console.log(response.data)
                this.items = response.data

                if(this.items.length > 0){
                    this.showAutocompleteResult = true
                    this.searchActive = 'active'
                }else{
                    this.showAutocompleteResult = false
                    this.searchActive = ''
                }

                this.autocLoader = false
              });
            }else{
              if(this.searchQuery.length < 1){
                this.overlay_mainForm = false
              }
              this.showAutocompleteResult = false
              if(this.searchQuery.length > 0){
                this.searchMin3char = true
              }
            }         
          },

          dateFormat(d, k){
            var months = [,'Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec']

            this.dateStartPopup = false
            this.dateEndPopup = false

            var ds = d.split('-')
            var dl = ds[2] + ' ' + months[parseInt(ds[1])] + ' ' + ds[0]

            if(k == 's'){
              //this.dateEndMin = d
              this.dateStartLabel = dl
            }else{
              this.dateEndLabel = dl
            }

            this.searchTo_ok()
          },

          showDatePicker(k){
              this.dateStartPopup = false
              this.dateEndPopup = false
              if(k == 's'){
                  this.dateStartPopup = true
              }else if(k == 'e'){
                  this.dateEndPopup = true
              }
          },

          searchQueryItemClicked(ty, id, v){              
              this.selectedType = ty
              this.selectedId = id
              this.searchQuery = v
              this.showAutocompleteResult = false
              this.overlay_mainForm = false
              this.ty = ty
              this.id = id
              axios.get(this.$store.state.apiPath +'/front?mod=cityUrl&ty='+ ty +'&id='+ id).then(response => {
                this.searchTo_url = '/' + response.data.Citykey + '-' + (response.data.Countryisocode).toLowerCase() + '-hotels.html'
                this.cityId = response.data.CityId
                this.searchTo_ok()
              })
          },

          searchCloseClicked(){
            this.showAutocompleteResult = false
            this.overlay_mainForm = false
            this.searchClose = false
            this.searchQuery = ''

            this.selectedType = ''
            this.selectedId = ''
          },

          searchQueryFocus(){
              document.getElementById("searchQuery").focus();
          }
        },

        beforeMount(){

            var d = new Date()

            this.ty = this.$route.query.ty
            this.id = this.$route.query.id
            this.searchQuery = this.$route.query.q
            this.dateStartPicker = this.$route.query.sd
            this.dateEndPicker = this.$route.query.ed

            var mo = (d.getMonth() + 1)
            if(mo < 10){
                mo = '0' + mo
            }
            var dt = d.getDate()
            if(dt < 10){
                dt = '0' + dt
            }

            if(this.dateStartPicker == undefined){
                this.dateStartPicker = d.getFullYear() + '-' + mo + '-' + dt               
            }
            this.dateFormat(this.dateStartPicker, 's')

            d.setDate(d.getDate() + 1)
            var mo = (d.getMonth() + 1)
            if(mo < 10){
                mo = '0' + mo
            }
            var dt = d.getDate()
            if(dt < 10){
                dt = '0' + dt
            }
            if(this.dateEndPicker == undefined){
                this.dateEndPicker = d.getFullYear() + '-' + mo + '-' + dt
            }
            this.dateFormat(this.dateEndPicker, 'e')


            axios.get(this.$store.state.apiPath +'/front?mod=cityUrl&ty='+ this.ty +'&id='+ this.id).then(response => {
              this.searchTo_url = response.data[0]
              this.searchTo_ok()
            })
        }
  }
</script>

<style type="text/css">
::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
    color: #ccc;
    opacity: 1; /* Firefox */
}

:-ms-input-placeholder { /* Internet Explorer 10-11 */
    color: #ccc;
    opacity: 1; 
}

::-ms-input-placeholder { /* Microsoft Edge */
    color: #ccc;
    opacity: 1; 
}

::-webkit-input-placeholder { /* Chrome/Opera/Safari */
    color: #ccc;
    opacity: 1; 
}
::-moz-placeholder { /* Firefox 19+ */
    color: #ccc;
    opacity: 1; 
}
:-ms-input-placeholder { /* IE 10+ */
    color: #ccc;
    opacity: 1; 
}
:-moz-placeholder { /* Firefox 18- */
    color: #ccc;
    opacity: 1; 
}
.container.mainForm_style2{
  padding-bottom: 5px!important;
}
#mainForm .overlay_mainForm{
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background: rgba(0,0,0,.4);
  z-index: 1;
}
#mainForm .input_wrapper{
  position: relative;
  background-color: white;
}
#mainForm .input_wrapper label{
  position: absolute;
  top: 5px;
  width: 100%;
  font-size: 92%;
  left: 30px;
  text-shadow: 1px 1px 1px #eee;
  color: #999;
}
#mainForm .input_wrapper>.icon{
  position: absolute;
  top: 36px;
  left: 6px;
  cursor: pointer;
}
#mainForm .v-icon.v-icon.v-icon--link{
  margin: 2px;
}
#mainForm .mainSearch .full{
  width: 100%;
}
#mainForm .mainSearch .half{
  width: 50%;
  float: left;
  border-left: 1px solid #999;
}
#mainForm .mainSearch .row{
   width: 50%;
   float: left;
   position: relative;
   margin-top: 5px;
}
#mainForm input:focus{
  outline: 0;
}
#mainForm .mainSearch .input{
  padding: 10px 6px 10px 35px;
  background: #fff; 
  font-size: 18px;
  color: #333;
  font-weight: bold;
  text-shadow: 1px 1px 1px #fff;
  width: 100%;
}
#mainForm .mainSearch input.input.searchQuery{
    padding-right: 50px;
}
#mainForm .mainSearch input.active{
  border: 1px solid #18b0f7!important;
  background-color: white!important;
}
#mainForm .mainSearch .searchInfo{
    position: absolute;
    top: 2px;
    z-index: 9;
    right: 0px;
    font-size: 80%;
    color: red;
    width: 90px;
}
#mainForm .mainSearch .searchInfo span{
  margin-top: 20px;
  display: table;
}
#mainForm .mainSearch .searchInfo button{
  right: -5px;
  float: right;
  border: none;
}
#mainForm #autocompleteResult {
    position: absolute;
    background-color: white;
    margin: auto;
    text-align: left;
    margin-top: -5px;
    left: 0px;
    border-top: none;
    padding: 5px 1px;
    z-index: 99;
    font-size: 14px;
    width: calc(100% - 0px);    
}
#mainForm #autocompleteResult .progress{
  position: relative;
}
#mainForm #autocompleteResult .progress-linear{
  margin-top: -6px;
  position: absolute;
}
#mainForm #autocompleteResult .item{
    display: block;
    padding: 5px 15px;
    position: relative;
    text-decoration: none;
    cursor: pointer;
}
#mainForm #autocompleteResult i.icon {
    font-size: 18px;
    margin-top: -6px;
}
#mainForm #autocompleteResult .item:hover{
  background: #EE595D;
}
#mainForm #autocompleteResult .item:hover span{
  color: #fff;
}
#mainForm #autocompleteResult .item.c span.i:after{
  content: "properties";
  color: #18b0f7;
  margin-left: 5px;
}
#mainForm #autocompleteResult .item:hover span.i:after{
  color: white;
}
#mainForm #autocompleteResult .item span.i {
    font-size: 80%;
    font-style: italic;
    position: absolute;
    right: 15px;
    top: 6px;
    color: #EE595D;
    background-color: white;
    padding-left: 10px;
}
#mainForm #autocompleteResult .item.h span.i{
  font-weight: bold;
  width: 60px;
}
#mainForm #autocompleteResult .item:hover span.i{
  color: white;
  background-color: #EE595D;
}

@media only screen and (max-width: 1023px) {
    #mainForm .mainSearch .row{
      width: 100%;
    }
    #mainForm .mainSearch .input_wrapper{
      border: 1px solid #999;
    }
    #mainForm .mainSearch .input_wrapper.half:nth-child(1){
      border-right: none;
    }
    #mainForm #autocompleteResult{
      /*border: 1px solid #999;*/
      border-top: none;
    }      
}
@media only screen and (max-width: 599px) {
  #mainForm .mainSearch .input_wrapper{
    border: 1px solid #999!important;
  }
  #mainForm .mainSearch .searchInfo{
    right: 1px;
  }
  #mainForm .mainSearch .searchInfo span{
    margin-top: 0;
  }
  #mainForm .mainSearch .searchInfo button{
    right: -14px;
  }
  #mainForm .mainSearch .input_wrapper.half{
    width: 49.5%;
  } 
  #mainForm .mainSearch .input_wrapper.half:nth-child(2){
    float: right;
  }
}  
</style>




<!-- #autocompleteResult .item.c span.v:before{
  content: "C";  
  margin-right: 9px;
}
#autocompleteResult .item.c span.i:after{
  content: "properties";
  color: #18b0f7;
  margin-left: 5px;
}
#autocompleteResult .item:hover span.i:after{
  color: white;
}
#autocompleteResult .item.h span.v:before{
  content: "H";
  margin-right: 9px;
}
#autocompleteResult .item.h span.i:before{
  content: "G";
  margin-right: 5px;
} -->