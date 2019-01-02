package stringsutil

// Concat concatenates string arrays into one
func Concat(params ...[]string) []string {
	x := make([]string, 0)
	if params != nil {
		for _, z := range params {
			x = append(x, z...)
		}
	}
	return x
}
