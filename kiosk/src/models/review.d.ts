export interface IReview {
  userDisplayName: string
  productSKU: string
  textReview: string
  rating: number
  anonymous: boolean
  createdAt: Date
}