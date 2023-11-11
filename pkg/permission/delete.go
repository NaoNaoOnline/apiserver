package permission

import "github.com/xh3b4sd/tracer"

func (p *Permission) DeleteLock() error {
	var err error

	{
		err = p.pol.DeleteLock()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
