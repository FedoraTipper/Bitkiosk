export interface IPasswordRequirement {
  regex: RegExp;
  minCount: number;
  maxCount: number;
  addedScore: number;
}

export interface IPasswordScore {
  score: number;
  requirementsMet: boolean;
  errorMessages: Array<string>;
}