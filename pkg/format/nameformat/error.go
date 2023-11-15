package nameformat

import (
	"github.com/xh3b4sd/tracer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Errorf(res string, fld string) tracer.Interface {
	cas := cases.Title(language.English)

	return &tracer.Error{
		Kind: cas.String(res) + cas.String(fld) + "FormatError",
		Desc: `The request expects the ` + res + ` ` + fld + ` to contain words, numbers or: . -. The ` + res + ` ` + fld + ` was found to contain invalid characters. Therefore the request failed.`,
	}
}
