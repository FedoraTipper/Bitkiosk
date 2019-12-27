package DBResult

type DBResult struct {
	Errors []error
	Info []string
	Warnings []string
}

func NewResult() (newResult *DBResult) {
	return &DBResult{
		Errors: []error{},
		Info:   []string{},
	}
}

func (r *DBResult) AddResult(newResult *DBResult) (result *DBResult) {
	result = r

	for _, err := range newResult.Errors {
		result = result.AddErrorToResult(err)
	}

	for _, info := range newResult.Info {
		result.Info = result.AddInfo(info)
	}

	return result
}

func (r *DBResult) AddInfo(newInfo string) (info []string) {
	info = r.Info
	return append(info, newInfo)
}

func (r *DBResult) AddErrorToResult(newError error) (result *DBResult) {
	result = r

	if newError == nil {
		return result
	}

	errs := result.Errors
	result.Errors = append(errs, newError)

	return result
}


func (r *DBResult) AddWarningToResult(newWarning string) (result *DBResult) {
	result = r

	if newWarning == "" {
		return result
	}

	currentWarnings := result.Warnings
	result.Warnings = append(currentWarnings, newWarning)

	return result
}

func (r *DBResult) IsError() bool {
	return len(r.Errors) > 0
}

func (r *DBResult) IsOk() bool {
	return !r.IsError()
}

func (r *DBResult) GetFirstError() error {
	if r.IsError() {
		return r.Errors[0]
	}

	return nil
}