package utility

func GetBoolMapFromStringSlice(slice []string) map[string]bool {
	st := make(map[string]bool)
	for _, sliceItem := range slice {
		st[sliceItem] = true
	}
	return st
}
