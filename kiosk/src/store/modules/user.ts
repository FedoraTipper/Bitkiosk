import {
  VuexModule,
  Module,
  getModule,
  MutationAction,
  Mutation, Action
} from "vuex-module-decorators";
import store from "@/store";
import { UserProfile } from "@/models/userprofile.ts";
import UserAPI from "@/modules/graphql/user/usergql";
import NotificationUtil from "@/utils/notification/notificationutil";

export interface IUserState {
  userProfile: UserProfile;
}

@Module({
  dynamic: true,
  namespaced: true,
  name: "User",
  store
})
class User extends VuexModule implements IUserState {
  userProfile: UserProfile = new UserProfile();

  @Mutation
  async setUserProfile(displayError: boolean) {
    new UserAPI().fetchUserProfile("").then(result => {
      this.userProfile = result;

      if (displayError) {
        if (this.userProfile.loggedIn !== true) {
          new NotificationUtil().displayError(
            "Unable to retrieve your user profile. :("
          );
        }
      }
    });
  }

  @Mutation
  async destroyUserSession() {
    this.userProfile = new UserProfile();
  }
}

export const UserModule = getModule(User);
