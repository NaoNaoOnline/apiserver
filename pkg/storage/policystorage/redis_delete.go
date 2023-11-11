package policystorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteBffr() error {
	// We use Simple.Delete to purge all buffered policy records, even though we
	// use Sorted.Create to persist them. All we want to do here is to get rid of
	// all buffered permission state. And Simple.Delete uses the redis command
	// DEL, which simply erases any key and its associated value.
	{
		_, err := r.red.Simple().Delete().Multi(keyfmt.PolicyBuffer)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *Redis) DeleteLock() error {
	{
		_, err := r.red.Simple().Delete().Multi(keyfmt.PolicyLock)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
