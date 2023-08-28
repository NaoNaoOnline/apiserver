package reactionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
)

func (r *Redis) Exists(inp objectid.String) bool {
	var err error

	var lis []*Object
	{
		lis, err = r.Search()
		if err != nil {
			panic(err)
		}
	}

	for _, x := range lis {
		if inp == x.Rctn {
			return true
		}
	}

	return false
}
