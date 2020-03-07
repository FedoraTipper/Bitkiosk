import { IUser } from "@/models/user.d.ts";
import { UserProfile } from "@/models/userprofile";

export default class User implements IUser {
  private _userProfile!: UserProfile;
  private _email!: string;
  private _role!: number;
  private _loggedIn: boolean = false;

  constructor() {}

  setUserFromResponseObject(obj: IUser) {
    this._role = obj.role;
    this._email = obj.email;

    let userProfile: UserProfile = new UserProfile();
    userProfile.setUserProfileFromResponse(obj.userProfile);
    this._userProfile = userProfile;

    this._loggedIn = true;
  }

  get userProfile(): UserProfile {
    return this._userProfile;
  }

  set userProfile(value: UserProfile) {
    this._userProfile = value;
  }

  get email(): string {
    return this._email;
  }

  set email(value: string) {
    this._email = value;
  }

  get role(): number {
    return this._role;
  }

  set role(value: number) {
    this._role = value;
  }

  get loggedIn(): boolean {
    return this._loggedIn;
  }

  set loggedIn(value: boolean) {
    this._loggedIn = value;
  }
}
