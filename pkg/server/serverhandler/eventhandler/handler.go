package eventhandler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
	Prm permission.Interface
}

type Handler struct {
	eve eventstorage.Interface
	log logger.Interface
	prm permission.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}

	return &Handler{
		eve: c.Eve,
		log: c.Log,
		prm: c.Prm,
	}
}

func inpLab(str string) []objectid.ID {
	var lis []objectid.ID

	for _, x := range strings.Split(str, ",") {
		if x != "" {
			lis = append(lis, objectid.ID(x))
		}
	}

	return lis
}

func inpDur(str string) time.Duration {
	sec, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return time.Duration(sec) * time.Second
}

func inpTim(str string) time.Time {
	sec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(sec, 0).UTC()
}

func outDur(dur time.Duration) string {
	return strconv.Itoa(int(dur.Seconds()))
}

func outLab(sco []objectid.ID) string {
	var str []string

	for _, x := range sco {
		str = append(str, string(x))
	}

	return strings.Join(str, ",")
}

func outTim(tim time.Time) string {
	return strconv.FormatInt(tim.Unix(), 10)
}
