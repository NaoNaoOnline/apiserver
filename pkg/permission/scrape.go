package permission

import (
	"math/big"
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/contract/policycontract"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (p *Permission) ScrapeRcrd(tas *task.Task, bud *budget.Budget) error {
	var err error

	var cid string
	var cnt string
	var rpc string
	{
		cid = tas.Meta.Get(objectlabel.PlcyChanid)
		cnt = tas.Meta.Get(objectlabel.PlcyCntrct)
		rpc = tas.Meta.Get(objectlabel.PlcyRpcUrl)
	}

	var eth *ethclient.Client
	{
		eth, err = ethclient.Dial(rpc)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var pcn *policycontract.Policy
	{
		pcn, err = policycontract.NewPolicy(common.HexToAddress(cnt), eth)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var upd bool

	// Try to lookup policy records onchain up to three times, in case
	// intermittend changes happened while we are searching for the full state
	// representation at current.
	for i := 0; i < 3; i++ {
		// Before starting our cursor based iteration we fetch the current block
		// recorded inside the Policy contract. We want to receive the same block
		// height before and after our complete search process.
		var fir *big.Int
		{
			fir, err = pcn.SearchBlocks(nil)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// The cursor based iteration does always start with the cursor zero.
		var cur *big.Int
		{
			cur = big.NewInt(0)
		}

		var lis []*policystorage.Object
		for {
			// For each iteration we try to claim an operational budget to make sure
			// that the task does not exceed its resource limits.
			{
				clm := bud.Claim(1)
				if clm != 1 {
					return tracer.Mask(taskBudgetLimitError)
				}
			}

			// Try to fetch a list of policy records using the current cursor and the
			// current block number. If the call succeeds we will receive a list of
			// policy records and the updated cursor for the next call. If the block
			// number became outdated meanwhile the call will fail.
			var rec []policycontract.TripleRecord
			{
				cur, rec, err = pcn.SearchRecord(nil, cur, fir)
				if err != nil {
					return tracer.Mask(err)
				}
			}

			for _, x := range rec {
				lis = append(lis, &policystorage.Object{
					Acce: int64(x.Acc.Uint64()),
					ChID: []int64{musNum(cid)},
					Memb: x.Mem.Hex(),
					Syst: int64(x.Sys.Uint64()),
				})
			}

			// As soon as we receive the final cursor zero the cursor based iteration
			// finished, implying that we received a complete snapshot of the policy
			// records inside the Policy smart contract.
			if cur.Uint64() == 0 {
				break
			}
		}

		// Once we gathered the most recent snapshot of the current state of policy
		// records inside the Policy smart contract we fetch the internally recorded
		// block height again.
		var sec *big.Int
		{
			sec, err = pcn.SearchBlocks(nil)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Our received snapshot is valid if, and only if, the first and the second
		// block number match exactly. If they do not match, changes have been made
		// inside the Policy contract while we were reading its internal state,
		// forcing us to fetch the whole state all over again once more.
		if fir.Uint64() != sec.Uint64() {
			continue
		}

		// At this point we must have received the valid representation of the
		// policy records on the blockchain we are responsible for tracking. So we
		// buffer our result via the policy storage implementation for distributed
		// use.
		{
			err = p.pol.CreateBffr(lis)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			upd = true
		}

		break
	}

	if !upd {
		return tracer.Mask(policyScrapeFailedError)
	}

	return nil
}

func musNum(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return num
}
