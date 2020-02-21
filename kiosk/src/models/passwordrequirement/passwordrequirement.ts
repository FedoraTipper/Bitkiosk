import { IPasswordRequirement } from "@/models/passwordrequirement/passwordrequirement.d.ts";

export default class PasswordRequirement implements IPasswordRequirement{
  private _regex: RegExp = new RegExp(``);
  private _addedScore: number = 0;
  private _maxCount: number = 0;
  private _minCount: number = 0;

  constructor(Regex: RegExp, addedScore: number, maxCount: number = -1, minCount: number = -1) {
    this._regex = Regex;
    this._addedScore = addedScore;
    this._maxCount = maxCount;
    this._minCount = minCount;
  }

  get regex(): RegExp {
    return this._regex;
  }

  set regex(value: RegExp) {
    this._regex = value;
  }

  get addedScore(): number {
    return this._addedScore;
  }

  set addedScore(value: number) {
    this._addedScore = value;
  }

  get maxCount(): number {
    return this._maxCount;
  }

  set maxCount(value: number) {
    this._maxCount = value;
  }

  get minCount(): number {
    return this._minCount;
  }

  set minCount(value: number) {
    this._minCount = value;
  }
}

