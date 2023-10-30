package keyfmt

func Strings(inp []string, fnc func(string) string) []string {
	var out []string

	for _, x := range inp {
		out = append(out, fnc(x))
	}

	return out
}
