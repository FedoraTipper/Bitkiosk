import Vue from "vue";
import VueRouter from "vue-router";
import routeDefinitions from "./routes";

Vue.use(VueRouter);

let routes: Array<any> = [];

for (let [key, value] of Object.entries(routeDefinitions)) {
  let newRoute = {
    component: value.component,
    name: value.name,
    path: value.path
  };

  routes.push(newRoute);
}

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
