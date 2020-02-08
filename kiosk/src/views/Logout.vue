<template>
  <div>
    Logging you out.
  </div>
</template>

<script>
import { UserModule } from "@/store/modules/user";
import AuthHandler from "@/modules/authentication/authhandler";
import NotificationUtil from "@/utils/notification/notificationutil";

export default {
  name: "Logout",
  components: {},
  mounted() {
    new AuthHandler()
      .Logout()
      .then(() => {
        UserModule.destroyUserSession();
        this.$router.push(this.routeDefinitions.home.path);
      })
      .catch(error => {
        new NotificationUtil().displayError("Unable to log you out.");
      });
  }
};
</script>
