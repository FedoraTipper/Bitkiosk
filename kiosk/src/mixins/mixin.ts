import routeDefinitions from "@/router/routes";
import config from "@/utils/config/config";
import {Component, Emit, Vue} from "vue-property-decorator";

@Component
export default class Mixin extends Vue {
  routeDefinitions: Object;
  config: Object;

  constructor() {
    super();
    this.config = config;
    this.routeDefinitions = routeDefinitions;
  }

  pushToPage(name: string) {
    this.pushToPageWithParams(name, {});
  }

  pushToPageWithParams(name: string, params: any) {
    console.log(params);
    if (this.$router.currentRoute.name != name && this.$router.currentRoute.params != params) {
      this.$router.push({ name: name, params: params });
    } else {
      this.$router.go(0);
    }

    this.$emit("page-change");
  }
}
