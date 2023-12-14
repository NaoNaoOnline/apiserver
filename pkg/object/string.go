package object

import (
	"fmt"
	"reflect"
	"strings"
)

func String(obj any) string {
	var typ reflect.Type
	{
		typ = reflect.TypeOf(obj)
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	var pkg string
	{
		pkg = typ.PkgPath()
	}

	var spl []string
	{
		spl = strings.Split(pkg, "/")
	}

	if len(spl) != 0 {
		pkg = spl[len(spl)-1]
	}

	return fmt.Sprintf("%s.%s", pkg, typ.Name())
}
