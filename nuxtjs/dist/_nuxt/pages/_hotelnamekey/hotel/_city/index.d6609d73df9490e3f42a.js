webpackJsonp([1],{"0A0B":function(t,e,a){"use strict";var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"hotel"}},[a("v-container",{staticClass:"mainForm_style2 blue lighten-5",attrs:{fluid:"","grid-list-xs":""}},[a("mainForm"),a("breadcrumb",{attrs:{data:t.breadcrumb}})],1),t.hotelLoader?a("v-container",{attrs:{"grid-list-sm":""}},[a("p",[t._v(" ")]),a("p",[t._v(" ")]),a("v-progress-linear",{staticStyle:{margin:"auto",width:"60%"},attrs:{indeterminate:!0,value:"15",height:"1",color:"grey"}}),a("br"),a("br"),a("v-progress-linear",{staticStyle:{margin:"auto",width:"30%"},attrs:{indeterminate:!0,value:"15",height:"1",color:"grey"}}),a("br"),a("br"),a("v-progress-linear",{staticStyle:{margin:"auto",width:"80%"},attrs:{indeterminate:!0,value:"15",height:"1",color:"grey"}}),a("br"),a("br"),a("v-layout",{attrs:{wrap:"",gallery:""}},[a("v-flex",{attrs:{xs12:"",sm6:""}},[a("div",{staticClass:"flex",staticStyle:{height:"400px","background-color":"#fafafa"}},[a("v-progress-linear",{staticStyle:{width:"80%",margin:"280px auto 0"},attrs:{indeterminate:!0,value:"15",height:"1",color:"black lighten-1"}})],1)]),a("v-flex",{attrs:{xs12:"",sm6:""}},[a("v-layout",{attrs:{wrap:""}},[a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{staticClass:"flex",staticStyle:{"background-color":"#fafafa",height:"200px"}},[a("v-progress-linear",{staticStyle:{width:"80%",margin:"180px auto 0"},attrs:{indeterminate:!0,value:"15",height:"1",color:"black lighten-1"}})],1)]),a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{staticClass:"flex",staticStyle:{"background-color":"#fafafa",height:"200px"}},[a("v-progress-linear",{staticStyle:{width:"80%",margin:"180px auto 0"},attrs:{indeterminate:!0,value:"15",height:"1",color:"black lighten-1"}})],1)]),a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{staticClass:"flex",staticStyle:{"background-color":"#fafafa",height:"200px"}},[a("v-progress-linear",{staticStyle:{width:"80%",margin:"180px auto 0"},attrs:{indeterminate:!0,value:"15",height:"1",color:"black lighten-1"}})],1)]),a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{staticClass:"flex",staticStyle:{"background-color":"#fafafa",height:"200px"}},[a("v-progress-linear",{staticStyle:{width:"80%",margin:"180px auto 0"},attrs:{indeterminate:!0,value:"15",height:"1",color:"black lighten-1"}})],1)])],1)],1)],1)],1):t._e(),t.hotelLoader?t._e():a("v-container",{attrs:{fluid:"","grid-list-xs":""}},[a("v-container",{attrs:{"grid-list-sm":""}},[a("v-flex",{staticStyle:{margin:"50px auto 0"},attrs:{xs11:"",md8:""}},[a("div",{staticClass:"display-1 text-xs-center orange--text"},[t._v(t._s(t.detail.Hotel_name))]),a("div",{staticClass:"mt-2",staticStyle:{display:"table",margin:"auto"}},[a("stars",{attrs:{Star_rating:t.detail.Star_rating}}),t._v("     ("),a("b",[t._v(t._s(t.detail.Number_of_reviews))]),t._v(" reviews)\n                ")],1),a("p",{staticClass:"body-1 address mt-2",attrs:{align:"center"}},[a("v-icon",{attrs:{small:""}},[t._v("location_on")]),a("span",{staticClass:"pt-2",staticStyle:{"margin-top":"9px",color:"#777"}},[t._v(t._s(t.detail.Addressline1)+", "+t._s(t.detail.City)+", "+t._s(t.detail.Countryisocode))]),t._v("   "),a("a",{staticClass:"viewonmap blue--text",staticStyle:{"text-decoration":"underline","font-size":"12px!important","white-space":"nowrap"},on:{click:t.openMap}},[t._v("view on map")])],1),t.overviewFull&&""!=t.detail.Overview?a("p",{staticClass:"subheading overview text-xs-center"},[t._v(t._s(t.detail.Overview)+" "),a("a",{on:{click:function(e){t.overviewFull=!1}}},[t._v("less")])]):t._e(),t.overviewFull||""==t.detail.Overview?t._e():a("p",{staticClass:"subheading overview text-xs-center"},[t._v(t._s(t.detail.Excerpt)+"... "),a("a",{on:{click:function(e){t.overviewFull=!0}}},[t._v("more")])]),a("p",{staticClass:"green--text my-4",attrs:{align:"center"}},[a("v-icon",{attrs:{color:"green"}},[t._v("wifi")]),t._v("   wifi    \n                    "),a("v-icon",{attrs:{color:"blue"}},[t._v("check")]),t._v("  \n                    "),a("nowrap",[t._v("order now pay later")]),t._v("    \n                    "),a("nowrap",[a("v-icon",{staticStyle:{"font-size":"16px","margin-top":"-6px"},attrs:{color:"pink"}},[t._v("fastfood")]),t._v("  \n                    include breakfast")],1)],1)])],1),a("v-container",{staticStyle:{"max-width":"1360px"},attrs:{fluid:"","grid-list-sm":""}},[a("v-layout",{attrs:{wrap:"",gallery:""}},[a("v-flex",{staticStyle:{"min-height":"250px"},attrs:{xs12:"",sm6:t.photosCount>1}},[a("div",{staticClass:"mainphoto",style:"height: 100%; background-image: url("+t.detail.Photo1+"); background-size: cover!important; background-color: #ddd; position: relative"},[a("div",{staticStyle:{position:"absolute",top:"0",right:"0","background-color":"rgba(0,0,0,.5)",color:"white","text-align":"center",padding:"5px 20px","margin-top":"20px"}},[a("b",{staticClass:"display-1"},[t._v(t._s(t.detail.Rating_average))]),a("br"),t.detail.Rating_average>=9?a("span",[t._v("Exceptional")]):t._e(),t.detail.Rating_average>=8&&t.detail.Rating_average<9?a("span",[t._v("Exellent")]):t._e(),t.detail.Rating_average>=7&&t.detail.Rating_average<8?a("span",[t._v("Very good")]):t._e(),t.detail.Rating_average>=6&&t.detail.Rating_average<7?a("span",[t._v("Good")]):t._e()]),a("div",{staticStyle:{position:"absolute",top:"100px",right:"0","font-weight":"bold","background-color":"rgba(238,89,93, .9)",color:"white","text-align":"center",padding:"5px 20px","margin-top":"20px"}},[a("span",[t._v("usd ")]),a("b",{staticClass:"headline"},[t._v(t._s(t.detail.Rates_from))])])])]),a("v-flex",{attrs:{xs12:"",sm6:""}},[a("v-layout",{attrs:{wrap:""}},[a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{style:"height: 200px; background-image: url("+t.detail.Photo2+"); background-size: cover!important; background-color: #ddd; position: relative"})]),a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{style:"height: 200px; background-image: url("+t.detail.Photo3+"); background-size: cover!important; background-color: #ddd; position: relative"})]),a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{style:"height: 200px; background-image: url("+t.detail.Photo4+"); background-size: cover!important; background-color: #ddd; position: relative"})]),a("v-flex",{attrs:{item:"",xs6:""}},[a("div",{style:"height: 200px; background-image: url("+t.detail.Photo1+"); background-size: cover!important; background-color: #ddd; position: relative"})])],1)],1)],1)],1),a("v-container",{attrs:{"grid-list-sm":"","mt-5":""}},[a("v-layout",{attrs:{row:"",wrap:""}},[a("v-flex",{attrs:{xs12:"",sm6:"",md7:"","pl-3":""}},[a("v-layout",{attrs:{row:"",wrap:""}},[a("v-flex",{staticStyle:{"border-bottom":"1px dashed #999","padding-bottom":"10px"},attrs:{xs12:"",headline:""}},[t._v("Nearest essentials")]),a("v-layout",{attrs:{row:"",wrap:"","ml-1":"","mt-4":""}},[a("v-flex",{attrs:{xs12:"",sm6:"","body-1":""}},[a("b",[t._v("Airports")]),a("br"),t._v("Ngurah Rai International Airport (DPS) ( 30.14 km )\n                                "),a("br"),t._v("Lombok International Airport (LOP) ( 114.59 km )")]),a("v-flex",{attrs:{xs12:"",sm6:"","body-1":""}},[a("b",[t._v("Drug stores")]),a("br"),t._v("Puri Mandela Pharmacy ( 8.24 km )")]),a("v-flex",{attrs:{xs12:"",sm6:"","body-1":""}},[a("b",[t._v("Cash withdrawal")]),a("br"),t._v("ATM ( 1.06 km )")]),a("v-flex",{attrs:{xs12:"",sm6:"","body-1":""}},[a("b",[t._v("Bus and rail stations")]),a("br"),t._v("Bangli Bus Terminal ( 11.85 km )\n                                "),a("br"),t._v("Ubung Bus Terminal ( 16.74 km )")]),a("v-flex",{attrs:{xs12:"",sm6:"","body-1":""}},[a("b",[t._v("Hospitals and clinics")]),a("br"),t._v("Praktek Dokter Dipayana ( 5.29 km )")]),a("v-flex",{attrs:{xs12:"",sm6:"","body-1":""}},[a("b",[t._v("Convenience stores")]),a("br"),t._v("Indomaret ( 7.11 km )")])],1),a("p",[t._v(" ")]),a("p",[t._v(" ")]),a("v-flex",{staticStyle:{"border-bottom":"1px dashed #999","padding-bottom":"10px"},attrs:{xs12:"","mb-3":"",headline:""}},[t._v("Hotel Reviews")]),t._l(t.reviews,function(e,i){return a("v-layout",{key:i,attrs:{"my-3":"",row:"",wrap:""}},[a("v-flex",{attrs:{xs3:"","text-xs-right":"","pr-2":"","mt-3":""}},[a("p",{staticClass:"title"},[a("b",[t._v("8.4")])]),a("p",{staticClass:"body-1 grey--text"},[t._v("\n                                    mellissa from Singapore\n                                    "),a("br"),t._v(" Couple\n                                    "),a("br"),t._v(" Suite Pool\n                                ")])]),a("v-flex",{attrs:{xs9:""}},[a("div",{staticClass:"body-1 grey--text"},[t._v("Stayed 3 nights in December 2016")]),a("div",{staticClass:"title my-2"},[a("em",[t._v("“Wonderful place with excellent service”")])]),a("div",{staticClass:"subheading"},[t._v("Enjoyed our stay. We had a pool villa which was lovely. A few problems with things not working were resolved quickly and efficiently. Staff were very helpful.")]),a("div",{staticClass:"body-1 grey--text"},[t._v("Reviewed December 20, 2016")])])],1)})],2)],1),a("v-flex",{attrs:{xs12:"",sm5:"",md4:"","offset-sm1":"",banner:""}},[a("div",{staticClass:"flex",staticStyle:{"min-height":"300px"}},[a("p",{staticClass:"blue--text",attrs:{align:"center"}},[a("a",{staticClass:"viewonmap blue--text",staticStyle:{"text-decoration":"underline","font-size":"12px!important"},on:{click:t.openMap}},[t._v("click view full map")])]),a("iframe",{staticStyle:{"background-color":"lightgreen"},attrs:{src:"http://localhost:8080/gmap",allowTransparency:"true",frameborder:"0",width:"100%",height:"300"}})]),a("div",{staticClass:"flex mt-3",staticStyle:{background:"url('/gallery/img-3.jpg') no-repeat","background-size":"100% 90vh","min-height":"300px"}})])],1)],1)],1),a("gmap",{attrs:{map:t.map,mapData:t.mapData}})],1)};i._withStripped=!0;var r={render:i,staticRenderFns:[]};e.a=r},"C+Hg":function(t,e,a){"use strict";var i=a("mtWM"),r=a.n(i);e.a={data:function(){return{breadcrumb:[],map:!1,mapData:[],overviewFull:!1,photosCount:4,detail:[],hotelLoader:!0,reviews:[{},{},{},{}]}},methods:{openMap:function(){this.map=!0,this.mapData={1:"satu",2:"dua"}}},beforeMount:function(){var t=this;this.hotelLoader=!0,r.a.get(this.$store.state.apiPath+"/front?mod=hotel&hotelnamekey="+this.$route.params.hotelnamekey+"&city_cc="+this.$route.params.city).then(function(e){t.detail=e.data,t.detail.Excerpt=t.detail.Overview.substr(0,200),""==t.detail.Photo2&&(t.detail.Photo2=t.detail.Photo1),""==t.detail.Photo3&&(t.detail.Photo3=t.detail.Photo1),""==t.detail.Photo4&&(t.detail.Photo4=t.detail.Photo1),t.hotelLoader=!1,t.breadcrumb={city:e.data.City,citykey:e.data.Citykey,country:e.data.Country,countryisocode_lower:e.data.Countryisocode_lower,hotel_name:e.data.Hotel_name,hotel_namekey:e.data.Hotel_namekey}})}}},OUav:function(t,e,a){(t.exports=a("FZ+f")(!1)).push([t.i,"#hotel .stars{float:left;margin-top:-2px}#hotel .stars i{font-size:20px}#hotel .overview a{font-size:90%;text-decoration:underline;color:#27a7f7}@media only screen and (min-width:600px){#hotel .banner{padding-right:10px!important}}",""])},TJf4:function(t,e,a){var i=a("OUav");"string"==typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);a("rjj0")("675d7e73",i,!1,{sourceMap:!1})},TdPg:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var i=a("C+Hg"),r=a("0A0B"),s=!1;var o=function(t){s||a("TJf4")},l=a("VU/8")(i.a,r.a,!1,o,null,null);l.options.__file="pages/_hotelnamekey/hotel/_city/index.vue",e.default=l.exports}});