package walletstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
)

type Interface interface {
	// CreateXtrn persists new wallet objects for the provided signatures, if
	// those signatures are valid.
	//
	//     @inp[0] the wallet objects providing signatures
	//     @out[0] the wallet objects mapped to their internal wallet IDs
	//
	CreateXtrn([]*Object) ([]*Object, error)

	// Delete purges the given wallet objects.
	//
	//     @inp[0] the wallet objects to delete
	//     @out[0] the list of operation states related to the purged wallet objects
	//
	Delete([]*Object) ([]objectstate.String, error)

	// SearchKind returns the wallet objects for the given user, matching the
	// given wallet kinds, e.g. eth.
	//
	//     @inp[0] the user ID used to search wallets
	//     @inp[1] the wallet kinds under which wallet objects are grouped together
	//     @out[0] the list of wallet objects for the given user, matching the given wallet kinds
	//
	SearchKind(objectid.String, []string) ([]*Object, error)

	// SearchWllt returns the wallet objects for the given user, matching the
	// given wallet IDs.
	//
	//     @inp[0] the user ID used to search wallets
	//     @inp[1] the wallet IDs to search for
	//     @out[0] the list of wallet objects for the given user, matching the given wallet IDs
	//
	SearchWllt(objectid.String, []objectid.String) ([]*Object, error)

	// Update modifies the existing wallet objects by solving the signature
	// verification challenge again. On success object.intern.last is set to the
	// time of execution.
	//
	//     @inp[0] the list of wallet objects to modify
	//     @out[0] the list of operation states related to the modified wallet objects
	//
	Update([]*Object) ([]objectstate.String, error)
}
