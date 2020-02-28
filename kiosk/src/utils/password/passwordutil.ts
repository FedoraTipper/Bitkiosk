import PasswordRequirement from "@/models/passwordrequirement/passwordrequirement";
import { IPasswordScore } from "@/models/passwordrequirement/passwordrequirement.d.ts";
import { vsprintf } from "sprintf-js";

export default class PasswordUtil {
  private readonly _passwordRequirements: Array<PasswordRequirement>;

  constructor() {
    this._passwordRequirements = this.generateDefaultPasswordRequirements();
  }

  calculatePasswordStrength(password: string): IPasswordScore {
    let score: number = 0;
    let requirementsMet: boolean = true;
    let errorMessages: Array<string> = new Array<string>();
    if (password == undefined || password.length < 1) {
      return <IPasswordScore>{ score: 0, requirementsMet: false, errorMessages: [] };
    }

    this._passwordRequirements.forEach((passwordRequirement: PasswordRequirement) => {
        let matches = password.match(passwordRequirement.regex);

        if (matches == null || matches.length < passwordRequirement.minCount) {
          requirementsMet = false;
          errorMessages.push(
            vsprintf(passwordRequirement.errorMessage, [
              passwordRequirement.minCount
            ])
          );
        }

        if (matches != null && matches.length > 0) {
          score += matches.length >= passwordRequirement.maxCount
              ? passwordRequirement.addedScore
              : Math.ceil(
                  passwordRequirement.addedScore * (matches.length / passwordRequirement.maxCount)
                );
        }
      }
    );

    return <IPasswordScore>{ score, requirementsMet, errorMessages };
  }

  generateDefaultPasswordRequirements(): Array<PasswordRequirement> {
    let requirementList: Array<PasswordRequirement> = new Array<PasswordRequirement>();
    requirementList.push(new PasswordRequirement(/[^\w\d]/g, 10,"%d special character [!@#^ etc]", 0, 1));
    requirementList.push(new PasswordRequirement(/[A-Z]/g, 10, "%d uppercase [A-Z]",0, 1));
    requirementList.push(new PasswordRequirement(/\d/g, 10, "%d digit [0-9]",0, 1));
    requirementList.push(new PasswordRequirement(/./g, 70, "a length of %d characters or more",28, 8));
    return requirementList;
  }
}
