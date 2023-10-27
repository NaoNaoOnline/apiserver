package rulehandler

import (
	"fmt"
	"strings"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Rul rulestorage.Interface
}

type Handler struct {
	log logger.Interface
	rul rulestorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rul == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rul must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		rul: c.Rul,
	}
}

func inpIDs(str string) []objectid.ID {
	var lis []objectid.ID

	for _, x := range strings.Split(str, ",") {
		if x != "" {
			lis = append(lis, objectid.ID(x))
		}
	}

	return lis
}

// func outIDs(ids []objectid.ID) string {
// 	var str []string

// 	for _, x := range ids {
// 		str = append(str, string(x))
// 	}

// 	return strings.Join(str, ",")
// }
