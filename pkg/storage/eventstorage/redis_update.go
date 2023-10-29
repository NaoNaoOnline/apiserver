package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateClck(obj []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range obj {
		// Track the new link click on the event object by incrementing its internal
		// counter.
		{
			obj[i].Clck.Data++
		}

		// Track the time of the last updated link click.
		{
			obj[i].Clck.Time = time.Now().UTC()
		}

		// Verify the modified event object to ensure the applied changes are not
		// rendering it invalid.
		{
			err := obj[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the modified event object is still valid after tracking the
		// new like, we update its normalized key-value pair.
		{
			err = r.red.Simple().Create().Element(eveObj(obj[i].Evnt), musStr(obj[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Updated)
		}
	}

	return out, nil
}
