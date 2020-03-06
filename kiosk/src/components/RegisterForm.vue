<template>
  <div>
    <div class="columns">
      <div class="column is-half">
        <b-field
          label="First Name"
          :type="firstNameError.length > 0 ? 'is-danger' : ''"
          :message="firstNameError"
        >
          <b-input type="text" v-model="firstName" />
        </b-field>
      </div>
      <div class="column is-half">
        <b-field
          label="Last Name"
          :type="lastNameError.length > 0 ? 'is-danger' : ''"
          :message="lastNameError"
        >
          <b-input type="text" v-model="lastName" />
        </b-field>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <b-field
          label="Email"
          :type="emailError.length > 0 ? 'is-danger' : ''"
          :message="emailError"
        >
          <b-input type="email" v-model="email" @click="console.log(email)" />
        </b-field>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <b-field
          label="Password"
          :type="passwordError.length > 0 ? 'is-danger' : ''"
          :message="passwordError"
        >
          <b-input v-model="password" type="password" icon="lock" />
        </b-field>
        <b-progress
          :value="passwordScore"
          :type="passwordType"
          v-if="this.password.length > 0"
        ></b-progress>
      </div>
    </div>
    <b-button type="is-primary" @click="performRegister">Sign Up</b-button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import EmailValidator from "email-validator";
import Authhandler from "@/modules/authentication/authhandler";
import PasswordUtil from "@/utils/password/passwordutil";
import { IPasswordScore } from "@/models/passwordrequirement/passwordrequirement.d.ts";
import NotificationUtil from "@/utils/notification/notificationutil";
import Usergql from "@/modules/graphql/user/usergql";
import { IRegisterDetails } from "@/models/authentication/authdetails";
import routes from '@/router/routes';

@Component
export default class RegisterForm extends Vue {
  firstName: string = "";
  firstNameError: string = "";
  lastName: string = "";
  lastNameError: string = "";
  password: string = "";
  passwordScore: number = 0;
  passwordType: string = "is-danger";
  email: string = "";
  passwordError: string = "";
  emailError: string = "";

  constructor() {
    super();
  }

  performRegister() {
    if (this.validateForm()) {
      new Usergql().registerUser({
          firstName: this.firstName,
          lastName: this.lastName,
          token: this.password,
          email: this.email,
          authMethodId: 1
        } as IRegisterDetails)
        .then(response => {
          if (response.success) {
            new NotificationUtil().displaySuccess("Successfully registered!");
            this.$router.push(routes.home.path);
          } else {
            new NotificationUtil().displayError(response.message);
          }
        });
    } else {
      new NotificationUtil().displayError(
        "Please recheck the registration input fields"
      );
    }
  }

  validateForm(): boolean {
    return (
      this.validateEmail(this.email) ||
      (this.validatePassword(this.password) &&
        this.validateFirstName(this.firstName)) ||
      this.validateLastName(this.lastName)
    );
  }

  @Watch("email", { immediate: false })
  validateEmail(val: string): boolean {
    this.emailError = !EmailValidator.validate(val)
      ? "Double check your email address."
      : "";
    return this.emailError.length < 1;
  }

  @Watch("password", { immediate: false })
  validatePassword(val: string): boolean {
    let newScore = {} as IPasswordScore;
    this.passwordError = "";
    if (val.length > 0) {
      newScore = new PasswordUtil().calculatePasswordStrength(val);
    } else {
      this.passwordError = "Please input a password.";
    }

    if (!newScore.requirementsMet && newScore.errorMessages) {
      let passwordError = "Your password requires";
      let errorMessages = newScore.errorMessages;
      for (let i = 0; i < errorMessages.length; i++) {
        passwordError += " " + errorMessages[i];

        if (i < errorMessages.length - 2) {
          passwordError += ",";
        } else if (i == errorMessages.length - 2) {
          passwordError += " and";
        } else if (i == errorMessages.length - 1) {
          passwordError += ".";
        }
      }
      this.passwordError = passwordError;
    }

    this.passwordScore = newScore.score;

    return this.passwordError.length > 0;
  }

  @Watch("passwordScore")
  setPasswordScoreColour(val: number): void {
    if (val >= 50) {
      this.passwordType = "is-success";
    } else if (val >= 35) {
      this.passwordType = "is-warning";
    } else {
      this.passwordType = "is-danger";
    }
  }

  @Watch("firstName")
  validateFirstName(val: string): boolean {
    this.firstNameError =
      val == undefined || val.length == 0 ? "Please input a first name." : "";
    return this.firstNameError.length <= 0;
  }

  @Watch("lastName")
  validateLastName(val: string): boolean {
    this.lastNameError =
      val == undefined || val.length == 0 ? "Please input a last name." : "";
    return this.lastNameError.length <= 0;
  }
}
</script>
