import Vue from "vue";
import VueRouter from "vue-router";
import routeDefinitions from "./routes";

Vue.use(VueRouter);

const routes = [
  {
    path: routeDefinitions.home.path,
    name: routeDefinitions.home.name,
    component: routeDefinitions.home.component
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  },
  {
    path: routeDefinitions.login.path,
    name: routeDefinitions.login.name,
    component: routeDefinitions.login.component
  },
  {
    path: routeDefinitions.signup.path,
    name: routeDefinitions.signup.name,
    component: routeDefinitions.signup.component
  },
  {
    path: routeDefinitions.logout.path,
    name: routeDefinitions.logout.name,
    component: routeDefinitions.logout.component
  },
  {
    path: routeDefinitions.productview.path,
    name: routeDefinitions.productview.name,
    component: routeDefinitions.productview.component
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
