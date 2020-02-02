<template>
  <section class="hero is-light is-fullheight-with-navbar is-bold">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-5 is-offset-4">
          <h3 class="title has-text-grey-dark">User Login</h3>
          <div class="box">
            <LoginForm :register="false" actionButtonText="Login" />
          </div>
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

<script>
import LoginForm from "@/components/LoginForm.vue";

export default {
  name: "Login",
  components: {
    LoginForm
  },
  mounted() {}
};
</script>
