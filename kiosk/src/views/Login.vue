<template>
  <section class="hero is-light is-fullheight-with-navbar is-bold">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="columns is-centered" style="margin-bottom: 2rem">
          <h3 class="title has-text-grey-dark">User Login</h3>
        </div>
        <div class="columns is-centered">
          <div class="box column is-4">
            <LoginForm/>
          </div>
        </div>
        <div class="columns is-centered" style="margin-top: 1rem">
          <p class="has-text-grey">
            <a :href="this.routeDefinitions.signup.path">Sign Up</a> &nbsp;·&nbsp;
            <a href="../">Forgot Password</a> &nbsp;·&nbsp;
            <a href="../">Need Help?</a>
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { AuthMixin } from "@/mixins/authmixin";
import LoginForm from "@/components/LoginForm.vue";
import Component, { mixins } from 'vue-class-component';

@Component({
  components: {
    LoginForm
  }
})
export default class Login extends mixins(AuthMixin) {
  constructor() {
    super();
  }
}
</script>
