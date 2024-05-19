<template>
  <NavBar />
  <router-view/>
  <PingBackend/>
</template>

<script>
import PingBackend from "@/components/PingBackend.vue";
import NavBar from "@/components/NavBar.vue";
export default {
  components: {
    PingBackend,
    NavBar
  }
}
</script>

<script setup>
import {cError, cInfo, cWarn} from "@/common/console/console"

cInfo("check console level")
cWarn("check console level")
cError("check console level")
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
