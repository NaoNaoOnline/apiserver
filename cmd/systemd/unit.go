package systemd

import (
	"fmt"
)

type Unit struct {
	cou int
	nam string
	tem string
}

func (u Unit) Cou() int {
	return u.cou
}

func (u Unit) Nam(i int) string {
	if u.cou > 1 {
		return fmt.Sprintf(u.nam, i)
	} else {
		return u.nam
	}
}

func (u Unit) Tem() string {
	return u.tem
}
