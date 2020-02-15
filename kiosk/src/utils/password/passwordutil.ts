import PasswordRequirement from "@/models/passwordrequirement/passwordrequirement";

export default class PasswordUtil {
  private _passwordRequirements: Array<PasswordRequirement>;

  constructor() {
    this._passwordRequirements = this.generateDefaultPasswordRequirements();
  }

  calculatePasswordStrength(password: string): number {
    let score: number = 0;

    console.log("222222222222222222222")

    if (password == undefined || password.length < 1) {
      return 0;
    }

    console.log("333333333333333333")

    this._passwordRequirements.forEach((passwordRequirement: PasswordRequirement) => {
      let regExp = new RegExp(passwordRequirement.regex);
      regExp = regExp.compile();
      let matches = regExp.exec(password);
      if (matches != null && matches.length < passwordRequirement.minCount) {
        score += passwordRequirement.addedScore;
      }
    });

    return score;
  }

  generateDefaultPasswordRequirements(): Array<PasswordRequirement> {
    let requirementList: Array<PasswordRequirement> = new Array<PasswordRequirement>();
    requirementList.push(new PasswordRequirement("[^\\w\\d]", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    // requirementList.push(new PasswordRequirement("", 20, 0, 0));
    return requirementList;
  }
}