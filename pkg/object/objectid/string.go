package objectid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xh3b4sd/tracer"
)

type String string

func New(tim time.Time) String {
	return String(fmt.Sprintf("%d%06d", tim.Unix(), rand.Intn(999999)))
}

func System() String {
	return "0"
}

func (s String) Float() float64 {
	f, e := strconv.ParseFloat(string(s), 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return f
}

func (s String) String() string {
	return string(s)
}
