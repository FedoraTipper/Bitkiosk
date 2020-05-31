import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Buefy from "buefy";
import "buefy/dist/buefy.css";
import mixin from "./mixins/mixin";
Vue.config.productionTip = false;

Vue.use(Buefy);

Vue.mixin(mixin);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
