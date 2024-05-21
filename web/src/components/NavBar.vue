<script>
export default {
  name: "NavBar"
}
</script>

<template>
  <ul class="nav justify-content-center">
    <li class="nav-item">
      <router-link class="nav-link active" aria-current="page" to="/">coderc</router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" to="/blog">Blog</router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" to="/link">Link</router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" to="/question">Question</router-link>
    </li>
  </ul>
</template>

<style scoped>

</style>