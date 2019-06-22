<template>
  <v-layout wrap mainFilter>
              <v-flex xs12 sm6 md4 :order-xs2="listFilter">
                      <v-subheader>Popular hotel brands</v-subheader>          
                      <v-card-text>
                            <v-checkbox
                              label="Tauzia Hotel Management"
                              color="red"
                              value="red"
                              hide-details
                            ></v-checkbox>
                            <v-checkbox
                              label="Accor Hotels"
                              color="red darken-3"
                              value="red darken-3"
                              hide-details
                            ></v-checkbox>
                            <v-checkbox
                              label="Inna Hotel Group"
                              color="red darken-3"
                              value="red darken-3"
                              hide-details
                            ></v-checkbox>

                            <v-checkbox
                              label="Tauzia Hotel Management"
                              color="red"
                              value="red"
                              hide-details
                            ></v-checkbox>
                            <v-checkbox
                              label="Accor Hotels"
                              color="red darken-3"
                              value="red darken-3"
                              hide-details
                            ></v-checkbox>
                            <v-checkbox
                              label="Inna Hotel Group"
                              color="red darken-3"
                              value="red darken-3"
                              hide-details
                            ></v-checkbox>
                      </v-card-text>                    
              </v-flex>

              <v-flex xs12 sm6 md4 :order-xs3="listFilter">
                  <v-subheader>Property facilities</v-subheader>
                  <v-card-text>
                    <v-switch
                      label="Swimming pool"
                      color="red"
                      value="red"
                      hide-details
                    ></v-switch>
                    <v-switch
                      label="Internet"
                      color="red darken-3"
                      value="red darken-3"
                      hide-details
                    ></v-switch>
                    <v-switch
                      label="Beach access"
                      color="red darken-3"
                      value="red darken-3"
                      hide-details
                    ></v-switch>            
                  </v-card-text>


                      <v-subheader>Accommodation class</v-subheader>
                      <v-card-text>
                        <v-radio-group>
                          <v-radio color="orange"
                            key="budget"
                            label="Budget"
                            value="1"
                          ></v-radio>
                          <v-radio color="orange"
                            key="luxury"
                            label="Luxury"
                            value="2"
                          ></v-radio>              
                        </v-radio-group>
                      </v-card-text>                    
              </v-flex>

              <v-flex xs12 md4 :order-xs1="listFilter">
                  <v-subheader>Guest rating</v-subheader>
                  <v-card-text pl-2>
                    <v-slider v-model="guest_rating" step="10" thumb-color="blue" thumb-label></v-slider>
                  </v-card-text>

                  <v-subheader>Price range</v-subheader>
                  <v-card-text>
                    <v-slider color="green"></v-slider>
                  </v-card-text>
              </v-flex>

              {{ dataFilter }}

              <p align="center" style="margin: auto;"><v-btn color="blue" dark @click="filter">Filter now</v-btn></p>

  </v-layout>
</template>

<script>

  import axios from 'axios'

  export default {

    data: () => ({
      guest_rating: 0,
      dataFilter: [],
      listFilter: false
    }),

        methods:{
          filter(){
            this.dataFilter = {'sd': this.$route.query.sd, 'ed': this.$route.query.ed, 'ty': this.$route.query.ty, 'id': this.$route.query.id, breakfast: true, wifi: true, 'guest_rating': this.guest_rating }

            this.$emit('loading', true)
            // axios.get(this.$store.state.apiPath +'/front?mod=hotels&city_cc=bangkok-th-hotels.html').then(response => {
            //   this.$emit('filterSubmit', response.data.Hotels)
            // })
          }
        }    
  }
</script>

<style type="text/css">
.mainFilter .card__text{
   padding: 5px 16px 20px;
}
.mainFilter .input-group{
   padding: 0!important;
}
.mainFilter>.flex{
  padding: 0 8px!important;
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
</style>