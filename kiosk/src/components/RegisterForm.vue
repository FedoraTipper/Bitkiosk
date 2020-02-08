<template>
  <div>
    <div class="columns">
      <div class="column">
        <b-field label="First Name">
          <b-input type="text" v-model="firstName" />
        </b-field>
      </div>
      <div class="column">
        <b-field label="Last Name">
          <b-input type="text" v-model="lastName" />
        </b-field>
      </div>
    </div>
    <div class="columns">
      <div class="column is-offset-2 is-center is-three-fifths">
        <b-field
          label="Email"
          :type="emailErrors.length != 0 ? 'is-danger' : ''"
          :message="emailErrors"
        >
          <b-input type="email" icon-pack="fas" icon="envelope">
            icon="email" v-model="email"></b-input>
        </b-field>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <b-field
          label="Password"
          :type="passwordErrors.length != 0 ? 'is-danger' : ''"
          :message="passwordErrors"
        >
          <b-input v-model="password" type="password" icon="lock" />
        </b-field>
      </div>
      <div class="column">
        <b-field
          label="Confirm Password"
          :type="passwordErrors.length != 0 ? 'is-danger' : ''"
          :message="passwordErrors"
        >
          <b-input v-model="password" type="password" />
        </b-field>
      </div>
    </div>
    <br />
    <b-button type="is-primary" @click="performRegister">Sign Up</b-button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import routes from "@/router/routes";
import Authhandler from "@/modules/authentication/authhandler";
import { UserModule } from "@/store/modules/user";

@Component
export default class RegisterForm extends Vue {
  firstName: string = "";
  lastName: string = "";
  dateOfBirth: string = "";
  password: string = "";
  email: string = "";
  passwordErrors: string[] = [];
  emailErrors: string[] = [];

  performRegister() {
    new Authhandler().Register();
  }
}
</script>
