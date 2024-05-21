<template>
  <ul class="list-group">
    <li v-for="(blog,index) in blogs" :key="index" class="list-group-item">
      <ContentCard :blog="blog"></ContentCard>
    </li>
  </ul>
</template>

<script>
// @ is an alias to /src
import ContentCard from "@/components/blog/ContentCard.vue";
import {mapGetters} from "vuex";

export default {
  name: 'HomeView',
  components: {
    ContentCard
  },
  computed: {
    ...mapGetters(['blogs'])
  },
  mounted() {
    console.log("home view mounted")
    this.$store.commit("setSingleBlog", {"title": "titlexx", "content": "content"})
  }
}
</script>
