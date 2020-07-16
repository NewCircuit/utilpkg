package stringutil

// MergeStrings formats two strings by setting the secondary string as a fallback.
func MergeStrings(a string, b string) string {
	if a == "" {
		return b
	}

	if b == "" {
		return a
	}

	return a + " (" + b + ")"
}
