package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) SearchRcrd() ([]*policycache.Record, error) {
	var err error

	var rec []*policycache.Record
	{
		rec = p.pol.SearchRcrd()
	}

	// Especially durring the program's startup sequence it may happen that no
	// policy records have been buffered and merged yet. So in order to prevent
	// invalid storage calls below we just return nil if there is in fact not a
	// single policy available right now.
	if len(rec) == 0 {
		return nil, nil
	}

	var add []string
	for _, x := range rec {
		add = append(add, x.Memb)
	}

	var use []objectid.ID
	{
		use, err = p.wal.SearchAddr(add)
		if walletstorage.IsWalletObjectNotFound(err) {
			// It may happen, especially during development or first platform
			// deployment, that there is only one policy record without an associated
			// wallet object. The redis implementation of the storage interfaces
			// returns "not found" errors if single objects cannot be found. In that
			// case we simply return the policy record that we have, without
			// augmenting it with a user ID.
			return rec, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for i := range rec {
		rec[i].User = use[i]
	}

	return rec, nil
}

func (p *Permission) SearchUser(use objectid.ID) ([]string, error) {
	var err error

	var wal []*walletstorage.Object
	{
		wal, err = p.wal.SearchKind(use, []string{"eth"})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var mem []string
	for _, x := range wal {
		if p.pol.ExistsMemb(x.Addr.Data) {
			mem = append(mem, x.Addr.Data)
		}
	}

	return mem, nil
}
