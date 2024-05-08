<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import {request} from '@/common/http/request'
import {backendUrl} from "@/config/backendUrl"
import {ref} from "vue"

let response = ref({})

function callback(resp) {
  response.value = resp.data
  console.log(resp)
}

function errorCallback(resp) {
  console.log(resp)
}

let req = new request("POST", backendUrl+"/ping", {}, callback, errorCallback)
req.send()

export default {
  name: 'HomeView',
  components: {
    HelloWorld,
  }
}
</script>
