import {VuexModule, Module, getModule, MutationAction} from "vuex-module-decorators";
import store from "@/store"
import {UserProfile} from "@/models/user";
import UserAPI from "@/modules/api/user/userapi";

export interface IUserState {
  userProfile: UserProfile | undefined;
}

@Module({
  dynamic: true,
  namespaced: true,
  name: "User",
  store
})
class User extends VuexModule implements IUserState {
  public get userProfile(): UserProfile | undefined {
    return this.userProfile;
  }

  @MutationAction({ mutate: ["userProfile"] })
  async setUserProfile(email: string) {
    let userProfile: UserProfile | null = null;
    new UserAPI().fetchUserProfile(email).then(result => {
      console.log(result);
      userProfile = result;
    });
    return { userProfile };
  }
}

export const UserModule = getModule(User);