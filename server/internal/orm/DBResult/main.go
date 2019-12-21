package DBResult

type DBResult struct {
	Errors []error
	Info []string
}

func (r *DBResult) AddResult(newResult *DBResult) (result *DBResult) {
	result = r

	for _, err := range r.Errors {
		result.Errors = result.AddError(err)
	}

	for _, info := range r.Info {
		result.Info = result.AddInfo(info)
	}

	return result
}

func (r *DBResult) AddInfo(newInfo string) (info []string) {
	info = r.Info
	return append(info, newInfo)
}

func (r *DBResult) AddError(newError error) (errs []error) {
	errs = r.Errors
	return append(errs, newError)
}

func (r *DBResult) IsError() bool {
	return len(r.Errors) > 0
}

func (r *DBResult) IsOk() bool {
	return !r.IsError()
}