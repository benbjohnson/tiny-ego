package template

func formatStringPtr(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
