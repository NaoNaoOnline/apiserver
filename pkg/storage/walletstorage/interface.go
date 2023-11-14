package walletstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
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

	// SearchAddr returns the user IDs for the given addresses.
	//
	//     @inp[0] the wallet addresses to search users
	//     @out[0] the list of user IDs matching the given wallet addresses
	//
	SearchAddr([]string) ([]objectid.ID, error)

	// SearchKind returns the wallet objects for the given user, matching the
	// given wallet kinds, e.g. eth.
	//
	//     @inp[0] the user ID used to search wallets
	//     @inp[1] the wallet kinds under which wallet objects are grouped together
	//     @out[0] the list of wallet objects for the given user, matching the given wallet kinds
	//
	SearchKind(objectid.ID, []string) ([]*Object, error)

	// SearchWllt returns the wallet objects for the given user, matching the
	// given wallet IDs.
	//
	//     @inp[0] the user ID used to search wallets
	//     @inp[1] the wallet IDs to search for
	//     @out[0] the list of wallet objects for the given user, matching the given wallet IDs
	//
	SearchWllt(objectid.ID, []objectid.ID) ([]*Object, error)

	// UpdatePtch modifies the existing wallet objects by applying the given
	// RFC6902 JSON-Patches to the underlying JSON documents. The list items are
	// used according to their respective indices, e.g. the second patch is
	// applied to the second object.
	//
	//     @inp[0] the list of wallet objects to modify
	//     @inp[1] the list of RFC6902 compliant JSON-Patches
	//     @out[0] the list of modified wallet objects
	//     @out[1] the list of operation states related to the modified wallet objects
	//
	UpdatePtch([]*Object, PatchSlicer) ([]*Object, []objectstate.String, error)

	// UpdateSign modifies the existing wallet objects by solving the signature
	// verification challenge again. On success object.intern.last is set to the
	// time of execution.
	//
	//     @inp[0] the list of wallet objects to modify
	//     @out[0] the list of modified wallet objects
	//     @out[1] the list of operation states related to the modified wallet objects
	//
	UpdateSign([]*Object) ([]*Object, []objectstate.String, error)
}
