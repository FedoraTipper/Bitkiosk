import { IUser } from "@/models/user.d.ts";

export interface IProduct {
    SKU: string
    name: string
    description: string
    shortDescription: string
    price: number
    stock: number
    rating: number
    reviewCount: number
    startDate: Date
    endDate: Date
    createdByAdmin: IUser
    createdAt: Date
    updatedAt: Date
}
