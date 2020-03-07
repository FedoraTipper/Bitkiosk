import { IUserProfile } from "@/models/userprofile.d.ts";

export interface IUser {
    email: string
    role: number
    UserProfile: IUserProfile
}