<template>
  <b-navbar>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path:  this.routeDefinitions.home.path}">
        <img
          src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
          alt="Lightweight UI components for Vue.js based on Bulma"
        />
      </b-navbar-item>
    </template>

    <template slot="start">
      <b-navbar-item @click="$router.push(routeDefinitions.home.path)">
        Home {{userFirstName}}
      </b-navbar-item>
    </template>

    <template slot="end">
      <b-navbar-item tag="div">
        <div class="buttons">
          <a class="button is-primary" @click="$router.push(routeDefinitions.signup.path)">
            <strong>Sign up</strong>
          </a>
          <a class="button is-light" @click="$router.push(routeDefinitions.login.path)">
            Log in
          </a>
        </div>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script lang="ts">
import {Component, Prop, PropSync, Vue} from "vue-property-decorator";
import { UserModule } from "@/store/modules/user";

@Component
export default class NavBar extends Vue{
  constructor() {
    super();
  }

  get userFirstName(){
    if (UserModule.userProfile !== undefined) {
      return UserModule.userProfile.firstName
    }
    return 'a'
  }

}

</script>
