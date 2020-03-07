import { IUser } from "@/models/user.d.ts";

export interface IProduct {
    SKU: string
    name: string
    description: string
    price: number
    stock: number
    startDate: Date
    endDate: Date
    createdByAdmin: IUser
    createdAt: Date
    updatedAt: Date
}
