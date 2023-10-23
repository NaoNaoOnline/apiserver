package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) BufferActv() (bool, error) {
	var err error

	var act []*policystorage.Object
	{
		act, err = p.pol.SearchActv()
		if simple.IsNotFound(err) {
			// fall through
		} else if err != nil {
			return false, tracer.Mask(err)
		}
	}

	if len(act) != 0 {
		err = p.cac.UpdateRcrd(act)
		if err != nil {
			return false, tracer.Mask(err)
		}

		return true, nil
	}

	return false, nil
}
