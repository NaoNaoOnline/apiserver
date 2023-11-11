package permission

import "github.com/xh3b4sd/tracer"

func (p *Permission) CreateLock() error {
	var err error

	{
		err = p.pol.CreateLock()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
