package handleformat

import "regexp"

var (
	regExp = regexp.MustCompile(`^[A-Za-z0-9._\-]+$`)
)

func Verify(str string) bool {
	return regExp.MatchString(str)
}
