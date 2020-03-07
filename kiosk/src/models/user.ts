import { IUser } from "@/models/user.d.ts";
import { UserProfile } from "@/models/userprofile";

export default class User implements IUser {
  private _UserProfile: UserProfile;
  private _email: string;
  private _role: number;

  constructor(UserProfile: UserProfile, email: string, role: number) {
    this._UserProfile = UserProfile;
    this._email = email;
    this._role = role;
  }

  get UserProfile(): UserProfile {
    return this._UserProfile;
  }

  set UserProfile(value: UserProfile) {
    this._UserProfile = value;
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
}
