<template>
      <nav class="pagination">
          <ul> 
              <li><router-link :to="'/' + cc + '-hotels-p' + (P - 1) + '.html'" v-if="P != 1">PREV</router-link></li>
              <li><router-link v-bind:class="{ active: P == 1 }" :to="'/' + cc + '-hotels.html'">1</router-link></li>
              <li v-if="P > 4 && Ps > PsMax"><a>...</a></li>
              <li v-for="i in Pss" :key="i"><router-link v-bind:class="{ active: P == (Pstart + i - 1) }" :to="'/' + cc + '-hotels-p' + (Pstart + i - 1) + '.html'">{{ Pstart + i - 1 }}</router-link></li>
              <li v-if="P < Ps - 6 && Ps > PsMax"><a>...</a></li>
              <li><router-link v-bind:class="{ active: P == Ps }" :to="'/' + cc + '-hotels-p' + Ps + '.html'"><b>{{ Ps }}</b></router-link></li>
              <li><router-link :to="'/' + cc + '-hotels-p' + (P + 1) + '.html'" v-if="P != Ps">NEXT</router-link></li>
          </ul>
      </nav>
</template>

<script>
  export default {
    props:{ cc: String, P: Number, Ps: Number },

    data: () => ({  	
      Pstart: 0,
      PsMax: 10,
      Pss: 0,
    }),

        methods: {

        },
         beforeMount(){
              // this.P = 3
              // this.Ps = 25

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
         }     
  }
</script>

<style type="text/css">
nav.pagination {
  width: 100%;
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
float: right;
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
</style>