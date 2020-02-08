import { IUserProfile } from "@/models/userprofile.d.ts";

export class UserProfile implements IUserProfile{
  private _dateOfBirth: String | undefined;
  private _email: String | undefined;
  private _firstName: String | undefined;
  private _lastName: String | undefined;
  private _loggedIn: boolean = false;
  private _role: Number | undefined;

  constructor() {}

  setUserProfileFromResponse(result: IUserProfile) {
    this._dateOfBirth = result.dateOfBirth;
    this._email = result.email;
    this._firstName = result.firstName;
    this._lastName = result.lastName;
    this._loggedIn = true;
    this._role = result.role;
  }

  setUserProfile(dateOfBirth: String, email: String, firstName: String, lastName: String, loggedIn: boolean, role: Number) {
    this._dateOfBirth = dateOfBirth;
    this._email = email;
    this._firstName = firstName;
    this._lastName = lastName;
    this._loggedIn = true;
    this._role = role;
  }

  get dateOfBirth(): String | undefined {
    return this._dateOfBirth;
  }

  set dateOfBirth(value: String | undefined) {
    this._dateOfBirth = value;
  }

  get email(): String | undefined {
    return this._email;
  }

  set email(value: String | undefined) {
    this._email = value;
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

  get loggedIn(): boolean {
    return this._loggedIn;
  }

  set loggedIn(value: boolean) {
    this._loggedIn = value;
  }

  get role(): Number | undefined {
    return this._role;
  }

  set role(value: Number | undefined) {
    this._role = value;
  }
}