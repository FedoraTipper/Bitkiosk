import Component from "vue-class-component";
import {UserModule} from "@/store/modules/user";
import Vue from 'vue';

@Component
export class AuthMixin extends Vue {
  constructor() {
    super();
  }

  created() {
    this.authOnLoad();
  }

  authOnLoad() {
    if (!UserModule.user.loggedIn) {
      UserModule.setUserProfile(false);
    }
  }
}
