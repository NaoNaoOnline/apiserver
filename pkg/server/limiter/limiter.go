package limiter

const (
	// Default is the default amount of response objects returned by any API using
	// the response limiter.
	Default = 1000
)

func Len(num int) int {
	if num > Default {
		return Default
	}

	return num
}

func Log(num int) bool {
	return num > Default
}
