import { IUserProfile } from "@/models/userprofile.d.ts";

export class UserProfile implements IUserProfile{
  private _firstName: String | undefined;
  private _lastName: String | undefined;

  constructor() {}

  setUserProfileFromResponse(result: IUserProfile) {
    this._firstName = result.firstName;
    this._lastName = result.lastName;
  }

  setUserProfile(firstName: String, lastName: String) {
    this._firstName = firstName;
    this._lastName = lastName;
  }

  get firstName(): String | undefined {
    return this._firstName;
  }

  set firstName(value: String | undefined) {
    this._firstName = value;
  }

  get lastName(): String | undefined {
    return this._lastName;
  }

  set lastName(value: String | undefined) {
    this._lastName = value;
  }
}