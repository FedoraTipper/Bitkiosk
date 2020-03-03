package passwordvalidator

import (
	"errors"
	"fmt"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"regexp"
)

type PasswordRequirement struct {
	Regex string
	MinCount int
	Message string
}

type PasswordValidator struct {
	password string
	PasswordRequirements []PasswordRequirement
}

var minLength, minSpecialCharacters, minDigits, minUppercase int

func init() {
	minLength = utils.MustGetInt("PASSWORD_REQUIREMENT_MIN_LENGTH")
	minSpecialCharacters = utils.MustGetInt("PASSWORD_REQUIREMENT_MIN_SPECIAL_CHARACTERS")
	minDigits = utils.MustGetInt("PASSWORD_REQUIREMENT_MIN_DIGITS")
	minUppercase = utils.MustGetInt("PASSWORD_REQUIREMENT_MIN_UPPERCASE")
}

func DefaultPasswordValidator(password string) PasswordValidator {
	return PasswordValidator{
		password:             password,
		PasswordRequirements: []PasswordRequirement{
			{`.`, 8, "a length of %d characters or more"},
			{`[^\w\d]`, 1,"%d special character [!@#^ etc]"},
			{`[A-Z]`, 1, "%d uppercase [A-Z]"},
			{`\d`, 1, "%d digit [0-9]"},
		},
	}
}

func (p PasswordValidator) Validate() []error {
	var errs []error

	for _, passRequirement := range p.PasswordRequirements  {
		regex, err := regexp.Compile(passRequirement.Regex)

		if err != nil {
			errs = append(errs, errors.New("Unable to perform password requirement"))
			logger.Error("Unable to parse regex: \n " + passRequirement.Regex)
			continue
		}

		if len(regex.FindAllString(p.password, -1)) < passRequirement.MinCount {
			errs = append(errs, errors.New(fmt.Sprintf(passRequirement.Message, passRequirement.MinCount)))
		}
	}

	return errs
}