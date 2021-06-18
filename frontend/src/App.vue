<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>
<script>
import Vue from 'vue'
import WebRTC from 'vue-webrtc'
import  * as io from 'socket.io-client'

window.io = io

// Vue.component(WebRTC.name, WebRTC)
Vue.use(WebRTC)

export default {
  name: 'App',
  components: {

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100%;
}
</style>
