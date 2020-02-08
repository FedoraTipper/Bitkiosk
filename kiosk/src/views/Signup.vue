<template>
  <section class="hero is-light is-fullheight-with-navbar is-bold">
    <div class="hero-body">
      <div class="container">
        <div class="column">
          <h3 class="title has-text-grey-dark">User Registration</h3>
          <div class="box">
            <RegisterForm/>
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
import RegisterForm from "@/components/RegisterForm.vue";

export default {
  name: "Register",
  components: {
    RegisterForm
  },
  mounted() {}
};
</script>
