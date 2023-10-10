package policystorage

import (
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) ExistsSyst(sys int64, usr objectid.ID) (bool, error) {
	var err error

	// For every access control lookup we must know the user's current wallet
	// addresses. So we fetch all of the wallets for the given user ID.
	var wal []*walletstorage.Object
	{
		wal, err = r.wal.SearchKind(usr, []string{"eth"})
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	for _, x := range wal {
		var cre int64
		{
			cre, err = r.red.Sorted().Metric().Count(polSys("cre", sys, x.Addr.Data))
			if err != nil {
				return false, tracer.Mask(err)
			}
		}

		var del int64
		{
			del, err = r.red.Sorted().Metric().Count(polSys("del", sys, x.Addr.Data))
			if err != nil {
				return false, tracer.Mask(err)
			}
		}

		if cre > 0 && cre > del {
			return true, nil
		}
	}

	return false, nil
}

func (r *Redis) ExistsMemb(usr objectid.ID) (bool, error) {
	var err error

	// For every access control lookup we must know the user's current wallet
	// addresses. So we fetch all of the wallets for the given user ID.
	var wal []*walletstorage.Object
	{
		wal, err = r.wal.SearchKind(usr, []string{"eth"})
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	for _, x := range wal {
		var cre int64
		{
			cre, err = r.red.Sorted().Metric().Count(polMem("cre", x.Addr.Data))
			if err != nil {
				return false, tracer.Mask(err)
			}
		}

		var del int64
		{
			del, err = r.red.Sorted().Metric().Count(polMem("del", x.Addr.Data))
			if err != nil {
				return false, tracer.Mask(err)
			}
		}

		if cre > 0 && cre > del {
			return true, nil
		}
	}

	return false, nil
}

func (r *Redis) ExistsAcce(sys int64, usr objectid.ID, acc int64) (bool, error) {
	var err error

	// For every access control lookup we must know the user's current wallet
	// addresses. So we fetch all of the wallets for the given user ID.
	var wal []*walletstorage.Object
	{
		wal, err = r.wal.SearchKind(usr, []string{"eth"})
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	for _, x := range wal {
		var cre bool
		{
			cre, err = r.red.Sorted().Exists().Value(polSys("cre", sys, x.Addr.Data), strconv.FormatInt(acc, 10))
			if err != nil {
				return false, tracer.Mask(err)
			}
		}

		var del bool
		{
			del, err = r.red.Sorted().Exists().Value(polSys("del", sys, x.Addr.Data), strconv.FormatInt(acc, 10))
			if err != nil {
				return false, tracer.Mask(err)
			}
		}

		if cre && !del {
			return true, nil
		}
	}

	return false, nil
}
