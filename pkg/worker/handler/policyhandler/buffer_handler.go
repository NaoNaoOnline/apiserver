package policyhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/contract/policycontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type BufferHandlerConfig struct {
	Cnt string
	Log logger.Interface
	Pol policycache.Interface
	Rpc string
}

type BufferHandler struct {
	cid int64
	eth *ethclient.Client
	log logger.Interface
	pcn *policycontract.Policy
	pol policycache.Interface
}

func NewBufferHandler(c BufferHandlerConfig) (*BufferHandler, int64) {
	if c.Cnt == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cnt must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Pol == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pol must not be empty", c)))
	}
	if c.Rpc == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rpc must not be empty", c)))
	}

	var err error

	var eth *ethclient.Client
	{
		eth, err = ethclient.Dial(c.Rpc)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var pcn *policycontract.Policy
	{
		pcn, err = policycontract.NewPolicy(common.HexToAddress(c.Cnt), eth)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var cid int64
	{
		big, err := eth.ChainID(context.Background())
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}

		cid = int64(big.Uint64())
	}

	var han *BufferHandler
	{
		han = &BufferHandler{
			cid: cid,
			eth: eth,
			log: c.Log,
			pcn: pcn,
			pol: c.Pol,
		}
	}

	return han, cid
}
