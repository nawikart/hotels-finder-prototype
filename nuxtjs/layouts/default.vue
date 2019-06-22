<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      fixed
      temporary
    >
      <v-list>
        <v-list-tile
          router
          :to="item.to"
          :key="i"
          v-for="(item, i) in items"
          exact
        >
          <v-list-tile-action>
            <v-icon v-html="item.icon"></v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title v-text="item.title"></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar style="position: absolute; z-index: 9; opacity: .8;">
      <v-container style="padding: 0 15px; margin: auto;">

        <div style="float: left;margin-top: 16px; margin-left: 14px;">
          <!-- <v-toolbar-side-icon @click="drawer = !drawer" class="hidden-sm-and-down"></v-toolbar-side-icon> -->

          <!-- &nbsp; &nbsp;  -->
          <router-link to="/" class="headline">
            HOTELS DEMO
          </router-link>
        </div>

        <div style="float: right;">
          <div id="desktop_menutop" class="hidden-sm-and-down">
            <router-link to="#">
              <v-btn flat color="blue-grey darken-1" @click="undercontruction_snackbar = true">Become a Host</v-btn>  
            </router-link>      
            <router-link to="#">
              <v-btn flat color="indigo accent-2" @click="undercontruction_snackbar = true">Sign In</v-btn>  
            </router-link>
            <router-link to="#">
              <v-btn outline large color="indigo accent-2" @click="undercontruction_snackbar = true">Create Account</v-btn>   
            </router-link>                                         
          </div>

          <!-- <span class="hidden-md-and-up">
              <v-btn outline color="blue">Book</v-btn>       
          </span> -->
        </div>
        
      </v-container>
    </v-toolbar>

    <v-content>
        <nuxt />
    </v-content> 

    <v-container id="allCities" fluid style="margin-top: 70px; background-color: #333; color: hsla(0,0%,100%,.5); font-size: 16px; font-weight: 300; padding: 50px 0 10px;">
      <p align="center">This website is made for experiment and research purposes.<br />
Still in maintenance, so please understand if there is some error here :-)</p>
      <v-layout raw>
          <v-flex md-2 px-1 text-xs-center>x</v-flex>
          <v-flex md-2 px-1 text-xs-center>x</v-flex>
          <v-flex md-2 px-1 text-xs-center>x</v-flex>
          <v-flex md-2 px-1 text-xs-center>x</v-flex>
          <v-flex md-2 px-1 text-xs-center>x</v-flex>
          <v-flex md-2 px-1 text-xs-center>x</v-flex>          
      </v-layout>
    </v-container>

    <v-snackbar
      v-model="undercontruction_snackbar"
      multi-line
      block
      color="blue"
      dark
      right
      top
    >
      Sorry, still undercontruction :-)
      <v-icon
        color="red darken-5"
        @click="undercontruction_snackbar = false"
      >
        close
      </v-icon>
    </v-snackbar>    
  </v-app>
</template>

<script>
  export default {
    data() {
      return {
        undercontruction_snackbar: false,
        map: false,
        drawer: false,
        fixed: false,
        fab: false,
        items: [
          { icon: 'home', title: 'Homepage', to: '/' },
          { icon: 'hotel', title: 'Hotel & Accommodations', to: '/hotel' },
          { icon: 'hotel', title: 'Tour Packages', to: '/car' },
          { icon: 'bubble_chart', title: 'News & Updates', to: '/#' },
          { icon: 'bubble_chart', title: 'Autocomplete', to: '/autocomplete' },
          { icon: 'bubble_chart', title: 'Test', to: '/test' }
        ],
        title: 'IOTRAVEL'
      }
    }
  }
</script>

<style>
  @font-face{font-family:mallory;src:url(/assets/fonts/mallory-light-webfont-v01.eot);src:url(/assets/fonts/mallory-light-webfont-v01.eot?#iefix) format("embedded-opentype"),url(/assets/fonts/mallory-light-webfont-v01.woff2) format("woff2"),url(/assets/fonts/mallory-light-webfont-v01.woff) format("woff"),url(/assets/fonts/mallory-light-webfont-v01.ttf) format("truetype");font-weight:500;font-style:normal}
  div{
    font-family: 'mallory', "Helvetica Neue", Helvetica, Arial, sans-serif;
  }  
  BODY{
    background-color: #999;
  }
  a{
    text-decoration: none;
  }
  a.underline{
    text-decoration: underline;
    color: #2196f3;
  }
  #app{
    background-color: #fff;
  }
  .topbar{
    /*position: absolute;*/
    z-index: 1;
    opacity: .8;
  }
/*  #__nuxt{
    max-width: 1440px;
    margin: auto;
  }*/
  #desktop_menutop{
    margin-top: 4px;
    float: right;
  }
  #desktop_menutop .btn__content{
    font-weight: bold;
  }
  #create .speed-dial {
    position: absolute;
  }

  #create .btn--floating {
    position: relative;
  }

  .mainForm_style2 #mainForm{
    padding: 75px 3% 0;
    margin: auto;
    width: 100%;
    display: table;
  }
  .mainForm_style2 #mainForm .input_wrapper{
    background-color: white;
    border: 1px solid #666;
    border-right: none;
    position: relative;
    /* z-index: 2; */
  }
  .mainForm_style2 #mainForm button{
    z-index: 3;
  }
  .mainForm_style2 #mainForm .input.half{
    border-left: none!important;
  }
  .mainForm_style2 #mainForm #autocompleteResult{
    border: 1px solid #666!important;
    border-top: none!important;
  }
  @media only screen and (max-width: 1023px) {
    .mainForm_style2 #mainForm .input_wrapper{
      border: 1px solid #666!important;
    }
      .mainForm_style2 #mainForm .input_wrapper.half:nth-child(1){
        border-right: none!important;
      }  
  }
  @media only screen and (max-width: 599px) {
    .mainForm_style2 #mainForm .input_wrapper,
    .mainForm_style2 #mainForm .input_wrapper.half:nth-child(1){
      border: 1px solid #666!important;
    } 
  }  
</style>