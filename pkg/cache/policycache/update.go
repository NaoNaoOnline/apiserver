package policycache

import (
	"slices"
	"sort"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/tracer"
)

func (m *Memory) UpdateRcrd(inp []*policystorage.Object) error {
	{
		m.mut.Lock()
		defer m.mut.Unlock()
	}

	if len(inp) == 0 {
		return tracer.Mask(policyBufferEmptyError)
	}

	{
		m.cac = []*policystorage.Object{}
		m.mem = map[string]struct{}{}
		m.rec = map[int64]map[string]*policystorage.Object{}
	}

	buf := map[int64][]*policystorage.Object{}
	for _, x := range inp {
		buf[x.ChID[0]] = append(buf[x.ChID[0]], x)
	}

	// For a reliable and stable merge result we need to process the buffered
	// policy records in order of their respective chain IDs. So we collect and
	// sort them, and then process the buffered policy records based on their
	// sorted order below.
	var cid []int64
	for k := range buf {
		cid = append(cid, k)
	}

	{
		slices.Sort(cid)
	}

	for _, x := range cid {
		for _, y := range buf[x] {
			var exi *policystorage.Object
			{
				exi = m.searchRcrd(y.Syst, y.Memb)
			}

			if exi != nil && exi.Acce == y.Acce {
				err := m.update(exi, y)
				if err != nil {
					return tracer.Mask(err)
				}
			} else {
				err := m.create(y)
				if err != nil {
					return tracer.Mask(err)
				}
			}
		}
	}

	// Sort aggregated policies by access with secondary priority.
	sort.SliceStable(m.cac, func(i, j int) bool {
		return m.cac[i].Acce < m.cac[j].Acce
	})

	// Sort aggregated policies by system with first priority.
	sort.SliceStable(m.cac, func(i, j int) bool {
		return m.cac[i].Syst < m.cac[j].Syst
	})

	return nil
}

func (m *Memory) update(exi *policystorage.Object, inp *policystorage.Object) error {
	// Add the new chain specific information to the existing policy object.
	{
		exi.ChID = append(exi.ChID, inp.ChID...)
	}

	// Ensure a sorted chain ID order.
	{
		slices.Sort(exi.ChID)
	}

	// After merging the existing policy object with the new record we need to
	// verify whether the policy object is still valid. This may help prevent
	// issues caused by event indexing or duplicated chain IDs.
	{
		err := exi.Verify()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// At this point the existing record should be updated through its pointer
	// reference. The unit tests verify that assumption.

	return nil
}
