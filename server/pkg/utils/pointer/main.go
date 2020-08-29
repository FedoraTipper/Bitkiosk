package pointer

func DereferenceInt(i *int) int {
	if i == nil {
		return 0
	}

	return *i
}