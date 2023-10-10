package policyhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/contract/policycontract"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Eth *ethclient.Client
	Log logger.Interface
	Pcn *policycontract.Policy
	Pol policystorage.Interface
}

type SystemHandler struct {
	eth *ethclient.Client
	log logger.Interface
	pcn *policycontract.Policy
	pol policystorage.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Eth == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eth must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Pcn == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pcn must not be empty", c)))
	}
	if c.Pol == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pol must not be empty", c)))
	}

	return &SystemHandler{
		eth: c.Eth,
		log: c.Log,
		pcn: c.Pcn,
		pol: c.Pol,
	}
}
