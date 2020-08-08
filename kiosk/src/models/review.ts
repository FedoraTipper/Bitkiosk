import { IReview } from "@/models/review.d.ts";

export default class Review implements IReview {
  private _userDisplayName!: string;
  private _productSKU!: string;
  private _textReview!: string;
  private _rating!: number;
  private _createdAt!: Date;

  constructor() {}

  setReviewFromResponseObject(obj: IReview) {
    console.log("setting obj")
    console.log(obj);
    this._userDisplayName = obj.userDisplayName;
    this._productSKU = obj.productSKU;
    this._textReview = obj.textReview;
    this._rating = obj.rating;
    this._createdAt = new Date(obj.createdAt);
  }

  get userDisplayName(): string {
    return this._userDisplayName;
  }

  set userDisplayName(value: string) {
    this._userDisplayName = value;
  }

  get productSKU(): string {
    return this._productSKU;
  }

  set productSKU(value: string) {
    this._productSKU = value;
  }

  get textReview(): string {
    return this._textReview;
  }

  set textReview(value: string) {
    this._textReview = value;
  }

  get rating(): number {
    return this._rating;
  }

  set rating(value: number) {
    this._rating = value;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  set createdAt(value: Date) {
    this._createdAt = value;
  }
}