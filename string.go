package utility

func ValidatePointerString(value *string) bool {
	if value == nil || *value == "" {
		return false
	}
	return true
}

func ValidateString(value string) bool {
	if value == "" {
		return false
	}
	return true
}