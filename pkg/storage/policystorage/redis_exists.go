package policystorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) ExistsLock() (bool, error) {
	var err error

	var val []string
	{
		val, err = r.red.Simple().Search().Multi(keyfmt.PolicyLock)
		if simple.IsNotFound(err) {
			return false, nil
		} else if err != nil {
			return false, tracer.Mask(err)
		}
	}

	if len(val) == 0 {
		return false, nil
	}

	if len(val) != 1 {
		return false, tracer.Mask(runtime.ExecutionFailedError)
	}

	return val[0] == "1", nil
}
