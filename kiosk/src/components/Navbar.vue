<template>
  <b-navbar
    :shadow="isViewOnHomePage()"
    :type="navBarType"
  >
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: this.routeDefinitions.home.name}">
        <img
          src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
          alt="Lightweight UI components for Vue.js based on Bulma"
        />
      </b-navbar-item>
    </template>

    <template slot="start">
      <b-navbar-item @click="pushToPage(routeDefinitions.home.name)">
        Home
      </b-navbar-item>
    </template>

    <template slot="end">
      <b-navbar-item id="loginNav" tag="div" v-if="!userLoggedIn">
        <div class="buttons">
          <a class="button is-primary" @click="pushToPage(routeDefinitions.signup.name)">
            <strong>Sign up</strong>
          </a>
          <a class="button is-light" @click="pushToPage(routeDefinitions.login.name)">
            Log in
          </a>
        </div>
      </b-navbar-item>
      <b-navbar-item id="logoutNav" tag="div" v-else>
        <strong style="margin-right: 10px">Welcome {{userFirstName}}</strong>
        <div class="buttons">
          <a class="button is-primary" @click="pushToPage(routeDefinitions.logout.name)">
            Log out
          </a>
        </div>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import { UserModule } from "@/store/modules/user";
import routes from "@/router/routes";

@Component
export default class NavBar extends Vue {
  constructor() {
    super();
  }

  private get userFirstName() {
    return UserModule.user.userProfile.firstName;
  }

  private get userLoggedIn() {
    return UserModule.user.loggedIn;
  }

  private get currentPath(): string {
    return this.$route.path;
  }

  private get navBarType() {
    if (this.isViewOnHomePage()) {
      return "is-white";
    }

    return "is-primary";
  }

  private isViewOnHomePage(): boolean {
    return this.currentPath == routes.home.path;
  }
}
</script>
