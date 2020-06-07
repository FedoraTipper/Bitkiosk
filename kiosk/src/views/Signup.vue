<template>
  <section class="hero is-light is-fullheight-with-navbar is-bold">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered" style="margin-bottom: 2rem">
          <h3 class="title has-text-grey-dark">User Registration</h3>
        </div>
        <div class="columns is-centered is-vertical">
          <div class="box column is-4 is-vertical">
            <RegisterForm />
          </div>
        </div>
        <div class="columns is-centered">
          <p class="has-text-grey" style="margin-top: 1rem">
            <a @click="pushToPage(routeDefinitions.login.name)">
              Have an account?
            </a>
            &nbsp;·&nbsp;
            <a @click="pushToPage(routeDefinitions.home.name)">Need Help?</a>
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { AuthMixin } from "@/mixins/authmixin";
import Component, { mixins } from "vue-class-component";
import RegisterForm from "@/components/RegisterForm.vue";

@Component({
  components: {
    RegisterForm
  }
})
export default class Login extends mixins(AuthMixin) {
  constructor() {
    super();
  }
}
</script>
