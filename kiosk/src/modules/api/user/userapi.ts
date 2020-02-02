import { UserProfile } from "../../../models/user";
import gqlfactory from "@/utils/gqlclient/gqlfactory";
import UserQueries from "@/modules/api/user/userqueries";
import NotificationUtil from "@/utils/notification/notificationutil";

export default class UserAPI {
  constructor() {}

  async fetchUserProfile(email: string): Promise<UserProfile | null> {
    return new Promise<UserProfile | null>(async resolve => {
      let userProfile: UserProfile | null = null;

      let GQLClient = new gqlfactory().newGQLClient();

      const inputData = {
        email: email
      };

      await GQLClient.request(UserQueries.getUserProfile, inputData)
        .then(response => {
          if (response) {
            // userProfile = JSON.parse(response['userProfile']);
            response = response['userProfile'];
            userProfile = <UserProfile>{
              firstName: response["firstName"],
              lastName: response["lastName"],
              email: email,
              dateOfBirth: response["dateOfBirth"],
              role: response["role"]
            };
          }
        })
        .catch(error => {
          console.log(error);
          new NotificationUtil().displayError(
            "Unable to retrieve your user profile. :("
          );
        });

      resolve(userProfile);
    });
  }
}
