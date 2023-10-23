package permission

import (
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) EnsureActv() error {
	var err error

	var buf bool
	{
		buf, err = p.BufferActv()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if !buf {
		err := p.emi.Scrape()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
