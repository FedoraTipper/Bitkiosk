import { IProduct } from "@/models/product.d.ts";
import User from "@/models/user.ts";

export default class Product implements IProduct {
  private _SKU!: string;
  private _description!: string;
  private _shortDescription!: string;
  private _name!: string;
  private _price!: number;
  private _stock!: number;
  private _rating!: number;
  private _reviewCount!: number;
  private _createdByAdmin!: User;
  private _startDate!: Date;
  private _endDate!: Date;
  private _createdAt!: Date;
  private _updatedAt!: Date;

  constructor() {}

  setProductFromResponseObject(obj: IProduct) {
    this._SKU = obj.SKU;
    this._name = obj.name;
    this._description = obj.description;
    this._shortDescription = obj.shortDescription;
    this._price = obj.price;
    this._stock = obj.stock;
    this._rating = obj.rating;
    this._reviewCount = obj.reviewCount;
    this._startDate = new Date(obj.startDate);
    this._endDate = new Date(obj.endDate);
    this._createdAt = new Date(obj.createdAt);
    this._updatedAt = new Date(obj.updatedAt);
  }

  get SKU(): string {
    return this._SKU;
  }

  set SKU(value: string) {
    this._SKU = value;
  }

  get description(): string {
    return this._description;
  }

  set description(value: string) {
    this._description = value;
  }

  get shortDescription(): string {
    if (this._shortDescription.length > 0)
      return this._shortDescription;

    return "No description";
  }

  set shortDescription(value: string) {
    this._shortDescription = value;
  }

  get name(): string {
    return this._name;
  }

  set name(value: string) {
    this._name = value;
  }

  get price(): number {
    return this._price;
  }

  set price(value: number) {
    this._price = value;
  }

  get stock(): number {
    return this._stock;
  }

  get rating(): number {
    return this._rating;
  }

  set rating(value: number) {
    this._rating = value;
  }

  get reviewCount(): number {
    return this._reviewCount;
  }

  set reviewCount(value: number) {
    this._reviewCount = value;
  }

  set stock(value: number) {
    this._stock = value;
  }

  get createdByAdmin(): User {
    return this._createdByAdmin;
  }

  set createdByAdmin(value: User) {
    this._createdByAdmin = value;
  }

  get startDate(): Date {
    return this._startDate;
  }

  set startDate(value: Date) {
    this._startDate = value;
  }

  get endDate(): Date {
    return this._endDate;
  }

  set endDate(value: Date) {
    this._endDate = value;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  set createdAt(value: Date) {
    this._createdAt = value;
  }

  get updatedAt(): Date {
    return this._updatedAt;
  }

  set updatedAt(value: Date) {
    this._updatedAt = value;
  }

  public getReviewDisplay(): string {
    if (this.reviewCount == 0) {
      return "No Reviews";
    }

    let display = this.reviewCount + " Review";
    if (this.reviewCount > 1)
      display += "s";
    return display;
  }
}
