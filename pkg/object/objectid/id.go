package objectid

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/xh3b4sd/tracer"
)

type ID string

func Paging(num int64) ID {
	return ID(fmt.Sprintf("%d%06d", num, 0))
}

func Random(num int64) ID {
	return ID(fmt.Sprintf("%d%06d", num, rand.Intn(999999)))
}

func System() ID {
	return "1"
}

func (i ID) Float() float64 {
	f, e := strconv.ParseFloat(string(i), 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return f
}

func (i ID) Int() int64 {
	n, e := strconv.ParseInt(string(i), 10, 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return n
}

func (i ID) String() string {
	return string(i)
}
