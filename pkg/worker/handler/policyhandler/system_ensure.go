package policyhandler

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/contract/policycontract"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var lcm int64
	var lcs int64
	var ldm int64
	var lds int64
	{
		lcm, err = h.updateCrMe(tas, bud)
		if err != nil {
			return tracer.Mask(err)
		}
		lcs, err = h.updateCrSy(tas, bud)
		if err != nil {
			return tracer.Mask(err)
		}
		ldm, err = h.updateDeMe(tas, bud)
		if err != nil {
			return tracer.Mask(err)
		}
		lds, err = h.updateDeSy(tas, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		tas.Sync.Set(objectlabel.PlcyLatest, fmt.Sprintf("%d,%d,%d,%d", lcm, lcs, ldm, lds))
	}

	return nil
}

func (h *SystemHandler) createPlcy(raw types.Log, kin string, sys *big.Int, mem common.Address, acc *big.Int) error {
	var err error

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var hdr *types.Header
	{
		hdr, err = h.eth.HeaderByHash(ctx, raw.BlockHash)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var tnx *types.Transaction
	{
		tnx, _, err = h.eth.TransactionByHash(ctx, raw.BlockHash)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var frm common.Address
	{
		frm, err = h.eth.TransactionSender(ctx, tnx, raw.BlockHash, raw.TxIndex)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var obj *policystorage.Object
	{
		obj = &policystorage.Object{
			Acce: int64(acc.Uint64()),
			Blck: []int64{int64(raw.BlockNumber)},
			ChID: []int64{int64(tnx.ChainId().Uint64())},
			From: []string{frm.Hex()},
			Hash: []string{raw.BlockHash.Hex()},
			Kind: kin,
			Memb: mem.Hex(),
			Syst: int64(sys.Uint64()),
			Time: []time.Time{time.Unix(int64(hdr.Time), 0)},
		}
	}

	{
		_, err = h.pol.Create([]*policystorage.Object{obj})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *SystemHandler) updateCrMe(tas *task.Task, bud *budget.Budget) (int64, error) {
	var err error

	var spl []string
	{
		spl = strings.Split(tas.Sync.Get(objectlabel.PlcyLatest), ",")
	}

	if len(spl) != 4 {
		return 0, tracer.Mask(policyLatestInvalidError)
	}

	var lcm int64
	{
		lcm, err = strconv.ParseInt(spl[0], 10, 64)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	var ite *policycontract.PolicyCreateMemberIterator
	{
		opt := &bind.FilterOpts{
			// Start configures the "FromBlock" log filter. In case the events are
			// being indexed the very first time, lcm will be 0 and filtering starts
			// from the earliest possible block related to the referenced contract.
			Start: uint64(lcm),
		}

		ite, err = h.pcn.FilterCreateMember(opt, nil, nil, nil)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	{
		defer ite.Close()
	}

	for ite.Next() {
		{
			clm := bud.Claim(1)
			if clm != 1 {
				return lcm, nil
			}
		}

		{
			lcm = int64(ite.Event.Raw.BlockNumber)
		}

		{
			err := h.createPlcy(ite.Event.Raw, "CreateMember", ite.Event.Sys, ite.Event.Mem, ite.Event.Acc)
			if err != nil {
				return 0, tracer.Mask(err)
			}
		}
	}

	{
		err := ite.Error()
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	return lcm, nil
}

func (h *SystemHandler) updateCrSy(tas *task.Task, bud *budget.Budget) (int64, error) {
	var err error

	var spl []string
	{
		spl = strings.Split(tas.Sync.Get(objectlabel.PlcyLatest), ",")
	}

	if len(spl) != 4 {
		return 0, tracer.Mask(policyLatestInvalidError)
	}

	var lcs int64
	{
		lcs, err = strconv.ParseInt(spl[1], 10, 64)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	var ite *policycontract.PolicyCreateSystemIterator
	{
		opt := &bind.FilterOpts{
			// Start configures the "FromBlock" log filter. In case the events are
			// being indexed the very first time, lcs will be 0 and filtering starts
			// from the earliest possible block related to the referenced contract.
			Start: uint64(lcs),
		}

		ite, err = h.pcn.FilterCreateSystem(opt, nil, nil, nil)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	{
		defer ite.Close()
	}

	for ite.Next() {
		{
			clm := bud.Claim(1)
			if clm != 1 {
				return lcs, nil
			}
		}

		{
			lcs = int64(ite.Event.Raw.BlockNumber)
		}

		{
			err := h.createPlcy(ite.Event.Raw, "CreateSystem", ite.Event.Sys, ite.Event.Mem, ite.Event.Acc)
			if err != nil {
				return 0, tracer.Mask(err)
			}
		}
	}

	{
		err := ite.Error()
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	return lcs, nil
}

func (h *SystemHandler) updateDeMe(tas *task.Task, bud *budget.Budget) (int64, error) {
	var err error

	var spl []string
	{
		spl = strings.Split(tas.Sync.Get(objectlabel.PlcyLatest), ",")
	}

	if len(spl) != 4 {
		return 0, tracer.Mask(policyLatestInvalidError)
	}

	var ldm int64
	{
		ldm, err = strconv.ParseInt(spl[2], 10, 64)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	var ite *policycontract.PolicyDeleteMemberIterator
	{
		opt := &bind.FilterOpts{
			// Start configures the "FromBlock" log filter. In case the events are
			// being indexed the very first time, ldm will be 0 and filtering starts
			// from the earliest possible block related to the referenced contract.
			Start: uint64(ldm),
		}

		ite, err = h.pcn.FilterDeleteMember(opt, nil, nil, nil)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	{
		defer ite.Close()
	}

	for ite.Next() {
		{
			clm := bud.Claim(1)
			if clm != 1 {
				return ldm, nil
			}
		}

		{
			ldm = int64(ite.Event.Raw.BlockNumber)
		}

		{
			err := h.createPlcy(ite.Event.Raw, "DeleteMember", ite.Event.Sys, ite.Event.Mem, ite.Event.Acc)
			if err != nil {
				return 0, tracer.Mask(err)
			}
		}
	}

	{
		err := ite.Error()
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	return ldm, nil
}

func (h *SystemHandler) updateDeSy(tas *task.Task, bud *budget.Budget) (int64, error) {
	var err error

	var spl []string
	{
		spl = strings.Split(tas.Sync.Get(objectlabel.PlcyLatest), ",")
	}

	if len(spl) != 4 {
		return 0, tracer.Mask(policyLatestInvalidError)
	}

	var lds int64
	{
		lds, err = strconv.ParseInt(spl[3], 10, 64)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	var ite *policycontract.PolicyDeleteSystemIterator
	{
		opt := &bind.FilterOpts{
			// Start configures the "FromBlock" log filter. In case the events are
			// being indexed the very first time, lds will be 0 and filtering starts
			// from the earliest possible block related to the referenced contract.
			Start: uint64(lds),
		}

		ite, err = h.pcn.FilterDeleteSystem(opt, nil, nil, nil)
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	{
		defer ite.Close()
	}

	for ite.Next() {
		{
			clm := bud.Claim(1)
			if clm != 1 {
				return lds, nil
			}
		}

		{
			lds = int64(ite.Event.Raw.BlockNumber)
		}

		{
			err := h.createPlcy(ite.Event.Raw, "DeleteSystem", ite.Event.Sys, ite.Event.Mem, ite.Event.Acc)
			if err != nil {
				return 0, tracer.Mask(err)
			}
		}
	}

	{
		err := ite.Error()
		if err != nil {
			return 0, tracer.Mask(err)
		}
	}

	return lds, nil
}
