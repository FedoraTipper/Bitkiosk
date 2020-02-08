<template>
  <div>
    <b-field
      label="Email"
      :type="emailErrors.length != 0 ? 'is-danger' : ''"
      :message="emailErrors"
    >
      <b-input type="email" v-model="email" />
    </b-field>
    <b-field
      label="Password"
      :type="passwordErrors.length != 0 ? 'is-danger' : ''"
      :message="passwordErrors"
    >
      <b-input v-model="password" type="password" />
    </b-field>
    <b-checkbox v-model="rememberMe">
      Remember me
    </b-checkbox>
    <br />
    <b-button type="is-primary" @click="performLoginAction">Login</b-button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import routes from "@/router/routes";
import Authhandler from "@/modules/authentication/authhandler";
import { UserModule } from "@/store/modules/user";

@Component
export default class LoginForm extends Vue {
  password: string = "";
  email: string = "";
  passwordErrors: string[] = [];
  emailErrors: string[] = [];
  rememberMe: boolean = false;

  performLoginAction() {
    const details = {
      identification: this.email,
      token: this.password,
      authMethodId: 1
    };
    new Authhandler().Login(details).then(result => {
      if (result) {
        UserModule.setUserProfile(true);
        this.$router.push(routes.home.path);
      }
    });
  }
}
</script>
