package subscriptionscrapehandler

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/NaoNaoOnline/apiserver/pkg/contract/subscriptioncontract"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	crefmt = "onchain state for creator addresses %v does not match offchain state %v"
	unifmt = "onchain state for subscription timestamp %d does not match offchain state %d"
	vldfmt = "creator addresses %v do not match criteria of legitimate content creators"
)

func (h *ScrapeHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var sid string
	{
		sid = tas.Sync.Get(objectlabel.SubsObject)
	}

	var cnt string
	var rpc string
	{
		cnt = tas.Meta.Get(objectlabel.SubsCntrct)
		rpc = tas.Meta.Get(objectlabel.SubsRpcUrl)
	}

	var sob []*subscriptionstorage.Object
	{
		sob, err = h.sub.SearchSubs([]objectid.ID{objectid.ID(sid)})
		if subscriptionstorage.IsSubscriptionObjectNotFound(err) {
			h.log.Log(
				context.Background(),
				"level", "warning",
				"message", "stopped processing task",
				"object", sid,
				"reason", "subscription object not found",
			)

			return nil
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	var eth *ethclient.Client
	{
		eth, err = ethclient.Dial(rpc)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var scn *subscriptioncontract.Subscription
	{
		scn, err = subscriptioncontract.NewSubscription(common.HexToAddress(cnt), eth)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var sec *big.Int
	{
		sec, err = scn.GetSubUni(nil, big.NewInt(sob[0].Rcvr.Int()))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var add [3]common.Address
	{
		add, err = scn.GetSubRec(nil, big.NewInt(sob[0].Rcvr.Int()))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var uni int64
	{
		uni = int64(sec.Uint64())
	}

	var cre []string
	{
		cre = filAdd(add)
	}

	// There are potentially multiple tasks for multiple blockchain networks.
	// Any given task may or may not find a registered subscription onchain. We
	// assume that the chain we are looking at has no registered subscription if
	// the subscription timestamp and the creator addresses are empty. In such a
	// case we return here because the task has nothing more to do.
	if uni == 0 && len(cre) == 0 {
		return nil
	}

	// Ensure the onchain and offchain state of subscription timestamps match. And
	// if they don't, then finalize the subscription at hand with a failure
	// reason.
	if uni != sob[0].Unix.Unix() {
		sob[0].Fail = fmt.Sprintf(unifmt, uni, sob[0].Unix.Unix())
		sob[0].Stts = objectstate.Failure
	}

	// Ensure the onchain and offchain state of creator addresses match. And if
	// they don't, then finalize the subscription at hand with a failure reason.
	if len(cre) != len(sob[0].Crtr) || !generic.All(cre, sob[0].Crtr) {
		sob[0].Fail = fmt.Sprintf(crefmt, cre, sob[0].Crtr)
		sob[0].Stts = objectstate.Failure
	}

	var vld []bool
	{
		vld, err = h.sub.VerifyAddr(cre)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if !vldAdd(vld) {
		sob[0].Fail = fmt.Sprintf(vldfmt, cre)
		sob[0].Stts = objectstate.Failure
	}

	// Mark the subscription object in redis as valid since all checks passed.
	// Note that all code branches from above flow down here, so it is important
	// to not overwrite any failure state assigned.
	if sob[0].Stts != objectstate.Failure {
		sob[0].Stts = objectstate.Success
	}

	{
		_, err = h.sub.UpdateObct([]*subscriptionstorage.Object{sob[0]})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

// filAdd takes a list of 3 address types representing the onchain state of the
// subscription contract, and returns a list of strings representing the valid
// addresses from the given set.
func filAdd(add [3]common.Address) []string {
	var str []string

	for _, x := range add {
		// Anything with 20 leading zeros, in real life, is, very unlikely, a real
		// address.
		if !strings.HasPrefix(x.Hex(), "0x00000000000000000000") {
			str = append(str, x.Hex())
		}
	}

	return str
}

func vldAdd(vld []bool) bool {
	for _, x := range vld {
		if !x {
			return false
		}
	}

	return true
}
