package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
)

// TODO read event IDs from these in CreateRule
func eveRea(rob *rulestorage.Object) []string {
	var key []string

	if rob.Kind == "cate" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.EveCat(x))
		}
	}

	if rob.Kind == "evnt" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.EveEve(x))
		}
	}

	if rob.Kind == "host" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.EveHos(x))
		}
	}

	if rob.Kind == "user" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.EveUse(x))
		}
	}

	return key
}

func eveRul(str string) string {
	return keyfmt.EveRul(objectid.ID(str))
}

// TODO write event IDs to these in CreateEvnt
func eveWri(eob *eventstorage.Object) []string {
	var key []string

	for _, x := range append(eob.Bltn, eob.Cate...) {
		key = append(key, keyfmt.EveCat(x))
	}

	{
		key = append(key, keyfmt.EveEve(eob.Evnt))
	}

	for _, x := range eob.Host {
		key = append(key, keyfmt.EveHos(x))
	}

	{
		key = append(key, keyfmt.EveUse(eob.User))
	}

	return key
}

func fmtFnc[T string | objectid.ID](lis []T, fnc func(T) string) []string {
	var key []string

	for _, x := range lis {
		key = append(key, fnc(x))
	}

	return key
}

// TODO read rule IDs from these in CreateEvnt
func rulRea(eob *eventstorage.Object) []string {
	var key []string

	for _, x := range append(eob.Bltn, eob.Cate...) {
		key = append(key, keyfmt.RulCat(x))
	}

	{
		key = append(key, keyfmt.RulEve(eob.Evnt))
	}

	for _, x := range eob.Host {
		key = append(key, keyfmt.RulHos(x))
	}

	{
		key = append(key, keyfmt.RulUse(eob.User))
	}

	return key
}

// TODO write rule IDs to these in CreateRule
func rulWri(rob *rulestorage.Object) []string {
	var key []string

	if rob.Kind == "cate" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.RulCat(x))
		}
	}

	if rob.Kind == "evnt" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.RulEve(x))
		}
	}

	if rob.Kind == "host" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.RulHos(x))
		}
	}

	if rob.Kind == "user" {
		for _, x := range append(rob.Excl, rob.Incl...) {
			key = append(key, keyfmt.RulUse(x))
		}
	}

	return key
}
