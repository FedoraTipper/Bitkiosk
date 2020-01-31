import routeDefinitions from "@/router/routes";
import {Component, Vue} from "vue-property-decorator";

@Component
export default class Mixin extends Vue {
  routeDefinitions: Object;
  constructor() {
    super();
    this.routeDefinitions = routeDefinitions;
  }
}