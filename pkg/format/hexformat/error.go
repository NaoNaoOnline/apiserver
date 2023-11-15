package hexformat

import (
	"github.com/xh3b4sd/tracer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Errorf(res string, fld string) tracer.Interface {
	cas := cases.Title(language.English)

	return &tracer.Error{
		Kind: cas.String(res) + cas.String(fld) + "FormatError",
		Desc: "The request expects the " + res + " " + fld + " to be in hex format including 0x prefix. The " + res + " " + fld + " was not found to be in hex format including 0x prefix. Therefore the request failed.",
	}
}
