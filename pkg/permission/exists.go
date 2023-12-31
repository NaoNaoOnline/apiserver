package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) ExistsAcce(sys int64, uid objectid.ID, acc int64) (bool, error) {
	var err error

	var mem []string
	{
		mem, err = p.SearchUser(uid)
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

func (p *Permission) ExistsMemb(uid objectid.ID) (bool, error) {
	var err error

	var mem []string
	{
		mem, err = p.SearchUser(uid)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	// SearchUser returns the SMA members for the given user ID. So as long as mem
	// is not empty, the given user is considered a policy member.
	return len(mem) != 0, nil
}

func (p *Permission) ExistsSyst(sys int64, uid objectid.ID) (bool, error) {
	var err error

	var mem []string
	{
		mem, err = p.SearchUser(uid)
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
