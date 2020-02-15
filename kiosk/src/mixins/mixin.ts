import routeDefinitions from "@/router/routes";
import config from "@/utils/config/config";
import {Component, Vue} from "vue-property-decorator";

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
      this.$router.push(path);
    } else {
      this.$router.go(0);
    }
  }
}