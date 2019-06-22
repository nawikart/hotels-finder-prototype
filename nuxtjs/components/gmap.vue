<template>
    <v-dialog v-model="map" fullscreen hide-overlay transition="dialog-bottom-transition" scrollable>
      <v-card tile>
        <v-toolbar card dark color="primary">
          <v-toolbar-title>Hotel Villa Resorts</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark flat v-on:click="map = false">Close</v-btn>
          </v-toolbar-items>
          <v-btn icon dark v-on:click="map = false">
            <v-icon>close</v-icon>
          </v-btn>
        </v-toolbar>

          <div id="map" style="width: 100%; height: calc(100vh - 70px);">
              
              <!-- {{ mapData }} -->

              <iframe src="http://localhost:8080/gmap" 
                      allowTransparency="true" 
                      style="background-color:lightgreen;" 
                      width="100%" height="100%">
              </iframe>
          </div>
      </v-card>
    </v-dialog> 	
</template>

<script>
  export default {
    props:{ map: Boolean, mapData: Array }, 
  }
</script>