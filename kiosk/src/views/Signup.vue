<template>
  <section class="hero is-light is-fullheight-with-navbar is-bold">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-5 is-offset-4">
          <h3 class="title has-text-grey-dark">User Registration</h3>
          <div class="box">
            <LoginForm :register="true" actionButtonText="Sign Up" />
          </div>
          <p class="has-text-grey">
            <a :href="this.routeDefinitions.login.path">Have an account?</a> &nbsp;·&nbsp;
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
  name: "Register",
  components: {
    LoginForm
  },
  mounted() {}
};
</script>
