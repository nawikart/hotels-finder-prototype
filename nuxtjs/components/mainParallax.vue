<template><div id="mainParallax">

    <div class="overlay_mainForm" v-if="overlay_mainForm" v-on:click="searchCloseClicked()"></div>
    <v-parallax :src="$store.state.mediaPath +'/media/images/main.jpg'">
    </v-parallax>

    <v-container fluid class="mainForm_container">
        <v-flex class="mainForm_wrapper">
            <div class="mainForm">
                <p class="headline font-weight-light" align="center">Find your ideal hotel at the lowest prices.</p>
                <!-- <span class="subheading">Get the best prices on 1,900,000+ properties, worldwide</span>       -->      
                <mainForm></mainForm>                       
            </div> 
        </v-flex>
    </v-container>

</div></template>

<script>
  export default {
    data: () => ({
      overlay_mainForm: false,
    })    
  }
</script>

<style type="text/css">
#mainParallax .mainForm_container{
    padding: 0; top: 220px; position: absolute;
}
#mainParallax .parallax{
  height: 400px!important;
  background-color: #18b0f7;  
}
#mainParallax .parallax__content{
  padding: 0;
}
#mainParallax .mainForm_wrapper{
  width: 90%; max-width: 1024px; margin: auto; display: table; position: relative; z-index: 1;  
}
#mainParallax .mainForm{
  margin: auto;
  padding: 20px!important;
  background-color: rgba(255,255,255,.7)!important;
  border: 1px solid white;
  color: #222;
}
@media only screen and (max-width: 1023px) {
    #mainParallax .mainForm_container{
        top: 190px;
    }
    #mainParallax .parallax{
        height: 450px!important;
    }    
}
@media only screen and (max-width: 599px) { 
  #mainParallax .mainForm_container{
      top: 140px;
  }
  #mainParallax .parallax{
    height: 540px!important;
  }
  #mainParallax .mainForm_wrapper{
    width: 100%!important;
    /*top: 120px!important;*/
  }
  #mainParallax .mainForm{
    padding: 10px!important;
  }  
}  
</style>