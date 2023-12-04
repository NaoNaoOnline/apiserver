package liststorage

import (
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the list description complies with the expected format.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Crtd = now
			inp[i].List = objectid.Random(objectid.Time(now))
			inp[i].Desc.Data = strings.TrimSpace(inp[i].Desc.Data)
		}

		// Once we know the list description is valid, we create the normalized
		// key-value pair so that we can search for list objects using their IDs.
		{
			err = r.red.Simple().Create().Element(lisObj(inp[i].List), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all lists they created.
		{
			err = r.red.Sorted().Create().Score(lisUse(inp[i].User), inp[i].List.String(), inp[i].List.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
