package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (f *Feed) SearchEvnt(rid objectid.ID, pag [2]int) ([]objectid.ID, error) {
	var err error

	var val []string
	{
		val, err = f.red.Sorted().Search().Order(keyfmt.EveRul(rid), pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return objectid.IDs(val), nil
}

func (f *Feed) SearchFeed(lid objectid.ID, pag [2]int) ([]objectid.ID, error) {
	var err error

	var val []string
	{
		val, err = f.red.Sorted().Search().Order(keyfmt.EveFee(lid), pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return objectid.IDs(val), nil
}

func (f *Feed) SearchList(eid objectid.ID, pag [2]int) ([]objectid.ID, error) {
	var err error

	var rid []objectid.ID
	{
		rid, err = f.SearchRule(eid, pag)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// there might not be any rule IDs for this list
	if len(rid) == 0 {
		return nil, nil
	}

	var val []string
	{
		val, err = f.red.Sorted().Search().Union(objectid.Fnc(rid, keyfmt.LisRul)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return objectid.IDs(val), nil
}

func (f *Feed) SearchRule(eid objectid.ID, pag [2]int) ([]objectid.ID, error) {
	var err error

	var val []string
	{
		val, err = f.red.Sorted().Search().Order(keyfmt.RulEve(eid), pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return objectid.IDs(val), nil
}
