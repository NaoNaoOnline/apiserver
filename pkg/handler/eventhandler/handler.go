package eventhandler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
}

type Handler struct {
	eve eventstorage.Interface
	log logger.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Handler{
		eve: c.Eve,
		log: c.Log,
	}
}

func inpCat(str string) []scoreid.String {
	var lis []scoreid.String

	for _, x := range strings.Split(str, ",") {
		lis = append(lis, scoreid.String(x))
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

	return time.Unix(sec, 0)
}

func outCat(sco []scoreid.String) string {
	var str []string

	for _, x := range sco {
		str = append(str, string(x))
	}

	return strings.Join(str, ",")
}

func outDur(dur time.Duration) string {
	return strconv.Itoa(int(dur.Seconds()))
}

func outTim(tim time.Time) string {
	return strconv.FormatInt(tim.Unix(), 10)
}
