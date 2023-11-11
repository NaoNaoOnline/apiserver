package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) ExistsAcce(sys int64, use objectid.ID, acc int64) (bool, error) {
	var err error

	var mem []string
	{
		mem, err = p.SearchUser(use)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	for _, x := range mem {
		if p.cac.ExistsAcce(sys, x, acc) {
			return true, nil
		}
	}

	return false, nil
}

func (p *Permission) ExistsLock() (bool, error) {
	var err error

	var exi bool
	{
		exi, err = p.pol.ExistsLock()
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return exi, nil
}

func (p *Permission) ExistsMemb(use objectid.ID) (bool, error) {
	var err error

	var mem []string
	{
		mem, err = p.SearchUser(use)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	for _, x := range mem {
		if p.cac.ExistsMemb(x) {
			return true, nil
		}
	}

	return false, nil
}

func (p *Permission) ExistsSyst(sys int64, use objectid.ID) (bool, error) {
	var err error

	var mem []string
	{
		mem, err = p.SearchUser(use)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	for _, x := range mem {
		if p.cac.ExistsSyst(sys, x) {
			return true, nil
		}
	}

	return false, nil
}
