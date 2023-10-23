package policystorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateActv(inp []*Object) error {
	var err error

	for _, x := range inp {
		err := x.Verify()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var byt []byte
	{
		byt, err = json.Marshal(inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = r.red.Simple().Create().Element(keyfmt.PolicyActive, string(byt))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *Redis) CreateBffr(inp []*Object) error {
	var err error

	{
		if len(inp) == 0 {
			return tracer.Mask(policyRecordEmptyError)
		}
	}

	// At first we need to validate the given input object and, amongst others,
	// whether the buffered record originates from a single chain.
	for _, x := range inp {
		err := x.Verify()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cid int64
	{
		cid = inp[0].ChID[0]
	}

	for _, x := range inp {
		if x.ChID[0] != cid {
			return tracer.Mask(policyChIDInvalidError)
		}
		if len(x.ChID) > 1 {
			return tracer.Mask(policyChIDLimitError)
		}
	}

	for _, x := range inp {
		err = r.red.Sorted().Create().Score(keyfmt.PolicyBuffer, musStr(x), float64(x.ChID[0]))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
