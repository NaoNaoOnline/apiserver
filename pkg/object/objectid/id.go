package objectid

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/xh3b4sd/tracer"
)

type ID string

func Random(num int64) ID {
	return ID(fmt.Sprintf("%d%06d", num, rand.Intn(999999)))
}

func System() ID {
	return "0"
}

func (i ID) Float() float64 {
	f, e := strconv.ParseFloat(string(i), 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return f
}

func (i ID) String() string {
	return string(i)
}
