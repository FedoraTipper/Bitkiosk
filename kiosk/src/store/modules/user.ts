import {
  VuexModule,
  Module,
  getModule,
  MutationAction,
  Mutation, Action
} from "vuex-module-decorators";
import store from "@/store";
import User from "@/models/user";
import UserAPI from "@/modules/graphql/user/usergql";
import NotificationUtil from "@/utils/notification/notificationutil";

export interface IUserState {
  user: User;
}

@Module({
  dynamic: true,
  namespaced: true,
  name: "User",
  store
})
class UserStore extends VuexModule implements IUserState {
  user: User = new User();

  @Mutation
  async setUserProfile(displayError: boolean) {
    new UserAPI().fetchUserProfile("").then(result => {
      this.user = result;

      if (displayError) {
        if (this.user.loggedIn !== true) {
          new NotificationUtil().displayError(
            "Unable to retrieve your user profile. :("
          );
        }
      }
    });
  }

  @Mutation
  async destroyUserSession() {
    this.user = new User();
  }
}

export const UserModule = getModule(UserStore);
