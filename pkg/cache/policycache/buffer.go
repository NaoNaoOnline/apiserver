package policycache

import (
	"github.com/xh3b4sd/tracer"
)

func (m *Memory) Buffer(rec []*Record) error {
	{
		m.mut.Lock()
		defer m.mut.Unlock()
	}

	{
		if len(rec) == 0 {
			return tracer.Mask(policyRecordEmptyError)
		}
	}

	for i := range rec {
		// At first we need to validate the given input object and, amongst others,
		// whether the buffered record originates from a single chain.
		{
			err := rec[i].Verify()
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	var cid int64
	{
		cid = rec[0].ChID[0]
	}

	for i := range rec {
		if len(rec[i].ChID) > 1 {
			return tracer.Mask(policyChIDLimitError)
		}
		if rec[i].ChID[0] != cid {
			return tracer.Mask(policyChIDInvalidError)
		}
	}

	{
		m.buf[cid] = rec
	}

	return nil
}
