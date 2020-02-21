import PasswordRequirement from "@/models/passwordrequirement/passwordrequirement";
import { IPasswordScore } from "@/models/passwordrequirement/passwordrequirement.d.ts";

export default class PasswordUtil {
  private readonly _passwordRequirements: Array<PasswordRequirement>;

  constructor() {
    this._passwordRequirements = this.generateDefaultPasswordRequirements();
  }

  calculatePasswordStrength(password: string): IPasswordScore {
    let score: number = 0;
    let requirementsMet: boolean = true;

    if (password == undefined || password.length < 1) {
      return <IPasswordScore>{ score: 0, requirementsMet: false };
    }

    this._passwordRequirements.forEach((passwordRequirement: PasswordRequirement) => {
        let matches = password.match(passwordRequirement.regex);
        if (matches != null && matches.length >= passwordRequirement.minCount) {
          score += passwordRequirement.addedScore;
        }
      }
    );

    return <IPasswordScore>{ score, requirementsMet };
  }

  generateDefaultPasswordRequirements(): Array<PasswordRequirement> {
    let requirementList: Array<PasswordRequirement> = new Array<PasswordRequirement>();
    requirementList.push(new PasswordRequirement(/[^\w\d]/g, 20, 0, 1));
    requirementList.push(new PasswordRequirement(/[A-Z]/g, 15, 0, 1));
    requirementList.push(new PasswordRequirement(/\d/g, 15, 0, 1));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    return requirementList;
  }
}
