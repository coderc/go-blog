import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import BlogView from "@/views/BlogView.vue";
import LinkView from "@/views/LinkView.vue";
import QuestionView from "@/views/QuestionView.vue";

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: "/blog",
    name: "blog",
    component: BlogView
  },
  {
    path: "/link",
    name: "link",
    component: LinkView
  },
  {
    path: "/question",
    name: "question",
    component: QuestionView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
