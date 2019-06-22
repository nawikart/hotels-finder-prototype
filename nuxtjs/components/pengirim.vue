<template>
  <div>
    <button color="success" @click="helloworld">klik disini untuk helloworld</button>
     
  </div>
</template>


<script>
export default {
 
  methods:{
    helloworld(){
      this.$emit('clickedx', 'Hello world nawiiiii') //this emmit itu kirim dari CHILD ke root component/page
    }
  }

}
</script>