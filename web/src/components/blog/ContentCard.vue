<template>
  <div class="card text-center container-sm">
    <div class="card-header">
      Featured
    </div>
    <div class="card-body">
      <h5 class="card-title">{{ blog.title }}</h5>
      <p class="card-text">{{ blog.content }}</p>
      <button @click="commitBlog" class="btn btn-primary">commit</button>
      <div>blogCount: {{ blogCount }}</div>
    </div>
    <div class="card-footer text-muted">
      2 days ago
    </div>
  </div>
</template>

<script>
import {mapGetters} from "vuex";

export default {
  name: "ContentCard",
  props: {
    blog: Object,
  },
  computed: {
    ...mapGetters([
        'blogCount'
    ])
  },
  methods: {
    commitBlog() {
      this.$store.commit("setSingleBlog", {"title": "title", "content": "content"})
    }
  }
}
</script>

<style scoped>

</style>