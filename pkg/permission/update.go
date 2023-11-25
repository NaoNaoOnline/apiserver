package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) UpdateActv() error {
	var err error

	var buf []*policystorage.Object
	{
		buf, err = p.pol.SearchBffr()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = p.cac.UpdateRcrd(buf)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var act []*policystorage.Object
	{
		act = p.cac.SearchRcrd()
	}

	{
		err = p.pol.CreateActv(act)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = p.pol.DeleteBffr()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = p.loc.Delete(objectlabel.PlcyLocker)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
