package policycache

import (
	"slices"

	"github.com/xh3b4sd/tracer"
)

func (m *Memory) Update() error {
	{
		m.mut.Lock()
		defer m.mut.Unlock()
	}

	if len(m.buf) == 0 {
		return tracer.Mask(policyBufferEmptyError)
	}

	{
		m.cac = []*Record{}
		m.mem = map[string]struct{}{}
		m.rec = map[int64]map[string]*Record{}
	}

	// For a reliable and stable merge result we need to process the buffered
	// policy records in order of their respective chain IDs. So we collect and
	// sort them, and then process the buffered policy records based on their
	// sorted order below.
	var cid []int64
	for k := range m.buf {
		cid = append(cid, k)
	}

	{
		slices.Sort(cid)
	}

	for _, x := range cid {
		for _, y := range m.buf[x] {
			var exi *Record
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

	{
		m.buf = map[int64][]*Record{}
	}

	return nil
}

func (m *Memory) update(exi *Record, inp *Record) error {
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
