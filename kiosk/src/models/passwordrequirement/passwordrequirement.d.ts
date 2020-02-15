export interface IPasswordRequirement {
  regex: string;
  minCount: number;
  maxCount: number;
  addedScore: number;
}