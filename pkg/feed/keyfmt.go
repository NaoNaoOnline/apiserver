package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
)

// eveRea computes storage keys that allow us to search for event IDs in
// CreateRule that are associated to the given rule object.
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

// eveWri computes storage keys that allow us to persist event IDs in CreateEvnt
// that are associated to the given event object.
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

// rulRea computes storage keys that allow us to search for rule IDs in
// CreateEvnt that are associated to the given event object.
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

// rulWri computes storage keys that allow us to persist rule IDs in CreateRule
// that are associated to the given rule object.
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
