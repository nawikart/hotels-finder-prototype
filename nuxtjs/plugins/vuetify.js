import Vue from 'vue'
import Vuetify from 'vuetify'
import colors from 'vuetify/es5/util/colors'

Vue.use(Vuetify, {
  theme: {
    primary: '#121212', // a color that is not in the material colors palette
    accent: colors.grey.darken3,
    secondary: colors.amber.darken3,
    info: colors.teal.lighten1,
    warning: colors.amber.base,
    error: colors.deepOrange.accent4,
    success: colors.green.accent3
  }
})


import mainParallax from '@/components/mainParallax'
Vue.component('mainParallax', mainParallax)

import mainForm from '@/components/mainForm'
Vue.component('mainForm', mainForm)

import breadcrumb from '@/components/breadcrumb'
Vue.component('breadcrumb', breadcrumb)

import gmap from '@/components/gmap'
Vue.component('gmap', gmap)


// LISTING
import listing from '@/components/listing'
Vue.component('listing', listing)

import hotelinListing from '@/components/listing/hotel'
Vue.component('hotelinListing', hotelinListing)

import listingList from '@/components/listing/list'
Vue.component('listingList', listingList)

import listingSmGrid from '@/components/listing/sm-grid'
Vue.component('listingSmGrid', listingSmGrid)	

import listingMdGrid from '@/components/listing/md-grid'
Vue.component('listingMdGrid', listingMdGrid)	

import listingPagination from '@/components/listing/pagination'
Vue.component('listingPagination', listingPagination)	

import stars from '@/components/small/stars'
Vue.component('stars', stars) 


import frontHotels from '@/components/frontHotels'
Vue.component('frontHotels', frontHotels)

import contentCarousel from '@/components/contentCarousel'
Vue.component('contentCarousel', contentCarousel)

import mainCarousel from '@/components/mainCarousel'
Vue.component('mainCarousel', mainCarousel)

import mainCities from '@/components/mainCities'
Vue.component('mainCities', mainCities)

import whyUs from '@/components/whyUs'
Vue.component('whyUs', whyUs)


import filterby from '@/components/filterby'
Vue.component('filterby', filterby)