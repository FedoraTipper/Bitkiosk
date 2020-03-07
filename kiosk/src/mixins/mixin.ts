import routeDefinitions from "@/router/routes";
import config from "@/utils/config/config";
import { Component, Vue } from "vue-property-decorator";
import logger from "vuex/dist/logger";

@Component
export default class Mixin extends Vue {
  routeDefinitions: Object;
  config: Object;
  constructor() {
    super();
    this.config = config;
    this.routeDefinitions = routeDefinitions;
  }

  pushToPage(path: string) {
    if (this.$router.currentRoute.path != path) {
      this.$router.push({ path: path });
    } else {
      this.$router.go(0);
    }
  }

  pushToPageWithParams(name: string, params: any) {
    console.log(params);
    if (this.$router.currentRoute.name != name && this.$router.currentRoute.params != params) {
      this.$router.push({ name: name, params: params });
    } else {
      this.$router.go(0);
    }
  }
}
