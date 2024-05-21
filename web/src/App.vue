<template>
  <NavBar />
  <router-view/>
  <PingBackend/>
  <p class="text-center"><a href="https://beian.miit.gov.cn" target="_blank" class="text-body-tertiary">ICP备案号：鄂ICP备2024055858号-1</a></p>
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

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav {
  color: #42b983;
}
</style>
