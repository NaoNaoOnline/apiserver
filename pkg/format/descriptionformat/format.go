package descriptionformat

import "regexp"

var (
	regExp = regexp.MustCompile(`^[A-Za-z0-9,.:\-\+'"!$%&#\s]+$`)
)

func Verify(str string) bool {
	return regExp.MatchString(str)
}
