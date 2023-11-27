package subscriptionscrapehandler

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/NaoNaoOnline/apiserver/pkg/contract/subscriptioncontract"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *ScrapeHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var sid string
	{
		sid = tas.Meta.Get(objectlabel.SubsObject)
	}

	var sub []*subscriptionstorage.Object
	{
		sub, err = h.sub.SearchSubs([]objectid.ID{objectid.ID(sid)})
		// TODO stop processing if sub not found, otherwise the task gets stuck forever
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cid string
	var cnt string
	var rpc string
	{
		cid = tas.Meta.Get(objectlabel.SubsChanid)
		cnt = tas.Meta.Get(objectlabel.SubsCntrct)
		rpc = tas.Meta.Get(objectlabel.SubsRpcUrl)
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

	var unx *big.Int
	{
		unx, err = scn.GetSubUnx(nil, common.HexToAddress(sub[0].Sbsc))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var add [3]common.Address
	{
		add, err = scn.GetSubAdd(nil, common.HexToAddress(sub[0].Sbsc))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var uni int64
	{
		uni = int64(unx.Uint64())
	}

	var cre []string
	{
		cre = addStr(add)
	}

	// There are potentially multiple tasks for multiple blockchain networks.
	// Any given task may or may not find a registered subscription onchain. We
	// assume that the chain we are looking at has no registered subscription if
	// the subscription timestamp and the creator addresses are empty. In such a
	// case we return here because the task has nothing more to do.
	if uni == 0 && len(cre) == 0 {
		return nil
	}

	if cid != fmt.Sprintf("%d", sub[0].ChID) {
		// TODO update subscription object valid flag with failure reason
		return nil
	}

	if uni != sub[0].Unix.Unix() {
		// TODO update subscription object valid flag with failure reason
		return nil
	}

	if len(add) != len(sub[0].Crtr) {
		// TODO update subscription object valid flag with failure reason
		return nil
	}

	if !generic.All(cre, sub[0].Crtr) {
		// TODO update subscription object valid flag with failure reason
		return nil
	}

	// TODO validate whether creator addresses represent valid content creators.
	//
	//     https://github.com/NaoNaoOnline/issues/issues/6
	//

	// We modify the data in Task.Sync to forward the validity of the current
	// subscription we are processing.
	{
		tas.Sync.Set(objectlabel.SubsVerify, strconv.FormatBool(true))
	}

	return nil
}

func addStr(add [3]common.Address) []string {
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
