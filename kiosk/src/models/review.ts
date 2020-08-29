import { IReview } from "@/models/review.d.ts";

export const RATING_TEXT: string[] = ['Very bad', 'Bad', 'Good', 'Very good', 'Awesome'];

export default class Review implements IReview {
  private _userDisplayName!: string;
  private _productSKU!: string;
  private _textReview!: string;
  private _rating!: number;
  private _anonymous!: boolean;
  private _createdAt!: Date;

  constructor() {}

  setReviewFromResponseObject(obj: IReview) {
    this._userDisplayName = obj.userDisplayName;
    this._productSKU = obj.productSKU;
    this._textReview = obj.textReview;
    this._rating = obj.rating;
    this._anonymous = obj.anonymous;
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

  get anonymous(): boolean {
    return this._anonymous;
  }

  set anonymous(value: boolean) {
    this._anonymous = value;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  set createdAt(value: Date) {
    this._createdAt = value;
  }
}