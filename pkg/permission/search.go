package permission

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) SearchActv() ([]*policystorage.Object, error) {
	var err error

	var obj []*policystorage.Object
	{
		obj = p.cac.SearchRcrd()
	}

	// Especially durring the program's startup sequence it may happen that no
	// policy records have been buffered and merged yet. So in order to prevent
	// invalid storage calls below we just return nil if there is in fact not a
	// single policy available right now.
	if len(obj) == 0 {
		return nil, nil
	}

	var add []string
	for _, x := range obj {
		add = append(add, x.Memb)
	}

	var use []objectid.ID
	{
		use, err = p.wal.SearchAddr(add)
		if walletstorage.IsWalletObjectNotFound(err) {
			// It may happen, especially during development or first platform
			// deployment, that there is only one policy record without an associated
			// wallet object. The redis implementation of the storage interface
			// returns "not found" errors if single objects cannot be found. In that
			// case we simply return the policy record that we have, without
			// augmenting it with a user ID.
			return obj, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for i := range obj {
		obj[i].User = use[i]
	}

	return obj, nil
}

func (p *Permission) SearchUser(use objectid.ID) ([]string, error) {
	var err error

	// wal will result in a list of all wallet objects owned by the given user, if
	// any.
	var wal []*walletstorage.Object
	{
		wal, err = p.wal.SearchKind(use, []string{"eth"})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var mem []string
	for _, x := range wal {
		// Using PolicyCache.ExistsMemb here is critical for Permission.ExistsMemb
		// to work correctly. Only wallet addresses recorded in the policy contract
		// should be considered policy members.
		if x.HasLab(objectlabel.WalletModeration) && p.cac.ExistsMemb(x.Addr.Data) {
			mem = append(mem, x.Addr.Data)
		}
	}

	return mem, nil
}

func (p *Permission) SearchTime() (time.Time, error) {
	var err error

	var tim time.Time
	{
		tim, err = p.pol.SearchTime()
		if err != nil {
			return time.Time{}, tracer.Mask(err)
		}
	}

	return tim, nil
}
