<style scoped>
  p, button {
    margin-left: 25px;
  }
  #startCapture {
    background-color: darkseagreen;
  }
  #close {
    background-color: mistyrose;
  }
</style>
<template>
  <modal name="captureSettings"
         :width="300"
         :height="200"
         @before-open="beforeOpen"
         @before-close="beforeClose">
    <h4 align="center">Capture Settings</h4>
    <p>Device:<b>{{interfaceName}}</b></p>
    <p>IP4 address:<b>{{ip4Address}}</b></p>
    <p>IP6 address:<b>{{ip6Address}}</b></p>
    <button id="startCapture" @click="captureStart()">Start capture!</button>
    <button id="close" @click="$modal.hide('captureSettings')">Close</button>
  </modal>
</template>
<script>
  export default {
    name: 'CaptureSettings',
    data () {
      return {
        interfaceName: '',
        ip4Address: '',
        ip6Address: ''
      }
    },
    methods: {
      beforeOpen (event) {
        console.log(event)
        this.interfaceName = event.params.interfaceName
        this.ip4Address = event.params.ip4Address
        this.ip6Address = event.params.ip6Address
      },
      beforeClose (event) {
        console.log(event)
      },
      captureStart () {
        this.$router.push({name: 'capture', params: { interfaceName: this.interfaceName }})
        this.$modal.hide('captureSettings')
      }
    }
  }
</script>
