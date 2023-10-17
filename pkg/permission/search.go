package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) SearchMemb(add []string) ([]objectid.ID, error) {
	var err error

	var use []objectid.ID
	{
		use, err = p.wal.SearchAddr(add)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return use, nil
}

func (p *Permission) SearchRcrd() ([]*policycache.Record, error) {
	var err error

	var rec []*policycache.Record
	{
		rec = p.pol.SearchRcrd()
	}

	var add []string
	for _, x := range rec {
		add = append(add, x.Memb)
	}

	var use []objectid.ID
	{
		use, err = p.SearchMemb(add)
		if err != nil {
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
