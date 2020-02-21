<template>
  <div>
    <div class="columns">
      <div class="column is-half">
        <b-field label="First Name">
          <b-input type="text" v-model="firstName" />
        </b-field>
      </div>
      <div class="column is-half">
        <b-field label="Last Name">
          <b-input type="text" v-model="lastName" />
        </b-field>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <b-field
          label="Email"
          :type="emailErrors.length != 0 ? 'is-danger' : ''"
          :message="emailErrors"
        >
          <b-input type="email" v-model="email" @click="console.log(email)"/>
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
        <b-progress :value="passwordScore" :type="passwordType" v-if="this.password.length > 0 "></b-progress>
      </div>
    </div>
    <b-button type="is-primary" @click="performRegister">Sign Up</b-button>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue, Watch} from "vue-property-decorator";
import EmailValidator from "email-validator";
import Authhandler from "@/modules/authentication/authhandler";
import PasswordUtil from "@/utils/password/passwordutil";
import { IPasswordScore } from "@/models/passwordrequirement/passwordrequirement.d.ts";

@Component
export default class RegisterForm extends Vue {
  firstName: string = "";
  lastName: string = "";
  dateOfBirth: string = "";
  password: string = "";
  passwordScore: number = 0;
  passwordType: string = "is-danger";
  email: string = "";
  passwordErrors: string[] = [];
  emailErrors: string[] = [];

  private readonly _PasswordUtil: PasswordUtil = new PasswordUtil();

  constructor() {
    super();
    console.log(this._PasswordUtil);
  }

  performRegister() {
    if (this.validateForm()) {
      new Authhandler().Register();
    }
  }

  validateForm(): boolean {
    return this.validateEmail(this.email) && this.validatePassword(this.password);
  }

  @Watch('email', {immediate: false})
  validateEmail(val: string) : boolean{
    if (!EmailValidator.validate(val)) {
      this.emailErrors = ["Double check your email address."];
    } else {
      this.emailErrors = [];
    }
    return this.emailErrors.length < 1;
  }

  @Watch('password', {immediate: false})
  validatePassword(val: string): boolean {
    let newScore = {} as IPasswordScore;
    if (val != undefined && val.length > 0) {
      newScore = new PasswordUtil().calculatePasswordStrength(val);
    }
    console.log(newScore);
    if (!newScore.requirementsMet) {
      this.passwordErrors = ["The password should consist of one uppercase A-Z, number 0-9 and special character [!@#$%]."];
    } else {
      this.passwordErrors = [];
    }

    this.passwordScore = newScore.score;

    return this.emailErrors.length < 1;
  }

  @Watch('passwordScore')
  setPasswordScoreColour(val: number): void {
    if (val >= 50) {
      this.passwordType = "is-success";
    } else if (val >= 35) {
      this.passwordType = "is-warning";
    } else {
      this.passwordType = "is-danger";
    }
  }
}
</script>
