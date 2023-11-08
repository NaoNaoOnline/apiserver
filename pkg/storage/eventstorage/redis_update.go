package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateClck(use objectid.ID, obj []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range obj {
		// Ensure user clicks are not counted on events that have already happened.
		{
			if obj[i].Happnd() {
				out = append(out, objectstate.Dropped)
				continue
			}
		}

		// Ensure user clicks are not counted twice.
		{
			exi, err := r.red.Sorted().Exists().Score(clkEve(obj[i].Evnt), use.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if exi {
				out = append(out, objectstate.Dropped)
				continue
			}
		}

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

		// We use a sorted set to store all the user IDs that have clicked on an
		// event link. This internal data structure is used to prevent counting
		// duplicates on this particular metric.
		{
			err = r.red.Sorted().Create().Score(clkEve(obj[i].Evnt), use.String(), use.Float())
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
